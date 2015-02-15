package main

import (
	"encoding/json"
	"fmt"
	"github.com/bernljung/go-tts"
	"log"
	"net/http"
	"time"
)

var toSay []tts.TTS

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

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

func queueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		go queue(r.URL.Query()["tl"][0], r.URL.Query()["q"][0])
		fmt.Fprint(w, Response{"success": true, "message": "Queued"})
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	go start()
	http.HandleFunc("/queue", queueHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
