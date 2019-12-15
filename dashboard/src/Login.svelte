<script>
    import * as api from './api.js'

    export let gotoMain

    var alias

    function login() {
        api.setAlias(alias)
        gotoMain.call()
    }
</script>

<style>
#title {
    color: white;
    text-align: center
}
div {
    display: flex;
    flex-direction: column;
    margin-top: 2em;
    margin-left: 40%;
    margin-right: 40%;
    padding: 10pt;
}
</style>

<div class="card">
    <h1 id="title">Cloudsourced</h1>
    <!-- <p>Username</p>
    <input>
    <p>Password</p>
    <input> -->
    <p>Choose an alias:</p>
    <input bind:value={alias}>
    <br>
    <!-- <p style="color:red">This login is for demonstration purposes only.  <strong>Whatever you type here will be sent in plaintext via HTTP.</strong></p> -->
    <button on:click={login}>Submit</button>
</div>