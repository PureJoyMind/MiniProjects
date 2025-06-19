package excercises

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", readErr)
			continue
		}
		resp.Body.Close()
		fmt.Printf("%s", body)

	}
}

func FetchWithCopy() {
	const (
		prefix = "https://"
	)

	for _, arg := range os.Args[1:] {

		if !strings.HasPrefix(arg, prefix) {
			arg = prefix + arg
		}

		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		_, readErr := io.Copy(os.Stdout, resp.Body)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", readErr)
			continue
		}
		resp.Body.Close()
		fmt.Printf("status: %s\n", resp.Status)

		fmt.Println(arg)
	}
}

func StartServerOnPort(port int32) {
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	})

	server := fmt.Sprintf("localhost:%d", port)
	log.Fatal(http.ListenAndServe(server, nil))
}
