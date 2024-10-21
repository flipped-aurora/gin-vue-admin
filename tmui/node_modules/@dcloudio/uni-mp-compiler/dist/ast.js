"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseStringLiteral = exports.isTrueExpr = exports.isUndefined = exports.parseParam = exports.parseExpr = exports.createVIfSpreadElement = exports.createVIfConditionalExpression = exports.createVIfProperty = exports.createObjectExpression = exports.createSpreadElement = exports.createObjectProperty = exports.createIdentifier = void 0;
const shared_1 = require("@vue/shared");
const parser_1 = require("@babel/parser");
const types_1 = require("@babel/types");
const compiler_core_1 = require("@vue/compiler-core");
const codegen_1 = require("./codegen");
function createIdentifier(name) {
    return (0, types_1.identifier)(name);
}
exports.createIdentifier = createIdentifier;
function createObjectProperty(name, value) {
    return (0, types_1.objectProperty)((0, types_1.identifier)(name), value);
}
exports.createObjectProperty = createObjectProperty;
function createSpreadElement(argument) {
    return (0, types_1.spreadElement)(argument);
}
exports.createSpreadElement = createSpreadElement;
function createObjectExpression(properties) {
    return (0, types_1.objectExpression)(properties);
}
exports.createObjectExpression = createObjectExpression;
function createVIfProperty(condition, { id }) {
    return (0, types_1.objectProperty)((0, types_1.identifier)(id.next()), condition);
}
exports.createVIfProperty = createVIfProperty;
function createVIfConditionalExpression({ condition, properties, }) {
    return (0, types_1.conditionalExpression)(condition, (0, types_1.objectExpression)(properties), (0, types_1.objectExpression)([]));
}
exports.createVIfConditionalExpression = createVIfConditionalExpression;
function createVIfSpreadElement(vIfScope) {
    return (0, types_1.spreadElement)(createVIfConditionalExpression(vIfScope));
}
exports.createVIfSpreadElement = createVIfSpreadElement;
// function numericLiteralToArrayExpr(num: number) {
//   const elements: NumericLiteral[] = []
//   for (let i = 0; i < num; i++) {
//     elements.push(numericLiteral(i + 1))
//   }
//   return arrayExpression(elements)
// }
function parseExpr(code, context, node) {
    if (!(0, shared_1.isString)(code)) {
        node = code;
        code = (0, codegen_1.genExpr)(code);
    }
    try {
        return (0, parser_1.parseExpression)(code, {
            plugins: context.expressionPlugins,
        });
    }
    catch (e) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_INVALID_EXPRESSION, node && node.loc, undefined, '\n' + code + '\n' + e.message));
    }
}
exports.parseExpr = parseExpr;
function parseParam(code, context, node) {
    const { params: [expr], } = parseExpr(`(${code})=>{}`, context, node);
    return expr;
}
exports.parseParam = parseParam;
function isUndefined(expr) {
    return (0, types_1.isIdentifier)(expr) && expr.name === 'undefined';
}
exports.isUndefined = isUndefined;
function isTrueExpr(expr) {
    if ((0, types_1.isNullLiteral)(expr)) {
        return false;
    }
    if ((0, types_1.isStringLiteral)(expr) ||
        (0, types_1.isNumericLiteral)(expr) ||
        (0, types_1.isBooleanLiteral)(expr) ||
        (0, types_1.isBigIntLiteral)(expr) ||
        (0, types_1.isDecimalLiteral)(expr)) {
        return !!expr.value;
    }
    return true;
}
exports.isTrueExpr = isTrueExpr;
function parseStringLiteral(expr) {
    if ((0, types_1.isIdentifier)(expr)) {
        return (0, types_1.stringLiteral)(expr.name);
    }
    if ((0, types_1.isStringLiteral)(expr)) {
        return (0, types_1.stringLiteral)(expr.value);
    }
    return (0, types_1.stringLiteral)('');
}
exports.parseStringLiteral = parseStringLiteral;
