package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiagozs/go-cleancodegen/internal/app/config"
)

var (
	Project string
	Path    string
)

func variablesHandlers(cmd *cobra.Command) {

	opts := []config.ConfigOpts{}

	if Project != "" {
		opts = append(opts, config.CfgProject(Project))
	}

	if Path != "" {
		opts = append(opts, config.CfgPath(Path))
	}

	config.NewCfgParams(opts...)
}
