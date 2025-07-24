package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserStore manages user data (in-memory for this example)
type UserStore struct {
	mu     sync.RWMutex
	users  map[int]*User
	nextID int
}

// NewUserStore creates a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// CreateUser adds a new user
func (s *UserStore) CreateUser(name, email string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	s.users[s.nextID] = user
	s.nextID++
	return user
}

// GetUser retrieves a user by ID
func (s *UserStore) GetUser(id int) (*User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

// GetAllUsers returns all users
func (s *UserStore) GetAllUsers() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// UpdateUser updates an existing user
func (s *UserStore) UpdateUser(id int, name, email string) (*User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, false
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	return user, true
}

// DeleteUser removes a user
func (s *UserStore) DeleteUser(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if exists {
		delete(s.users, id)
	}
	return exists
}

// UserRequest represents the request body for creating/updating users
type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// APIServer represents our HTTP server
type APIServer struct {
	store  *UserStore
	router *mux.Router
}

// NewAPIServer creates a new API server
func NewAPIServer() *APIServer {
	server := &APIServer{
		store:  NewUserStore(),
		router: mux.NewRouter(),
	}
	server.setupRoutes()
	return server
}

// setupRoutes configures all the API routes
func (s *APIServer) setupRoutes() {
	// Middleware
	s.router.Use(s.loggingMiddleware)
	s.router.Use(s.corsMiddleware)
	s.router.Use(s.jsonMiddleware)

	// API routes
	api := s.router.PathPrefix("/api/v1").Subrouter()
	
	// User endpoints
	api.HandleFunc("/users", s.handleGetUsers).Methods("GET")
	api.HandleFunc("/users", s.handleCreateUser).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleGetUser).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleUpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", s.handleDeleteUser).Methods("DELETE")

	// Health check
	api.HandleFunc("/health", s.handleHealth).Methods("GET")

	// Static file serving (for a simple frontend)
	s.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
}

// Middleware functions

func (s *APIServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Create a custom ResponseWriter to capture status code
		wrapper := &responseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		
		next.ServeHTTP(wrapper, r)
		
		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapper.statusCode, duration)
	})
}

func (s *APIServer) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

// responseWriterWrapper wraps http.ResponseWriter to capture status code
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Handler functions

func (s *APIServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now(),
		"service":   "user-api",
	}
	s.writeJSON(w, http.StatusOK, response)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.store.GetAllUsers()
	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"users": users,
		"count": len(users),
	})
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, exists := s.store.GetUser(id)
	if !exists {
		s.writeError(w, http.StatusNotFound, "User not found", fmt.Sprintf("User with ID %d does not exist", id))
		return
	}

	s.writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validation
	if req.Name == "" || req.Email == "" {
		s.writeError(w, http.StatusBadRequest, "Validation failed", "Name and email are required")
		return
	}

	user := s.store.CreateUser(req.Name, req.Email)
	s.writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	user, exists := s.store.UpdateUser(id, req.Name, req.Email)
	if !exists {
		s.writeError(w, http.StatusNotFound, "User not found", fmt.Sprintf("User with ID %d does not exist", id))
		return
	}

	s.writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	deleted := s.store.DeleteUser(id)
	if !deleted {
		s.writeError(w, http.StatusNotFound, "User not found", fmt.Sprintf("User with ID %d does not exist", id))
		return
	}

	s.writeJSON(w, http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

// Helper functions

func (s *APIServer) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func (s *APIServer) writeError(w http.ResponseWriter, status int, error, message string) {
	response := ErrorResponse{
		Error:   error,
		Message: message,
	}
	s.writeJSON(w, status, response)
}

// Start starts the server
func (s *APIServer) Start(port string) error {
	log.Printf("Starting server on port %s", port)
	log.Printf("Health check: http://localhost%s/api/v1/health", port)
	log.Printf("API endpoints:")
	log.Printf("  GET    /api/v1/users")
	log.Printf("  POST   /api/v1/users")
	log.Printf("  GET    /api/v1/users/{id}")
	log.Printf("  PUT    /api/v1/users/{id}")
	log.Printf("  DELETE /api/v1/users/{id}")
	
	return http.ListenAndServe(port, s.router)
}

func main() {
	// Create and configure server
	server := NewAPIServer()

	// Add some sample data
	server.store.CreateUser("Alice Johnson", "alice@example.com")
	server.store.CreateUser("Bob Smith", "bob@example.com")
	server.store.CreateUser("Charlie Brown", "charlie@example.com")

	// Start server
	port := ":8080"
	if err := server.Start(port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

/*
Example API calls using curl:

# Get all users
curl -X GET http://localhost:8080/api/v1/users

# Get specific user
curl -X GET http://localhost:8080/api/v1/users/1

# Create new user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"David Wilson","email":"david@example.com"}'

# Update user
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Smith","email":"alice.smith@example.com"}'

# Delete user
curl -X DELETE http://localhost:8080/api/v1/users/1

# Health check
curl -X GET http://localhost:8080/api/v1/health
*/
