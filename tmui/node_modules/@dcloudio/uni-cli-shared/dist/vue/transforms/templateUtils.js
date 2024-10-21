"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseUrl = exports.isDataUrl = exports.isExternalUrl = exports.isRelativeUrl = void 0;
const url_1 = require("url");
const shared_1 = require("@vue/shared");
function isRelativeUrl(url) {
    const firstChar = url.charAt(0);
    return firstChar === '.' || firstChar === '~' || firstChar === '@';
}
exports.isRelativeUrl = isRelativeUrl;
const externalRE = /^(https?:)?\/\//;
function isExternalUrl(url) {
    return externalRE.test(url);
}
exports.isExternalUrl = isExternalUrl;
const dataUrlRE = /^\s*data:/i;
function isDataUrl(url) {
    return dataUrlRE.test(url);
}
exports.isDataUrl = isDataUrl;
/**
 * Parses string url into URL object.
 */
function parseUrl(url) {
    const firstChar = url.charAt(0);
    if (firstChar === '~') {
        const secondChar = url.charAt(1);
        url = url.slice(secondChar === '/' ? 2 : 1);
    }
    return parseUriParts(url);
}
exports.parseUrl = parseUrl;
/**
 * vuejs/component-compiler-utils#22 Support uri fragment in transformed require
 * @param urlString an url as a string
 */
function parseUriParts(urlString) {
    // A TypeError is thrown if urlString is not a string
    // @see https://nodejs.org/api/url.html#url_url_parse_urlstring_parsequerystring_slashesdenotehost
    return (0, url_1.parse)((0, shared_1.isString)(urlString) ? urlString : '', false, true);
}
