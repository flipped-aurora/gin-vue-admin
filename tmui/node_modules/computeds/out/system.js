"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.trigger = exports.cleanupDepEffect = exports.track = exports.depsMap = exports.resetEffect = exports.pauseEffect = exports.resetTracking = exports.pauseTracking = exports.activeTrackers = void 0;
exports.activeTrackers = [];
let pauseEffectStack = 0;
const pausedTrackers = [];
const pausedEffects = [];
function pauseTracking() {
    pausedTrackers.push(exports.activeTrackers);
    exports.activeTrackers = [];
}
exports.pauseTracking = pauseTracking;
function resetTracking() {
    exports.activeTrackers = pausedTrackers.pop();
}
exports.resetTracking = resetTracking;
function pauseEffect() {
    pauseEffectStack++;
}
exports.pauseEffect = pauseEffect;
function resetEffect() {
    pauseEffectStack--;
    while (!pauseEffectStack && pausedEffects.length) {
        pausedEffects.shift().effect();
    }
}
exports.resetEffect = resetEffect;
exports.depsMap = new WeakMap();
const trackerRegistry = new FinalizationRegistry(trackToken => {
    const deps = exports.depsMap.get(trackToken);
    if (deps) {
        for (const dep of deps) {
            dep.delete(trackToken);
        }
        deps.length = 0;
    }
});
function track(dep) {
    if (exports.activeTrackers.length) {
        const tracker = exports.activeTrackers[exports.activeTrackers.length - 1];
        if (!tracker.trackToken) {
            if (tracker.effect) {
                tracker.trackToken = tracker;
            }
            else {
                tracker.trackToken = new WeakRef(tracker);
                trackerRegistry.register(tracker, tracker.trackToken, tracker);
            }
            exports.depsMap.set(tracker.trackToken, []);
        }
        const trackToken = tracker.trackToken;
        const deps = exports.depsMap.get(trackToken);
        if (deps) {
            if (dep.get(tracker) !== tracker.trackId) {
                dep.set(tracker, tracker.trackId);
                const oldDep = deps[tracker.depsLength];
                if (oldDep !== dep) {
                    if (oldDep) {
                        cleanupDepEffect(oldDep, tracker);
                    }
                    deps[tracker.depsLength++] = dep;
                }
                else {
                    tracker.depsLength++;
                }
            }
        }
    }
}
exports.track = track;
function cleanupDepEffect(dep, tracker) {
    const trackId = dep.get(tracker);
    if (trackId !== undefined && tracker.trackId !== trackId) {
        dep.delete(tracker);
    }
}
exports.cleanupDepEffect = cleanupDepEffect;
function trigger(dep, dirtyLevel) {
    pauseEffect();
    for (const trackToken of dep.keys()) {
        const tracker = trackToken.deref();
        if (!tracker) {
            continue;
        }
        if (tracker.dirtyLevel < dirtyLevel &&
            (!tracker.runnings || dirtyLevel !== 2 /* DirtyLevels.ComputedValueDirty */)) {
            const lastDirtyLevel = tracker.dirtyLevel;
            tracker.dirtyLevel = dirtyLevel;
            if (lastDirtyLevel === 0 /* DirtyLevels.NotDirty */ &&
                (!tracker.queryings || dirtyLevel !== 2 /* DirtyLevels.ComputedValueDirty */)) {
                tracker.spread();
                if (tracker.effect) {
                    pausedEffects.push(tracker);
                }
            }
        }
    }
    resetEffect();
}
exports.trigger = trigger;
