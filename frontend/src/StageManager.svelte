<script lang="ts">
    import Registrations from "./components/Registrations/Registrations.svelte";
    import * as Models from "./../bindings/astroproject/astro/services/models";
    import { SelectedCompetition, CurrentStage } from "./store";
    import { onMount } from "svelte";
    import { GetStageKind } from "./../bindings/astroproject/astro/services/session";

    let currentStage: Models.Stage;
    let StageType: Models.StageKind;

    onMount(async () => {
        SelectedCompetition.subscribe(
            (competition: Models.Competition | undefined) => {
                if (competition) {
                    currentStage =
                        competition.CompetitionStages[
                            `${competition.CompetitionCurrentStageID}`
                        ];
                    CurrentStage.set(currentStage);

                    GetStageKind(
                        competition.CompetitionID,
                        currentStage.StageID
                    ).then((stageKind) => {
                        StageType = stageKind;
                    });
                }
            }
        );
    });
</script>

{#if StageType == Models.StageKind.REGISTRATIONS}
    <Registrations />
{:else}
    <p class="error-message">Please select a competition.</p>
{/if}

<style>
    .error-message {
        color: red;
        text-align: center;
    }
</style>
