package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/unix", getUnixTime)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("[%s] Server closed\n", formattedTime())
	} else {
		fmt.Printf("[%s] Error while starting server: %s\n", formattedTime(), err)
		os.Exit(1)
	}
}
func getRoot(w http.ResponseWriter, r *http.Request) {
	log(r)
	if _, err := io.WriteString(w,
		fmt.Sprintf("Unix time is %d UTC\n\nMade by @Dannode36",
			time.Now().UTC().Unix())); err != nil {
		return
	}
}
func getUnixTime(w http.ResponseWriter, r *http.Request) {
	log(r)
	if _, err := io.WriteString(w,
		fmt.Sprintf("%d",
			time.Now().UTC().Unix())); err != nil {
		return
	}
}
func log(r *http.Request) {
	if !strings.Contains(r.URL.String(), "/favicon.ico") {
		fmt.Printf("[%s] %s -> %s %s \n", formattedTime(), r.RemoteAddr, r.Method, r.URL)
	}
}
func formattedTime() string {
	return time.Now().Format("02-01-2006 15:04:05")
}
