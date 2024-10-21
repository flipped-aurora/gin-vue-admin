import type { Dep } from './dep';
import type { Tracker, TrackToken } from './tracker';
export declare const enum DirtyLevels {
    NotDirty = 0,
    ComputedValueMaybeDirty = 1,
    ComputedValueDirty = 2,
    Dirty = 3
}
export declare let activeTrackers: Tracker[];
export declare function pauseTracking(): void;
export declare function resetTracking(): void;
export declare function pauseEffect(): void;
export declare function resetEffect(): void;
export declare const depsMap: WeakMap<TrackToken, Dep[]>;
export declare function track(dep: Dep): void;
export declare function cleanupDepEffect(dep: Dep, tracker: Tracker): void;
export declare function trigger(dep: Dep, dirtyLevel: DirtyLevels): void;
