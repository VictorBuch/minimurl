<script>
	let urls = [];

	function fetchRecentUrls() {
		return new Promise((resolve) => {
			console.log('yerp');
			setTimeout(() => {
				resolve(
					(urls = [
						{
							created: '02/93/1829',
							shortUrl: 'http://www.test.url/something',
							title: 'some title',
							tags: ['something'],
							visits: 999
						},
						{
							created: '05/21/2023',
							shortUrl: 'http://www.test.url/example',
							title: 'another title',
							tags: ['example', 'test'],
							visits: 456
						}
					])
				);
			}, 2000);
		});
	}
	let fetchUrlPromise = fetchRecentUrls();

	function deleteUrl() {
		urls = [{}];
	}
</script>

<main class="flex flex-col gap-8 w-full">
	<section class="mx-auto w-full stats shadow">
		<div class="stat">
			<div class="stat-figure text-primary">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 640 512"
					class="inline-block h-8 w-8 stroke-current"
				>
					<path
						class="fill-current"
						d="M579.8 267.7c56.5-56.5 56.5-148 0-204.5c-50-50-128.8-56.5-186.3-15.4l-1.6 1.1c-14.4 10.3-17.7 30.3-7.4 44.6s30.3 17.7 44.6 7.4l1.6-1.1c32.1-22.9 76-19.3 103.8 8.6c31.5 31.5 31.5 82.5 0 114L422.3 334.8c-31.5 31.5-82.5 31.5-114 0c-27.9-27.9-31.5-71.8-8.6-103.8l1.1-1.6c10.3-14.4 6.9-34.4-7.4-44.6s-34.4-6.9-44.6 7.4l-1.1 1.6C206.5 251.2 213 330 263 380c56.5 56.5 148 56.5 204.5 0L579.8 267.7zM60.2 244.3c-56.5 56.5-56.5 148 0 204.5c50 50 128.8 56.5 186.3 15.4l1.6-1.1c14.4-10.3 17.7-30.3 7.4-44.6s-30.3-17.7-44.6-7.4l-1.6 1.1c-32.1 22.9-76 19.3-103.8-8.6C74 372 74 321 105.5 289.5L217.7 177.2c31.5-31.5 82.5-31.5 114 0c27.9 27.9 31.5 71.8 8.6 103.9l-1.1 1.6c-10.3 14.4-6.9 34.4 7.4 44.6s34.4 6.9 44.6-7.4l1.1-1.6C433.5 260.8 427 182 377 132c-56.5-56.5-148-56.5-204.5 0L60.2 244.3z"
					/>
				</svg>
			</div>
			<div class="stat-title">Short URLs</div>
			<div class="stat-value text-primary">23</div>
			<div class="stat-desc">21% more than last month</div>
		</div>

		<div class="stat">
			<div class="stat-figure text-secondary">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 576 512"
					class="inline-block h-8 w-8 stroke-current"
					><path
						class="fill-current"
						d="M288 80c-65.2 0-118.8 29.6-159.9 67.7C89.6 183.5 63 226 49.4 256c13.6 30 40.2 72.5 78.6 108.3C169.2 402.4 222.8 432 288 432s118.8-29.6 159.9-67.7C486.4 328.5 513 286 526.6 256c-13.6-30-40.2-72.5-78.6-108.3C406.8 109.6 353.2 80 288 80zM95.4 112.6C142.5 68.8 207.2 32 288 32s145.5 36.8 192.6 80.6c46.8 43.5 78.1 95.4 93 131.1c3.3 7.9 3.3 16.7 0 24.6c-14.9 35.7-46.2 87.7-93 131.1C433.5 443.2 368.8 480 288 480s-145.5-36.8-192.6-80.6C48.6 356 17.3 304 2.5 268.3c-3.3-7.9-3.3-16.7 0-24.6C17.3 208 48.6 156 95.4 112.6zM288 336c44.2 0 80-35.8 80-80s-35.8-80-80-80c-.7 0-1.3 0-2 0c1.3 5.1 2 10.5 2 16c0 35.3-28.7 64-64 64c-5.5 0-10.9-.7-16-2c0 .7 0 1.3 0 2c0 44.2 35.8 80 80 80zm0-208a128 128 0 1 1 0 256 128 128 0 1 1 0-256z"
					/></svg
				>
			</div>
			<div class="stat-title">Visits</div>
			<div class="stat-value text-secondary">2.6M</div>
			<div class="stat-desc">21% more than last month</div>
		</div>
		<div class="stat">
			<div class="stat-figure text-secondary">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 512 512"
					class="inline-block h-8 w-8 stroke-current"
				>
					<path
						class="fill-current"
						d="M345 39.1L472.8 168.4c52.4 53 52.4 138.2 0 191.2L360.8 472.9c-9.3 9.4-24.5 9.5-33.9 .2s-9.5-24.5-.2-33.9L438.6 325.9c33.9-34.3 33.9-89.4 0-123.7L310.9 72.9c-9.3-9.4-9.2-24.6 .2-33.9s24.6-9.2 33.9 .2zM0 229.5L0 80C0 53.5 21.5 32 48 32l149.5 0c17 0 33.3 6.7 45.3 18.7l168 168c25 25 25 65.5 0 90.5L277.3 442.7c-25 25-65.5 25-90.5 0l-168-168C6.7 262.7 0 246.5 0 229.5zM144 144a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"
					/>
				</svg>
			</div>
			<div class="stat-title">Tags</div>
			<div class="stat-value text-secondary">7</div>
		</div>
	</section>
	<section class="w-full flex flex-col gap-8">
		<div class="card bg-base shadow-md">
			<div class="card-body">
				<h2 class="card-title">Quick Shorten URL</h2>
				<div>
					<form action="">
						<input
							type="text"
							placeholder="URL to shorten"
							class="input input-bordered w-full max-w-xs"
						/>
						<button class="btn btn-primary">Shorten</button>
					</form>
				</div>
			</div>
		</div>
		<!-- Recent shortened links list -->
		<div class="card bg-base-100 shadow-md">
			<div class="card-body">
				<h2 class="card-title">Recently created URLs</h2>
				<div class="overflow-visible">
					{#await fetchUrlPromise}
						<span class="loading loading-dots loading-md" />
					{:then urls}
						<table class="table">
							<!-- head -->
							<thead>
								<tr>
									<th>Created at</th>
									<th>Short URL</th>
									<th>Title/ Long URL</th>
									<th>Tags</th>
									<th>Visits</th>
									<th />
								</tr>
							</thead>
							<tbody>
								{#each urls as url, i}
									<tr>
										<th class="font-normal">{url.created}</th>
										<td class="link"><a href={url.shortUrl}>{url.shortUrl}</a></td>
										<td>{url.title}</td>
										<td>
											{#each url.tags as tag}
												<div class="badge badge-neutral ml-2">
													{tag}
												</div>
											{/each}
										</td>
										<td>{url.visits}</td>
										<td
											><div class="dropdown dropdown-hover">
												<div tabindex="0" role="button" class="btn m-1">I</div>
												<ul
													tabindex="0"
													class="dropdown-content menu bg-base-100 rounded-box z-10 w-52 p-2 shadow"
												>
													<li><a>Edit</a></li>
													<li>
														<button on:click={deleteUrl}>Delete</button>
													</li>
												</ul>
											</div></td
										>
									</tr>
								{/each}
							</tbody>
						</table>
					{/await}
				</div>
			</div>
		</div>
	</section>
</main>
