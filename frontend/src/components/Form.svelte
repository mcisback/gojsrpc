<script>

import GoJsRpc from "../rpc/GoJsRpc";

const rpc = new GoJsRpc("http://localhost:3000/gorpc")

let form = {
    a: 0,
    b: 0
}

let sum = null

async function onSubmit() {
    sum = await rpc.sum(form)

    console.log('sum: ', sum)
}

</script>

<form on:submit|preventDefault={onSubmit}>

    <input type="number" name="a" id="a" bind:value={form.a}>

    <br>

    <input type="number" name="b" id="b" bind:value={form.b}>

    <br>

    <button type="submit">
        Calcola Somma
    </button>

    {#if sum}
        <p>
            La Somma è: {sum}
        </p>
    {/if}

</form>