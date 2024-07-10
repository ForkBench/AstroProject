<script lang="ts">
    import { onMount } from "svelte";

    import NavBar from "./components/NavBar.svelte";
    import Registrations from "./components/Registrations.svelte";
    import * as Session from "../bindings/changeme/astro/services/session";
    import { Competition } from "../bindings/changeme/astro/services/models";

    let competitions: Competition[] = [];

    onMount(async () => {
        competitions = await Session.GetCompetitions();
    });
</script>

<div class="container">
    <NavBar />
    {#if competitions.length === 0}
        <p>Loading...</p>
    {:else}
        <Registrations competition={competitions[0]} />
    {/if}
</div>

<style>
    /* Put your standard CSS here */
</style>
