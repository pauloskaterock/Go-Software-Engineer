package storage

import (
    "sync"
    "todo-backend/models"
)

type TodoStorage struct {
    todos map[int]models.Todo
    mu    sync.RWMutex
    nextID int
}

func NewTodoStorage() *TodoStorage {
    return &TodoStorage{
        todos:  make(map[int]models.Todo),
        nextID: 1,
    }
}

func (s *TodoStorage) GetAll() []models.Todo {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    todos := make([]models.Todo, 0, len(s.todos))
    for _, todo := range s.todos {
        todos = append(todos, todo)
    }
    return todos
}

func (s *TodoStorage) Create(todo models.Todo) models.Todo {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    todo.ID = s.nextID
    todo.CreatedAt = time.Now()
    s.todos[todo.ID] = todo
    s.nextID++
    return todo
}

func (s *TodoStorage) Update(id int, updated models.Todo) (models.Todo, bool) {
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
    return models.Todo{}, false
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