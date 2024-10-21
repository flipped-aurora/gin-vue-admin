"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computedSet = void 0;
const computed_1 = require("../computed");
function computedSet(getter) {
    return (0, computed_1.computed)((oldValue) => {
        const newValue = getter();
        if (oldValue?.size === newValue.size && [...oldValue].every(c => newValue.has(c))) {
            return oldValue;
        }
        return newValue;
    });
}
exports.computedSet = computedSet;
