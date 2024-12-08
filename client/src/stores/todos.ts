import { writable } from 'svelte/store';

export const API_BASE_URL = 'http://localhost:8080';

export type Todo = {
	id: number;
	task: string;
	done: boolean;
};

const todosStore = writable<Todo[]>([]);

// Fetch todos from the backend
export async function fetchTodos() {
	try {
		const response = await fetch(`${API_BASE_URL}/todos`);
		if (!response.ok) throw new Error('Failed to fetch todos');
		const data: Todo[] = await response.json();
		todosStore.set(data);
	} catch (error) {
		console.error('Error fetching todos:', error);
	}
}

// Add a new todo
export async function addTodo(task: string) {
	try {
		const response = await fetch(`${API_BASE_URL}/todos`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ task, done: false })
		});
		if (!response.ok) throw new Error('Failed to add todo');
		await fetchTodos(); // Refresh the list
	} catch (error) {
		console.error('Error adding todo:', error);
	}
}

// Toggle a todo's done status
export async function toggleTodo(id: number, done: boolean) {
	try {
		const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ done })
		});
		if (!response.ok) throw new Error('Failed to toggle todo');
		await fetchTodos(); // Refresh the list
	} catch (error) {
		console.error('Error toggling todo:', error);
	}
}

// Delete a todo
export async function deleteTodo(id: number) {
	try {
		const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
			method: 'DELETE'
		});
		if (!response.ok) throw new Error('Failed to delete todo');
		await fetchTodos(); // Refresh the list
	} catch (error) {
		console.error('Error deleting todo:', error);
	}
}

export { todosStore as todos };
