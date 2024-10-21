"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniCssScopedPlugin = exports.uniRemoveCssScopedPlugin = exports.addScoped = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const constants_1 = require("../../constants");
const preprocess_1 = require("../../preprocess");
const parse_1 = require("../../vue/parse");
const utils_1 = require("../../utils");
const debugScoped = (0, debug_1.default)('uni:scoped');
const SCOPED_RE = /<style\s[^>]*scoped[^>]*>/i;
function addScoped(code) {
    return code.replace(/(<style\b[^><]*)>/gi, (str, $1) => {
        if ($1.includes('scoped')) {
            return str;
        }
        return `${$1} scoped>`;
    });
}
exports.addScoped = addScoped;
function removeScoped(code) {
    if (!SCOPED_RE.test(code)) {
        return code;
    }
    return code.replace(/(<style.*)scoped(.*>)/gi, '$1$2');
}
function uniRemoveCssScopedPlugin({ filter } = { filter: () => false }) {
    return {
        name: 'uni:css-remove-scoped',
        enforce: 'pre',
        transform(code, id) {
            if (!filter(id))
                return null;
            debugScoped(id);
            return {
                code: removeScoped(code),
                map: null,
            };
        },
    };
}
exports.uniRemoveCssScopedPlugin = uniRemoveCssScopedPlugin;
function uniCssScopedPlugin({ filter } = { filter: () => false }) {
    return {
        name: 'uni:css-scoped',
        enforce: 'pre',
        transform(code, id) {
            if (!filter(id))
                return null;
            debugScoped(id);
            return {
                code: addScoped(code),
                map: null,
            };
        },
        // 仅 h5
        handleHotUpdate(ctx) {
            if (!constants_1.EXTNAME_VUE.includes(path_1.default.extname(ctx.file))) {
                return;
            }
            const scoped = !(0, utils_1.isAppVue)(ctx.file);
            debugScoped('hmr', ctx.file);
            const oldRead = ctx.read;
            ctx.read = async () => {
                let code = await oldRead();
                // hotUpdate preprocess
                if (code.includes('#endif')) {
                    code = (0, preprocess_1.preJs)((0, preprocess_1.preHtml)(code));
                }
                if (scoped) {
                    code = addScoped(code);
                }
                // 处理 block, wxs 等
                return (0, parse_1.parseVueCode)(code).code;
            };
        },
    };
}
exports.uniCssScopedPlugin = uniCssScopedPlugin;
