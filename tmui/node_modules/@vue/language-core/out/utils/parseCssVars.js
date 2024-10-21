"use strict";
// https://github.com/vuejs/core/blob/main/packages/compiler-sfc/src/cssVars.ts#L47-L61
Object.defineProperty(exports, "__esModule", { value: true });
exports.clearComments = exports.parseCssVars = void 0;
const vBindCssVarReg = /\bv-bind\(\s*(?:'([^']+)'|"([^"]+)"|([^'"][^)]*))\s*\)/g;
const commentReg1 = /\/\*([\s\S]*?)\*\//g;
const commentReg2 = /\/\/([\s\S]*?)\n/g;
function* parseCssVars(styleContent) {
    styleContent = clearComments(styleContent);
    const matchs = styleContent.matchAll(vBindCssVarReg);
    for (const match of matchs) {
        if (match.index !== undefined) {
            const matchText = match[1] ?? match[2] ?? match[3];
            if (matchText !== undefined) {
                const offset = match.index + styleContent.slice(match.index).indexOf(matchText);
                yield { offset, text: matchText };
            }
        }
    }
}
exports.parseCssVars = parseCssVars;
function clearComments(css) {
    return css
        .replace(commentReg1, match => `/*${' '.repeat(match.length - 4)}*/`)
        .replace(commentReg2, match => `//${' '.repeat(match.length - 3)}\n`);
}
exports.clearComments = clearComments;
//# sourceMappingURL=parseCssVars.js.map