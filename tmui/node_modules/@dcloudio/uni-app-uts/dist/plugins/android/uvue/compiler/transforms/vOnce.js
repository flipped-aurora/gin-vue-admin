"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformOnce = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const seen = new WeakSet();
const transformOnce = (node, context) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT && (0, compiler_core_1.findDir)(node, 'once', true)) {
        if (seen.has(node) || context.inVOnce) {
            return;
        }
        seen.add(node);
        context.inVOnce = true;
        context.helper(compiler_core_1.SET_BLOCK_TRACKING);
        return () => {
            context.inVOnce = false;
            const cur = context.currentNode;
            if (cur.codegenNode) {
                cur.codegenNode = context.cache(cur.codegenNode, true /* isVNode */);
            }
        };
    }
};
exports.transformOnce = transformOnce;
