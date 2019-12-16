<script>
    import JobCard from './JobCard.svelte'
    import * as api from './api.js'

    export let gotoJobedit

    var queuedJobs = null
    var runningJobs = null
    var failedJobs = null
    var finishedJobs = null
    var yourJobs = null

    async function refreshJobs() {
        queuedJobs = await api.getQueuedJobs()
        runningJobs = await api.getRunningJobs()
        failedJobs = await api.getFailedJobs()
        finishedJobs = await api.getFinishedJobs()
        yourJobs = await api.getJobsByOwner(api.alias)
    }
    setInterval(refreshJobs, 5000)

    refreshJobs()
</script>

<style>
.autoflow {
    display: grid;
    grid-auto-flow: row;
    gap: 2pt;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))
}
</style>

<h1>Your jobs</h1>
<div class="card">
{#if yourJobs == null || yourJobs.length == 0}
    You aren't running any jobs.
{:else}
    {#each yourJobs as job}
        <JobCard
            name={job.name}
            owner={job.owner}
            description={job.description}
            status={job.status}
            info="Web"
            submitdate={job.submitdate} />
    {/each}
{/if}
<button on:click={gotoJobedit}>New Job</button>
</div>

<h1>Queue</h1>

<h2>Waiting</h2>
<div class="autoflow">
{#if queuedJobs !== null}
    {#each queuedJobs as job}
        <JobCard
            name={job.name}
            owner={job.owner}
            description={job.description}
            status={job.status}
            info="Web"
            submitdate={job.submitdate} />
    {/each}
{:else}
    <p>No jobs.</p>
{/if}
</div>

<h2>In Progress</h2>
<div class="autoflow">
{#if runningJobs !== null}
    {#each runningJobs as job}
        <JobCard
            name={job.name}
            owner={job.owner}
            description={job.description}
            status={job.status}
            info="Web"
            submitdate={job.submitdate} />
    {/each}
{:else}
    <p>No jobs.</p>
{/if}
</div>

<h2>Failed</h2>
<div class="autoflow">
{#if failedJobs !== null}
    {#each failedJobs as job}
        <JobCard
            name={job.name}
            owner={job.owner}
            description={job.description}
            status={job.status}
            info="Web"
            submitdate={job.submitdate} />
    {/each}
{:else}
    <p>No jobs.</p>
{/if}
</div>

<h2>Finished</h2>
<div class="autoflow">
{#if finishedJobs !== null}
    {#each finishedJobs as job}
        <JobCard
            name={job.name}
            owner={job.owner}
            description={job.description}
            status={job.status}
            info="Web"
            submitdate={job.submitdate} />
    {/each}
{:else}
    <p>No jobs.</p>
{/if}
</div>