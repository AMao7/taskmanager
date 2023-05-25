<script>
	import { onMount } from 'svelte';
  
	let tasks = [];
  
	onMount(async () => {
	  const res = await fetch('http://localhost:5050/api/task');
	  tasks = await res.json();
	  console.log(tasks); // Log tasks to the console for debugging
	});
  </script>
  
  <ul>
	{#each tasks as task, index (index)}
    <li>{task.Content}</li>
  {/each}
  </ul>
  