<script lang="ts">
    import { onMount } from "svelte";
    import * as Models from "../../../bindings/changeme/astro/services/models";
    import * as Session from "../../../bindings/changeme/astro/services/session";
    import { SelectedCompetition } from "../../store";
    import { GenerateRandomPlayer } from "../../../bindings/changeme/astro/services/player";
    import swal from "sweetalert";

    let players: (Models.Player | null)[] = [];
    let filteredPlayers: (Models.Player | null)[] = [];
    let competition: Models.Competition | undefined;

    SelectedCompetition.subscribe((value) => {
        competition = value;
        loadPlayers();
    });

    async function loadPlayers() {
        if (competition === undefined) {
            return;
        }

        Session.GetAllPlayersFromCompetition(competition.CompetitionID).then(
            (result) => {
                players = result;
            }
        );
    }

    onMount(async () => {
        await loadPlayers();
        sortBy = { key: "PlayerInitialRank", asc: true };
    });

    function getPlayerValueByKey(player: Models.Player | null, key: string) {
        if (player === null) {
            return "";
        }

        switch (key) {
            case "PlayerInitialRank":
                return player.PlayerInitialRank;
            case "PlayerFirstname":
                return player.PlayerFirstname;
            case "PlayerLastname":
                return player.PlayerLastname;
            case "PlayerClub":
                return player.PlayerClub;
            default:
                return "";
        }
    }

    let sortBy = { key: "PlayerInitialRank", asc: true };

    $: players = players.sort((a, b) => {
        let aValue = getPlayerValueByKey(a, sortBy.key);
        let bValue = getPlayerValueByKey(b, sortBy.key);
        if (aValue === null || bValue === null) {
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

    let searchTerms = "";

    $: filteredPlayers = players.filter((player) => {
        if (player === null) {
            return [];
        }

        return (
            player.PlayerFirstname.toLowerCase().includes(
                searchTerms.toLowerCase()
            ) ||
            player.PlayerLastname.toLowerCase().includes(
                searchTerms.toLowerCase()
            ) ||
            player.PlayerClub?.club_name
                .toLowerCase()
                .includes(searchTerms.toLowerCase())
        );
    });
</script>

<div class="registration-container">
    <input
        type="text"
        bind:value={searchTerms}
        placeholder="Search for a player"
    />
    {#if competition == undefined}
        <p>No competition selected</p>
    {:else}
        <table>
            <thead>
                <tr>
                    <th
                        on:click={() => {
                            sortBy = {
                                key: "PlayerInitialRank",
                                asc: !sortBy.asc,
                            };
                        }}>Initial Rank</th
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
                            sortBy = {
                                key: "PlayerLastname",
                                asc: !sortBy.asc,
                            };
                        }}>Lastname</th
                    >
                    <th
                        on:click={() => {
                            sortBy = { key: "PlayerClub", asc: !sortBy.asc };
                        }}>Club</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each filteredPlayers as player}
                    {#if player != null}
                        <tr>
                            <td
                                on:dblclick={async () => {
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

                                    if (newRank) {
                                        player.PlayerInitialRank =
                                            parseInt(newRank);

                                        if (isNaN(player.PlayerInitialRank)) {
                                            await swal(
                                                "Error",
                                                "Rank must be a number",
                                                "error"
                                            );
                                            return;
                                        }

                                        if (competition != undefined)
                                            await Session.UpdateCompetitionPlayer(
                                                competition.CompetitionID,
                                                player
                                            ).then(() => {
                                                loadPlayers();
                                            });
                                    }
                                }}>{player.PlayerInitialRank}</td
                            >
                            <td
                                on:dblclick={async () => {
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

                                    if (newFirstname) {
                                        player.PlayerFirstname = newFirstname;
                                        if (competition != undefined)
                                            await Session.UpdateCompetitionPlayer(
                                                competition.CompetitionID,
                                                player
                                            ).then(() => {
                                                loadPlayers();
                                            });
                                    }
                                }}>{player.PlayerFirstname}</td
                            >
                            <td
                                on:dblclick={async () => {
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

                                    if (newLastname) {
                                        player.PlayerLastname = newLastname;
                                        if (competition != undefined)
                                            await Session.UpdateCompetitionPlayer(
                                                competition.CompetitionID,
                                                player
                                            ).then(() => {
                                                loadPlayers();
                                            });
                                    }
                                }}>{player.PlayerLastname}</td
                            >
                            <td
                                >{#if player.PlayerClub}{player.PlayerClub
                                        .club_name}{:else}Sans Nom{/if}</td
                            >
                        </tr>
                    {/if}
                {/each}
                <tr id="add-player-tr">
                    <td colspan="4">
                        <button
                            on:click={async () => {
                                var player = await GenerateRandomPlayer();

                                if (competition != undefined)
                                    await Session.AddPlayerToCompetition(
                                        competition.CompetitionID,
                                        player
                                    ).then(() => {
                                        loadPlayers();
                                    });
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
</style>
