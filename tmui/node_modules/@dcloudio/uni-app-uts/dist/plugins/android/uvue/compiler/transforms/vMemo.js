"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createFunctionExpression = exports.transformMemo = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const seen = new WeakSet();
const transformMemo = (node, context) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
        const dir = (0, compiler_core_1.findDir)(node, 'memo');
        if (!dir || seen.has(node)) {
            return;
        }
        seen.add(node);
        return () => {
            const codegenNode = node.codegenNode ||
                context.currentNode.codegenNode;
            if (codegenNode && codegenNode.type === compiler_core_1.NodeTypes.VNODE_CALL) {
                // non-component sub tree should be turned into a block
                if (node.tagType !== compiler_core_1.ElementTypes.COMPONENT) {
                    (0, compiler_core_1.convertToBlock)(codegenNode, context);
                }
                const fn = createFunctionExpression(undefined, codegenNode);
                fn.returnType = 'VNode';
                node.codegenNode = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_1.WITH_MEMO), [
                    dir.exp,
                    fn,
                    `_cache`,
                    String(context.cached++),
                ]);
            }
        };
    }
};
exports.transformMemo = transformMemo;
function createFunctionExpression(params, returns = undefined, newline = false, isSlot = false, loc = compiler_core_1.locStub) {
    return {
        type: compiler_core_1.NodeTypes.JS_FUNCTION_EXPRESSION,
        params,
        returns,
        newline,
        isSlot,
        loc,
    };
}
exports.createFunctionExpression = createFunctionExpression;
