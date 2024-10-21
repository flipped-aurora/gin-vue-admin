"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.isJsFile = exports.cleanUrl = exports.hashRE = exports.queryRE = exports.isInternalRequest = exports.ENV_PUBLIC_PATH = exports.CLIENT_PUBLIC_PATH = exports.VALID_ID_PREFIX = exports.FS_PREFIX = exports.isImportRequest = exports.parseVueRequest = void 0;
const shared_1 = require("@vue/shared");
const path_1 = __importDefault(require("path"));
const constants_1 = require("../../constants");
function parseVueRequest(id) {
    const [filename, rawQuery] = id.split(`?`, 2);
    const query = Object.fromEntries(new URLSearchParams(rawQuery));
    if (query.vue != null) {
        query.vue = true;
    }
    if (query.src != null) {
        query.src = true;
    }
    if (query.index != null) {
        query.index = Number(query.index);
    }
    if (query.raw != null) {
        query.raw = true;
    }
    return {
        filename,
        query,
    };
}
exports.parseVueRequest = parseVueRequest;
const importQueryRE = /(\?|&)import=?(?:&|$)/;
const isImportRequest = (url) => importQueryRE.test(url);
exports.isImportRequest = isImportRequest;
/**
 * Prefix for resolved fs paths, since windows paths may not be valid as URLs.
 */
exports.FS_PREFIX = `/@fs/`;
/**
 * Prefix for resolved Ids that are not valid browser import specifiers
 */
exports.VALID_ID_PREFIX = `/@id/`;
exports.CLIENT_PUBLIC_PATH = `/@vite/client`;
exports.ENV_PUBLIC_PATH = `/@vite/env`;
const internalPrefixes = [
    exports.FS_PREFIX,
    exports.VALID_ID_PREFIX,
    exports.CLIENT_PUBLIC_PATH,
    exports.ENV_PUBLIC_PATH,
];
const InternalPrefixRE = new RegExp(`^(?:${internalPrefixes.join('|')})`);
const isInternalRequest = (url) => InternalPrefixRE.test(url);
exports.isInternalRequest = isInternalRequest;
exports.queryRE = /\?.*$/;
exports.hashRE = /#.*$/;
const cleanUrl = (url) => url.replace(exports.hashRE, '').replace(exports.queryRE, '');
exports.cleanUrl = cleanUrl;
function isJsFile(id) {
    const { filename, query } = parseVueRequest(id);
    const isJs = constants_1.EXTNAME_JS_RE.test(filename);
    if (isJs) {
        return true;
    }
    const isVueJs = constants_1.EXTNAME_VUE.includes(path_1.default.extname(filename)) &&
        (!query.vue ||
            query.setup ||
            (0, shared_1.hasOwn)(query, 'lang.ts') ||
            (0, shared_1.hasOwn)(query, 'lang.js') ||
            (0, shared_1.hasOwn)(query, 'lang.uts'));
    if (isVueJs) {
        return true;
    }
    return false;
}
exports.isJsFile = isJsFile;
