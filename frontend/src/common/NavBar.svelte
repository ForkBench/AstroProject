<script lang="ts">
    import { Competition } from "../../bindings/changeme/astro/services/models";
    import { randomColor, rightBrighness } from "../Util";
    import {
        GetCompetitions,
        AddCompetition,
        RemoveCompetition,
    } from "../../bindings/changeme/astro/services/session";
    import { onMount } from "svelte";
    import { SelectedCompetition } from "../store";

    let competitions: Competition[] = [];

    async function loadCompetitions() {
        competitions = await GetCompetitions();

        // Emit a signal to the parent component
        dispatchEvent(new CustomEvent("need-to-update", {}));
    }

    onMount(loadCompetitions);
</script>

<div class="nav-bar">
    <div class="competition-container">
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        {#each competitions as competition}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div
                class="competition-el"
                style="background-color: {randomColor(rightBrighness)};"
                on:click={() => {
                    SelectedCompetition.set(competition);
                }}
            >
                <span>{competition.CompetitionName}</span>
                <button
                    on:click={async () => {
                        await RemoveCompetition(competition.CompetitionID);
                        loadCompetitions();

                        SelectedCompetition.set(undefined);
                    }}>X</button
                >
            </div>
        {/each}
    </div>
    <div class="competition-modifier">
        <button
            on:click={async () => {
                const newCompetition = await AddCompetition(
                    "Nouvelle compet",
                    "U20",
                    "Foil"
                );
                loadCompetitions();
            }}>Add competition</button
        >
    </div>
</div>

<style>
    .nav-bar {
        padding: 50px 10px 25px;
        display: grid;
        /* 70% for competitions, 30% for button */
        grid-template-columns: 70% 30%;
        gap: 10px;
        border-bottom: solid #003566 3px;
    }

    .competition-container {
        display: grid;
        overflow-x: scroll;
        grid-auto-flow: column;
        gap: 10px;
    }

    .competition-modifier {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .competition-el {
        min-width: 20vw;
        display: flex;
        justify-content: space-between;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 5px;
    }
</style>
