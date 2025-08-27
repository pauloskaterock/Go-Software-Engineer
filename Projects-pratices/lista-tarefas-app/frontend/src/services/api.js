import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const todoService = {
  getAllTodos: () => api.get('/todos'),
  createTodo: (title) => api.post('/todos', { title }),
  updateTodo: (id, updates) => api.put(`/todos/${id}`, updates),
  deleteTodo: (id) => api.delete(`/todos/${id}`),
};

export default api;


