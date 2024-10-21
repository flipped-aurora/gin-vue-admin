"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isCompoundExpressionNode = exports.isSimpleExpressionNode = exports.isDirectiveNode = exports.isAttributeNode = exports.isElementNode = exports.parseVue = exports.createCallExpression = exports.createIdentifier = exports.createLiteral = exports.isReference = exports.isExportSpecifier = exports.isMethodDefinition = exports.isMemberExpression = exports.isCallExpression = exports.isAssignmentExpression = exports.isIdentifier = exports.isProperty = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const compiler_dom_1 = require("@vue/compiler-dom");
const isProperty = (node) => node.type === 'Property';
exports.isProperty = isProperty;
const isIdentifier = (node) => node.type === 'Identifier';
exports.isIdentifier = isIdentifier;
const isAssignmentExpression = (node) => node.type === 'AssignmentExpression';
exports.isAssignmentExpression = isAssignmentExpression;
const isCallExpression = (node) => node.type === 'CallExpression';
exports.isCallExpression = isCallExpression;
const isMemberExpression = (node) => node.type === 'MemberExpression';
exports.isMemberExpression = isMemberExpression;
const isMethodDefinition = (node) => node.type === 'MethodDefinition';
exports.isMethodDefinition = isMethodDefinition;
const isExportSpecifier = (node) => node.type === 'ExportSpecifier';
exports.isExportSpecifier = isExportSpecifier;
const isReference = (node, parent) => {
    if ((0, exports.isMemberExpression)(node)) {
        return !node.computed && (0, exports.isReference)(node.object, node);
    }
    if ((0, exports.isIdentifier)(node)) {
        if ((0, exports.isMemberExpression)(parent))
            return parent.computed || node === parent.object;
        // `bar` in { bar: foo }
        if ((0, exports.isProperty)(parent) && node !== parent.value)
            return false;
        // `bar` in `class Foo { bar () {...} }`
        if ((0, exports.isMethodDefinition)(parent))
            return false;
        // `bar` in `export { foo as bar }`
        if ((0, exports.isExportSpecifier)(parent) && node !== parent.local)
            return false;
        return true;
    }
    return false;
};
exports.isReference = isReference;
function createLiteral(value) {
    return {
        type: 'Literal',
        value,
        raw: `'${value}'`,
    };
}
exports.createLiteral = createLiteral;
function createIdentifier(name) {
    return {
        type: 'Identifier',
        name,
    };
}
exports.createIdentifier = createIdentifier;
function createCallExpression(callee, args) {
    return {
        type: 'CallExpression',
        callee,
        arguments: args,
    };
}
exports.createCallExpression = createCallExpression;
function parseVue(code, errors) {
    return (0, compiler_dom_1.parse)(code, {
        isNativeTag: () => true,
        isPreTag: () => true,
        parseMode: 'sfc',
        onError: (e) => {
            errors.push(e);
        },
    });
}
exports.parseVue = parseVue;
function isElementNode(node) {
    return node.type === compiler_core_1.NodeTypes.ELEMENT;
}
exports.isElementNode = isElementNode;
function isAttributeNode(node) {
    return node.type === compiler_core_1.NodeTypes.ATTRIBUTE;
}
exports.isAttributeNode = isAttributeNode;
function isDirectiveNode(node) {
    return node.type === compiler_core_1.NodeTypes.DIRECTIVE;
}
exports.isDirectiveNode = isDirectiveNode;
function isSimpleExpressionNode(node) {
    return node.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION;
}
exports.isSimpleExpressionNode = isSimpleExpressionNode;
function isCompoundExpressionNode(node) {
    return node.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION;
}
exports.isCompoundExpressionNode = isCompoundExpressionNode;
