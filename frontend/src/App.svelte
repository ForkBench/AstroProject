<script lang="ts">
    import { onMount } from "svelte";

    import NavBar from "./components/NavBar.svelte";
    import Registrations from "./components/Registrations.svelte";
    import * as Session from "../bindings/changeme/astro/services/session";
    import { Competition } from "../bindings/changeme/astro/services/models";
    import NewPlayerForm from "./components/NewPlayerForm.svelte";

    let competitions: Competition[] = [];

    let competitionID = 0;

    onMount(async () => {
        competitions = await Session.GetCompetitions();
    });
</script>

<div class="container">
    <NavBar />
    <NewPlayerForm />
    <!-- {#if competitions.length <= competitionID}
        <p>Loading...</p>
    {:else}
        <Registrations competition={competitions[competitionID]} />
    {/if} -->
</div>

<style>
    /* Put your standard CSS here */
</style>
