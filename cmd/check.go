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

	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(x)

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

		limit := 5
		sem := make(chan struct{}, limit)

		for _, site := range args {

			sem <- struct{}{}
			wg.Add(1)

			go func(url string) {

				defer func() {
					<-sem
				}()

				checkUrl(url, &wg)

			}(site)

		}

		wg.Wait() // tells the main to wait for routines to finish
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
