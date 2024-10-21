"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getEscapedCssVarName = exports.cssVarNameEscapeSymbolsRE = exports.getEscapedPropName = exports.propNameEscapeSymbolsRE = exports.joinPaths = exports.normalizePath = exports.createGetCanonicalFileName = exports.getId = exports.getImportedName = exports.toRuntimeTypeString = exports.isCallOf = exports.unwrapTSNode = exports.isLiteralNode = exports.concatStrings = exports.resolveObjectKey = exports.resolveDefineCode = exports.UNKNOWN_TYPE = void 0;
const path_1 = __importDefault(require("path"));
const compiler_dom_1 = require("@vue/compiler-dom");
exports.UNKNOWN_TYPE = 'Unknown';
function resolveDefineCode(componentType) {
    return componentType === 'app'
        ? `defineApp`
        : componentType === 'page'
            ? `defineComponent` //`definePage`
            : `defineComponent`;
}
exports.resolveDefineCode = resolveDefineCode;
function resolveObjectKey(node, computed) {
    switch (node.type) {
        case 'StringLiteral':
        case 'NumericLiteral':
            return String(node.value);
        case 'Identifier':
            if (!computed)
                return node.name;
    }
    return undefined;
}
exports.resolveObjectKey = resolveObjectKey;
function concatStrings(strs) {
    return strs.filter((s) => !!s).join(', ');
}
exports.concatStrings = concatStrings;
function isLiteralNode(node) {
    return node.type.endsWith('Literal');
}
exports.isLiteralNode = isLiteralNode;
function unwrapTSNode(node) {
    if (compiler_dom_1.TS_NODE_TYPES.includes(node.type)) {
        return unwrapTSNode(node.expression);
    }
    else {
        return node;
    }
}
exports.unwrapTSNode = unwrapTSNode;
function isCallOf(node, test) {
    return !!(node &&
        test &&
        node.type === 'CallExpression' &&
        node.callee.type === 'Identifier' &&
        (typeof test === 'string'
            ? node.callee.name === test
            : test(node.callee.name)));
}
exports.isCallOf = isCallOf;
function toRuntimeTypeString(types) {
    return types.length > 1 ? `[${types.join(', ')}]` : types[0] || `Object`;
}
exports.toRuntimeTypeString = toRuntimeTypeString;
function getImportedName(specifier) {
    if (specifier.type === 'ImportSpecifier')
        return specifier.imported.type === 'Identifier'
            ? specifier.imported.name
            : specifier.imported.value;
    else if (specifier.type === 'ImportNamespaceSpecifier')
        return '*';
    return 'default';
}
exports.getImportedName = getImportedName;
function getId(node) {
    return node.type === 'Identifier'
        ? node.name
        : node.type === 'StringLiteral'
            ? node.value
            : null;
}
exports.getId = getId;
const identity = (str) => str;
const fileNameLowerCaseRegExp = /[^\u0130\u0131\u00DFa-z0-9\\/:\-_\. ]+/g;
const toLowerCase = (str) => str.toLowerCase();
function toFileNameLowerCase(x) {
    return fileNameLowerCaseRegExp.test(x)
        ? x.replace(fileNameLowerCaseRegExp, toLowerCase)
        : x;
}
/**
 * We need `getCanonicalFileName` when creating ts module resolution cache,
 * but TS does not expose it directly. This implementation is repllicated from
 * the TS source code.
 */
function createGetCanonicalFileName(useCaseSensitiveFileNames) {
    return useCaseSensitiveFileNames ? identity : toFileNameLowerCase;
}
exports.createGetCanonicalFileName = createGetCanonicalFileName;
// in the browser build, the polyfill doesn't expose posix, but defaults to
// posix behavior.
const normalize = (path_1.default.posix || path_1.default).normalize;
const windowsSlashRE = /\\/g;
function normalizePath(p) {
    return normalize(p.replace(windowsSlashRE, '/'));
}
exports.normalizePath = normalizePath;
exports.joinPaths = (path_1.default.posix || path_1.default).join;
/**
 * key may contain symbols
 * e.g. onUpdate:modelValue -> "onUpdate:modelValue"
 */
exports.propNameEscapeSymbolsRE = /[ !"#$%&'()*+,./:;<=>?@[\\\]^`{|}~\-]/;
function getEscapedPropName(key) {
    return exports.propNameEscapeSymbolsRE.test(key) ? JSON.stringify(key) : key;
}
exports.getEscapedPropName = getEscapedPropName;
exports.cssVarNameEscapeSymbolsRE = /[ !"#$%&'()*+,./:;<=>?@[\\\]^`{|}~]/g;
function getEscapedCssVarName(key, doubleEscape) {
    return key.replace(exports.cssVarNameEscapeSymbolsRE, (s) => doubleEscape ? `\\\\${s}` : `\\${s}`);
}
exports.getEscapedCssVarName = getEscapedCssVarName;
