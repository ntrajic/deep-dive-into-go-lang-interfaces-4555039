package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Event struct {
	Time    time.Time `json:"time,omitempty"`
	Message string    `json:"message,omitempty"`
}

type syncer interface {
	Sync() error
}

type Encoder struct {
	w io.Writer
	s syncer
}

type nopSyncer struct{}

func (nopSyncer) Sync() error { return nil }

func NewEncoder(w io.Writer) *Encoder {
	e := Encoder{w: w, s: nopSyncer{}}
	if s, ok := w.(syncer); ok {
		e.s = s
	}

	return &e
}

func (e *Encoder) Encode(evt Event) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	n, err := e.w.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("partial write (%d out of %d bytes)", n, len(data))
	}

	e.s.Sync()

	return nil
}

func main() {
	enc := NewEncoder(os.Stdout)
	evt := Event{
		Time:    time.Now().UTC(),
		Message: "elliot login",
	}
	enc.Encode(evt)
}
