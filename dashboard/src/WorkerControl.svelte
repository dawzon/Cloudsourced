<script>
    import { onMount } from 'svelte'
    import * as api from './api.js'

    let time = new Date()
    $: hours = time.getHours()
    $: minutes = time.getMinutes()
    $: seconds = time.setSeconds()

    onMount(() => {
        const interval = setInterval(() => {
            time = new Date()
        }, 1000)
        return () => {
            clearInterval(interval)
        }
    })

    var enabled = false;
    var currentJob
    var jobOutput = []

    function startWork() {
        enabled = true;
        
        currentJob = api.getWork(api.getAlias())

        //Hijack console.log() so ouput can be captured
        //Taken from https://stackoverflow.com/questions/11403107/capturing-javascript-console-log
        (function(){
            var oldLog = console.log;
            console.log = function (message) {
                jobOutput.push(message)
                oldLog.apply(console, arguments);
            };
        })();

        //eval()

        //submit work
    }
</script>

<style>
.output-display {
    max-height: 200pt;
    overflow-x: hidden;
    overflow-y: scroll;
}
</style>

<h2>Web Worker</h2>

<div class="card" style="width: 50%;">
    {#if enabled}
        Time remaining: {seconds} <br>
        Current Job: {#if currentJob} {currentJob.name} {:else} Hello world {/if}
        <h4>Output</h4>
        <div class="output-display">
        {#each jobOutput as line}
            {line}<br>
        {/each}
        </div>
    {:else}
        <button on:click={startWork}>Start working</button>
    {/if}
</div>