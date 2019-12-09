<script>
    var jobName = "Untitled"
    var description
    var platform
    var memory = 2048 //TODO I might be confusing megabytes and mebibytes

    //TODO get the actual files
    var files = ["swag.exe", "naenae.dll", "magic_numbers.txt"]
    var selectedFiles = []

    function toggleFile(file) {
        if(selectedFiles.includes(file))
            selectedFiles.splice(selectedFiles.indexOf(file), 1)
        else
            selectedFiles.push(file)
    }

    function submit() {
        //TODO submit the job
        console.log(memory)
        page = 0
    }
</script>

<style>
.memory {
    font-weight: bold;
    font-size: 1.25em
}
.code-flavor {
    color: aquamarine
}
</style>

<div class="card">
    <h2>Job Editor</h2>

    <table>
        <tr>
            <td>Job name</td>
            <td><input bind:value={jobName}/></td>
        </tr>
        <tr>
            <td>Description</td>
            <td><input bind:value={description}/></td>
        </tr>
        <tr>
            <td>Memory</td>
            <td class="memory">
                <input type="radio" bind:group={memory} value={512}> 512M
                <input type="radio" bind:group={memory} value={1024}> 1G
                <input type="radio" bind:group={memory} value={2048}> 2G
                <input type="radio" bind:group={memory} value={4096}> 4G
                <input type="radio" bind:group={memory} value={8192}> 8G
                <input type="radio" bind:group={memory} value={16384}> 16G
                <input type="radio" bind:group={memory} value={32768}> 32G
            </td>
        </tr>
        <tr>
            <td>Platform</td>
            <td>
                <select bind:value={platform}>
                    <option>Windows</option>
                    <option>Linux</option>
                    <option>Web</option>
                </select>
            </td>
        </tr>
        <tr>
            <td>Files</td>
            <td>
                {#if platform != "Web"}
                    {#each files as file, i}
                        <input type="checkbox" on:click={() => toggleFile(file)}> {file} <br>
                    {/each}
                {:else}
                    N/A
                {/if}
            </td>
        </tr>
        <tr>
            <td>Script</td>
            <td>
                <div style="display:flex; flex-direction:column;">
                    {#if platform == "Windows"}
                        <p class="code-flavor">cmd (batch)</p>
                    {:else if platform == "Linux"}
                        <p class="code-flavor">bash</p>
                    {:else if platform == "Web"}
                        <p class="code-flavor">Javascript</p>
                    {/if}
                    <textarea></textarea>
                </div>
            </td>
        </tr>
        <tr>
            <td></td>
            <td><button on:click={submit}>Submit</button></td>
        </tr>
    </table>
</div>