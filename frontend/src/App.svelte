<script lang="ts">
    import { Competition } from "../bindings/changeme/astro/services/models";
    import * as Session from "../bindings/changeme/astro/services/session";

    let competitions: Competition[] = [];

    function getCompetitions() {
        Session.GetCompetitions().then((data) => {
            competitions = data;
        });
    }
</script>

<div class="container">
    <h1>Competitions</h1>
    <button on:click={getCompetitions}>Get</button>
    <ul>
        {#each competitions as competition}
            <li>{competition.CompetitionName}</li>
        {/each}
    </ul>
</div>

<style>
    /* Put your standard CSS here */
</style>
