"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createVForArrowFunctionExpression = exports.createForLoopParams = exports.parseForExpression = exports.parseVForScope = exports.transformFor = exports.isForElementNode = void 0;
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const ast_1 = require("../ast");
const transform_1 = require("../transform");
const transformExpression_1 = require("./transformExpression");
const codegen_1 = require("../codegen");
const types_1 = require("@babel/types");
const utils_1 = require("./utils");
const runtimeHelpers_1 = require("../runtimeHelpers");
const vSlot_1 = require("./vSlot");
function isForElementNode(node) {
    return !!node.vFor;
}
exports.isForElementNode = isForElementNode;
exports.transformFor = (0, transform_1.createStructuralDirectiveTransform)('for', (node, dir, context) => {
    if (!dir.exp) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_FOR_NO_EXPRESSION, dir.loc));
        return;
    }
    const parseResult = parseForExpression(dir.exp, context);
    if (!parseResult) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_FOR_MALFORMED_EXPRESSION, dir.loc));
        return;
    }
    parseResult.tagType = node.tagType;
    const { addIdentifiers, removeIdentifiers } = context;
    const { source, value, key, index } = parseResult;
    if (context.prefixIdentifiers) {
        addIdentifiers(value);
        addIdentifiers(key);
        addIdentifiers(index);
    }
    const { currentScope: parentScope, scopes, popScope } = context;
    const sourceExpr = (0, ast_1.parseExpr)(source, context);
    const valueCode = (0, codegen_1.genExpr)(value);
    const valueExpr = (0, ast_1.parseParam)(valueCode, context, value);
    const valueAlias = parseAlias(valueExpr, valueCode, 'v' + scopes.vFor);
    const keyCode = (0, codegen_1.genExpr)(key);
    const keyExpr = (0, ast_1.parseParam)(keyCode, context, key);
    const keyAlias = parseAlias(keyExpr, keyCode, 'k' + scopes.vFor);
    const indexCode = (0, codegen_1.genExpr)(index);
    const indexExpr = (0, ast_1.parseParam)(indexCode, context, index);
    const indexAlias = parseAlias(indexExpr, indexCode, 'i' + scopes.vFor);
    // 先占位 vFor，后续更新 cloneSourceExpr 为 CallExpression
    const cloneSourceExpr = (0, types_1.cloneNode)(sourceExpr, false);
    const sourceAliasReferencedScope = (0, utils_1.findReferencedScope)(cloneSourceExpr, context.currentScope, 
    // vFor 嵌套时，始终保持嵌套关系，issues/3263
    false);
    // 寻找子节点中 if 指令作用域
    const vIfReferencedScope = findVIfReferencedScope(node, context.currentScope, context);
    // 取最近的作用域
    const referencedScope = vIfReferencedScope &&
        context.getScopeIndex(vIfReferencedScope) >
            context.getScopeIndex(sourceAliasReferencedScope)
        ? vIfReferencedScope
        : sourceAliasReferencedScope;
    const sourceAlias = (0, utils_1.rewriteExpression)(source, context, cloneSourceExpr, parentScope, 
    // 强制 rewrite，因为即使是字符串，数字，也要走 vFor 函数
    {
        property: true,
        ignoreLiteral: true,
        referencedScope,
    }).content;
    const sourceCode = `{{${sourceAlias}}}`;
    const vForData = {
        source,
        sourceExpr,
        sourceAlias,
        sourceCode,
        value,
        valueCode,
        valueExpr,
        valueAlias,
        key,
        keyCode,
        keyExpr,
        keyAlias,
        index,
        indexCode,
        indexExpr,
        indexAlias,
        node,
    };
    const vForScope = context.addVForScope({
        ...vForData,
        locals: findVForLocals(parseResult),
    });
    const vFor = {
        ...vForData,
    };
    const isScopedSlot = (0, transform_1.isScopedSlotVFor)(vForScope);
    node.vFor = vFor;
    scopes.vFor++;
    return () => {
        scopes.vFor--;
        if ((0, compiler_core_1.isTemplateNode)(node)) {
            node.children.some((c) => {
                if (c.type === compiler_core_1.NodeTypes.ELEMENT && !isForElementNode(c)) {
                    const key = (0, compiler_core_1.findProp)(c, 'key');
                    if (key) {
                        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_FOR_TEMPLATE_KEY_PLACEMENT, key.loc));
                        return true;
                    }
                }
            });
        }
        if (context.prefixIdentifiers) {
            value && removeIdentifiers(value);
            key && removeIdentifiers(key);
            index && removeIdentifiers(index);
        }
        (0, shared_1.extend)(clearExpr(cloneSourceExpr), isScopedSlot
            ? (0, vSlot_1.createVSlotCallExpression)(node
                .slotComponent, vForScope, context)
            : createVForCallExpression(vForScope, context));
        popScope();
    };
});
function clearExpr(expr) {
    Object.keys(expr).forEach((key) => {
        delete expr[key];
    });
    return expr;
}
function parseAlias(babelExpr, exprCode, fallback) {
    if ((0, types_1.isIdentifier)(babelExpr)) {
        return exprCode;
    }
    return fallback;
}
function parseVForScope(currentScope) {
    while (currentScope) {
        if ((0, transform_1.isVForScope)(currentScope) && !(0, transform_1.isScopedSlotVFor)(currentScope)) {
            return currentScope;
        }
        currentScope = currentScope.parent;
    }
}
exports.parseVForScope = parseVForScope;
function findVIfReferencedScope(node, currentScope, context) {
    if (!currentScope) {
        return;
    }
    const vForScope = parseVForScope(currentScope);
    if (!vForScope) {
        return;
    }
    if (!node.children.find((item) => checkVIfReferenced(item, vForScope, context))) {
        return findVIfReferencedScope(node, currentScope.parent, context);
    }
    return vForScope;
}
function checkVIfReferenced(node, vForScope, context) {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return false;
    }
    // 嵌套 for 不查找
    if ((0, compiler_core_1.findDir)(node, 'for')) {
        return false;
    }
    const ifDir = (0, compiler_core_1.findDir)(node, 'if');
    if (ifDir) {
        return checkDirReferenced(ifDir.exp, vForScope, context);
    }
    const elseIfDir = (0, compiler_core_1.findDir)(node, 'else-if');
    if (elseIfDir) {
        return checkDirReferenced(elseIfDir.exp, vForScope, context);
    }
    return !!node.children.find((item) => checkVIfReferenced(item, vForScope, context));
}
function checkDirReferenced(node, vForScope, context) {
    if (node) {
        const babelNode = (0, ast_1.parseExpr)(node, context);
        if (babelNode && (0, utils_1.isReferencedByIds)(babelNode, vForScope.locals)) {
            return true;
        }
    }
    return false;
}
function findVForLocals({ value, key, index }) {
    const ids = [];
    if (value) {
        findIds(value, ids);
    }
    if (key) {
        findIds(key, ids);
    }
    if (index) {
        findIds(index, ids);
    }
    return ids;
}
function findIds(exp, ids) {
    if ((0, shared_1.isString)(exp)) {
        ids.push(exp);
    }
    else if (exp.identifiers) {
        exp.identifiers.forEach((id) => ids.push(id));
    }
    else if (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
        ids.push(exp.content);
    }
}
const forAliasRE = /([\s\S]*?)\s+(?:in|of)\s+([\s\S]*)/;
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
        value: (0, compiler_core_1.createSimpleExpression)('v' + context.scopes.vFor),
        key: (0, compiler_core_1.createSimpleExpression)('k' + context.scopes.vFor),
        index: (0, compiler_core_1.createSimpleExpression)('i' + context.scopes.vFor),
        tagType: compiler_core_1.ElementTypes.ELEMENT,
    };
    if (context.prefixIdentifiers) {
        result.source = (0, transformExpression_1.processExpression)(result.source, context);
    }
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
            if (context.prefixIdentifiers) {
                result.key = (0, transformExpression_1.processExpression)(result.key, context, true);
            }
        }
        if (iteratorMatch[2]) {
            const indexContent = iteratorMatch[2].trim();
            if (indexContent) {
                result.index = createAliasExpression(loc, indexContent, exp.indexOf(indexContent, result.key
                    ? keyOffset + keyContent.length
                    : trimmedOffset + valueContent.length));
                if (context.prefixIdentifiers) {
                    result.index = (0, transformExpression_1.processExpression)(result.index, context, true);
                }
            }
        }
    }
    if (valueContent) {
        result.value = createAliasExpression(loc, valueContent, trimmedOffset);
        if (context.prefixIdentifiers) {
            result.value = (0, transformExpression_1.processExpression)(result.value, context, true);
        }
    }
    return result;
}
exports.parseForExpression = parseForExpression;
function createAliasExpression(range, content, offset) {
    return (0, compiler_core_1.createSimpleExpression)(content, false, (0, uni_cli_shared_1.getInnerRange)(range, offset, content.length));
}
function createForLoopParams({ value, key, index }, memoArgs = []) {
    return createParamsList([value, key, index, ...memoArgs]);
}
exports.createForLoopParams = createForLoopParams;
function createParamsList(args) {
    let i = args.length;
    while (i--) {
        if (args[i])
            break;
    }
    return args
        .slice(0, i + 1)
        .map((arg, i) => arg || (0, compiler_core_1.createSimpleExpression)(`_`.repeat(i + 1), false));
}
function createVForCallExpression(vForScope, context) {
    // let sourceExpr: Expression = vForScope.sourceExpr!
    // if (isNumericLiteral(sourceExpr)) {
    //   sourceExpr = numericLiteralToArrayExpr((sourceExpr as NumericLiteral).value)
    // }
    return (0, types_1.callExpression)((0, types_1.identifier)(context.helperString(runtimeHelpers_1.V_FOR)), [
        vForScope.sourceExpr,
        createVForArrowFunctionExpression(vForScope),
    ]);
}
function createVForArrowFunctionExpression({ valueExpr, keyExpr, indexExpr, properties, }) {
    const params = [];
    if (valueExpr) {
        params.push(valueExpr);
    }
    if (keyExpr) {
        params.push(keyExpr);
    }
    if (indexExpr) {
        params.push(indexExpr);
    }
    return (0, types_1.arrowFunctionExpression)(params, (0, types_1.blockStatement)([(0, types_1.returnStatement)((0, types_1.objectExpression)(properties))]));
}
exports.createVForArrowFunctionExpression = createVForArrowFunctionExpression;
