"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteSelfKey = exports.isSelfKey = void 0;
const compiler_core_1 = require("@vue/compiler-core");
function isSelfKey({ arg, exp }, vFor) {
    return (vFor &&
        arg &&
        exp &&
        arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
        arg.content === 'key' &&
        exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
        exp.content === vFor.valueAlias);
}
exports.isSelfKey = isSelfKey;
function rewriteSelfKey(dir) {
    ;
    dir.exp.content = '*this';
}
exports.rewriteSelfKey = rewriteSelfKey;
