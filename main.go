package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mbenaiss/fizzbuzz/internal/fizzbuzz"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/stats", statsEndpoint)
	http.HandleFunc("/", fizzbuzzEndpoint)

	log.Printf("http listening port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func fizzbuzzEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
	f := parseQuery(r.URL.Query())
	result := fizzbuzz.Get(f)
	j, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func statsEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
	result := fizzbuzz.GetStats()
	j, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseQuery(query map[string][]string) fizzbuzz.FizzBuzz {
	n1 := parseInt(getValue(query["n1"]))
	n2 := parseInt(getValue(query["n2"]))
	limit := parseInt(getValue(query["limit"]))
	str1 := getValue(query["str1"])
	str2 := getValue(query["str2"])

	return fizzbuzz.FizzBuzz{
		N1:    n1,
		N2:    n2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}
}

func getValue(q []string) string {
	if len(q) == 0 {
		return ""
	}
	return q[0]
}

func parseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int(n)
}
