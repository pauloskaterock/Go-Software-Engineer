import React, { useState, useEffect } from 'react';
import { todoService } from './services/api';
import TodoItem from './components/TodoItem'; // Import correto
import './App.css';

function App() {
  const [todos, setTodos] = useState([]);
  const [newTodoTitle, setNewTodoTitle] = useState('');

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await todoService.getAllTodos();
      setTodos(response.data);
    } catch (error) {
      console.error('Error fetching todos:', error);
    }
  };

  const handleCreateTodo = async (e) => {
    e.preventDefault();
    if (newTodoTitle.trim() === '') return;

    try {
      const response = await todoService.createTodo(newTodoTitle);
      setTodos([...todos, response.data]);
      setNewTodoTitle('');
    } catch (error) {
      console.error('Error creating todo:', error);
    }
  };

  const handleUpdateTodo = (updatedTodo) => {
    setTodos(todos.map(todo => 
      todo.id === updatedTodo.id ? updatedTodo : todo
    ));
  };

  const handleDeleteTodo = (id) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };

  return (
    <div className="app">
      <div className="todo-container">
        <h1>📝 Lista de Tarefas</h1>
        
        <form onSubmit={handleCreateTodo} className="todo-form">
          <input
            type="text"
            value={newTodoTitle}
            onChange={(e) => setNewTodoTitle(e.target.value)}
            placeholder="Adicione uma nova tarefa..."
            className="todo-input"
          />
          <button type="submit" className="add-btn">
            ➕ Adicionar
          </button>
        </form>

        <div className="todo-list">
          {todos.length === 0 ? (
            <p className="empty-state">Nenhuma tarefa encontrada. Adicione uma nova tarefa!</p>
          ) : (
            todos.map(todo => (
              <TodoItem
                key={todo.id}
                todo={todo}
                onUpdate={handleUpdateTodo}
                onDelete={handleDeleteTodo}
              />
            ))
          )}
        </div>

        <div className="todo-stats">
          <span>Total: {todos.length}</span>
          <span>Concluídas: {todos.filter(t => t.completed).length}</span>
        </div>
      </div>
    </div>
  );
}

export default App;