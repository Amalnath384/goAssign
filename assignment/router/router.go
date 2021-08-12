package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	. "project/assignment/pkg/handler/student"
)

func Router() *mux.Router {

	studentHandler := NewStudentService()
	r := mux.NewRouter()

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	r.HandleFunc("/students/{userid}/{token}", studentHandler.ListStudent).Methods("GET")
	r.HandleFunc("/students/{id}/{userid}/{token}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}/{userid}/{token}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}/{userid}/{token}", studentHandler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/users", studentHandler.CreateUser).Methods("POST")
	r.HandleFunc("/login/{id}", studentHandler.CreateToken).Methods("POST")

	http.ListenAndServe("127.0.0.1:8000", r)

	return r
}
