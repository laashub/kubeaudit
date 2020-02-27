package commands

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var autofixConfig struct {
	outFile string
}

func autofix(cmd *cobra.Command, args []string) {
	f, err := os.Create(autofixConfig.outFile)
	if err != nil {
		log.WithError(err).Fatal("Error opening out file")
	}

	report := getReport()
	err = report.Fix(f)
	if err != nil {
		log.WithError(err).Fatal("Error fixing manifest")
	}
}

var autofixCmd = &cobra.Command{
	Use:   "autofix",
	Short: "Automagically make a manifest secure",
	Long: `This command automatically fixes all identified security issues for a given manifest
(ie. all ERROR results generated by 'kubeaudit all'). If no output file is specified using the -o flag,
the source manifest will be modified.

Example usage:
kubeaudit autofix -f /path/to/yaml
kubeaudit autofix -f /path/to/yaml -o /path/for/fixed/yaml`,
	Run: autofix,
}

func init() {
	RootCmd.AddCommand(autofixCmd)
	autofixCmd.Flags().StringVarP(&autofixConfig.outFile, "outfile", "o", "", "file to write fixed manifest to")
}