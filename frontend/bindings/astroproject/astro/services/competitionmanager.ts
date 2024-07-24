// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * For now, competitionManager has a UserSession field, which is a pointer to a Session struct.
 * I guess it's better for now to keep it like this insted of having a Manager in Session struct.
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as structs$0 from "../structs/models.js";

export function AddPlayerToCompetition(competitionID: number, player: structs$0.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(310104718, competitionID, player) as any;
    return $resultPromise;
}

export function GetAllPlayersFromCompetition(competitionID: number): Promise<(structs$0.Player | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3364829354, competitionID) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function RemovePlayerFromCompetition(competitionID: number, player: structs$0.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2605805672, competitionID, player) as any;
    return $resultPromise;
}

export function UpdateCompetitionPlayer(competitionID: number, player: structs$0.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2983979889, competitionID, player) as any;
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = structs$0.Player.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
