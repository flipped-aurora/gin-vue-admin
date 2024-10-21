"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteConsole = exports.hasConsole = void 0;
const estree_walker_1 = require("estree-walker");
function hasConsole(content) {
    return content.includes('console.');
}
exports.hasConsole = hasConsole;
function rewriteConsole(ast, s, { fileName, startLine, startOffset, }) {
    (0, estree_walker_1.walk)(ast, {
        enter(node, _parent) {
            if (node.type === 'CallExpression' && node.loc) {
                const callee = node.callee;
                if (callee.type === 'MemberExpression') {
                    // console.log()
                    // UTSAndroid.consoleDebugError()
                    if (callee.object.type === 'Identifier' &&
                        callee.property.type === 'Identifier' &&
                        (callee.object.name === 'console' ||
                            callee.property.name === 'consoleDebugError')) {
                        const at = `${node.arguments.length ? ', " ' : '"'}at ${fileName}:${node.loc.start.line + startLine}"`;
                        s.appendRight(startOffset + node.end - 1, at);
                    }
                }
            }
        },
    });
}
exports.rewriteConsole = rewriteConsole;
