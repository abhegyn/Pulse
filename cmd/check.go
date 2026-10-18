package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [sites...]",
	Args:  cobra.ArbitraryArgs,
	Short: "Checks the status of the http request provided",
	Run: func(cmd *cobra.Command, args []string) {
		// contents of the command
		for _, site := range args {
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println("Request failed: ", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("for:", site)
			fmt.Println("status:", resp.Status)
			fmt.Println("status code: ", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
