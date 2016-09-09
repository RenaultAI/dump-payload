package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func handlerDump(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("Dump failed: %s\n", err)
	}
	log.Println(string(dump))

	writeJSON(w, r, dump)
}

func main() {
	http.HandleFunc("/", handlerDump)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func writeJSON(w http.ResponseWriter, r *http.Request, m interface{}) {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
