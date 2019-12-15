<script>
	export let name;
	import Login from './Login.svelte'
	import Queue from './Queue.svelte'
	import Nodes from './Nodes.svelte'
	import Editor from './Editor.svelte'
	import WorkerControl from './WorkerControl.svelte'
	import * as api from './api.js'

	const LOGIN = 0
	const MAIN = 1
	const JOBEDIT = 2
	var page = LOGIN

	function gotoLogin() { page = LOGIN }
	function gotoMain() { page = MAIN }
	function gotoJobedit() { page = JOBEDIT }

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
<Login gotoMain={gotoMain}/>
{/if}

{#if page == MAIN}
<!-- <h1 style="padding-left: 15pt">Your Name - InstanceIP</h1> -->
<div class="row">
	<!-- <div class="column" style="flex-basis: 25%"><YourJobs gotoJobedit={gotoJobedit}/></div> -->
	<div class="column" style="flex-basis: 100%"><Queue gotoJobedit={gotoJobedit}/></div>
	<div class="column" style="flex-basis: 25%"><Nodes/></div>
</div>
<WorkerControl/>
{/if}

{#if page == JOBEDIT}
<Editor gotoMain={gotoMain} alias="REPLACE_THIS"/>
{/if}