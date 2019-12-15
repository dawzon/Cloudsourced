<script>
    import * as api from './api.js'

    export let alias

    var enabled = false;
    var currentJob
    var jobOutput = []

    function startWork() {
        enabled = true;
        
        //Taken from https://stackoverflow.com/questions/11403107/capturing-javascript-console-log
        (function(){
            var oldLog = console.log;
            console.log = function (message) {
                jobOutput.push(message)
                oldLog.apply(console, arguments);
            };
        })();
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
        Current Job: {#if currentJob} {currentJob} {:else} Hello world {/if}
        <h4>Output</h4>
        <div class="output-display">
        <!-- {#each jobOutput as line}
            {line}<br>
        {/each} -->
        Hello world!
        </div>
    {:else}
        <button on:click={startWork}>Start working</button>
    {/if}
</div>