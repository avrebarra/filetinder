<script>
	export let name;

	import FilePreviewer from "./FilePreviewer.svelte"
	import EmptyState from "./EmptyState.svelte"

	let files = [1]
</script>

<div class="container grid-lg">
	<div style="margin-top: 40px;">
		<h1>FileTinder <small class="label">UI</small></h1>
		<p class="text-gray" style="margin-top: -15px;"><em>Tinderify your files. Keep what you liked, delete what you hate (with caution)</em></p>
	</div>

	{#if files.length <= 0}
	<!-- Empty State -->
	<div class="columns">
		<EmptyState/>
	</div>
	{:else}
	<!-- Non Empty State -->
	<div class="columns">
		<div class="column col-9">
			<!--  Previewer -->
			<FilePreviewer fileuri="/assets/sample-image.jpeg"/>

			<div style="margin-top: 20px; text-align: center;">
				<div class="d-inline-block">
					<ul class="pagination">
						<li class="page-item disabled">
							<a href="/" tabindex="-1">Previous</a>
						</li>
						<li class="page-item active">
							<a href="/">1</a>
						</li>
						<li class="page-item">
							<a href="/">2</a>
						</li>
						<li class="page-item">
							<a href="/">3</a>
						</li>
						<li class="page-item">
							<span>...</span>
						</li>
						<li class="page-item">
							<a href="/">12</a>
						</li>
						<li class="page-item">
							<a href="/">Next</a>
						</li>
					</ul>
				</div>
			</div>
		</div>
		
		<!--  Menu -->
		<div class="column col-3">
			<ul class="menu">
				<li class="divider" data-content="MARK FILE"></li>
				
				<!-- Menu Entries -->
				<li class="menu-item">
					<a href="/">
						<i class="icon icon-stop"></i> Add to Delete list
					</a>
					<div class="menu-badge">
					<label class="label label-primary">2</label>
					</div>
				</li>
			</ul>

			<br>

			<ul class="menu">
				<li class="divider" data-content="RUN FUNCTIONS"></li>

				<!-- Menu Entries -->
				<li class="menu-item">
					<a href="/">
					<i class="icon icon-link"></i> Delete All Marked
					</a>
					
				</li>
			</ul>
		</div>
	</div>
	{/if}
</div>
