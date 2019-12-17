<script>
    import { onMount } from 'svelte'
    import * as api from './api.js'

    var enabled = false;
    var currentJob
    var jobOutput = []

    //Borrowed from https://stackoverflow.com/questions/951021/what-is-the-javascript-version-of-sleep
    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    async function startWork() {
        
        currentJob = await api.getWork();
        if(currentJob == null) {
            alert("No work available.")
            return;
        }

        enabled = true;
        await sleep(1000);

        //Hijack console.log() so ouput can be captured
        //Taken from https://stackoverflow.com/questions/11403107/capturing-javascript-console-log
        (function(){
            var oldLog = console.log;
            console.log = function (message) {
                jobOutput.push(message)
                oldLog.apply(console, arguments);
            };
        })();

        //Run the job!
        eval(currentJob.script)

        await sleep(1000);

        var allOutput = jobOutput.join("\n")
        api.submitWork(currentJob.id, api.getAlias(), allOutput)

        enabled = false
        alert("Job complete!")
    }

    function timeLeft() {
        return endTime - new Date()
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
        <!-- Time remaining: {secondsLeft} <br> -->
        Current Job: {#if currentJob} {currentJob.name} {:else} - {/if}
        <h4>Output</h4>
        <!-- <div class="output-display">
        {#each jobOutput as line}
            {line}<br>
        {/each}
        </div> -->
    {:else}
        <button on:click={startWork}>Start working</button>
    {/if}
</div>