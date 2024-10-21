"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteConsoleExpr = void 0;
const magic_string_1 = __importDefault(require("magic-string"));
const utils_1 = require("../utils");
function rewriteConsoleExpr(method, id, filename, code, sourceMap = false) {
    filename = (0, utils_1.normalizePath)(filename);
    const re = /(console\.(log|info|debug|warn|error))\(([^)]+)\)/g;
    const locate = getLocator(code);
    const s = new magic_string_1.default(code);
    let match;
    while ((match = re.exec(code))) {
        const [, expr, type] = match;
        s.overwrite(match.index, match.index + expr.length + 1, method + `('${type}','at ${filename}:${locate(match.index).line + 1}',`);
    }
    if (s.hasChanged()) {
        return {
            code: s.toString(),
            map: sourceMap ? s.generateMap({ hires: true }) : { mappings: '' },
        };
    }
    return { code, map: { mappings: '' } };
}
exports.rewriteConsoleExpr = rewriteConsoleExpr;
function getLocator(source) {
    const originalLines = source.split('\n');
    const lineOffsets = [];
    for (let i = 0, pos = 0; i < originalLines.length; i++) {
        lineOffsets.push(pos);
        pos += originalLines[i].length + 1;
    }
    return function locate(index) {
        let i = 0;
        let j = lineOffsets.length;
        while (i < j) {
            const m = (i + j) >> 1;
            if (index < lineOffsets[m]) {
                j = m;
            }
            else {
                i = m + 1;
            }
        }
        const line = i - 1;
        const column = index - lineOffsets[line];
        return { line, column };
    };
}
