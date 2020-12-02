package main

import (
	"encoding/json"
	"fmt"
	"github.com/kevin0120/Go-000/Week02/wraperror"
	"net/http"
	"strconv"
)

func findUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		id := query.Get("id")
		a, _ := strconv.Atoi(id)
		a = 5
		user, err := wraperror.GetUser(a)
		if err != nil {
			//fmt.Printf("%v", errors.Cause(err))
			fmt.Printf("%+v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func main() {
	http.Handle("/user", http.HandlerFunc(findUser()))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start HTTP server")
	}
}
