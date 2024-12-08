<script lang="ts">
	import type { Todo } from '../stores/todos';
	export let todo: Todo;
	export let onToggle: (id: number) => void;
	export let onDelete: (id: number) => void;
</script>

<div class="todo-item">
	<label>
		<input type="checkbox" checked={todo.done} on:change={() => onToggle(todo.id)} />
		<span class:completed={todo.done}>{todo.task}</span>
	</label>
	<button on:click={() => onDelete(todo.id)}>Delete</button>
</div>

<style>
	.completed {
		text-decoration: line-through;
		color: gray;
	}
	.todo-item {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-bottom: 0.5rem;
	}
</style>
