package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func checkUrl(x string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(x)

	start := time.Now()

	if err != nil {
		fmt.Println("Request failed: ", err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	fmt.Println("for:", x)
	fmt.Println("status:", resp.Status, "status code: ", resp.StatusCode, "time taken: ", elapsed)
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [sites...]",
	Args:  cobra.ArbitraryArgs,
	Short: "Checks the status of the http request provided",
	Run: func(cmd *cobra.Command, args []string) {

		var wg sync.WaitGroup

		for _, site := range args {

			wg.Add(1)
			go checkUrl(site, &wg)

		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
