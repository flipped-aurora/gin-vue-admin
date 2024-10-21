"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createVirtualHostClass = exports.rewriteClass = exports.findStaticClassIndex = exports.isClassBinding = void 0;
const types_1 = require("@babel/types");
const compiler_core_1 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const ast_1 = require("../ast");
const codegen_1 = require("../codegen");
const runtimeHelpers_1 = require("../runtimeHelpers");
const utils_1 = require("./utils");
function isClassBinding({ arg }) {
    return (arg && arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && arg.content === 'class');
}
exports.isClassBinding = isClassBinding;
function findStaticClassIndex(props) {
    return props.findIndex((prop) => prop.name === 'class');
}
exports.findStaticClassIndex = findStaticClassIndex;
function rewriteClass(index, classBindingProp, props, virtualHost, context) {
    const expr = classBindingProp.exp
        ? (0, ast_1.parseExpr)(classBindingProp.exp, context)
        : undefined;
    let classBindingExpr;
    if (expr) {
        classBindingExpr = expr;
        if ((0, types_1.isObjectExpression)(expr)) {
            classBindingExpr = createClassBindingByObjectExpression(rewriteClassObjectExpression(expr, classBindingProp.loc, context));
        }
        else if ((0, types_1.isArrayExpression)(expr)) {
            classBindingExpr = createClassBindingByArrayExpression(rewriteClassArrayExpression(expr, context));
        }
        else {
            classBindingExpr = (0, ast_1.parseExpr)(rewriteClassExpression(classBindingProp.exp, context).content, context);
        }
    }
    else if (virtualHost) {
        classBindingExpr = (0, types_1.arrayExpression)([]);
    }
    else {
        return;
    }
    const staticClassPropIndex = findStaticClassIndex(props);
    if (staticClassPropIndex > -1) {
        const staticClass = props[staticClassPropIndex].value
            .content;
        if (staticClass) {
            if (!(0, types_1.isArrayExpression)(classBindingExpr)) {
                classBindingExpr = (0, types_1.arrayExpression)([classBindingExpr]);
            }
            const staticClassLiterals = parseStaticClass(staticClass);
            if (index > staticClassPropIndex) {
                classBindingExpr.elements.unshift(...staticClassLiterals);
            }
            else {
                classBindingExpr.elements.push(...staticClassLiterals);
            }
        }
    }
    if (virtualHost) {
        if (!(0, types_1.isArrayExpression)(classBindingExpr)) {
            classBindingExpr = (0, types_1.arrayExpression)([classBindingExpr]);
        }
        classBindingExpr.elements.push((0, types_1.identifier)(utils_1.VIRTUAL_HOST_CLASS));
    }
    if (!context.miniProgram.class.array) {
        classBindingExpr = parseClassBindingArrayExpr(classBindingExpr);
    }
    classBindingProp.exp = (0, compiler_core_1.createSimpleExpression)((0, codegen_1.genBabelExpr)(classBindingExpr));
}
exports.rewriteClass = rewriteClass;
function createVirtualHostClass(props, context) {
    const classBindingProp = (0, uni_cli_shared_1.createBindDirectiveNode)('class', '');
    delete classBindingProp.exp;
    rewriteClass(0, classBindingProp, props, true, context);
    return classBindingProp;
}
exports.createVirtualHostClass = createVirtualHostClass;
/**
 * 目前 mp-toutiao, mp-alipay, mp-lark 不支持数组绑定class，故统一转换为字符串相加
 * @param classBindingExpr
 * @returns
 */
function parseClassBindingArrayExpr(classBindingExpr) {
    if (!(0, types_1.isArrayExpression)(classBindingExpr)) {
        return classBindingExpr;
    }
    let binaryExpr;
    classBindingExpr.elements.forEach((expr) => {
        if ((0, types_1.isArrayExpression)(expr)) {
            expr = parseClassBindingArrayExpr(expr);
        }
        if (!binaryExpr) {
            binaryExpr = (0, types_1.parenthesizedExpression)(expr);
        }
        else {
            binaryExpr = (0, types_1.binaryExpression)('+', (0, types_1.binaryExpression)('+', binaryExpr, (0, types_1.stringLiteral)(' ')), expr);
        }
    });
    return binaryExpr;
}
function parseStaticClass(staticClass) {
    // 已经在 parse 阶段格式化了多余空格等
    return staticClass.split(' ').map((clazz) => (0, types_1.stringLiteral)(clazz));
}
function rewriteClassExpression(expr, context) {
    return (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
        context.helperString(runtimeHelpers_1.NORMALIZE_CLASS) + '(',
        expr,
        ')',
    ]), context);
}
function rewriteClassArrayExpression(expr, context) {
    expr.elements.forEach((prop, index) => {
        if (!(0, types_1.isStringLiteral)(prop)) {
            const code = (0, codegen_1.genBabelExpr)((0, types_1.arrayExpression)([(0, types_1.isSpreadElement)(prop) ? prop.argument : prop]));
            expr.elements[index] = (0, types_1.identifier)(rewriteClassExpression((0, compiler_core_1.createSimpleExpression)(code.slice(1, -1), false), context).content);
        }
    });
    return expr;
}
function rewriteClassObjectExpression(expr, loc, context) {
    expr.properties.forEach((prop, index) => {
        if ((0, types_1.isSpreadElement)(prop)) {
            // <view :class="{...obj}"/>
            // <view class="{{[a]}}"/>
            const newExpr = (0, utils_1.rewriteSpreadElement)(runtimeHelpers_1.NORMALIZE_CLASS, prop, loc, context);
            if (newExpr) {
                expr.properties[index] = (0, types_1.objectProperty)(newExpr, (0, types_1.booleanLiteral)(true), true);
            }
        }
        else if ((0, types_1.isObjectProperty)(prop)) {
            const { key, value, computed } = prop;
            if (computed) {
                // {[handle(computedKey)]:1} => {[a]:1}
                prop.key = (0, utils_1.parseExprWithRewrite)((0, codegen_1.genBabelExpr)(key), loc, context, key);
            }
            if ((0, utils_1.isStaticLiteral)(value)) {
                return;
            }
            else {
                const newExpr = (0, utils_1.parseExprWithRewriteClass)((0, codegen_1.genBabelExpr)(value), loc, context, value);
                if (newExpr) {
                    prop.value = newExpr;
                }
            }
        }
    });
    return expr;
}
function createClassBindingByArrayExpression(expr) {
    const elements = [];
    expr.elements.forEach((prop) => {
        if ((0, types_1.isStringLiteral)(prop) || (0, types_1.isIdentifier)(prop)) {
            elements.push(prop);
        }
    });
    return (0, types_1.arrayExpression)(elements);
}
function createClassBindingByObjectExpression(expr) {
    const elements = [];
    expr.properties.forEach((prop) => {
        if ((0, types_1.isObjectProperty)(prop)) {
            const { value } = prop;
            if ((0, ast_1.isUndefined)(value) || (0, types_1.isPrivateName)(prop.key)) {
                // remove {a:undefined}
                return;
            }
            if ((0, types_1.isLiteral)(value)) {
                // {a:true,b:1,c:0} => ['a','b']
                if ((0, ast_1.isTrueExpr)(value)) {
                    elements.push(prop.computed
                        ? prop.key
                        : (0, ast_1.parseStringLiteral)(prop.key));
                }
                return;
            }
            elements.push((0, types_1.logicalExpression)('&&', value, prop.computed ? prop.key : (0, ast_1.parseStringLiteral)(prop.key)));
        }
    });
    return (0, types_1.arrayExpression)(elements);
}
