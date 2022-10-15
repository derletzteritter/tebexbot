package main

import "net/http"

// Sadly, I need to pay for Tebex Premium to do this.
func serve() {
	http.HandleFunc("/purchase", func(w http.ResponseWriter, r *http.Request) {
	})

	http.ListenAndServe(":6000", nil)
}
