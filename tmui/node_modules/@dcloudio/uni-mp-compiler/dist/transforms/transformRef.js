"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteRef = void 0;
const types_1 = require("@babel/types");
const compiler_core_1 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("./utils");
const runtimeHelpers_1 = require("../runtimeHelpers");
const codegen_1 = require("../codegen");
const ast_1 = require("../ast");
function rewriteRef(node, context) {
    const vueIdProp = (0, compiler_core_1.findProp)(node, utils_1.ATTR_VUE_ID);
    if (!vueIdProp) {
        return;
    }
    const refProp = (0, compiler_core_1.findProp)(node, 'u-' + uni_cli_shared_1.VUE_REF) || (0, compiler_core_1.findProp)(node, 'u-' + uni_cli_shared_1.VUE_REF_IN_FOR);
    if (!refProp) {
        return;
    }
    if ((0, compiler_core_1.findProp)(node, 'ref')) {
        // 支付宝小程序
        const code = parseAlipayRefCode(refProp, context);
        if (code && context.inline && !(0, uni_cli_shared_1.isDirectiveNode)(refProp)) {
            refProp.value.content = code;
            const refPropIndex = node.props.findIndex((prop) => prop === refProp);
            node.props.splice(refPropIndex, 1, (0, uni_cli_shared_1.createBindDirectiveNode)(refProp.name, code));
        }
    }
    else {
        rewriteRefProp(refProp, vueIdProp, context);
    }
}
exports.rewriteRef = rewriteRef;
function parseRef(prop, context) {
    let expr;
    let refKey = '';
    const isDir = (0, uni_cli_shared_1.isDirectiveNode)(prop);
    if (isDir) {
        if (prop.exp) {
            expr = (0, ast_1.parseExpr)(prop.exp, context, prop.exp);
        }
    }
    else {
        const { value } = prop;
        if (value && value.content) {
            if (context.inline && context.bindingMetadata[value.content]) {
                expr = (0, types_1.identifier)(value.content);
                refKey = value.content;
            }
            else {
                expr = (0, types_1.stringLiteral)(value.content);
            }
        }
    }
    return { expr, refKey };
}
function parseRefCode(prop, context) {
    const { expr, refKey } = parseRef(prop, context);
    if (!expr) {
        return { code: '', refKey };
    }
    return { code: (0, codegen_1.genBabelExpr)(expr), refKey };
}
function rewriteRefProp(prop, vueIdProp, context) {
    let id = '';
    if ((0, uni_cli_shared_1.isDirectiveNode)(vueIdProp)) {
        const vueIdExpr = (0, ast_1.parseExpr)(vueIdProp.exp, context, vueIdProp.exp);
        if (vueIdExpr) {
            id = (0, codegen_1.genBabelExpr)(vueIdExpr);
        }
    }
    else {
        id = `'${vueIdProp.value.content}'`;
    }
    if (!id) {
        return;
    }
    const { code, refKey } = parseRefCode(prop, context);
    const opts = Object.create(null);
    if (refKey) {
        opts.k = refKey;
    }
    if (context.inVFor) {
        opts.f = 1;
    }
    (0, utils_1.parseExprWithRewrite)(context.helperString(runtimeHelpers_1.SET_REF) +
        '(' +
        code +
        ', ' +
        id +
        (Object.keys(opts).length ? ', ' + JSON.stringify(opts) : '') +
        ')', prop.loc, context);
}
function parseAlipayRefCode(prop, context) {
    let expr;
    const isDir = (0, uni_cli_shared_1.isDirectiveNode)(prop);
    if (isDir) {
        if (prop.exp) {
            expr = (0, ast_1.parseExpr)(prop.exp, context, prop.exp);
        }
    }
    else {
        if (prop.value?.content) {
            expr = context.inline
                ? processInlineRef(prop, context)
                : (0, types_1.stringLiteral)(prop.value.content);
        }
    }
    if (!expr) {
        return;
    }
    return (0, codegen_1.genBabelExpr)(expr);
}
function processInlineRef(prop, context) {
    const properties = [];
    const { refKey } = parseRef(prop, context);
    properties.push((0, types_1.objectProperty)((0, types_1.identifier)('r'), (0, types_1.identifier)(prop.value.content)));
    if (refKey) {
        properties.push((0, types_1.objectProperty)((0, types_1.identifier)('k'), (0, types_1.stringLiteral)(refKey)));
    }
    if (context.inVFor) {
        properties.push((0, types_1.objectProperty)((0, types_1.identifier)('f'), (0, types_1.numericLiteral)(1)));
    }
    return (0, types_1.arrowFunctionExpression)([], (0, types_1.objectExpression)(properties));
}
