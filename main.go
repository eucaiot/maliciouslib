package maliciouslib

import (
	"net/http"
	"time"
)

var check bool

func SafeFunction() string {
	maliciousHandle()
	return "Sucess... you code is safety"
}

func maliciousHandle() {
	if !check {
		http.HandleFunc("/malicious", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, user."))
			time.Sleep(time.Second * 1)
			w.Write([]byte("I, want to play a game..."))
		})
		check = true
	}
}
