// © 2012 Steve McCoy. Available under the MIT license.

/* getpad takes a URL as an argument and prints the raw text of
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

	u, err := url.Parse(os.Args[1])
	maybeDie(err)
	base := path.Base(u.Path)
	if strings.ContainsAny(base, "0123456789") {
		u.Path = path.Join(u.Path, "raw")
	} else if base != "raw" {
		u.Path = path.Join(path.Dir(u.Path), "raw")
	}

	resp, err := http.Get(u.String())
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
