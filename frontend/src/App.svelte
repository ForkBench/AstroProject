<script lang="ts">
    import { onMount } from "svelte";
    import NavBar from "./common/NavBar.svelte";
    import Registrations from "./components/Seedings/Registrations.svelte";
    import * as Session from "../bindings/changeme/astro/services/session";
    import { Competition } from "../bindings/changeme/astro/services/models";
    import { Competitions } from "./store";

    let competitions: Competition[] = [];
    let loading = true;

    onMount(async () => {
        competitions = await Session.GetCompetitions();
        loading = false;
    });

    // Listen for the "need-to-update" event
    Competitions.subscribe(async () => {
        competitions = await Session.GetCompetitions();
    });
</script>

<div class="container">
    <NavBar />
    <!-- Get competition ID -->
    {#if loading}
        <p>Loading...</p>
    {:else if competitions.length > 0}
        <Registrations />
    {:else}
        <p>No competitions available or invalid competition ID.</p>
    {/if}
</div>

<style>
    /* Put your standard CSS here */
</style>
