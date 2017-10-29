/*
	Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst.
	Use it instead of ioutil.ReadAll to copy the response bodyto os.Stdout without requiring
	a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Check all necessary parameters before starting
	if (len(os.Args[1:]) == 0) || len(os.Args[2:]) == 0 || len(os.Args[3:]) != 0 {
		fmt.Printf("Usage:" + os.Args[0] + " [url] [file]\n")
	} else {

		url := os.Args[1]
		resp, err := http.Get(url)

		// if there is failure, app should show error.
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		} else {
			// if get http is okay, read all contents
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			// if there is any error occured, display and exit
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: reading %s: %v\n", err)
				os.Exit(1)
			} else {
				// if there is no any error from reading the contents,
				// we will prepare to write contents in file
				var filename string = os.Args[2]
				err := ioutil.WriteFile(filename, b, 0777)

				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: writeing %s: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(resp.Status)
			}
		}
	}
}
