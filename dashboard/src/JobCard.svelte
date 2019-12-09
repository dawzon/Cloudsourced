<script>
    export let name
    export let owner
    export let description
    export let status
    export let info
    export let submitdate

    var statusString
    switch(status) {
        case 0:
            statusString = "⏸ Waiting"
            break;
        case 1:
            statusString = "▶ Running"
            break;
        case 2:
            statusString = "✗ Failed"
            break;
        case 3:
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
    Started {submitdate}, <br>
    __:__:__ elapsed.<br>
</div>