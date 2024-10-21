"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.findVueComponentImports = void 0;
const path_1 = __importDefault(require("path"));
const es_module_lexer_1 = require("es-module-lexer");
const shared_1 = require("@vue/shared");
const types_1 = require("@babel/types");
const parser_1 = require("@babel/parser");
const constants_1 = require("../constants");
const utils_1 = require("../utils");
/**
 * 暂时没用
 * @param source
 * @param importer
 * @param resolve
 * @returns
 */
async function findVueComponentImports(source, importer, resolve) {
    await es_module_lexer_1.init;
    let imports = [];
    // strip UTF-8 BOM
    if (source.charCodeAt(0) === 0xfeff) {
        source = source.slice(1);
    }
    try {
        imports = (0, es_module_lexer_1.parse)(source)[0];
    }
    catch (e) {
        console.error(e);
    }
    if (!imports.length) {
        return [];
    }
    const rewriteImports = [];
    for (let i = 0; i < imports.length; i++) {
        const importSpecifier = imports[i];
        const { n } = importSpecifier;
        if (!n) {
            continue;
        }
        const extname = path_1.default.extname(n);
        // 仅处理没有后缀，或后缀是.vue,.nvue的文件
        if (extname && !constants_1.EXTNAME_VUE.includes(extname)) {
            continue;
        }
        const res = await resolve(n, importer);
        if (!res) {
            continue;
        }
        if (constants_1.EXTNAME_VUE_RE.test(res.id)) {
            const expr = (0, parser_1.parse)(source.slice(importSpecifier.ss, importSpecifier.se), {
                plugins: (0, utils_1.normalizeParsePlugins)(res.id),
                sourceType: 'module',
            }).program.body[0];
            if ((0, types_1.isImportDeclaration)(expr) && expr.specifiers.length === 1) {
                const importDefaultSpecifier = expr.specifiers[0];
                if (!(0, types_1.isImportDefaultSpecifier)(importDefaultSpecifier)) {
                    continue;
                }
                rewriteImports.push((0, shared_1.extend)(importSpecifier, {
                    n: res.id,
                    i: importDefaultSpecifier.local.name,
                }));
            }
        }
    }
    return rewriteImports;
}
exports.findVueComponentImports = findVueComponentImports;
