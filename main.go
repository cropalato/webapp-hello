//
// main.go
// Copyright (C) 2022 rmelo <Ricardo Melo <rmelo@ludia.com>>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func clientIP(w http.ResponseWriter, r *http.Request) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		return
	}
	forward := r.Header.Get("X-Forwarded-For")
	fmt.Fprintf(w, "<p>IP: %s</p>", ip)
	fmt.Fprintf(w, "<p>Port: %s</p>", port)
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward)
	fmt.Println("/client-ip")
}

func serverIP(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintf(w, "<p>Server IP: %s</p>", ipnet.IP.String())
			}
		}
	}
	fmt.Println("/server-ip")
}

func getVarEnv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("<br>%s<br>", os.Getenv("SECRET")))
	fmt.Println("/get-secret")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
	fmt.Println("/health_check")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/client-ip", clientIP)
	http.HandleFunc("/server-ip", serverIP)
	http.HandleFunc("/get-secret", getVarEnv)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
