package main

import (
	"net"
	"net/http"
	"os"

	"github.com/tenntenn/natureremo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	hostport := net.JoinHostPort("", port)

	ncli := natureremo.NewClient(os.Getenv("NATUREREMO_TOKEN"))
	sch := NewScheduler(ncli)
	webhookToken := os.Getenv("WEBHOOK_TOKEN")
	server := NewServer(sch, webhookToken)

	http.ListenAndServe(hostport, server)
}
