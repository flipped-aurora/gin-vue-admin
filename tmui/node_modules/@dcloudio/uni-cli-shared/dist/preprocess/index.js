"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.preUVueJson = exports.preUVueCss = exports.preUVueHtml = exports.preUVueJs = exports.preNVueJson = exports.preNVueCss = exports.preNVueHtml = exports.preNVueJs = exports.preJson = exports.preCss = exports.preHtml = exports.preJs = exports.initPreContext = void 0;
const context_1 = require("./context");
/* eslint-disable no-restricted-globals */
const { preprocess } = require('../../lib/preprocess');
var context_2 = require("./context");
Object.defineProperty(exports, "initPreContext", { enumerable: true, get: function () { return context_2.initPreContext; } });
function preJs(jsCode) {
    if (process.env.UNI_APP_X === 'true') {
        return preUVueJs(jsCode);
    }
    return preprocess(jsCode, (0, context_1.getPreVueContext)(), { type: 'js' });
}
exports.preJs = preJs;
function preHtml(htmlCode) {
    if (process.env.UNI_APP_X === 'true') {
        return preUVueHtml(htmlCode);
    }
    return preprocess(htmlCode, (0, context_1.getPreVueContext)(), { type: 'html' });
}
exports.preHtml = preHtml;
exports.preCss = preJs;
exports.preJson = preJs;
function preNVueJs(jsCode) {
    return preprocess(jsCode, (0, context_1.getPreNVueContext)(), { type: 'js' });
}
exports.preNVueJs = preNVueJs;
function preNVueHtml(htmlCode) {
    return preprocess(htmlCode, (0, context_1.getPreNVueContext)(), { type: 'html' });
}
exports.preNVueHtml = preNVueHtml;
exports.preNVueCss = preNVueJs;
exports.preNVueJson = preNVueJs;
function preUVueJs(jsCode) {
    return preprocess(jsCode, (0, context_1.getPreUVueContext)(), { type: 'js' });
}
exports.preUVueJs = preUVueJs;
function preUVueHtml(htmlCode) {
    return preprocess(htmlCode, (0, context_1.getPreUVueContext)(), { type: 'html' });
}
exports.preUVueHtml = preUVueHtml;
exports.preUVueCss = preUVueJs;
exports.preUVueJson = preUVueJs;
