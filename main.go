package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mbenaiss/fizzbuzz/internal/fizzbuzz"
	"github.com/mbenaiss/fizzbuzz/internal/repository"
)

func main() {
	port := os.Getenv("PORT")
	dbPath := os.Getenv("DB")

	rep, err := repository.New(dbPath)
	if err != nil {
		log.Fatalf("unable to initialize database %+v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/stats/", statsEndpoint(rep))
	mux.Handle("/", fizzbuzzEndpoint(rep))

	log.Printf("http listening port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func fizzbuzzEndpoint(rep *repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		p, err := json.Marshal(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//save the result into database
		err = rep.UpsertQuery(string(p))
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
	})
}

func statsEndpoint(rep *repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
		f, h, err := rep.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := struct {
			Fizzbuzz *fizzbuzz.FizzBuzz `json:"params"`
			Hits     int                `json:"hits"`
		}{
			Fizzbuzz: f,
			Hits:     h,
		}
		j, err := json.Marshal(response)
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
	})
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
