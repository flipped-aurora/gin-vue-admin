"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseFilterNames = exports.missingModuleName = exports.parseRenderjs = exports.isRenderjs = exports.isSjs = exports.isWxs = void 0;
const url_1 = require("./vite/utils/url");
const WXS_RE = /vue&type=wxs/;
function isWxs(id) {
    return WXS_RE.test(id);
}
exports.isWxs = isWxs;
const SJS_RE = /vue&type=sjs/;
function isSjs(id) {
    return SJS_RE.test(id);
}
exports.isSjs = isSjs;
const RENDERJS_RE = /vue&type=renderjs/;
function isRenderjs(id) {
    return RENDERJS_RE.test(id);
}
exports.isRenderjs = isRenderjs;
function parseRenderjs(id) {
    if (isWxs(id) || isRenderjs(id) || isSjs(id)) {
        const { query, filename } = (0, url_1.parseVueRequest)(id);
        return {
            type: query.type,
            name: query.name,
            filename,
        };
    }
    return {
        type: '',
        name: '',
        filename: '',
    };
}
exports.parseRenderjs = parseRenderjs;
function missingModuleName(type, code) {
    return `<script module="missing module name" lang="${type}">
${code}
</script>`;
}
exports.missingModuleName = missingModuleName;
const moduleRE = /module=["'](.*?)["']/;
function parseFilterNames(lang, code) {
    const names = [];
    const scriptTags = code.match(/<script\b[^>]*>/gm);
    if (!scriptTags) {
        return names;
    }
    const langRE = new RegExp(`lang=["']${lang}["']`);
    scriptTags.forEach((scriptTag) => {
        if (langRE.test(scriptTag)) {
            const matches = scriptTag.match(moduleRE);
            if (matches) {
                names.push(matches[1]);
            }
        }
    });
    return names;
}
exports.parseFilterNames = parseFilterNames;
