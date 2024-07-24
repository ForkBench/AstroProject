// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Create as $Create} from "@wailsio/runtime";

/**
 * Category : Age category
 */
export enum Category {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero = 0,

    /**
     * 0
     */
    U7 = 0,

    /**
     * 1
     */
    U9 = 1,

    /**
     * 2
     */
    U11 = 2,

    /**
     * 3
     */
    U13 = 3,

    /**
     * 4
     */
    U15 = 4,

    /**
     * 5
     */
    U17 = 5,

    /**
     * 6
     */
    U20 = 6,

    /**
     * 7
     */
    SENIOR = 7,

    /**
     * 8
     */
    VETERAN = 8,
};

export class Club {
    "club_id": number;
    "club_name": string;

    /** Creates a new Club instance. */
    constructor($$source: Partial<Club> = {}) {
        if (!("club_id" in $$source)) {
            this["club_id"] = 0;
        }
        if (!("club_name" in $$source)) {
            this["club_name"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Club instance from a string or object.
     */
    static createFrom($$source: any = {}): Club {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new Club($$parsedSource as Partial<Club>);
    }
}

/**
 * Competition : Competition details
 */
export class Competition {
    /**
     * 255 competitions max
     */
    "CompetitionID": number;
    "CompetitionName": string;
    "CompetitionCategory": Category;
    "CompetitionWeapon": Weapon;
    "CompetitionState": State;
    "CompetitionMaxStageNumber": number;
    "CompetitionStages": { [_: `${number}`]: Stage | null };
    "CompetitionPlayers": { [_: `${number}`]: Player | null };
    "CompetitionCurrentStageID": number;

    /** Creates a new Competition instance. */
    constructor($$source: Partial<Competition> = {}) {
        if (!("CompetitionID" in $$source)) {
            this["CompetitionID"] = 0;
        }
        if (!("CompetitionName" in $$source)) {
            this["CompetitionName"] = "";
        }
        if (!("CompetitionCategory" in $$source)) {
            this["CompetitionCategory"] = (0 as Category);
        }
        if (!("CompetitionWeapon" in $$source)) {
            this["CompetitionWeapon"] = (0 as Weapon);
        }
        if (!("CompetitionState" in $$source)) {
            this["CompetitionState"] = (0 as State);
        }
        if (!("CompetitionMaxStageNumber" in $$source)) {
            this["CompetitionMaxStageNumber"] = 0;
        }
        if (!("CompetitionStages" in $$source)) {
            this["CompetitionStages"] = {};
        }
        if (!("CompetitionPlayers" in $$source)) {
            this["CompetitionPlayers"] = {};
        }
        if (!("CompetitionCurrentStageID" in $$source)) {
            this["CompetitionCurrentStageID"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Competition instance from a string or object.
     */
    static createFrom($$source: any = {}): Competition {
        const $$createField6_0 = $$createType0;
        const $$createField7_0 = $$createType3;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("CompetitionStages" in $$parsedSource) {
            $$parsedSource["CompetitionStages"] = $$createField6_0($$parsedSource["CompetitionStages"]);
        }
        if ("CompetitionPlayers" in $$parsedSource) {
            $$parsedSource["CompetitionPlayers"] = $$createField7_0($$parsedSource["CompetitionPlayers"]);
        }
        return new Competition($$parsedSource as Partial<Competition>);
    }
}

export class Nation {
    "nation_id": number;
    "nation_name": string;
    "nation_code": string;

    /** Creates a new Nation instance. */
    constructor($$source: Partial<Nation> = {}) {
        if (!("nation_id" in $$source)) {
            this["nation_id"] = 0;
        }
        if (!("nation_name" in $$source)) {
            this["nation_name"] = "";
        }
        if (!("nation_code" in $$source)) {
            this["nation_code"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Nation instance from a string or object.
     */
    static createFrom($$source: any = {}): Nation {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new Nation($$parsedSource as Partial<Nation>);
    }
}

/**
 * Player : Person details
 * 
 * PlayerID is 16 bits, 4 are for the competitionID and 12 are for the playerID (4095 players max)
 */
export class Player {
    /**
     * More than 255 players
     */
    "PlayerID": number;
    "PlayerFirstname": string;
    "PlayerLastname": string;
    "PlayerNation": Nation | null;
    "PlayerRegion": Region | null;
    "PlayerClub": Club | null;
    "PlayerInitialRank": number;

    /** Creates a new Player instance. */
    constructor($$source: Partial<Player> = {}) {
        if (!("PlayerID" in $$source)) {
            this["PlayerID"] = 0;
        }
        if (!("PlayerFirstname" in $$source)) {
            this["PlayerFirstname"] = "";
        }
        if (!("PlayerLastname" in $$source)) {
            this["PlayerLastname"] = "";
        }
        if (!("PlayerNation" in $$source)) {
            this["PlayerNation"] = null;
        }
        if (!("PlayerRegion" in $$source)) {
            this["PlayerRegion"] = null;
        }
        if (!("PlayerClub" in $$source)) {
            this["PlayerClub"] = null;
        }
        if (!("PlayerInitialRank" in $$source)) {
            this["PlayerInitialRank"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Player instance from a string or object.
     */
    static createFrom($$source: any = {}): Player {
        const $$createField3_0 = $$createType5;
        const $$createField4_0 = $$createType7;
        const $$createField5_0 = $$createType9;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("PlayerNation" in $$parsedSource) {
            $$parsedSource["PlayerNation"] = $$createField3_0($$parsedSource["PlayerNation"]);
        }
        if ("PlayerRegion" in $$parsedSource) {
            $$parsedSource["PlayerRegion"] = $$createField4_0($$parsedSource["PlayerRegion"]);
        }
        if ("PlayerClub" in $$parsedSource) {
            $$parsedSource["PlayerClub"] = $$createField5_0($$parsedSource["PlayerClub"]);
        }
        return new Player($$parsedSource as Partial<Player>);
    }
}

export class Region {
    "region_id": number;
    "region_name": string;

    /** Creates a new Region instance. */
    constructor($$source: Partial<Region> = {}) {
        if (!("region_id" in $$source)) {
            this["region_id"] = 0;
        }
        if (!("region_name" in $$source)) {
            this["region_name"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Region instance from a string or object.
     */
    static createFrom($$source: any = {}): Region {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new Region($$parsedSource as Partial<Region>);
    }
}

/**
 * Stage : Stage details
 */
export type Stage = any;

/**
 * Stage Kind
 */
export enum StageKind {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero = 0,

    /**
     * 0
     */
    REGISTRATIONS = 0,

    /**
     * 1
     */
    POOLS = 1,

    /**
     * 2
     */
    POOLSRESULTS = 2,

    /**
     * 3
     */
    DIRECTELIMINATION = 3,

    /**
     * 4
     */
    FINALRANKING = 4,

    /**
     * 5
     */
    UNKNOWN = 5,
};

/**
 * State : Competition state
 */
export enum State {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero = 0,

    /**
     * 0
     */
    REGISTERING = 0,

    /**
     * 1
     */
    STARTED = 1,

    /**
     * 2
     */
    FINISHED = 2,

    /**
     * 3
     */
    LOCKED = 3,

    /**
     * 4
     */
    IDLE = 4,
};

/**
 * Weapon : Weapon used in competition
 */
export enum Weapon {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero = 0,

    /**
     * 0
     */
    FOIL = 0,

    /**
     * 1
     */
    EPEE = 1,

    /**
     * 2
     */
    SABRE = 2,
};

// Private type creation functions
const $$createType0 = $Create.Map($Create.Any, $Create.Any);
const $$createType1 = Player.createFrom;
const $$createType2 = $Create.Nullable($$createType1);
const $$createType3 = $Create.Map($Create.Any, $$createType2);
const $$createType4 = Nation.createFrom;
const $$createType5 = $Create.Nullable($$createType4);
const $$createType6 = Region.createFrom;
const $$createType7 = $Create.Nullable($$createType6);
const $$createType8 = Club.createFrom;
const $$createType9 = $Create.Nullable($$createType8);
