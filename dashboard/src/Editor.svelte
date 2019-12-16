<script>
    import * as api from './api.js'

    export let gotoMain

    const timeoutTip = "This is the amount of time to wait before the job will be considered abandoned.  This is to prevent your job from being lost if the worker goes offline during job execution."

    var jobName = "Untitled"
    var description
    var platform
    //var memory = 2048 //TODO I might be confusing megabytes and mebibytes
    var script
    var timeout = 10

    //var files = ["swag.exe", "naenae.dll", "magic_numbers.txt"]
    //var selectedFiles = []

    // function toggleFile(file) {
    //     if(selectedFiles.includes(file))
    //         selectedFiles.splice(selectedFiles.indexOf(file), 1)
    //     else
    //         selectedFiles.push(file)
    // }

    function submit() {
        //TODO set owner
        api.submitJob(jobName, description, platform, script, api.getAlias(), timeout*60)
        //TODO tell the user about the result
        gotoMain.call()
    }

    function cancel() {
        gotoMain.call()
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
.container {
    display: flex;
    justify-content: center;
}
</style>

<div class="container">
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
        <!-- <tr>
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
        </tr> -->
        <tr>
            <td>Platform</td>
            <td>
                <select bind:value={platform}>
                    <!-- <option>Windows</option>
                    <option>Linux</option> -->
                    <option>Web</option>
                </select>
            </td>
        </tr>
        <!-- <tr>
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
        </tr> -->
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
                    <textarea bind:value={script}></textarea>
                </div>
            </td>
        </tr>
        <tr>
            <td>Timeout <a on:click="{() => alert(timeoutTip)}">[ ? ]</a></td>
            <td><input type="number" bind:value={timeout}> Minutes</td>
        </tr>
        <tr>
            <td></td>
            <td> <button on:click={submit}>Submit</button> <button on:click={cancel}>Cancel</button> </td>
        </tr>
    </table>
</div>
</div>