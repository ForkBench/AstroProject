<script lang="ts">
    import { onMount } from "svelte";
    import * as Models from "../../bindings/changeme/astro/services/models";
    import * as Session from "../../bindings/changeme/astro/services/session";
    import swal from "sweetalert";

    export let competition: Models.Competition;

    let clubs: Models.Club[] = [];

    let playerFirstname: string = "",
        playerLastname: string = "",
        playerClub: string = "";

    // Fetch /resources/clubs
    async function loadClubs() {
        var json = await fetch("/resources/clubs.json").then((r) => r.json());
        clubs = json.map((c: any) => new Models.Club(c));
    }

    onMount(loadClubs);
</script>

<div class="form">
    <label for="player-firstname">Firstname</label>
    <input
        type="text"
        id="player-firstname"
        name="player-firstname"
        placeholder="John"
        required
        bind:value={playerFirstname}
    />

    <label for="player-lastname">Lastname</label>
    <input
        type="text"
        id="player-lastname"
        name="player-lastname"
        placeholder="Doe"
        required
        bind:value={playerLastname}
    />

    <label for="player-club">Club</label>
    <input
        type="text"
        id="player-club"
        name="player-club"
        placeholder="Club"
        list="clubs"
        required
        bind:value={playerClub}
    />

    <button
        type="submit"
        on:click={async () => {
            // First check if the club exists
            let club = clubs.find((c) => c.club_name == playerClub);

            if (club === undefined) {
                await swal("Error", "Club does not exist", "error");
                return;
            }

            // Check if the name is not empty
            if (playerFirstname == "" || playerLastname == "") {
                await swal("Error", "Name cannot be empty", "error");
                return;
            }

            // Create the player
            let player = new Models.Player({
                PlayerFirstname: playerFirstname,
                PlayerLastname: playerLastname,
                PlayerClub: club,
            });

            // Send the player to the backend
            Session.AddPlayerToCompetition(competition.CompetitionID, player);

            // Clear the form
            playerFirstname = "";
            playerLastname = "";
            playerClub = "";
        }}>Add player</button
    >

    <datalist id="clubs">
        {#each clubs as club}
            <option value={club.club_name} />
        {/each}
    </datalist>
</div>

<style>
    .form {
        max-width: 30%;
        margin: 0 auto;
        border-radius: 5px;
        padding: 40px;
        display: flex;
        flex-direction: column;
        align-items: center;
        background: #e9ecef;
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
