package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [sites...]",
	Args:  cobra.ArbitraryArgs,
	Short: "Checks the status of the http request provided",
	Run: func(cmd *cobra.Command, args []string) {

		client := http.Client{
			Timeout: 5 * time.Second,
		}
		// instead of calling using the Get directly
		// we use Get from Client as it helps us setup timeout

		for _, site := range args {

			start := time.Now()
			resp, err := client.Get(site)
			if err != nil {
				fmt.Println("Request failed: ", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("for:", site)
			elapsed := time.Since(start)
			fmt.Println("status:", resp.Status, "status code: ", resp.StatusCode, "time taken: ", elapsed)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
