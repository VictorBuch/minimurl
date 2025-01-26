<script>
	import '../app.css';
	import Navbar from '$lib/components/navbar.svelte';
	import Sidebar from '$lib/components/sidebar.svelte';
</script>

<main class="w-screen h-screen flex flex-col">
	<Navbar />
	<section class="flex w-full h-full">
		<Sidebar />

		<section class="p-12 w-full">
			<slot />
		</section>
	</section>
</main>
