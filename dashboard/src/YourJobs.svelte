<script>
    import JobCard from './JobCard.svelte'
    import * as api from './api.js'
</script>

<style>
</style>

<h1>Your jobs</h1>
<div class="card">
{#await api.getJobsByOwner("noobmaster69")}
<p>...</p>
{:then yourJobs}
    {#if yourJobs.length == 0}
        You aren't running any jobs.
    {:else}
        {#each yourJobs as job}
            <JobCard
                name={job.name}
                owner={job.owner}
                description={job.description}
                status={job.status}
                info="TODO__"
                submitdate={job.submitdate} />
        {/each}
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}
    
    <button>New Job</button>
</div>