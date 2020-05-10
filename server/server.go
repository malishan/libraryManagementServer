package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"project/libraryManagementServer/server/handler"
	"project/libraryManagementServer/server/middleware"
	"project/libraryManagementServer/util"
)

// Run - start the service console app
func Run() {
	r := mux.NewRouter()
	r = r.PathPrefix("/myshelf").Subrouter()

	r.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	r.HandleFunc("/signin", handler.Login).Methods(http.MethodPost)

	// r.HandleFunc("/", home).Methods(http.MethodGet)
	// r.HandleFunc("/addBook", addBook).Methods(http.MethodPost)
	// r.HandleFunc("/getBook/{id}/{name}", getBookBySerialNo).Methods(http.MethodGet)
	// r.HandleFunc("/getBooksByName/{name}", getBooksByName).Methods(http.MethodGet)
	// r.HandleFunc("/getAllBooks", getAllBooks).Methods(http.MethodGet)
	// r.HandleFunc("/updateBook", updateBook).Methods(http.MethodPost)
	// r.HandleFunc("/delete", deleteBook).Methods(http.MethodDelete)

	server := middleware.SetServer(r)

	util.Fatalln(server.ListenAndServe())
}
