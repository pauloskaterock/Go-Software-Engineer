package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"  // Importação adicionada
    "sync"
    "time"
)

// Definições de modelos diretamente no main.go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
}

type TodoRequest struct {
    Title string `json:"title"`
}

// Storage em memória
type TodoStorage struct {
    todos  map[int]Todo
    mu     sync.RWMutex
    nextID int
}

func NewTodoStorage() *TodoStorage {
    return &TodoStorage{
        todos:  make(map[int]Todo),
        nextID: 1,
    }
}

func (s *TodoStorage) GetAll() []Todo {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    todos := make([]Todo, 0, len(s.todos))
    for _, todo := range s.todos {
        todos = append(todos, todo)
    }
    return todos
}

func (s *TodoStorage) Create(todo Todo) Todo {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    todo.ID = s.nextID
    todo.CreatedAt = time.Now()
    s.todos[todo.ID] = todo
    s.nextID++
    return todo
}

func (s *TodoStorage) Update(id int, updated Todo) (Todo, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if todo, exists := s.todos[id]; exists {
        if updated.Title != "" {
            todo.Title = updated.Title
        }
        todo.Completed = updated.Completed
        s.todos[id] = todo
        return todo, true
    }
    return Todo{}, false
}

func (s *TodoStorage) Delete(id int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if _, exists := s.todos[id]; exists {
        delete(s.todos, id)
        return true
    }
    return false
}

// Handlers
func enableCORS(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getTodos(w http.ResponseWriter, r *http.Request, storage *TodoStorage) {
    enableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    todos := storage.GetAll()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request, storage *TodoStorage) {
    enableCORS(&w)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    var todoReq TodoRequest
    if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    todo := Todo{
        Title:     todoReq.Title,
        Completed: false,
    }
    
    createdTodo := storage.Create(todo)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdTodo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, storage *TodoStorage) {
    enableCORS(&w)
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
    
    var updatedTodo Todo
    if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if todo, exists := storage.Update(id, updatedTodo); exists {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)
    } else {
        http.Error(w, "Todo not found", http.StatusNotFound)
    }
}

func deleteTodo(w http.ResponseWriter, r *http.Request, storage *TodoStorage) {
    enableCORS(&w)
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
    
    if storage.Delete(id) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Todo not found", http.StatusNotFound)
    }
}

func main() {
    // Inicializar storage
    todoStorage := NewTodoStorage()
    
    // Configurar rotas
    http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getTodos(w, r, todoStorage)
        case http.MethodPost:
            createTodo(w, r, todoStorage)
        case http.MethodOptions:
            enableCORS(&w)
            w.WriteHeader(http.StatusOK)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            updateTodo(w, r, todoStorage)
        case http.MethodDelete:
            deleteTodo(w, r, todoStorage)
        case http.MethodOptions:
            enableCORS(&w)
            w.WriteHeader(http.StatusOK)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    // Iniciar servidor
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}