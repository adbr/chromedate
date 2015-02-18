// 2015-02-18 Adam Bryt

// Program chromedate konwertuje date z formatu używanego w chrome
// bookmarks. Chrome w pliku bookmarks, w polach "date_added" używa
// daty w formacie Windows NT timestamp, czyli liczby mikrosekund od
// 1601-01-01.
//
// Sposób użycia:
//
//	chromedate timestamp
//		timestamp: liczba mikrosekund od 1601-01-01`
//
// Przykład:
//
//	chromedate 13068596981313705
//	2015-02-16 21:49:41.313705 +0000 UTC
//
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	microsecond = 1
	milisecond  = 1000 * microsecond
	second      = 1000 * milisecond // liczba mikrosekund w sekundzie
)

const usageStr = `usage: chromedate timestamp
	timestamp: liczba mikrosekund od 1601-01-01`

func usage() {
	fmt.Fprintln(os.Stderr, usageStr)
	os.Exit(1)
}

// chromedate konwertuje timestamp ts do standardowej wartości time.Time.
// Argument ts jest liczbą mikrosekund od 1601-01-01 (Windows NT timestamp).
func chromedate(ts int64) time.Time {
	s := ts / second
	us := ts % second
	ns := us * 1000
	d := time.Date(1601, 1, 1, 0, 0, int(s), int(ns), time.UTC)
	return d
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	s := os.Args[1]
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	d := chromedate(ts)
	fmt.Println(d)
}
