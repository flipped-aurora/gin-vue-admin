"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.signal = void 0;
const system_1 = require("./system");
const dep_1 = require("./dep");
function signal(oldValue) {
    const dep = new dep_1.Dep();
    const fn = (() => {
        (0, system_1.track)(dep);
        return oldValue;
    });
    fn.markDirty = () => {
        (0, system_1.trigger)(dep, 3 /* DirtyLevels.Dirty */);
    };
    fn.set = (newValue) => {
        if (!Object.is(oldValue, oldValue = newValue)) {
            fn.markDirty();
        }
    };
    return fn;
}
exports.signal = signal;
