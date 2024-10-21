"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseProgram = void 0;
const parser_1 = require("@babel/parser");
const utils_1 = require("../utils");
function parseProgram(code, importer, { babelParserPlugins }) {
    return (0, parser_1.parse)(code, {
        plugins: (0, utils_1.normalizeParsePlugins)(importer, babelParserPlugins),
        sourceType: 'module',
    }).program;
}
exports.parseProgram = parseProgram;
