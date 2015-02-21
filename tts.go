package main

import (
	"log"
	"net/url"
	"os/exec"
)

const tts_url = "http://translate.google.com/translate_tts"

type TTS struct {
	Tl, Q string
}

func (t TTS) query() string {
	v := url.Values{}
	v.Set("tl", t.Tl)
	v.Add("q", t.Q)
	v.Add("ie", "UTF-8")
	query := "?" + v.Encode()
	return query
}

func (t TTS) Play() {
	log.Println("Query:", t.Q)
	log.Println("Command:", "mpg123", t.Link())
	out, err := exec.Command("mpg123", t.Link()).Output()
	log.Println("command:", out, err)
}

func (t TTS) Link() string {
	return tts_url + t.query()
}
