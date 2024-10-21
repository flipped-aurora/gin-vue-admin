import type { TrackToken } from './tracker';
export declare class Dep extends Map<TrackToken, number> {
    computed?: (() => void) | undefined;
    constructor(computed?: (() => void) | undefined);
}
