package routes

import (
	"log"
	"net/http"
	api "sorts/api"

	"github.com/gorilla/mux"
)

func Routes() {

	//fmt.Println(srt(slice()))
	r := mux.NewRouter()
	r.HandleFunc("/quicksort", api.Quicksorthandler).Methods("POST")
	r.HandleFunc("/repetitions", api.Repetitions).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))

	// func slice() []int {
	// 	var slc []int
	// 	var temp int
	// 	fmt.Println("Enter the elements of the array")
	// 	for {
	// 		fmt.Scan(&temp)
	// 		if temp == 0 {
	// 			break
	// 		}
	// 		slc = append(slc, temp)
	// 	}
	// 	return slc
	// }
}
