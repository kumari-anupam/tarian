package ebpf

import (
	"bytes"
	"encoding/binary"

	"github.com/aquasecurity/libbpfgo"
	"go.uber.org/zap"

	_ "embed"
)

var bpfObjName = "capture_exec.bpf.o"

//go:embed capture_exec.bpf.o
var captureExecBpfObj []byte

type BpfExecEvent struct {
	Pid      uint32
	Comm     [80]uint8
	Filename [1024]uint8
}

type BpfCaptureExec struct {
	shouldClose bool

	bpfEventsChan  chan []byte
	execEventsChan chan BpfExecEvent

	bpfModule     *libbpfgo.Module
	bpfProg       *libbpfgo.BPFProg
	bpfRingBuffer *libbpfgo.RingBuffer

	logger *zap.SugaredLogger
}

func NewBpfCaptureExec() (*BpfCaptureExec, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	b := &BpfCaptureExec{
		bpfEventsChan:  make(chan []byte, 1000),
		execEventsChan: make(chan BpfExecEvent, 1000),
		logger:         l.Sugar(),
	}

	err = b.loadBpfObject()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *BpfCaptureExec) SetLogger(l *zap.SugaredLogger) {
	b.logger = l
}

func (b *BpfCaptureExec) loadBpfObject() error {
	var err error
	b.bpfModule, err = libbpfgo.NewModuleFromBuffer(captureExecBpfObj, bpfObjName)
	if err != nil {
		return err
	}

	b.bpfModule.BPFLoadObject()

	b.bpfRingBuffer, err = b.bpfModule.InitRingBuf("events", b.bpfEventsChan)
	if err != nil {
		return err
	}

	b.bpfProg, err = b.bpfModule.GetProgram("enter_execve")
	if err != nil {
		return err
	}

	_, err = b.bpfProg.AttachTracepoint("syscalls", "sys_enter_execve")
	if err != nil {
		return err
	}

	return nil
}

func (b *BpfCaptureExec) Start() {
	b.bpfRingBuffer.Start()

	for {
		evt := <-b.bpfEventsChan

		if b.shouldClose {
			break
		}

		var bpfExecEvent BpfExecEvent
		if err := binary.Read(bytes.NewBuffer(evt), binary.LittleEndian, &bpfExecEvent); err != nil {
			b.logger.Errorw("error parsing ringbuf event", "err", err)
			continue
		}

		b.execEventsChan <- bpfExecEvent
	}
}

func (b *BpfCaptureExec) Close() {
	b.shouldClose = true
	b.bpfRingBuffer.Close()
	b.bpfModule.Close()
}

func (b *BpfCaptureExec) GetExecEventsChannel() chan BpfExecEvent {
	return b.execEventsChan
}