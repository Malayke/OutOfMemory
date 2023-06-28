package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func OutOfMemoryGood1(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()
	MaxValue := 6
	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil || sink < 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if sink > MaxValue {
		return
	}
	result := make([]string, sink)
	for i := 0; i < sink; i++ {
		result[i] = fmt.Sprintf("Item %d", i+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func OutOfMemoryGood2(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()
	MaxValue := 6
	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil || sink < 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if sink <= MaxValue {
		result := make([]string, sink)
		for i := 0; i < sink; i++ {
			result[i] = fmt.Sprintf("Item %d", i+1)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func OutOfMemoryGood3(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()
	MaxValue := 6
	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil || sink < 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if sink > MaxValue {
		sink = MaxValue
		result := make([]string, sink)
		for i := 0; i < sink; i++ {
			result[i] = fmt.Sprintf("Item %d", i+1)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func OutOfMemoryGood4(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()
	MaxValue := 6
	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil || sink < 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if sink > MaxValue {
		sink = MaxValue
	} else {
		tmp := sink
		sink = tmp
	}
	result := make([]string, sink)
	for i := 0; i < sink; i++ {
		result[i] = fmt.Sprintf("Item %d", i+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func OutOfMemoryGood5(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()
	MaxValue := 6
	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil || sink < 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if sink < 0 || sink > MaxValue {
		sink = MaxValue
	}
	result := make([]string, sink)
	for i := 0; i < sink; i++ {
		result[i] = fmt.Sprintf("Item %d", i+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func OutOfMemoryBad(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query()

	sourceStr := source.Get("n")
	sink, err := strconv.Atoi(sourceStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := make([]string, sink)
	for i := 0; i < sink; i++ {
		result[i] = fmt.Sprintf("Item %d", i+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/oom", OutOfMemoryBad)
	http.HandleFunc("/oom_good1", OutOfMemoryGood1)
	http.HandleFunc("/oom_good2", OutOfMemoryGood2)
	http.HandleFunc("/oom_good3", OutOfMemoryGood3)
	http.HandleFunc("/oom_good4", OutOfMemoryGood4)
	http.HandleFunc("/oom_good5", OutOfMemoryGood5)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
