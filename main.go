package main

import (
	"flag"
	"fmt"
	"github.com/bernljung/go-tts"
	"log"
	"net/http"
	"time"
)

var toSay []tts.TTS

func start() {
	for range time.Tick(100 * time.Millisecond) {
		if len(toSay) > 0 {
			toSay[0].Play()
			toSay = toSay[1:]
		}
	}
}

func queue(tl, q string) {
	t := tts.TTS{Tl: tl, Q: q}
	toSay = append(toSay, t)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		q := r.FormValue("q")
		tl := r.FormValue("tl")
		if q != "" && (tl == "sv" || tl == "en") {
			go queue(tl, q)
			fmt.Fprint(w, Response{Success: true, Message: "Queued"})
		} else {
			fmt.Fprint(w, Response{Success: false, Message: "You know what you did... I need q and tl."})
		}
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	port := flag.Int("port", 8000, "port number")
	flag.Parse()
	go start()
	http.HandleFunc("/post", handler)

	message := fmt.Sprintf("Starting server on :%v", *port)
	log.Println(message)
	address := fmt.Sprintf(":%v", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
