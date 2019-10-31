<script>
    function login() {
        alert('oh yeah yeah')
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
    margin-left: 15em;
    margin-right: 15em;
    border-style: solid;
    border-radius: 1em;
    border-color: white;
    padding: 1em;
}
</style>

<div>
    <h1 id="title">Cloudsourced</h1>
    <input>
    <input>
    <button on:click={login()}>Submit</button>
</div>