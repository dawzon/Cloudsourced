<script>
    import JobCard from './JobCard.svelte'
    import * as api from './api.js'

    // let tempjobs = [
    //     { name: "My Job", description: "Okay, this is epic.", user: "Dawson", status: "▶ RUNNING", submitdate: "12:46 AM - November 11, 2019" },
    //     { name: "My Job but in 4k", description: "Okay, this is even more epic.", user: "Dawson", status: "▶ RUNNING on [node link]", submitdate: "1:15 AM - November 11, 2019" },
    //     { name: "The title of this one is really long to test the CSS layout so yeah this is a test"},
    //     {},
    //     {},
    //     {},
    //     {}
    // ]
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
    {#each runningJobs as job}
        <JobCard name={job.name} owner={job.owner} description={job.description} status={job.status} info="TODO__" submitdate={job.submitdate} />
    {/each}
{:catch error}
<p>Something bad happened: {error.message}</p>
{/await}
</div>

<h2>Waiting</h2>
<JobCard name="Start the game already" owner="Dawson2" description="This job hasn't run, for some reason" status="WAITING" info="2ManyResources4U" submitdate="yesterday, maybe"/>

<h2>Failed</h2>