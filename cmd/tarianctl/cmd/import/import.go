package importCommand

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/kube-tarian/tarian/cmd/tarianctl/cmd/flags"
	"github.com/kube-tarian/tarian/cmd/tarianctl/util"
	"github.com/kube-tarian/tarian/pkg/log"
	"github.com/kube-tarian/tarian/pkg/tarianctl/client"
	"github.com/kube-tarian/tarian/pkg/tarianpb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type importCommand struct {
	globalFlags *flags.GlobalFlags
	logger      *logrus.Logger
}

// importCmd represents the import command
func NewImportCommand(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &importCommand{
		globalFlags: globalFlags,
		logger:      log.GetLogger(),
	}

	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import resources to the Tarian Server.",
		Long:  "Import resources to the Tarian Server.",
		RunE:  cmd.run,
	}

	return importCmd
}

func (c *importCommand) run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		err := errors.New("specify file paths to import")
		return fmt.Errorf("import: %w", err)
	}

	files := []*os.File{}
	for _, path := range args {
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("import: failed to open file %s: %w", path, err)
		}

		files = append(files, f)
	}

	opts, err := util.ClientOptionsFromCliContext(c.logger, c.globalFlags)
	if err != nil {
		return fmt.Errorf("import: %w", err)
	}

	client, err := client.NewConfigClient(c.globalFlags.ServerAddr, opts...)
	if err != nil {
		return fmt.Errorf("import: %w", err)
	}

	for _, f := range files {
		err := c.importFile(f, client)
		c.logger.Warn(err)
		f.Close()
	}
	return nil
}

func (c *importCommand) importFile(f *os.File, client tarianpb.ConfigClient) error {
	decoder := yaml.NewDecoder(f)
	imported := 0

	for {
		var constraint tarianpb.Constraint
		err := decoder.Decode(&constraint)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("import: failed to decode yaml: %w", err)
		}

		req := &tarianpb.AddConstraintRequest{Constraint: &constraint}
		response, err := client.AddConstraint(context.Background(), req)
		if err != nil {
			return fmt.Errorf("import: failed to add constraint: %w", err)
		}

		if response.GetSuccess() {
			imported++
		}
	}

	if imported > 0 {
		c.logger.Info("Imported constraint successfully")
	} else {
		c.logger.Warn("No constraint was imported")
	}
	return nil
}