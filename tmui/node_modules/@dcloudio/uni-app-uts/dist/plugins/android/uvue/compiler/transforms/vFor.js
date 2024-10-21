"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createForLoopParams = exports.parseForExpression = exports.processFor = exports.transformFor = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const errors_1 = require("../errors");
const compiler_core_2 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const runtimeHelpers_1 = require("../runtimeHelpers");
const transformExpression_1 = require("./transformExpression");
// import { validateBrowserExpression } from '../validateExpression'
const shared_1 = require("@vue/shared");
const transform_1 = require("../transform");
const utils_1 = require("../utils");
exports.transformFor = (0, transform_1.createStructuralDirectiveTransform)('for', (node, dir, context) => {
    const { helper, removeHelper } = context;
    return processFor(node, dir, context, (forNode) => {
        // create the loop render function expression now, and add the
        // iterator on exit after all children have been traversed
        const renderExp = (0, compiler_core_1.createCallExpression)(helper(runtimeHelpers_1.RENDER_LIST), [
            forNode.source,
        ]);
        const isTemplate = (0, compiler_core_2.isTemplateNode)(node);
        const memo = (0, compiler_core_1.findDir)(node, 'memo');
        const keyProp = (0, compiler_core_2.findProp)(node, `key`);
        const keyExp = keyProp &&
            (keyProp.type === compiler_core_1.NodeTypes.ATTRIBUTE
                ? (0, compiler_core_1.createSimpleExpression)(keyProp.value.content, true)
                : keyProp.exp);
        const keyProperty = keyProp ? (0, compiler_core_1.createObjectProperty)(`key`, keyExp) : null;
        // if (!__BROWSER__ && isTemplate) {
        if (isTemplate) {
            // #2085 / #5288 process :key and v-memo expressions need to be
            // processed on `<template v-for>`. In this case the node is discarded
            // and never traversed so its binding expressions won't be processed
            // by the normal transforms.
            if (memo) {
                memo.exp = (0, transformExpression_1.processExpression)(memo.exp, context);
            }
            if (keyProperty && keyProp.type !== compiler_core_1.NodeTypes.ATTRIBUTE) {
                keyProperty.value = (0, transformExpression_1.processExpression)(keyProperty.value, context);
            }
        }
        const isStableFragment = forNode.source.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
            forNode.source.constType > compiler_core_1.ConstantTypes.NOT_CONSTANT;
        const fragmentFlag = isStableFragment
            ? shared_1.PatchFlags.STABLE_FRAGMENT
            : keyProp
                ? shared_1.PatchFlags.KEYED_FRAGMENT
                : shared_1.PatchFlags.UNKEYED_FRAGMENT;
        forNode.codegenNode = (0, compiler_core_1.createVNodeCall)(context, helper(runtimeHelpers_1.FRAGMENT), undefined, renderExp, fragmentFlag +
            // (__DEV__ ? ` /* ${PatchFlagNames[fragmentFlag]} */` : ``),
            ` /* ${shared_1.PatchFlagNames[fragmentFlag]} */`, undefined, undefined, true /* isBlock */, !isStableFragment /* disableTracking */, false /* isComponent */, node.loc);
        return () => {
            // finish the codegen now that all children have been traversed
            let childBlock;
            const { children } = forNode;
            // check <template v-for> key placement
            // if ((__DEV__ || !__BROWSER__) && isTemplate) {
            if (isTemplate) {
                node.children.some((c) => {
                    if (c.type === compiler_core_1.NodeTypes.ELEMENT) {
                        const key = (0, compiler_core_2.findProp)(c, 'key');
                        if (key) {
                            context.onError((0, errors_1.createCompilerError)(33 /* ErrorCodes.X_V_FOR_TEMPLATE_KEY_PLACEMENT */, key.loc));
                            return true;
                        }
                    }
                });
            }
            const needFragmentWrapper = children.length !== 1 || children[0].type !== compiler_core_1.NodeTypes.ELEMENT;
            const slotOutlet = (0, compiler_core_2.isSlotOutlet)(node)
                ? node
                : isTemplate &&
                    node.children.length === 1 &&
                    (0, compiler_core_2.isSlotOutlet)(node.children[0])
                    ? node.children[0] // api-extractor somehow fails to infer this
                    : null;
            if (slotOutlet) {
                // <slot v-for="..."> or <template v-for="..."><slot/></template>
                childBlock = slotOutlet.codegenNode;
                if (isTemplate && keyProperty) {
                    // <template v-for="..." :key="..."><slot/></template>
                    // we need to inject the key to the renderSlot() call.
                    // the props for renderSlot is passed as the 3rd argument.
                    (0, compiler_core_2.injectProp)(childBlock, keyProperty, context);
                }
            }
            else if (needFragmentWrapper) {
                // <template v-for="..."> with text or multi-elements
                // should generate a fragment block for each loop
                childBlock = (0, compiler_core_1.createVNodeCall)(context, helper(runtimeHelpers_1.FRAGMENT), keyProperty ? (0, compiler_core_1.createObjectExpression)([keyProperty]) : undefined, node.children, shared_1.PatchFlags.STABLE_FRAGMENT +
                    // (__DEV__
                    //   ? ` /* ${PatchFlagNames[PatchFlags.STABLE_FRAGMENT]} */`
                    //   : ``),
                    ` /* ${shared_1.PatchFlagNames[shared_1.PatchFlags.STABLE_FRAGMENT]} */`, undefined, undefined, true, undefined, false /* isComponent */);
            }
            else {
                // Normal element v-for. Directly use the child's codegenNode
                // but mark it as a block.
                childBlock = children[0]
                    .codegenNode;
                if (isTemplate && keyProperty) {
                    (0, compiler_core_2.injectProp)(childBlock, keyProperty, context);
                }
                if (childBlock.isBlock !== !isStableFragment) {
                    if (childBlock.isBlock) {
                        // switch from block to vnode
                        removeHelper(runtimeHelpers_1.OPEN_BLOCK);
                        removeHelper((0, compiler_core_2.getVNodeBlockHelper)(false, childBlock.isComponent));
                    }
                    else {
                        // switch from vnode to block
                        removeHelper((0, compiler_core_2.getVNodeHelper)(false, childBlock.isComponent));
                    }
                }
                childBlock.isBlock = !isStableFragment;
                if (childBlock.isBlock) {
                    helper(runtimeHelpers_1.OPEN_BLOCK);
                    helper((0, compiler_core_2.getVNodeBlockHelper)(false, childBlock.isComponent));
                }
                else {
                    helper((0, compiler_core_2.getVNodeHelper)(false, childBlock.isComponent));
                }
            }
            if (memo) {
                const loop = (0, compiler_core_1.createFunctionExpression)(createForLoopParams(forNode.parseResult, (0, compiler_core_1.createSimpleExpression)(`_cached`)));
                loop.body = (0, compiler_core_1.createBlockStatement)([
                    (0, compiler_core_1.createCompoundExpression)([
                        `const _memo: Array<any | null> = (`,
                        memo.exp,
                        `)`,
                    ]),
                    (0, compiler_core_1.createCompoundExpression)([
                        `if (_cached != null`,
                        ...(keyExp ? [` && _cached.key === `, keyExp] : []),
                        ` && ${context.helperString(compiler_core_1.IS_MEMO_SAME)}(_cached, _memo)) return _cached`,
                    ]),
                    (0, compiler_core_1.createCompoundExpression)([`const _item = `, childBlock]),
                    (0, compiler_core_1.createSimpleExpression)(`_item.memo = _memo`),
                    (0, compiler_core_1.createSimpleExpression)(`return _item`),
                ]);
                renderExp.arguments.push(loop, (0, compiler_core_1.createSimpleExpression)(`_cache`), (0, compiler_core_1.createSimpleExpression)(String(context.cached++)));
            }
            else {
                renderExp.arguments.push((0, compiler_core_1.createFunctionExpression)(createForLoopParams(forNode.parseResult, (0, compiler_core_1.createSimpleExpression)(`_cached`)), childBlock, true /* force newline */));
            }
        };
    });
});
// target-agnostic transform used for both Client and SSR
function processFor(node, dir, context, processCodegen) {
    if (!dir.exp) {
        context.onError((0, errors_1.createCompilerError)(31 /* ErrorCodes.X_V_FOR_NO_EXPRESSION */, dir.loc));
        return;
    }
    const parseResult = parseForExpression(
    // can only be simple expression because vFor transform is applied
    // before expression transform.
    dir.exp, context);
    if (!parseResult) {
        context.onError((0, errors_1.createCompilerError)(32 /* ErrorCodes.X_V_FOR_MALFORMED_EXPRESSION */, dir.loc));
        return;
    }
    const { addIdentifiers, removeIdentifiers, scopes } = context;
    const { source, value, key, index } = parseResult;
    if (key === undefined) {
        parseResult.key = {
            constType: 2,
            content: '__key',
            isStatic: false,
            loc: value?.loc,
            type: 4,
        };
    }
    if (index === undefined) {
        parseResult.index = {
            constType: 2,
            content: '__index',
            isStatic: false,
            loc: value?.loc,
            type: 4,
        };
    }
    parseResult.finalized = true;
    const forNode = {
        type: compiler_core_1.NodeTypes.FOR,
        loc: dir.loc,
        source,
        valueAlias: value,
        keyAlias: key,
        objectIndexAlias: index,
        parseResult,
        children: (0, compiler_core_2.isTemplateNode)(node) ? node.children : [node],
    };
    context.replaceNode(forNode);
    // bookkeeping
    scopes.vFor++;
    // if (!__BROWSER__ && context.prefixIdentifiers) {
    if (context.prefixIdentifiers) {
        // scope management
        // inject identifiers to context
        value && addIdentifiers(value);
        key && addIdentifiers(key);
        index && addIdentifiers(index);
    }
    const onExit = processCodegen && processCodegen(forNode);
    return () => {
        scopes.vFor--;
        // if (!__BROWSER__ && context.prefixIdentifiers) {
        if (context.prefixIdentifiers) {
            value && removeIdentifiers(value);
            key && removeIdentifiers(key);
            index && removeIdentifiers(index);
        }
        if (onExit)
            onExit();
    };
}
exports.processFor = processFor;
const forAliasRE = /([\s\S]*?)\s+(?:in|of)\s+([\s\S]*)/;
// This regex doesn't cover the case if key or index aliases have destructuring,
// but those do not make sense in the first place, so this works in practice.
const forIteratorRE = /,([^,\}\]]*)(?:,([^,\}\]]*))?$/;
const stripParensRE = /^\(|\)$/g;
function parseForExpression(input, context) {
    const loc = input.loc;
    const exp = input.content;
    const inMatch = exp.match(forAliasRE);
    if (!inMatch)
        return;
    const [, LHS, RHS] = inMatch;
    const result = {
        source: createAliasExpression(loc, RHS.trim(), exp.indexOf(RHS, LHS.length)),
        value: undefined,
        key: undefined,
        index: undefined,
        finalized: false,
    };
    if (!utils_1.__BROWSER__ && context.prefixIdentifiers) {
        result.source = (0, transformExpression_1.processExpression)(result.source, context);
    }
    // if (__DEV__ && __BROWSER__) {
    //   validateBrowserExpression(result.source as SimpleExpressionNode, context)
    // }
    let valueContent = LHS.trim().replace(stripParensRE, '').trim();
    const trimmedOffset = LHS.indexOf(valueContent);
    const iteratorMatch = valueContent.match(forIteratorRE);
    if (iteratorMatch) {
        valueContent = valueContent.replace(forIteratorRE, '').trim();
        const keyContent = iteratorMatch[1].trim();
        let keyOffset;
        if (keyContent) {
            keyOffset = exp.indexOf(keyContent, trimmedOffset + valueContent.length);
            result.key = createAliasExpression(loc, keyContent, keyOffset);
            // if (!__BROWSER__ && context.prefixIdentifiers) {
            if (context.prefixIdentifiers) {
                result.key = (0, transformExpression_1.processExpression)(result.key, context, true);
            }
            // if (__DEV__ && __BROWSER__) {
            //   validateBrowserExpression(
            //     result.key as SimpleExpressionNode,
            //     context,
            //     true
            //   )
            // }
        }
        if (iteratorMatch[2]) {
            const indexContent = iteratorMatch[2].trim();
            if (indexContent) {
                result.index = createAliasExpression(loc, indexContent, exp.indexOf(indexContent, result.key
                    ? keyOffset + keyContent.length
                    : trimmedOffset + valueContent.length));
                // if (!__BROWSER__ && context.prefixIdentifiers) {
                if (context.prefixIdentifiers) {
                    result.index = (0, transformExpression_1.processExpression)(result.index, context, true);
                }
                // if (__DEV__ && __BROWSER__) {
                //   validateBrowserExpression(
                //     result.index as SimpleExpressionNode,
                //     context,
                //     true
                //   )
                // }
            }
        }
    }
    if (valueContent) {
        result.value = createAliasExpression(loc, valueContent, trimmedOffset);
        // if (!__BROWSER__ && context.prefixIdentifiers) {
        if (context.prefixIdentifiers) {
            result.value = (0, transformExpression_1.processExpression)(result.value, context, true);
        }
        // if (__DEV__ && __BROWSER__) {
        //   validateBrowserExpression(
        //     result.value as SimpleExpressionNode,
        //     context,
        //     true
        //   )
        // }
    }
    return result;
}
exports.parseForExpression = parseForExpression;
function createAliasExpression(range, content, offset) {
    return (0, compiler_core_1.createSimpleExpression)(content, false, (0, uni_cli_shared_1.getInnerRange)(range, offset, content.length));
}
function createForLoopParams({ value, key, index }, memoArgs) {
    return createParamsList([value, key, index, memoArgs]);
}
exports.createForLoopParams = createForLoopParams;
const paramNames = ['__value', '__key', '__item', '_cached'];
function createParamsList(args) {
    let i = args.length;
    while (i--) {
        if (args[i])
            break;
    }
    return args
        .slice(0, i + 1)
        .map((arg, i) => arg || (0, compiler_core_1.createSimpleExpression)(paramNames[i], false));
}
