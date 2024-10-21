"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.effect = void 0;
const tracker_1 = require("./tracker");
function effect(fn) {
    const tracker = new tracker_1.Tracker(() => { }, () => {
        if (tracker.dirty) {
            tracker.track(fn);
        }
    });
    tracker.track(fn);
    return tracker;
}
exports.effect = effect;
