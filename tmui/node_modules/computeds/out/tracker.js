"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Tracker = void 0;
const system_1 = require("./system");
class Tracker {
    constructor(spread, effect) {
        this.spread = spread;
        this.effect = effect;
        this.dirtyLevel = 3 /* DirtyLevels.Dirty */;
        this.trackId = 0;
        this.runnings = 0;
        this.queryings = 0;
        this.depsLength = 0;
    }
    get dirty() {
        if (this.dirtyLevel === 1 /* DirtyLevels.ComputedValueMaybeDirty */) {
            this.dirtyLevel = 0 /* DirtyLevels.NotDirty */;
            if (this.trackToken) {
                const deps = system_1.depsMap.get(this.trackToken);
                if (deps) {
                    this.queryings++;
                    (0, system_1.pauseTracking)();
                    for (const dep of deps) {
                        if (dep.computed) {
                            dep.computed();
                            if (this.dirtyLevel >= 2 /* DirtyLevels.ComputedValueDirty */) {
                                break;
                            }
                        }
                    }
                    (0, system_1.resetTracking)();
                    this.queryings--;
                }
            }
        }
        return this.dirtyLevel >= 2 /* DirtyLevels.ComputedValueDirty */;
    }
    track(fn) {
        try {
            system_1.activeTrackers.push(this);
            this.runnings++;
            preCleanup(this);
            return fn();
        }
        finally {
            postCleanup(this);
            this.runnings--;
            system_1.activeTrackers.pop();
            if (!this.runnings) {
                this.dirtyLevel = 0 /* DirtyLevels.NotDirty */;
            }
        }
    }
    reset() {
        preCleanup(this);
        postCleanup(this);
        this.dirtyLevel = 3 /* DirtyLevels.Dirty */;
    }
    deref() {
        return this;
    }
}
exports.Tracker = Tracker;
function preCleanup(tracker) {
    tracker.trackId++;
    tracker.depsLength = 0;
}
function postCleanup(tracker) {
    if (tracker.trackToken) {
        const deps = system_1.depsMap.get(tracker.trackToken);
        if (deps && deps.length > tracker.depsLength) {
            for (let i = tracker.depsLength; i < deps.length; i++) {
                (0, system_1.cleanupDepEffect)(deps[i], tracker);
            }
            deps.length = tracker.depsLength;
        }
    }
}
