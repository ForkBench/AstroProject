<script lang="ts">
    import { onMount } from "svelte";
    import NavBar from "./common/NavBar.svelte";
    import Registrations from "./components/Registrations/Registrations.svelte";
    import * as Session from "../bindings/astroproject/astro/services/session";
    import { Competition } from "../bindings/astroproject/astro/structs/models";
    import { Competitions } from "./store";
    import StageManager from "./StageManager.svelte";

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
        <StageManager />
    {:else}
        <p class="error-message">Please create a competition.</p>
    {/if}
</div>

<style>
    .error-message {
        color: red;
        text-align: center;
    }
</style>
