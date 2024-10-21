"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var cssFontSizeKeywords = require('css-font-size-keywords');
function isSize(value) {
    return /^[\d\.]/.test(value)
        || value.indexOf('/') !== -1
        || cssFontSizeKeywords.indexOf(value) !== -1;
}
exports.isSize = isSize;
