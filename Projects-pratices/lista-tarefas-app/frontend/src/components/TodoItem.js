import React, { useState } from 'react';
import { todoService } from '../services/api';

const TodoItem = ({ todo, onUpdate, onDelete }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editTitle, setEditTitle] = useState(todo.title);

  const handleToggleComplete = async () => {
    try {
      const response = await todoService.updateTodo(todo.id, {
        ...todo,
        completed: !todo.completed
      });
      onUpdate(response.data);
    } catch (error) {
      console.error('Error updating todo:', error);
    }
  };

  const handleEdit = async () => {
    if (isEditing && editTitle.trim() !== '') {
      try {
        const response = await todoService.updateTodo(todo.id, {
          ...todo,
          title: editTitle
        });
        onUpdate(response.data);
        setIsEditing(false);
      } catch (error) {
        console.error('Error updating todo:', error);
      }
    } else {
      setIsEditing(!isEditing);
    }
  };

  const handleDelete = async () => {
    try {
      await todoService.deleteTodo(todo.id);
      onDelete(todo.id);
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  };

  return (
    <div className={`todo-item ${todo.completed ? 'completed' : ''}`}>
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={handleToggleComplete}
        className="todo-checkbox"
      />
      
      {isEditing ? (
        <input
          type="text"
          value={editTitle}
          onChange={(e) => setEditTitle(e.target.value)}
          onBlur={handleEdit}
          onKeyPress={(e) => e.key === 'Enter' && handleEdit()}
          className="todo-edit-input"
          autoFocus
        />
      ) : (
        <span className="todo-title" onDoubleClick={handleEdit}>
          {todo.title}
        </span>
      )}
      
      <div className="todo-actions">
        <button onClick={handleEdit} className="edit-btn">
          {isEditing ? '💾' : '✏️'}
        </button>
        <button onClick={handleDelete} className="delete-btn">
          🗑️
        </button>
      </div>
    </div>
  );
};

export default TodoItem;