package main

import (
	"encoding/json"
	"fmt"
	"github.com/bernljung/tts"
	"log"
	"net/http"
	"time"
)

const filename = "temp.mp3"

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
			if ok := toSay[0].Download(); ok {
				toSay[0].Play()
				toSay = toSay[1:]
			}
		}
	}
}

func queue(tl, q string) {
	t := tts.TTS{Filename: filename, Tl: tl, Q: q}
	toSay = append(toSay, t)
}

func queueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		go queue(r.URL.Query()["tl"][0], r.URL.Query()["q"][0])
		fmt.Fprint(rw, Response{"success": true, "message": "Queued"})
	} else {
		fmt.Fprint(rw, Response{"success": false, "message": "Use post"})
	}
}

func main() {
	go start()
	http.HandleFunc("/queue", queueHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}