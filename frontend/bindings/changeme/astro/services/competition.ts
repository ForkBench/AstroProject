// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * Competition : Competition details
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function AddPlayer(player: $models.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3220417277, player) as any;
    return $resultPromise;
}

export function AddPlayerToStage(player: $models.Player, stage: $models.Stage): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2532203882, player, stage) as any;
    return $resultPromise;
}

export function AddStage(stage: $models.Stage): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2261622806, stage) as any;
    return $resultPromise;
}

export function FinishCompetition(): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(826570137) as any;
    return $resultPromise;
}

export function InitCompetition(): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(582268878) as any;
    return $resultPromise;
}

export function RemovePlayer(player: $models.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1491269932, player) as any;
    return $resultPromise;
}

export function RemovePlayerFromStage(player: $models.Player, stage: $models.Stage): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(362483458, player, stage) as any;
    return $resultPromise;
}

export function StartCompetition(): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3322759170) as any;
    return $resultPromise;
}

export function String(): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2560236652) as any;
    return $resultPromise;
}

export function UpdatePlayer(player: $models.Player | null): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(132875429, player) as any;
    return $resultPromise;
}
