import { DirtyLevels } from './system';
export type TrackToken = WeakRef<Tracker> | Tracker;
export declare class Tracker {
    spread: () => void;
    effect?: (() => void) | undefined;
    trackToken?: TrackToken;
    dirtyLevel: DirtyLevels;
    trackId: number;
    runnings: number;
    queryings: number;
    depsLength: number;
    constructor(spread: () => void, effect?: (() => void) | undefined);
    get dirty(): boolean;
    track<T>(fn: () => T): T;
    reset(): void;
    deref(): this;
}
