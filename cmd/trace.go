/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP",
	Long:  "Trace IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide an IP address")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// {
//   "ip": "1.1.1.1",
//   "hostname": "one.one.one.one",
//   "anycast": true,
//   "city": "Los Angeles",
//   "region": "California",
//   "country": "US",
//   "loc": "34.0522,-118.2437",
//   "org": "AS13335 Cloudflare, Inc.",
//   "postal": "90076",
//   "timezone": "America/Los_Angeles",
//   "readme": "https://ipinfo.io/missingauth"
// }

type IPData struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

var result IPData

func showData(ipaddr string) {
	url := "https://ipinfo.io/" + ipaddr + "/geo"
	data, err := getData(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func getData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
