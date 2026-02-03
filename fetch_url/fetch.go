package fetchurl

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchUrl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
