package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var RootCommand = &cobra.Command{
	Use:   "root",
	Short: "Default command",
	Long:  "Default command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("please specify command")
	}}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Errorf("could not execute command: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initCobraConfig)
}

func initCobraConfig() {

}
