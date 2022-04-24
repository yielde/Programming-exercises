package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/index", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe("0.0.0.0:8888", mux); err != nil {
		log.Fatal("Http server start failed", err.Error())
	}
}

func setEnv() {
	os.Setenv("VERSION", "v0.0.1")
}

func getEnv() string {
	return os.Getenv("VERSION")
}

func index(w http.ResponseWriter, r *http.Request) {
	setEnv()
	version := getEnv()
	fmt.Println("OS VERSION ->", version)
	w.Header().Set("VERSION", version)
	for k, v := range r.Header {
		fmt.Printf("key: %v, value: %v\n", k, v)
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	io.WriteString(w, "<h1>Hello World</h1>")
	accessStat := getAcessIP(r)
	fmt.Printf("CLIENT_IP: %s -> Respose code: %d\n", accessStat, http.StatusOK)
}

func getAcessIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("status", "200")
	io.WriteString(w, "<h1>Access Succeed!</h1>")
}
