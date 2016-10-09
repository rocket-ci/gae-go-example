package main

import (
	"net/http"
	"log"
	"os"
	"net/http/httputil"
	"fmt"
	"os/signal"
	"syscall"
	"net"
)

func routes() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/checkin", handle)
	return
}

func server(addr string) (listener net.Listener, ch chan error) {
	ch = make(chan error)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		mux := routes()
		ch <- http.Serve(listener, mux)
	}()

	return
}

func main() {
	if os.Getenv("SLACK_TOKEN") == ""{
		log.Fatal("env SLACK_TOKEN is not available")
	}
	listener, ch := server(":8080")
	fmt.Println("Server started at", listener.Addr())

	// シグナルハンドリング (Ctrl + C)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		log.Println(<-sig)
		listener.Close()
	}()

	log.Println(<-ch)}

func handle(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumb error " + err.Error() )
	}
	if dump != nil {
		log.Print(string(dump))
	}
	switch r.Method {
	case "POST":
		token := os.Getenv("SLACK_TOKEN")
		slack := NewSlack(token)
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		dispatcher, err := NewCheckinDispatcher(r, slack)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			dispatcher.Dispatch()
			w.WriteHeader(http.StatusCreated)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
