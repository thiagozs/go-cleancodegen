package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagozs/go-cleancodegen/internal/app/application/usecase"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Genereate a folders structure for a clean code project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		variablesHandlers(cmd)

		if err := usecase.CreateScaffold(); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().StringVar(&Project, "project", "project", "Project name")
	generateCmd.PersistentFlags().StringVar(&Path, "path", "/tmp", "Path of target of project what will be create")

	generateCmd.MarkPersistentFlagRequired("project")
	generateCmd.MarkPersistentFlagRequired("path")

}
