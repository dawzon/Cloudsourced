<script>
    import JobCard from './JobCard.svelte'
    import * as api from './api.js'
</script>

<style>
.autoflow {
    display: grid;
    grid-auto-flow: row;
    gap: 5pt;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))
}
</style>

<h1>Queue</h1>

<div class=row>
    <button>By status</button>
    <button>By submission date</button>
</div>


<h2>In Progress</h2>
<div class="autoflow">
{#await api.getRunningJobs()}
    <p>...</p>
{:then runningJobs}
    {#if runningJobs !== null}
        {#each runningJobs as job}
            <JobCard
                name={job.name}
                owner={job.owner}
                description={job.description}
                status={job.status}
                info="TODO__"
                submitdate={job.submitdate} />
        {/each}
    {:else}
        <p>No jobs.</p>
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}
</div>

<h2>Waiting</h2>
{#await api.getQueuedJobs()}
    <p>...</p>
{:then queuedJobs}
    {#if queuedJobs !== null}
        {#each queuedJobs as job}
            <JobCard
                name={job.name}
                owner={job.owner}
                description={job.description}
                status={job.status}
                info="TODO__"
                submitdate={job.submitdate} />
        {/each}
    {:else}
        <p>No jobs.</p>
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}

<h2>Failed</h2>
{#await api.getFailedJobs()}
    <p>...</p>
{:then failedJobs}
    {#if failedJobs !== null}
        {#each failedJobs as job}
            <JobCard
                name={job.name}
                owner={job.owner}
                description={job.description}
                status={job.status}
                info="TODO__"
                submitdate={job.submitdate} />
        {/each}
    {:else}
        <p>No jobs.</p>
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}

<h2>Finished</h2>
{#await api.getFinishedJobs()}
    <p>...</p>
{:then finishedJobs}
    {#if finishedJobs !== null}
        {#each finishedJobs as job}
            <JobCard
                name={job.name}
                owner={job.owner}
                description={job.description}
                status={job.status}
                info="TODO__"
                submitdate={job.submitdate} />
        {/each}
    {:else}
        <p>No jobs.</p>
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}