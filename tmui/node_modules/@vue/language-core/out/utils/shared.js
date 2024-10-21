"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.hyphenateAttr = exports.hyphenateTag = exports.getSlotsPropertyName = void 0;
const shared_1 = require("@vue/shared");
function getSlotsPropertyName(vueVersion) {
    return vueVersion < 3 ? '$scopedSlots' : '$slots';
}
exports.getSlotsPropertyName = getSlotsPropertyName;
var shared_2 = require("@vue/shared");
Object.defineProperty(exports, "hyphenateTag", { enumerable: true, get: function () { return shared_2.hyphenate; } });
function hyphenateAttr(str) {
    let hyphencase = (0, shared_1.hyphenate)(str);
    // fix https://github.com/vuejs/core/issues/8811
    if (str.length && str[0] !== str[0].toLowerCase()) {
        hyphencase = '-' + hyphencase;
    }
    return hyphencase;
}
exports.hyphenateAttr = hyphenateAttr;
//# sourceMappingURL=shared.js.map