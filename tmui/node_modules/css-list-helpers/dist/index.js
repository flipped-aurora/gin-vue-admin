"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/**
 * Splits a CSS declaration value (shorthand) using provided separators
 * as the delimiters.
 */
function split(
/**
 * A CSS declaration value (shorthand).
 */
value, 
/**
 * Any number of separator characters used for splitting.
 */
separators, _a) {
    var _b = (_a === void 0 ? {} : _a).last, last = _b === void 0 ? false : _b;
    if (typeof value !== 'string') {
        throw new TypeError('expected a string');
    }
    if (!Array.isArray(separators)) {
        throw new TypeError('expected a string array of separators');
    }
    if (typeof last !== 'boolean') {
        throw new TypeError('expected a Boolean value for options.last');
    }
    var array = [];
    var current = '';
    var splitMe = false;
    var func = 0;
    var quote = false;
    var escape = false;
    for (var _i = 0, value_1 = value; _i < value_1.length; _i++) {
        var char = value_1[_i];
        if (quote) {
            if (escape) {
                escape = false;
            }
            else if (char === '\\') {
                escape = true;
            }
            else if (char === quote) {
                quote = false;
            }
        }
        else if (char === '"' || char === '\'') {
            quote = char;
        }
        else if (char === '(') {
            func += 1;
        }
        else if (char === ')') {
            if (func > 0) {
                func -= 1;
            }
        }
        else if (func === 0) {
            if (separators.indexOf(char) !== -1) {
                splitMe = true;
            }
        }
        if (splitMe) {
            if (current !== '') {
                array.push(current.trim());
            }
            current = '';
            splitMe = false;
        }
        else {
            current += char;
        }
    }
    if (last || current !== '') {
        array.push(current.trim());
    }
    return array;
}
exports.split = split;
/**
 * Splits a CSS declaration value (shorthand) using whitespace characters
 * as the delimiters.
 */
function splitBySpaces(
/**
 * A CSS declaration value (shorthand).
 */
value) {
    var spaces = [' ', '\n', '\t'];
    return split(value, spaces);
}
exports.splitBySpaces = splitBySpaces;
/**
 * Splits a CSS declaration value (shorthand) using commas as the delimiters.
 */
function splitByCommas(
/**
 * A CSS declaration value (shorthand).
 */
value) {
    var comma = ',';
    return split(value, [comma], { last: true });
}
exports.splitByCommas = splitByCommas;
