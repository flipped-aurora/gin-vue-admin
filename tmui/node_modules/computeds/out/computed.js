"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computed = void 0;
const tracker_1 = require("./tracker");
const system_1 = require("./system");
const dep_1 = require("./dep");
function computed(getter) {
    let oldValue;
    const tracker = new tracker_1.Tracker(() => (0, system_1.trigger)(dep, 1 /* DirtyLevels.ComputedValueMaybeDirty */));
    const fn = () => {
        (0, system_1.track)(dep);
        if (tracker.dirty
            && !Object.is(oldValue, oldValue = tracker.track(() => getter(oldValue)))) {
            (0, system_1.trigger)(dep, 2 /* DirtyLevels.ComputedValueDirty */);
        }
        return oldValue;
    };
    const dep = new dep_1.Dep(fn);
    return fn;
}
exports.computed = computed;
