<script>
	export let name;
	import Login from './Login.svelte'
	import Queue from './Queue.svelte'
	import Nodes from './Nodes.svelte'
	import YourJobs from './YourJobs.svelte'
	import Editor from './Editor.svelte'
	import * as api from './api.js'

	const LOGIN = 0
	const MAIN = 1
	const JOBEDIT = 2
	var page = MAIN
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Work+Sans&display=swap');
:global(body) {
	background-color: black;
	color: white;
	font-family: 'Work Sans', sans-serif;
}
:global(.card) {
    display: inline-block;
    border-style: solid;
    border-color: white;
    border-radius: 4pt;
    padding: 10pt;
    margin: 8pt;
}
:global(.row) {
	display: flex;
	flex-direction: row;
}
.column {
	display: flex;
	flex-direction: column;
	margin: 10pt;
}
</style>

<!-- This is bascially just a janky version of routing -->
{#if page == LOGIN}
This is the login page
{/if}

{#if page == MAIN}
<div class="row">
	<div class="column" style="flex-basis: 25%"><YourJobs/></div>
	<div class="column" style="flex-basis: 100%"><Queue/></div>
	<div class="column" style="flex-basis: 25%"><Nodes/></div>
</div>
{/if}

{#if page == JOBEDIT}
<Editor/>
{/if}