package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	var (
		serverName string
		ip         string
	)
	flag.StringVar(&serverName, "server", "", "The name of the server to search the panel for")
	flag.StringVar(&ip, "ip", "", "The IP address of the server")
	flag.Parse()

	if serverName == "" || ip == "" {
		fmt.Println("Please specify the server name and IP address")
		return
	} else if !strings.Contains(ip, ".") {
		fmt.Println("Please specify a valid IP address")
		return
	} else if strings.Contains(ip, ":") {
		strings.TrimSuffix(ip, ":")
	} else if strings.Contains(ip, "/") {
		strings.TrimSuffix(ip, "/")
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 8100; i <= 8990; i += 100 {
		for j := 0; j <= 9; j++ {
			port := i + j*10
			fmt.Printf("Searching for the panel on port %d\n", port)
			resp, err := client.Get(fmt.Sprintf("http://%v:%d/index.html", ip, port))
			if err != nil {
				fmt.Printf("%d timed out\n", port)
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			if strings.Contains(string(body), fmt.Sprintf("Login to %s", serverName)) {
				fmt.Printf("Found %v's panel on port %d\n", serverName, port)
				return
			}
		}
	}
}
