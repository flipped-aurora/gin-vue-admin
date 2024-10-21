"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteDebugError = exports.hasDebugError = void 0;
const estree_walker_1 = require("estree-walker");
const DEBUG_ERROR_RE = /(JSON\.parse)|(codeURI)/;
function hasDebugError(content) {
    return DEBUG_ERROR_RE.test(content);
}
exports.hasDebugError = hasDebugError;
function rewriteDebugError(ast, s, options) {
    (0, estree_walker_1.walk)(ast, {
        enter(node, _parent) {
            if (node.type === 'CallExpression' && node.loc) {
                const callee = node.callee;
                if (callee.type === 'MemberExpression') {
                    // JSON.parse || JSON.parseObject || JSON.parseArray
                    //
                    if (callee.object.type === 'Identifier' &&
                        callee.property.type === 'Identifier' &&
                        callee.object.name === 'JSON') {
                        const name = callee.property.name;
                        if (name === 'parse' ||
                            name === 'parseObject' ||
                            name === 'parseArray') {
                            wrapConsoleDebugError(s, node, options);
                        }
                    }
                }
                else if (callee.type === 'Identifier' &&
                    METHODS.includes(callee.name)) {
                    wrapConsoleDebugError(s, node, options);
                }
            }
        },
    });
}
exports.rewriteDebugError = rewriteDebugError;
function wrapConsoleDebugError(s, node, { fileName, startLine, startOffset }) {
    s.appendLeft(startOffset + node.start, `UTSAndroid.consoleDebugError(`);
    s.appendRight(startOffset + node.end, `)`);
    const at = `, " at ${fileName}:${node.loc.start.line + startLine}"`;
    s.appendLeft(startOffset + node.end, at);
}
const METHODS = [
    'decodeURI',
    'decodeURIComponent',
    'encodeURI',
    'encodeURIComponent',
];
