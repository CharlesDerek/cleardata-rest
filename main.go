package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"unicode"
)

func main() {

	r := mux.NewRouter().SkipClean(false).UseEncodedPath()
	r.HandleFunc("/v1/get/{stringVal}", getString).Methods("GET")
	r.HandleFunc("/v1/get/{stringVal}/", getString).Methods("GET")

	// Documentation page
	r.HandleFunc("/", root).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func getString(w http.ResponseWriter, r *http.Request) {
	// Get the string value from the URL
	stringVal := mux.Vars(r)["stringVal"]
	result := containsAllLetters(stringVal)

	// Marshal JSON object to string
	resultJSON, err := json.Marshal(result)
	if err != nil {
		panic("Failed to marshal JSON object")
	}

	// Write result json to response writer
	w.Write(resultJSON)

}

// Determines if string contains alphabet(n+1)
func containsAllLetters(s string) bool {

	// Check for invalid inputs.
	if len(s) == 0 {
		return false
	}

	// Create a map to store letters seen.
	lettersSeen := make(map[rune]bool)

	containsOneOfEachLetter := false
	
	// Iterate/traverse over each character.
	for _, c := range s {

		// Check if character is a letter in alphabet (unicode map using pkg)
		if unicode.IsLetter(c) {
			// Add letter to map of letters seen
			lettersSeen[c] = true
		}

	}

	// Check if map has 26 entries.
	if len(lettersSeen) == 26 {
		// All letters present.
		containsOneOfEachLetter = true
	}
	
	return containsOneOfEachLetter
}