package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID       int    "json:id"
	FistName string "json:firstname"
	LastName string "json:lastname"
	Age      int    "jason:age "
}

func main() {
	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)

		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FistName: "Murat", LastName: "UÇAR", Age: 27},
			User{ID: 1, FistName: "Rekabetçi", LastName: "KUR", Age: 57},
			User{ID: 1, FistName: "Ahmet", LastName: "AT", Age: 54},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
