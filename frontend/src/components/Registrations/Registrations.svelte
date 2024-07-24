<script lang="ts">
    /*
    
    Only stage where we don't care about the stage data : everything is already in the competition object.

    TODO: Try if headers can be components
    */
    import { onMount } from "svelte";
    import * as Models from "../../../bindings/astroproject/astro/structs/models";
    import { GetCompetition } from "../../../bindings/astroproject/astro/services/session";
    import { SelectedCompetition, CurrentStage } from "../../store";
    import {
        GenerateRandomID,
        GenerateRandomPlayer,
    } from "../../../bindings/astroproject/astro/services/personmanager";
    import {
        GetPlayersFromCompetitionStage,
        AddPlayerToCompetitionStage,
        UpdateStagePlayer,
    } from "../../../bindings/astroproject/astro/services/stagemanager";
    import {
        UpdateCompetitionPlayer,
        AddPlayerToCompetition,
        IsPlayerIDFreeInCompetition,
    } from "../../../bindings/astroproject/astro/services/competitionmanager";
    import swal from "sweetalert";
    import { getNationFlag, getNationFlatAlt } from "../../Util";

    // Variables that will be used in the component
    let players: Models.Player[] = [];
    let filteredPlayers: Models.Player[] = [];
    let competition: Models.Competition;
    let stage: Models.Stage;
    let sortBy = { key: "", asc: true };
    let searchTerms = "";

    // Listen to any change in the selected competition
    SelectedCompetition.subscribe((value) => {
        competition = value;
    });

    // Listen to any change in the current stage
    CurrentStage.subscribe((value) => {
        stage = value;
        loadPlayers();
    });

    // Lod the stage and competition
    async function loadStageAndCompetition() {
        if (competition === undefined) {
            return;
        }

        await GetCompetition(competition.CompetitionID).then((result) => {
            if (result !== null) {
                SelectedCompetition.set(result);
            }
        });
    }

    // Load the players from the selected competition
    async function loadPlayers() {
        if (competition === undefined) {
            return;
        }

        GetPlayersFromCompetitionStage(
            competition.CompetitionID,
            stage.StageID
        ).then((result) => {
            players = result;
        });
    }

    // When the page is loaded, load the players, and set the default sorting
    onMount(async () => {
        await loadPlayers();

        // By default, sort by initial rank
        sortBy = { key: "PlayerInitialRank", asc: true };
    });

    // Function to get the value of a player by key
    function getPlayerValueByKey(player: Models.Player, key: string) {
        switch (key) {
            case "PlayerInitialRank":
                return player.PlayerInitialRank;
            case "PlayerFirstname":
                return player.PlayerFirstname;
            case "PlayerLastname":
                return player.PlayerLastname;
            case "PlayerClub":
                return player.PlayerClub.club_name;
            case "PlayerRegion":
                return player.PlayerRegion.region_name;
            case "PlayerCountry":
                return player.PlayerNation.nation_name;
            default:
                return "";
        }
    }

    // To make players always sorted
    $: players = players.sort((a, b) => {
        let aValue = getPlayerValueByKey(a, sortBy.key);
        let bValue = getPlayerValueByKey(b, sortBy.key);

        if (aValue === undefined || bValue === undefined) {
            return 0;
        }

        if (aValue < bValue) {
            return sortBy.asc ? -1 : 1;
        }
        if (aValue > bValue) {
            return sortBy.asc ? 1 : -1;
        }
        return 0;
    });

    // Filter the players by search terms
    $: filteredPlayers = players.filter((player) => {
        return (
            player.PlayerFirstname.toLowerCase().includes(
                searchTerms.toLowerCase()
            ) ||
            player.PlayerLastname.toLowerCase().includes(
                searchTerms.toLowerCase()
            ) ||
            player.PlayerClub.club_name
                .toLowerCase()
                .includes(searchTerms.toLowerCase())
        );
    });

    async function updatePlayer(player: Models.Player) {
        // Update the player in the competition
        await UpdateCompetitionPlayer(competition.CompetitionID, player);

        // Update the player in the stage
        await UpdateStagePlayer(
            competition.CompetitionID,
            stage.SeedingStageID,
            player
        );

        loadPlayers();
    }
</script>

