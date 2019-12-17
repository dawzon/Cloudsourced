<script>
    export let name
    export let owner
    export let description
    export let status
    export let info
    export let date
    export let output

    const WAITING = 0;
    const RUNNING = 1;
    const FINISHED = 2;

    var statusString
    switch(status) {
        case WAITING:
            statusString = "⏸ Waiting"
            break;
        case RUNNING:
            statusString = "▶ Running"
            break;
        // case 2:
        //     statusString = "✗ Failed"
        //     break;
        case FINISHED:
            statusString = "✓ Finished"
            break;
    }
</script>

<style>
.status_running {
    color: greenyellow
}
.jobtype {
    color: brown
}
</style>

<div class="card">
    <strong>{name}</strong> by <a>{owner}</a> <br>
    {description} <br>
    <span class="status_running">{statusString}</span> <br>
    <span class="jobtype">{info}</span> <br>
    {#if status == WAITING}
        Submitted
    {:else if status == RUNNING}
        Started
    {:else if status == FINISHED}
        Finished
    {/if}
    {date} <br>
    <!-- __:__:__ elapsed.<br> -->
    {#if status == FINISHED}
        <a on:click="{() => alert(output)}">[ View result ]</a>
    {/if}
</div>