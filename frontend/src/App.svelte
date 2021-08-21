<script lang="ts">

	async function test() {
		var res = await fetch("http://localhost:6996/api/dir", {
			method: 'POST',
			body: JSON.stringify({
				path: "/home/bruno/projects/lolover_testroot"
			})
		})
		console.log(await res.json());
	}
	
</script>

<main>
	<button on:click="{test}">test</button>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>