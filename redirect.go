package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	destination := os.Getenv("DESTINATION")
	status := os.Getenv("STATUS")
	if status == "" {
		status = "307"
	}
	code, err := strconv.Atoi(status)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := &http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, destination+req.URL.Path, code)
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