<div class="registration-container">
    <input
        type="text"
        bind:value={searchTerms}
        placeholder="Search for a player"
    />
    {#if competition == undefined}
        <p class="error-message">Please select a competition.</p>
        <!-- TODO: Uniformize it -->
    {:else}
        <table>
            <thead>
                <tr>
                    <!-- 
                    The following code is used to sort the players by the column clicked.                    
                    -->
                    <th
                        on:click={() => {
                            sortBy = { key: "PlayerPresent", asc: !sortBy.asc };
                        }}
                    >
                        Present
                    </th>
                    <th
                        on:click={() => {
                            sortBy = {
                                key: "PlayerInitialRank",
                                asc: !sortBy.asc,
                            };
                        }}>Rank</th
                    >
                    <th
                        on:click={() => {
                            sortBy = {
                                key: "PlayerLastname",
                                asc: !sortBy.asc,
                            };
                        }}>Lastname</th
                    >
                    <th
                        on:click={() => {
                            sortBy = {
                                key: "PlayerFirstname",
                                asc: !sortBy.asc,
                            };
                        }}>Firstname</th
                    >
                    <th
                        on:click={() => {
                            sortBy = { key: "PlayerClub", asc: !sortBy.asc };
                        }}>Club</th
                    >
                    <th
                        on:click={() => {
                            sortBy = { key: "PlayerRegion", asc: !sortBy.asc };
                        }}>Region</th
                    >
                    <th
                        on:click={() => {
                            sortBy = { key: "PlayerCountry", asc: !sortBy.asc };
                        }}>Country</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each filteredPlayers as player}
                    <!-- Classes bellow help improving the UI -->
                    <tr
                        class:gold={player.PlayerInitialRank == 1}
                        class:silver={player.PlayerInitialRank == 2}
                        class:bronze={player.PlayerInitialRank == 3}
                    >
                        <td>
                            <!-- TODO: input checkbox -->
                            <!-- <input
                                    type="checkbox"
                                    checked={player.PlayerPresent}
                                    on:change={async () => {
                                        player.PlayerPresent =
                                            !player.PlayerPresent;
                                        if (competition != undefined)
                                            await UpdateCompetitionPlayer(
                                                competition.CompetitionID,
                                                player
                                            ).then(() => {
                                                loadPlayers();
                                            });
                                    }}
                                /> -->
                        </td>
                        <td
                            class="el first"
                            on:dblclick={async () => {
                                // Show a prompt to the user to change the rank
                                let newRank = await swal({
                                    content: {
                                        element: "input",
                                        attributes: {
                                            placeholder: "New rank",
                                            type: "number",
                                        },
                                    },
                                    buttons: {
                                        cancel: true,
                                        confirm: true,
                                    },
                                });

                                // Check if the user confirmed the prompt
                                if (newRank) {
                                    player.PlayerInitialRank =
                                        parseInt(newRank);

                                    // Update the player in the competition
                                    if (competition != undefined) {
                                        updatePlayer(player);
                                    }
                                }
                            }}>{player.PlayerInitialRank}</td
                        >
                        <td
                            class="el"
                            on:dblclick={async () => {
                                // Show a prompt to the user to change the lastname
                                let newLastname = await swal({
                                    content: {
                                        element: "input",
                                        attributes: {
                                            placeholder: "New lastname",
                                        },
                                    },
                                    buttons: {
                                        cancel: true,
                                        confirm: true,
                                    },
                                });

                                // Check if the user confirmed the prompt
                                if (newLastname) {
                                    player.PlayerLastname = newLastname;
                                    if (competition != undefined) {
                                        updatePlayer(player);
                                    }
                                }
                            }}>{player.PlayerLastname.toLocaleUpperCase()}</td
                        >
                        <td
                            class="el"
                            on:dblclick={async () => {
                                // Show a prompt to the user to change the firstname
                                let newFirstname = await swal({
                                    content: {
                                        element: "input",
                                        attributes: {
                                            placeholder: "New firstname",
                                        },
                                    },
                                    buttons: {
                                        cancel: true,
                                        confirm: true,
                                    },
                                });

                                // Check if the user confirmed the prompt
                                if (newFirstname) {
                                    player.PlayerFirstname = newFirstname;
                                    if (competition != undefined) {
                                        updatePlayer(player);
                                    }
                                }
                            }}>{player.PlayerFirstname}</td
                        >
                        <td class="el"
                            >{#if player.PlayerClub}{player.PlayerClub
                                    .club_name}{:else}Sans Nom{/if}</td
                        >
                        <td class="el"
                            >{#if player.PlayerRegion}{player.PlayerRegion
                                    .region_name}{:else}Sans Nom{/if}</td
                        >
                        <td class="flag-container el last"
                            ><img
                                src={getNationFlag(player.PlayerNation)}
                                alt="Player's flag : {getNationFlatAlt(
                                    player.PlayerNation
                                )}"
                                class="flag"
                            /></td
                        >
                    </tr>
                {/each}
                <tr id="add-player-tr">
                    <td colspan="4">
                        <button
                            on:click={async () => {
                                if (competition != undefined) {
                                    // Generate a random player
                                    var player = await GenerateRandomPlayer(
                                        competition.CompetitionID
                                    );

                                    // Set the initial rank of the player
                                    player.PlayerInitialRank =
                                        players.length + 1;

                                    while (
                                        !IsPlayerIDFreeInCompetition(
                                            competition.CompetitionID,
                                            player.PlayerID
                                        )
                                    ) {
                                        player.PlayerID =
                                            await GenerateRandomID(
                                                competition.CompetitionID
                                            );
                                    }

                                    await AddPlayerToCompetitionStage(
                                        competition.CompetitionID,
                                        stage.SeedingStageID,
                                        player
                                    );

                                    await AddPlayerToCompetition(
                                        competition.CompetitionID,
                                        player
                                    );

                                    loadStageAndCompetition();
                                }
                            }}
                        >
                            Add Player
                        </button></td
                    >
                </tr>
            </tbody>
        </table>
    {/if}
</div>

<style>
    .registration-container {
        margin-top: 5vh;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    input {
        width: 200px;
        margin-bottom: 10px;
        padding: 10px;
        font-size: 1em;
        border-radius: 10px;
        border: solid 1px #003566;
    }

    .gold .el {
        background-color: #e9c46a;
    }

    .silver .el {
        background-color: #adb5bd;
    }

    .bronze .el {
        background-color: #f28482;
    }

    .error-message {
        color: red;
        text-align: center;
    }

    .flag-container {
        text-align: left;
    }

    .flag {
        width: 20px;
        height: 20px;
        vertical-align: middle;
    }
</style>
