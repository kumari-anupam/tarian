package protoqueue

import (
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var logger *zap.SugaredLogger

func init() {
	l, err := zap.NewProduction()

	if err != nil {
		panic("Can not create logger")
	}

	logger = l.Sugar()
}

func SetLogger(l *zap.SugaredLogger) {
	logger = l
}

type JetStreamConnection struct {
	NATSConn  *nats.Conn
	JSContext nats.JetStreamContext
}

type JetStream struct {
	URL          string
	Options      []nats.Option
	StreamName   string
	Conn         JetStreamConnection
	Subscription *nats.Subscription
	channel      chan any
}

func NewJetstream(url string, options []nats.Option, streamName string) (*JetStream, error) {
	channel := make(chan any, 1000)
	return &JetStream{
		URL:        url,
		Options:    options,
		StreamName: streamName,
		channel:    channel,
	}, nil
}

func (j *JetStream) Connect() error {
	nc, err := nats.Connect(j.URL, j.Options...)
	if err != nil {
		logger.Errorw("failed to connect to NATS server", zap.Error(err))
		return err
	}

	logger.Infow("connected to NATS server", zap.Error(err))

	jetStreamContext, err := nc.JetStream()
	if err != nil {
		logger.Errorw("failed to get jetstream context", zap.Error(err))
		return err
	}

	logger.Infow("successfully get jetstream context", zap.Error(err))

	j.Conn = JetStreamConnection{NATSConn: nc, JSContext: jetStreamContext}
	return nil
}

func (j *JetStream) Init(streamConfig nats.StreamConfig) error {
	if err := j.CreateStreamIfNotExist(streamConfig); err != nil {
		return err
	}

	if _, err := j.CreateConsumer(); err != nil {
		return err
	}

	return j.CreateSubscription()
}

func (j *JetStream) CreateStreamIfNotExist(streamConfig nats.StreamConfig) error {
	if j.Conn.JSContext == nil {
		return errors.New("can not create stream due to nil connection")
	}

	var err error

	streamInfo, err := j.Conn.JSContext.StreamInfo(j.StreamName)
	if streamInfo != nil && err == nil {
		logger.Infow("not going to create stream as it already exists", "stream", j.StreamName)
		return nil
	}

	if err != nil && err != nats.ErrStreamNotFound {
		logger.Warnw("error calling jetstream StreamInfo", "stream", j.StreamName, "err", err)
	}

	logger.Infow("will use this config to create stream", "stream", j.StreamName, "config", streamConfig)

	_, err = j.Conn.JSContext.AddStream(&streamConfig)
	if err != nil {
		errStr := fmt.Sprintf("error while creating stream %s. %s", j.StreamName, err)
		return fmt.Errorf(errStr)
	}

	logger.Infow("stream created", "stream", j.StreamName)
	return nil
}

func (j *JetStream) CreateConsumer() (*nats.ConsumerInfo, error) {
	return j.Conn.JSContext.AddConsumer(j.StreamName, &nats.ConsumerConfig{
		Durable:        j.StreamName + "-TODO",
		DeliverSubject: j.StreamName + "-DeliverSubject",
		DeliverGroup:   j.StreamName + "-TODO",
		AckPolicy:      nats.AckExplicitPolicy,
	})
}

func (j *JetStream) CreateSubscription() error {
	subscription, err := j.Conn.NATSConn.QueueSubscribeSync(j.StreamName+"-DeliverSubject", j.StreamName+"-TODO")
	j.Subscription = subscription

	return err
}

func (j *JetStream) Publish(queuedMessage proto.Message) error {
	data, err := proto.Marshal(queuedMessage)
	if err != nil {
		return err
	}

	_, err = j.Conn.JSContext.Publish(j.StreamName, data)
	if err != nil {
		return err
	}

	return nil
}

func (j *JetStream) NextMessage(message proto.Message) (proto.Message, error) {
	msg, err := j.Subscription.NextMsg(1 * time.Hour)

	if errors.Is(err, nats.ErrTimeout) {
		return nil, fmt.Errorf("no message in the queue until timeout is reached")
	}

	if err != nil {
		return nil, err
	}

	msg.Ack()
	err = proto.Unmarshal(msg.Data, message)

	return message, err
}