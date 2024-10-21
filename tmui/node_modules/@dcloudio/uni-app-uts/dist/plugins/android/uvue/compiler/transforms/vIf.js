"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.processIf = exports.transformIf = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const transform_1 = require("../transform");
const errors_1 = require("../errors");
const transformExpression_1 = require("./transformExpression");
exports.transformIf = (0, transform_1.createStructuralDirectiveTransform)(/^(if|else|else-if)$/, (node, dir, context) => {
    return processIf(node, dir, context, (ifNode, branch, isRoot) => {
        // #1587: We need to dynamically increment the key based on the current
        // node's sibling nodes, since chained v-if/else branches are
        // rendered at the same depth
        const siblings = context.parent.children;
        let i = siblings.indexOf(ifNode);
        let key = 0;
        while (i-- >= 0) {
            const sibling = siblings[i];
            if (sibling && sibling.type === compiler_core_1.NodeTypes.IF) {
                key += sibling.branches.length;
            }
        }
        // Exit callback. Complete the codegenNode when all children have been
        // transformed.
        return () => {
            if (isRoot) {
                ifNode.codegenNode = createCodegenNodeForBranch(branch, key, context);
            }
            else {
                // attach this branch's codegen node to the v-if root.
                const parentCondition = getParentCondition(ifNode.codegenNode);
                parentCondition.alternate = createCodegenNodeForBranch(branch, key + ifNode.branches.length - 1, context);
            }
        };
    });
});
// target-agnostic transform used for both Client and SSR
function processIf(node, dir, context, processCodegen) {
    if (dir.name !== 'else' &&
        (!dir.exp || !dir.exp.content.trim())) {
        const loc = dir.exp ? dir.exp.loc : node.loc;
        context.onError((0, errors_1.createCompilerError)(28 /* ErrorCodes.X_V_IF_NO_EXPRESSION */, dir.loc));
        dir.exp = (0, compiler_core_1.createSimpleExpression)(`true`, false, loc);
    }
    if (context.prefixIdentifiers && dir.exp) {
        // dir.exp can only be simple expression because vIf transform is applied
        // before expression transform.
        dir.exp = (0, transformExpression_1.processExpression)(dir.exp, context);
    }
    if (dir.name === 'if') {
        const branch = createIfBranch(node, dir);
        const ifNode = {
            type: compiler_core_1.NodeTypes.IF,
            loc: node.loc,
            branches: [branch],
        };
        context.replaceNode(ifNode);
        if (processCodegen) {
            return processCodegen(ifNode, branch, true);
        }
    }
    else {
        // locate the adjacent v-if
        const siblings = context.parent.children;
        let i = siblings.indexOf(node);
        while (i-- >= -1) {
            const sibling = siblings[i];
            if (sibling && sibling.type === compiler_core_1.NodeTypes.COMMENT) {
                context.removeNode(sibling);
                continue;
            }
            if (sibling &&
                ((sibling.type === compiler_core_1.NodeTypes.TEXT && !sibling.content.trim().length) ||
                    // handle tag text but type = NodeTypes.ELEMENT
                    (sibling.type === compiler_core_1.NodeTypes.ELEMENT &&
                        sibling.tag === 'text' &&
                        !sibling.children[0].content.trim().length))) {
                context.removeNode(sibling);
                continue;
            }
            if (sibling && sibling.type === compiler_core_1.NodeTypes.IF) {
                // Check if v-else was followed by v-else-if
                if (dir.name === 'else-if' &&
                    sibling.branches[sibling.branches.length - 1].condition === undefined) {
                    context.onError((0, errors_1.createCompilerError)(30 /* ErrorCodes.X_V_ELSE_NO_ADJACENT_IF */, node.loc));
                }
                // move the node to the if node's branches
                context.removeNode();
                const branch = createIfBranch(node, dir);
                // check if user is forcing same key on different branches
                const key = branch.userKey;
                if (key) {
                    sibling.branches.forEach(({ userKey }) => {
                        if (isSameKey(userKey, key)) {
                            context.onError((0, errors_1.createCompilerError)(29 /* ErrorCodes.X_V_IF_SAME_KEY */, branch.userKey.loc));
                        }
                    });
                }
                sibling.branches.push(branch);
                const onExit = processCodegen && processCodegen(sibling, branch, false);
                // since the branch was removed, it will not be traversed.
                // make sure to traverse here.
                (0, transform_1.traverseNode)(branch, context);
                // call on exit
                if (onExit)
                    onExit();
                // make sure to reset currentNode after traversal to indicate this
                // node has been removed.
                context.currentNode = null;
            }
            else {
                context.onError((0, errors_1.createCompilerError)(30 /* ErrorCodes.X_V_ELSE_NO_ADJACENT_IF */, node.loc));
            }
            break;
        }
    }
}
exports.processIf = processIf;
function createIfBranch(node, dir) {
    const isTemplateIf = node.tagType === compiler_core_1.ElementTypes.TEMPLATE;
    return {
        type: compiler_core_1.NodeTypes.IF_BRANCH,
        loc: node.loc,
        condition: dir.name === 'else' ? undefined : dir.exp,
        children: isTemplateIf && !(0, compiler_core_1.findDir)(node, 'for') ? node.children : [node],
        userKey: (0, compiler_core_1.findProp)(node, `key`),
        isTemplateIf,
    };
}
function createCodegenNodeForBranch(branch, keyIndex, context) {
    if (branch.condition) {
        return (0, compiler_core_1.createConditionalExpression)(branch.condition, createChildrenCodegenNode(branch, keyIndex, context), 
        // make sure to pass in asBlock: true so that the comment node call
        // closes the current block.
        (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_1.CREATE_COMMENT), ['"v-if"', 'true']));
    }
    else {
        return createChildrenCodegenNode(branch, keyIndex, context);
    }
}
function createChildrenCodegenNode(branch, keyIndex, context) {
    const { helper } = context;
    const keyProperty = (0, compiler_core_1.createObjectProperty)(`key`, (0, compiler_core_1.createSimpleExpression)(`${keyIndex}`, false, compiler_core_1.locStub, compiler_core_1.ConstantTypes.CAN_HOIST));
    const { children } = branch;
    const firstChild = children[0];
    const needFragmentWrapper = children.length !== 1 || firstChild.type !== compiler_core_1.NodeTypes.ELEMENT;
    if (needFragmentWrapper) {
        if (children.length === 1 && firstChild.type === compiler_core_1.NodeTypes.FOR) {
            // optimize away nested fragments when child is a ForNode
            const vnodeCall = firstChild.codegenNode;
            (0, compiler_core_1.injectProp)(vnodeCall, keyProperty, context);
            return vnodeCall;
        }
        else {
            let patchFlag = shared_1.PatchFlags.STABLE_FRAGMENT;
            let patchFlagText = shared_1.PatchFlagNames[shared_1.PatchFlags.STABLE_FRAGMENT];
            // check if the fragment actually contains a single valid child with
            // the rest being comments
            return (0, compiler_core_1.createVNodeCall)(context, helper(compiler_core_1.FRAGMENT), (0, compiler_core_1.createObjectExpression)([keyProperty]), children, patchFlag + ` /* ${patchFlagText} */`, undefined, undefined, true, false, false /* isComponent */, branch.loc);
        }
    }
    else {
        const ret = firstChild.codegenNode;
        const vnodeCall = (0, compiler_core_1.getMemoedVNodeCall)(ret);
        // Change createVNode to createBlock.
        if (vnodeCall.type === compiler_core_1.NodeTypes.VNODE_CALL) {
            (0, compiler_core_1.convertToBlock)(vnodeCall, context);
        }
        // inject branch key
        (0, compiler_core_1.injectProp)(vnodeCall, keyProperty, context);
        return ret;
    }
}
function isSameKey(a, b) {
    if (!a || a.type !== b.type) {
        return false;
    }
    if (a.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
        if (a.value.content !== b.value.content) {
            return false;
        }
    }
    else {
        // directive
        const exp = a.exp;
        const branchExp = b.exp;
        if (exp.type !== branchExp.type) {
            return false;
        }
        if (exp.type !== compiler_core_1.NodeTypes.SIMPLE_EXPRESSION ||
            exp.isStatic !== branchExp.isStatic ||
            exp.content !== branchExp.content) {
            return false;
        }
    }
    return true;
}
function getParentCondition(node) {
    while (true) {
        if (node.type === compiler_core_1.NodeTypes.JS_CONDITIONAL_EXPRESSION) {
            if (node.alternate.type === compiler_core_1.NodeTypes.JS_CONDITIONAL_EXPRESSION) {
                node = node.alternate;
            }
            else {
                return node;
            }
        }
        else if (node.type === compiler_core_1.NodeTypes.JS_CACHE_EXPRESSION) {
            node = node.value;
        }
    }
}
