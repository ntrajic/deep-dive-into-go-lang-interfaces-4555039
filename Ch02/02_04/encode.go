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

type Encoder struct {
	w io.Writer
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

	return nil
}

func main() {
	enc := Encoder{os.Stdout}
	evt := Event{
		Time:    time.Now().UTC(),
		Message: "elliot login",
	}
	enc.Encode(evt)
}
