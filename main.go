// © 2012 Steve McCoy. Available under the MIT license.

/* getapad takes a URL as an argument and prints the raw text of
the paste on standard output.
*/
package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		os.Stderr.WriteString("I need the URL.\n")
		os.Exit(1)
	}

	url := os.Args[1]
	if !strings.HasSuffix(url, "/raw") {
		url += "/raw" // TODO(mccoyst): Be less dumb; omit the other possible views
	}

	resp, err := http.Get(url)
	maybeDie(err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	maybeDie(err)
}

func maybeDie(err error) {
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
