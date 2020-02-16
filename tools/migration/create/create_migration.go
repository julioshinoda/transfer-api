package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func main() {
	migrationTitle := flag.String("title", "", "Migration title")
	flag.Parse()

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	title, _, _ := transform.String(t, *migrationTitle)

	title = strings.Replace(title, " ", "_", -1)

	nsec := time.Now().UnixNano() / int64(time.Millisecond) // number of nanoseconds since January 1, 1970 UTC
	err := ioutil.WriteFile("db/migrations/"+strconv.Itoa(int(nsec))+"_"+title+".up.sql", []byte(""), 0777)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Println("Created file:", "db/migrations/"+strconv.Itoa(int(nsec))+"_"+title+".up.sql")
}
