"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var unquote = require('unquote');
var systemFontKeywords = require('css-system-font-keywords');
var fontWeightKeywords = require('css-font-weight-keywords');
var fontStyleKeywords = require('css-font-style-keywords');
var fontStretchKeywords = require('css-font-stretch-keywords');
var cssListHelpers = require("css-list-helpers");
var helpers = require("./helpers");
var errorPrefix = '[parse-css-font]';
var firstDeclarations = [
    'style',
    'weight',
    'stretch',
    'variant',
];
function parseCSSFont(value) {
    if (typeof value !== 'string') {
        throw error('Expected a string.', TypeError);
    }
    if (value === '') {
        throw error('Cannot parse an empty string.');
    }
    if (systemFontKeywords.indexOf(value) !== -1) {
        return { system: value };
    }
    var font = {
        lineHeight: 'normal',
        stretch: '',
        style: '',
        variant: '',
        weight: '',
    };
    var consumers = [style, weight, stretch, variant];
    var tokens = cssListHelpers.splitBySpaces(value);
    nextToken: for (var token = tokens.shift(); !!token; token = tokens.shift()) {
        if (token === 'normal') {
            continue;
        }
        for (var _i = 0, consumers_1 = consumers; _i < consumers_1.length; _i++) {
            var consume = consumers_1[_i];
            if (consume(token)) {
                continue nextToken;
            }
        }
        var parts = cssListHelpers.split(token, ['/']);
        font.size = parts[0];
        if (!!parts[1]) {
            font.lineHeight = parseLineHeight(parts[1]);
        }
        else if (tokens[0] === '/') {
            tokens.shift();
            font.lineHeight = parseLineHeight(tokens.shift());
        }
        if (!tokens.length) {
            throw error('Missing required font-family.');
        }
        font.family = cssListHelpers.splitByCommas(tokens.join(' ')).map(unquote);
        for (var _a = 0, firstDeclarations_1 = firstDeclarations; _a < firstDeclarations_1.length; _a++) {
            var name_1 = firstDeclarations_1[_a];
            font[name_1] = font[name_1] || 'normal';
        }
        return font;
    }
    throw error('Missing required font-size.');
    function style(token) {
        if (fontStyleKeywords.indexOf(token) === -1) {
            return;
        }
        if (font.style) {
            throw error('Font style already defined.');
        }
        return (font.style = token);
    }
    function weight(token) {
        if (fontWeightKeywords.indexOf(token) === -1) {
            return;
        }
        if (font.weight) {
            throw error('Font weight already defined.');
        }
        return (font.weight = token);
    }
    function stretch(token) {
        if (fontStretchKeywords.indexOf(token) === -1) {
            return;
        }
        if (font.stretch) {
            throw error('Font stretch already defined.');
        }
        return (font.stretch = token);
    }
    function variant(token) {
        return (!helpers.isSize(token) &&
            (font.variant = font.variant ? [font.variant, token].join(' ') : token));
    }
}
exports.default = parseCSSFont;
function error(message, ErrorType) {
    if (ErrorType === void 0) { ErrorType = Error; }
    return new ErrorType(errorPrefix + " " + message);
}
function parseLineHeight(value) {
    var parsed = parseFloat(value);
    if (parsed.toString() === value) {
        return parsed;
    }
    return value;
}
// @ts-ignore
module.exports = Object.assign(exports.default, exports);
