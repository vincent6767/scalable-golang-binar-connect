package scalablewithgolang

import (
	"net/http"
	"encoding/json"
	"io"
)

const (
	MaxLength = 256000
)

type Payload struct {}

type PayloadCollection struct {
	Payloads []*Payload
}

var payloadChan = make(chan *Payload, 200)

// UploadToSomewhere simulates the data upload
func (p *Payload) UploadToSomewhere() bool {
	var sum int
	for i := 0; i < 1000000; i++ {
		sum++
	}
	return true
}

// EventHandler handles the event endpoint.
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, payload := range content.Payloads {
		go payload.UploadToSomewhere()   // <----- What do you think will happen?
	}
	w.WriteHeader(http.StatusOK)
}

func EventHandlerChannel(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, payload := range content.Payloads {
		payloadChan <- payload  // <----- The execution will be blocked when the channel hit 100. You're welcome
	}
	w.WriteHeader(http.StatusOK)
}

func handlePayload(pc chan *Payload) {
	for {
		select {
		case payload := <- pc:
			payload.UploadToSomewhere()
		}
	}
}