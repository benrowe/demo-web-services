package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
)

var ServeCommand = &cobra.Command{
	Short: "Serve http",
	Long:  "Serve the authenication endpoints",
	Run: func(cmd *cobra.Command, args []string) {
		http.ListenAndServe(":80", nil)
	},
}

func init() {
	RootCommand.AddCommand(ServeCommand)
}
