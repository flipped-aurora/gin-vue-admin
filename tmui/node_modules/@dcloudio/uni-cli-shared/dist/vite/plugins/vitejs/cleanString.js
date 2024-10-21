"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.emptyCssComments = void 0;
const utils_1 = require("./utils");
const blankReplacer = (s) => ' '.repeat(s.length);
function emptyCssComments(raw) {
    return raw.replace(utils_1.multilineCommentsRE, blankReplacer);
}
exports.emptyCssComments = emptyCssComments;
