package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJSON is a helper function to write a JSON response.
// It takes an HTTP response writer, status code, and any data type to convert into JSON and send to the client.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// Set the content type of the response to JSON.
	w.Header().Set("Content-type", "application/json")

	// Write the HTTP status code (e.g., 200 for OK, 400 for Bad Request).
	w.WriteHeader(status)

	// Convert the data (v) to JSON and write it to the response body.
	return json.NewEncoder(w).Encode(v)
}

// apiFunc is a custom function type used to define HTTP handlers that return an error.
type apiFunc func(http.ResponseWriter, *http.Request) error

// ApiError represents a simple structure for returning error messages as JSON.
type ApiError struct {
	Error string
}

// checkerr is a utility function to log fatal errors and stop the program if an error occurs.
func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


// makeHTTPHandleFunc takes an apiFunc and wraps it into an http.HandlerFunc.
// It ensures that errors returned by apiFunc are handled and converted into JSON error responses.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the API function (f), if an error is returned, write it as a JSON response with a Bad Request status.
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// APIServer is a structure representing the API server.
// It holds the address (listenAddr) that the server will bind to.
type APIServer struct {
	listenAddr string
	store Storage
}

// NewApiServer is a constructor that creates and returns a new APIServer instance with the specified address.
func NewApiServer(listenAddr string,store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}


// Run starts the API server and listens for incoming HTTP requests.
func (s *APIServer) Run() {
	// Create a new Gorilla Mux router to handle routes.
	router := mux.NewRouter()

	// Define routes and associate them with their respective handlers.
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount)) // For creating, deleting, or listing accounts.
	router.HandleFunc("/account{id}", makeHTTPHandleFunc(s.handleGetAccountByID)) // For retrieving account details by ID.

	// Log that the server is running and bind it to the specified address.
	log.Println("Json api running on:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// handleAccount handles requests to the "/account" route. It supports multiple HTTP methods (GET, POST, DELETE).
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	// If the request method is GET, call the handleGetAccount function.
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	// If the request method is POST, call the handleCreateAccount function (not implemented yet).
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	// If the request method is DELETE, call the handleDeleteAccount function (not implemented yet).
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	// If the method is not allowed, return an error message.
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts,err:=s.store.GetAccounts()
	if err!=nil{
		return err
	}
	return WriteJSON(w,http.StatusOK,accounts)
}
// handleGetAccount handles retrieving a single account by its ID from the URL.
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	// Extract the account ID from the URL path variables.
	id := mux.Vars(r)["id"]
	
	// Print the account ID to the console (for debugging purposes).
	fmt.Println(id)

	// Create a dummy account and return it as a JSON response.
	return WriteJSON(w, http.StatusOK, &Account{})
}


// handleCreateAccount handles the creation of new accounts. (Functionality not yet implemented).
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq:=new(CreateAccountRequest)
	if err:=json.NewDecoder(r.Body).Decode(createAccountReq);err!=nil{
		return err
	}

	account:=NewAccount(createAccountReq.FirstName,createAccountReq.LastName)

	if err:=s.store.CreateAccount(account);err!=nil{
		return err
	}
	return WriteJSON(w,http.StatusOK,account)
}

// handleDeleteAccount handles the deletion of accounts. (Functionality not yet implemented).
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request)error{
	return nil
}