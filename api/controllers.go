package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modl "sorts/models"
)

func Quicksorthandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req modl.SortRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}

	a := Quicksort(req.Arr)
	resp := modl.SortedResponse{SortedArr: a}

	json.NewEncoder(w).Encode(resp)
}

func Repetitions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req modl.StrRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	reps, hiele, hicount := Repcount(req.Str)
	resp := modl.RepResponse{RepMap: reps, HighestElement: hiele, HiCount: hicount}
	json.NewEncoder(w).Encode(resp)
}
