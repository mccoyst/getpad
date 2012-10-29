// © 2012 Steve McCoy. Available under the MIT license.

/* getapad takes a URL as an argument and prints the raw text of
the paste on standard output.
*/
package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		os.Stderr.WriteString("I need the URL.\n")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	maybeDie(err)
	base := path.Base(url.Path)
	if strings.ContainsAny(base, "0123456789") {
		url.Path = path.Join(url.Path, "raw")
	} else if base != "raw" {
		url.Path = path.Join(path.Dir(url.Path), "raw")
	}

	resp, err := http.Get(url.String())
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
