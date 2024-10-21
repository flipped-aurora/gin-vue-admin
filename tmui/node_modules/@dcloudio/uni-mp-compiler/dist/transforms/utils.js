"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.removeAttribute = exports.isStaticLiteral = exports.isReferencedByIds = exports.findReferencedScope = exports.rewriteExpression = exports.rewriteExpressionWithoutProperty = exports.parseExprWithRewriteClass = exports.parseExprWithRewrite = exports.rewirteWithHelper = exports.rewriteSpreadElement = exports.VIRTUAL_HOST_CLASS = exports.VIRTUAL_HOST_STYLE = exports.SCOPED_SLOT_IDENTIFIER = exports.ATTR_COM_TYPE = exports.ATTR_VUE_REF_IN_FOR = exports.ATTR_VUE_REF = exports.ATTR_VUE_PROPS = exports.ATTR_VUE_SLOTS = exports.ATTR_VUE_ID = void 0;
const types_1 = require("@babel/types");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const estree_walker_1 = require("estree-walker");
const ast_1 = require("../ast");
const codegen_1 = require("../codegen");
const transform_1 = require("../transform");
// v-i,v-s 不能在 quickapp-webview 中使用，估计是内部处理成了指令之类的
exports.ATTR_VUE_ID = 'u-i';
exports.ATTR_VUE_SLOTS = 'u-s';
exports.ATTR_VUE_PROPS = 'u-p';
exports.ATTR_VUE_REF = 'u-' + uni_cli_shared_1.VUE_REF;
exports.ATTR_VUE_REF_IN_FOR = 'u-' + uni_cli_shared_1.VUE_REF_IN_FOR;
exports.ATTR_COM_TYPE = 'u-t';
exports.SCOPED_SLOT_IDENTIFIER = '__SCOPED_SLOT__';
exports.VIRTUAL_HOST_STYLE = 'virtualHostStyle';
exports.VIRTUAL_HOST_CLASS = 'virtualHostClass';
function rewriteSpreadElement(name, expr, loc, context) {
    return rewirteWithHelper(name, expr.argument, loc, context);
}
exports.rewriteSpreadElement = rewriteSpreadElement;
function rewirteWithHelper(name, expr, loc, context) {
    return parseExprWithRewrite(context.helperString(name) + '(' + (0, codegen_1.genBabelExpr)(expr) + ')', loc, context);
}
exports.rewirteWithHelper = rewirteWithHelper;
function parseExprWithRewrite(code, loc, context, node) {
    return (0, ast_1.parseExpr)(rewriteExpression((0, compiler_core_1.createSimpleExpression)(code, false, loc), context, node), context);
}
exports.parseExprWithRewrite = parseExprWithRewrite;
function parseExprWithRewriteClass(code, loc, context, node) {
    // a?1:0
    return (0, ast_1.parseExpr)(rewriteExpression((0, compiler_core_1.createSimpleExpression)(code, false, loc), context, !(0, ast_1.isUndefined)(node)
        ? (0, types_1.conditionalExpression)(node, (0, types_1.numericLiteral)(1), (0, types_1.stringLiteral)(''))
        : node), context);
}
exports.parseExprWithRewriteClass = parseExprWithRewriteClass;
function rewriteExpressionWithoutProperty(node, context, babelNode, scope = context.currentScope) {
    return rewriteExpression(node, context, babelNode, scope, {
        property: false,
        ignoreLiteral: false,
    });
}
exports.rewriteExpressionWithoutProperty = rewriteExpressionWithoutProperty;
function rewriteExpression(node, context, babelNode, scope = context.currentScope, { property, ignoreLiteral, referencedScope, } = {
    property: true,
    ignoreLiteral: false,
}) {
    if (node.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && node.isStatic) {
        return node;
    }
    if (!babelNode) {
        const code = (0, codegen_1.genExpr)(node);
        babelNode = (0, ast_1.parseExpr)(code, context, node);
        if (!babelNode) {
            return (0, compiler_core_1.createSimpleExpression)(code);
        }
    }
    if (!ignoreLiteral && isStaticLiteral(babelNode)) {
        return node;
    }
    if ((0, ast_1.isUndefined)(babelNode)) {
        return (0, compiler_core_1.createSimpleExpression)('undefined', false, node.loc);
    }
    // wxs 等表达式
    if (context.filters?.length) {
        if (isReferencedByIds(babelNode, context.filters)) {
            return (0, compiler_core_1.createSimpleExpression)((0, codegen_1.genExpr)(node), false, node.loc);
        }
    }
    referencedScope = referencedScope || findReferencedScope(babelNode, scope);
    const id = referencedScope.id.next();
    if (property) {
        referencedScope.properties.push((0, types_1.objectProperty)((0, types_1.identifier)(id), babelNode));
    }
    // 在 v-for 中包含的 v-if 块，所有变量需要补充当前 v-for value 前缀
    if ((0, transform_1.isVIfScope)(referencedScope)) {
        if ((0, transform_1.isVForScope)(referencedScope.parentScope)) {
            return (0, compiler_core_1.createSimpleExpression)(referencedScope.parentScope.valueAlias + '.' + id);
        }
        return (0, compiler_core_1.createSimpleExpression)(id);
    }
    else if ((0, transform_1.isVForScope)(referencedScope)) {
        return (0, compiler_core_1.createSimpleExpression)(referencedScope.valueAlias + '.' + id);
    }
    return (0, compiler_core_1.createSimpleExpression)(id);
}
exports.rewriteExpression = rewriteExpression;
function findReferencedScope(node, scope, findReferenced = true) {
    if ((0, transform_1.isVIfScope)(scope)) {
        return scope;
    }
    else if ((0, transform_1.isVForScope)(scope)) {
        if (!findReferenced) {
            return scope;
        }
        if (isReferencedByIds(node, scope.locals)) {
            return scope;
        }
        return findReferencedScope(node, scope.parent, findReferenced);
    }
    return scope;
}
exports.findReferencedScope = findReferencedScope;
function isReferencedByIds(node, knownIds) {
    let referenced = false;
    (0, estree_walker_1.walk)(node, {
        enter(node, parent) {
            if (referenced) {
                return this.skip();
            }
            if (!(0, types_1.isIdentifier)(node)) {
                return;
            }
            if (knownIds.includes(node.name) &&
                (!parent || (0, types_1.isReferenced)(node, parent))) {
                referenced = true;
                return this.skip();
            }
        },
    });
    return referenced;
}
exports.isReferencedByIds = isReferencedByIds;
function isStaticLiteral(value) {
    return (0, types_1.isLiteral)(value) && !(0, types_1.isTemplateLiteral)(value);
}
exports.isStaticLiteral = isStaticLiteral;
function removeAttribute(node, name) {
    const index = node.props.findIndex((prop) => prop.name === name);
    if (index > -1) {
        node.props.splice(index, 1);
    }
}
exports.removeAttribute = removeAttribute;
