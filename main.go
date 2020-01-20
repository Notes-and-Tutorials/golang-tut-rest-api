// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type server struct{}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message": "hello world"}`))
// }

// func main() {
// 	s := &server{}
// 	http.Handle("/", s)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// type server struct{}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "get called"}`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"message": "post called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(`{"message": "put called"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "delete called"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"message": "not found"}`))
// 	}
// }

// func main() {
// 	s := &server{}
// 	http.Handle("/", s)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// type server struct{}

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "get called"}`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"message": "post called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(`{"message": "put called"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "delete called"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"message": "not found"}`))
// 	}
// }

// func main() {
// 	http.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// Gorilla Mux
// type server struct{}

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "get called"}`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"message": "post called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(`{"message": "put called"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "delete called"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"message": "not found"}`))
// 	}
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "get called"}`))
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message": "post called"}`))

// }

// // ...... put, delete

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNotFound)
// 	w.Write([]byte(`{"message": "not found"}`))
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", get).Methods(http.MethodGet)
// 	r.HandleFunc("/", post).Methods(http.MethodPost)
// 	// ... put, delete
// 	r.HandleFunc("/", notFound)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "get called"}`))
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message": "post called"}`))

// }

// // ...... put, delete

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNotFound)
// 	w.Write([]byte(`{"message": "not found"}`))
// }

// func main() {
// 	r := mux.NewRouter()
// 	api := r.PathPrefix("/api/v1").Subrouter()
// 	api.HandleFunc("", get).Methods(http.MethodGet)
// 	api.HandleFunc("", post).Methods(http.MethodPost)
// 	// ... put, delete
// 	r.HandleFunc("", notFound)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
