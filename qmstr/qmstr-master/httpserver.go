package main

import (
	"context"
	"fmt"
	"net/http"
)

var closeServer chan interface{}

func handleQuitRequest(w http.ResponseWriter, r *http.Request) {
	// nothing to do except quit:
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Bye now.\n")
	Info.Printf("handleQuitRequest: quit request received.")
	closeServer <- nil
}

func startHTTPServer() chan string {
	address := ":8080"
	server := &http.Server{Addr: address}
	http.HandleFunc("/quit", handleQuitRequest)

	Info.Printf("starting HTTP server on address %s", address)
	channel := make(chan string)
	go func() {
		err := server.ListenAndServe()
		server = nil
		if err == http.ErrServerClosed {
			channel <- fmt.Sprintf("startHTTPServer: server closed.")
		} else if err != nil {
			channel <- fmt.Sprintf("startHTTPServer: exiting with error: %s", err.Error())
		} else {
			channel <- "startHTTPServer: retreating coordinatedly."
		}
	}()

	closeServer = make(chan interface{})
	go func() {
		<-closeServer
		Info.Printf("shutting down HTTP server on address %s", address)
		if server != nil {
			if err := server.Shutdown(context.Background()); err != nil {
				panic(err) // failure/timeout shutting down the server gracefully
			}
		} else {
			Log.Printf("stopHTTPServer: server shutdown requested, but server is not running.")
		}
		close(closeServer)
		closeServer = nil
	}()
	return channel
}
