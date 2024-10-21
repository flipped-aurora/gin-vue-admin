"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createVirtualHostStyle = exports.rewriteStyle = exports.findStaticStyleIndex = exports.isStyleBinding = void 0;
const types_1 = require("@babel/types");
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const runtimeHelpers_1 = require("../runtimeHelpers");
const ast_1 = require("../ast");
const codegen_1 = require("../codegen");
const utils_1 = require("./utils");
function isStyleBinding({ arg, exp }) {
    return (arg && arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && arg.content === 'style');
}
exports.isStyleBinding = isStyleBinding;
function findStaticStyleIndex(props) {
    return props.findIndex((prop) => prop.name === 'style');
}
exports.findStaticStyleIndex = findStaticStyleIndex;
function rewriteStyle(index, styleBindingProp, props, virtualHost, context) {
    const expr = styleBindingProp.exp
        ? (0, ast_1.parseExpr)(styleBindingProp.exp, context)
        : undefined;
    let styleBidingExpr = expr;
    if (expr) {
        if ((0, types_1.isObjectExpression)(expr)) {
            styleBidingExpr = createStyleBindingByObjectExpression(rewriteStyleObjectExpression(expr, styleBindingProp.loc, context));
        }
        else if ((0, types_1.isArrayExpression)(expr)) {
            styleBidingExpr = createStyleBindingByArrayExpression(rewriteStyleArrayExpression(expr, context));
        }
        else {
            styleBidingExpr = (0, ast_1.parseExpr)(rewriteStyleExpression(styleBindingProp.exp, context).content, context);
        }
        if (!styleBidingExpr) {
            return;
        }
    }
    else if (!virtualHost) {
        return;
    }
    const staticStylePropIndex = findStaticStyleIndex(props);
    if (staticStylePropIndex > -1) {
        const staticStyle = props[staticStylePropIndex].value
            .content;
        if (staticStyle.trim()) {
            if (styleBidingExpr) {
                if (index > staticStylePropIndex) {
                    styleBidingExpr = (0, types_1.binaryExpression)('+', addSemicolon((0, types_1.stringLiteral)(staticStyle)), styleBidingExpr);
                }
                else {
                    styleBidingExpr = (0, types_1.binaryExpression)('+', addSemicolon(styleBidingExpr), (0, types_1.stringLiteral)(staticStyle));
                }
            }
            else {
                styleBidingExpr = (0, types_1.stringLiteral)(staticStyle);
            }
        }
    }
    if (virtualHost) {
        styleBidingExpr = styleBidingExpr
            ? (0, types_1.binaryExpression)('+', addSemicolon(styleBidingExpr), (0, types_1.identifier)(utils_1.VIRTUAL_HOST_STYLE))
            : (0, types_1.identifier)(utils_1.VIRTUAL_HOST_STYLE);
    }
    styleBindingProp.exp = (0, compiler_core_1.createSimpleExpression)((0, codegen_1.genBabelExpr)(styleBidingExpr));
}
exports.rewriteStyle = rewriteStyle;
function createVirtualHostStyle(props, context) {
    const styleBindingProp = (0, uni_cli_shared_1.createBindDirectiveNode)('style', '');
    delete styleBindingProp.exp;
    rewriteStyle(0, styleBindingProp, props, true, context);
    return styleBindingProp;
}
exports.createVirtualHostStyle = createVirtualHostStyle;
function rewriteStyleExpression(expr, context) {
    return (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
        context.helperString(runtimeHelpers_1.STRINGIFY_STYLE) + '(',
        expr,
        ')',
    ]), context);
}
function rewriteStyleArrayExpression(expr, context) {
    expr.elements.forEach((prop, index) => {
        if (!(0, types_1.isStringLiteral)(prop)) {
            const code = (0, codegen_1.genBabelExpr)((0, types_1.arrayExpression)([(0, types_1.isSpreadElement)(prop) ? prop.argument : prop]));
            expr.elements[index] = (0, types_1.identifier)(rewriteStyleExpression((0, compiler_core_1.createSimpleExpression)(code.slice(1, -1), false), context).content);
        }
    });
    return expr;
}
function rewriteStyleObjectExpression(expr, loc, context) {
    expr.properties.forEach((prop, index) => {
        if ((0, types_1.isSpreadElement)(prop)) {
            // <view :style="{...obj}"/>
            // <view style="{{a}}"/>
            const newExpr = (0, utils_1.rewriteSpreadElement)(runtimeHelpers_1.STRINGIFY_STYLE, prop, loc, context);
            if (newExpr) {
                prop.argument = newExpr;
            }
        }
        else if ((0, types_1.isObjectProperty)(prop)) {
            const { key, value, computed } = prop;
            if (!(0, types_1.isPrivateName)(key)) {
                if (computed) {
                    // {[handle(computedKey)]:1} => {[a]:1}
                    const newExpr = (0, utils_1.rewirteWithHelper)(runtimeHelpers_1.HYPHENATE, key, loc, context);
                    if (newExpr) {
                        prop.key = newExpr;
                    }
                }
                else {
                    // {fontSize:'15px'} => {'font-size':'15px'}
                    prop.key = (0, ast_1.parseStringLiteral)(key);
                    prop.key.value = (0, shared_1.hyphenate)(prop.key.value) + ':';
                }
                // {fontSize:`${fontSize}px`} => {'font-size':a}
                if ((0, utils_1.isStaticLiteral)(value)) {
                    return;
                }
                else {
                    const newExpr = (0, utils_1.parseExprWithRewrite)((0, codegen_1.genBabelExpr)(value), loc, context, value);
                    if (newExpr) {
                        prop.value = newExpr;
                    }
                }
            }
        }
    });
    return expr;
}
function addSemicolon(expr) {
    return createBinaryExpression(expr, (0, types_1.stringLiteral)(';'));
}
function createBinaryExpression(left, right) {
    return (0, types_1.binaryExpression)('+', left, right);
}
function createStyleBindingByArrayExpression(expr) {
    let result;
    function concat(expr) {
        if (!result) {
            result = expr;
        }
        else {
            result = createBinaryExpression(addSemicolon(result), expr);
        }
    }
    expr.elements.forEach((prop) => {
        if ((0, types_1.isStringLiteral)(prop) || (0, types_1.isIdentifier)(prop)) {
            concat(prop);
        }
    });
    return result;
}
function createStyleBindingByObjectExpression(expr) {
    let result;
    function concat(expr) {
        if (!result) {
            result = expr;
        }
        else {
            result = createBinaryExpression(addSemicolon(result), expr);
        }
    }
    expr.properties.forEach((prop) => {
        if ((0, types_1.isSpreadElement)(prop)) {
            concat(prop.argument);
        }
        else if ((0, types_1.isObjectProperty)(prop)) {
            const { key, value } = prop;
            if (!(0, types_1.isPrivateName)(key)) {
                const expr = createBinaryExpression((0, types_1.isStringLiteral)(key)
                    ? key // 之前已经补充了:
                    : createBinaryExpression(key, (0, types_1.stringLiteral)(':')), value);
                concat(expr);
            }
        }
    });
    return result;
}
