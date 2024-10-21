"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.compile = exports.parse = exports.transformModel = exports.transformOn = exports.isForElementNode = exports.rewriteExpression = exports.genExpr = exports.findProp = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const compile_1 = require("./compile");
const parserOptions_1 = require("./parserOptions");
var compiler_core_2 = require("@vue/compiler-core");
Object.defineProperty(exports, "findProp", { enumerable: true, get: function () { return compiler_core_2.findProp; } });
var codegen_1 = require("./codegen");
Object.defineProperty(exports, "genExpr", { enumerable: true, get: function () { return codegen_1.genExpr; } });
var utils_1 = require("./transforms/utils");
Object.defineProperty(exports, "rewriteExpression", { enumerable: true, get: function () { return utils_1.rewriteExpression; } });
var vFor_1 = require("./transforms/vFor");
Object.defineProperty(exports, "isForElementNode", { enumerable: true, get: function () { return vFor_1.isForElementNode; } });
var vOn_1 = require("./transforms/vOn");
Object.defineProperty(exports, "transformOn", { enumerable: true, get: function () { return vOn_1.transformOn; } });
var vModel_1 = require("./transforms/vModel");
Object.defineProperty(exports, "transformModel", { enumerable: true, get: function () { return vModel_1.transformModel; } });
__exportStar(require("./runtimeHelpers"), exports);
function parse(template, options = {}) {
    return (0, compiler_core_1.baseParse)(template, (0, shared_1.extend)({}, parserOptions_1.parserOptions, options));
}
exports.parse = parse;
function compile(template, options = {}) {
    return (0, compile_1.baseCompile)(template, (0, shared_1.extend)({}, parserOptions_1.parserOptions, options, {
        directiveTransforms: (0, shared_1.extend)({}, options.directiveTransforms || {}),
    }));
}
exports.compile = compile;
