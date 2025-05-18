package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// user struct represents a bank account user
type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

// global variables
var (
	users = make(map[string]User) // map for the users
	mu    sync.Mutex              // mutex for safe access to user map
)

/*
This function manipulates the http headers to fix some issues associated with the response 
and request having different methods (one would have POST while the other would have GET)
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func handleCORS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling CORS for:", r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

func main() {
	// define routes
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handleCORS(w, r)
		registerHandler(w, r)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handleCORS(w, r)
		loginHandler(w, r)
	})
	http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		handleCORS(w, r)
		balanceHandler(w, r)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		handleCORS(w, r)
		deleteHandler(w, r)
	})
	http.HandleFunc("/modify", func(w http.ResponseWriter, r *http.Request) {
		handleCORS(w, r)
		modifyBalanceHandler(w, r)
	})

	// start server
	fmt.Println("Server running on localhost:8001")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

/*
This function handles the registering of the user
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Method not allowed in /register")
		fmt.Println("Http method", http.MethodPost)
		fmt.Println("request method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Error decoding register request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.Username]; exists {
		fmt.Println("User already exists:", user.Username)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	users[user.Username] = user
	fmt.Println("User registered successfully:", user.Username)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully")
}

/*
This function handles the login of the user. For simplicity, it just checks username and password
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Method not allowed in /login")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Println("Error decoding login request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[credentials.Username]
	if !exists || user.Password != credentials.Password {
		fmt.Println("Invalid login attempt for user:", credentials.Username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Println("Login successful for user:", credentials.Username)
	fmt.Fprintln(w, "Login successful")
}

/*
This function handles the balance checking of the user
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func balanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("Method not allowed in /balance")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[username]
	if !exists || user.Password != password {
		fmt.Println("Invalid credentials for balance check:", username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Println("Fetching balance for user:", username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"balance": user.Balance})
}
/*
This function handles deleting a user
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[username]
	if !exists || user.Password != password {
		fmt.Println("Invalid credentials for delete attempt:", username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	delete(users, username)
	fmt.Println("User deleted successfully:", username)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User deleted successfully")
}

/*
This function handles modifying the balance of a user
@param: http.ResponseWriter and pointer to http.Request
@return: none
*/
func modifyBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Username string  `json:"username"`
		Password string  `json:"password"`
		Amount   float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[data.Username]
	if !exists || user.Password != data.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	user.Balance = data.Amount
	users[data.Username] = user

	fmt.Println("Balance modified for user:", data.Username)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Balance modified successfully")
}