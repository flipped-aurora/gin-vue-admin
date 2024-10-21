"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computedArray = void 0;
const computed_1 = require("../computed");
function computedArray(arr, computedItem) {
    const length = (0, computed_1.computed)(() => arr().length);
    const keys = (0, computed_1.computed)(() => {
        const keys = [];
        for (let i = 0; i < length(); i++) {
            keys.push(String(i));
        }
        return keys;
    });
    const items = (0, computed_1.computed)((array) => {
        array ??= [];
        while (array.length < length()) {
            const index = array.length;
            const item = (0, computed_1.computed)(() => arr()[index]);
            array.push(computedItem(item, index));
        }
        if (array.length > length()) {
            array.length = length();
        }
        return array;
    });
    return new Proxy({}, {
        get(_, p, receiver) {
            if (p === 'length') {
                return length();
            }
            if (typeof p === 'string' && !isNaN(Number(p))) {
                return items()[Number(p)]?.();
            }
            return Reflect.get(items(), p, receiver);
        },
        has(_, p) {
            return Reflect.has(items(), p);
        },
        ownKeys() {
            return keys();
        },
    });
}
exports.computedArray = computedArray;
