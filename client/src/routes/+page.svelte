<script lang="ts">
	import { todos, fetchTodos, addTodo, toggleTodo, deleteTodo } from '../stores/todos';
	import { onMount } from 'svelte';

	let newTodo = '';

	// Fetch todos on page load
	onMount(fetchTodos);

	async function handleAddTodo() {
		if (newTodo.trim()) {
			await addTodo(newTodo.trim());
			newTodo = '';
		}
	}

	async function handleToggleTodo(id: number, done: boolean) {
		await toggleTodo(id, !done);
	}

	async function handleDeleteTodo(id: number) {
		await deleteTodo(id);
	}
</script>

<div class="todo-app">
	<h1>Todo App</h1>

	<form on:submit|preventDefault={handleAddTodo}>
		<input type="text" bind:value={newTodo} placeholder="Add a new task" />
		<button type="submit">Add</button>
	</form>

	<ul>
		{#each $todos as todo}
			<li>
				<label>
					<input
						type="checkbox"
						checked={todo.done}
						on:change={() => handleToggleTodo(todo.id, todo.done)}
					/>
					<span class:completed={todo.done}>{todo.task}</span>
				</label>
				<button on:click={() => handleDeleteTodo(todo.id)}>Delete</button>
			</li>
		{/each}
	</ul>
</div>

<style>
	.completed {
		text-decoration: line-through;
		color: gray;
	}

	.todo-app {
		width: 400px;
		margin: auto;
		text-align: center;
	}

	form {
		display: flex;
		gap: 0.5rem;
		margin-bottom: 1rem;
	}

	ul {
		list-style-type: none;
		padding: 0;
	}
</style>
