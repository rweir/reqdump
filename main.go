package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

//nolint:gochecknoglobals
var port = flag.String("p", "8080", "TCP port to bind to")

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Add("content-type", "text/plain")

	slog.InfoContext(ctx, "New connection", slog.String("remoteAddr", r.RemoteAddr), slog.String("proto", r.Proto))

	fmt.Fprintf(w, "Connection details:\n\n")

	fmt.Fprintf(w, "Peer adddress: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Protocol version: %s\n", r.Proto)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Headers:\n\n")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
			slog.InfoContext(ctx, "header", slog.String(name, value))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))

	//nolint:exhaustruct,gomnd
	srv := &http.Server{
		Addr:              ":" + *port,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           mux,
	}

	slog.Error("Listener stopped", slog.Any("err", srv.ListenAndServe()))
}
