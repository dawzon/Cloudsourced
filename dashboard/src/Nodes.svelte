<script>
    import * as api from './api.js'
    let activeNodes = 5
</script>

<style>
</style>

<h1>Nodes</h1>
{#await api.getActiveNodes()}
    <p>...</p>
{:then activeNodes}
    {#if activeNodes !== null}
        {activeNodes} active nodes
        {#each activeNodes as node}
            <div class="card">{node}</div>
        {/each}
    {:else}
        <p>No active nodes</p>
    {/if}
{:catch error}
    <p>Something bad happened: {error.message}</p>
{/await}