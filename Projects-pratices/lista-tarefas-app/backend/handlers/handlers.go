package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "todo-backend/models"
    "todo-backend/storage"
)

type TodoHandler struct {
    storage *storage.TodoStorage
}

func NewTodoHandler(storage *storage.TodoStorage) *TodoHandler {
    return &TodoHandler{storage: storage}
}

func (h *TodoHandler) EnableCORS(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
    h.EnableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    todos := h.storage.GetAll()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    h.EnableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    var todoReq models.TodoRequest
    if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    todo := models.Todo{
        Title:     todoReq.Title,
        Completed: false,
    }
    
    createdTodo := h.storage.Create(todo)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdTodo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    h.EnableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    idStr := r.URL.Path[len("/todos/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    var updatedTodo models.Todo
    if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if todo, exists := h.storage.Update(id, updatedTodo); exists {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)
    } else {
        http.Error(w, "Todo not found", http.StatusNotFound)
    }
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
    h.EnableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    idStr := r.URL.Path[len("/todos/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    if h.storage.Delete(id) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Todo not found", http.StatusNotFound)
    }
}