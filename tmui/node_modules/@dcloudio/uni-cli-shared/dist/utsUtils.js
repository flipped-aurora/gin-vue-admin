"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genUTSClassName = exports.matchUTSComponent = void 0;
const easycom_1 = require("./easycom");
const utils_1 = require("./utils");
function matchUTSComponent(tag) {
    const source = (0, easycom_1.matchEasycom)(tag);
    return !!(source && source.includes('uts-proxy'));
}
exports.matchUTSComponent = matchUTSComponent;
function genUTSClassName(fileName, prefix = 'Gen') {
    return (prefix +
        (0, utils_1.capitalize)((0, utils_1.camelize)(verifySymbol((0, utils_1.removeExt)((0, utils_1.normalizeNodeModules)(fileName)
            .replace(/[\/|_]/g, '-')
            .replace(/-+/g, '-'))))));
}
exports.genUTSClassName = genUTSClassName;
function isValidStart(c) {
    return !!c.match(/^[A-Za-z_-]$/);
}
function isValidContinue(c) {
    return !!c.match(/^[A-Za-z0-9_-]$/);
}
function verifySymbol(s) {
    const chars = Array.from(s);
    if (isValidStart(chars[0]) && chars.slice(1).every(isValidContinue)) {
        return s;
    }
    const buf = [];
    let hasStart = false;
    for (const c of chars) {
        if (!hasStart && isValidStart(c)) {
            hasStart = true;
            buf.push(c);
        }
        else if (isValidContinue(c)) {
            buf.push(c);
        }
    }
    if (buf.length === 0) {
        buf.push('_');
    }
    return buf.join('');
}
