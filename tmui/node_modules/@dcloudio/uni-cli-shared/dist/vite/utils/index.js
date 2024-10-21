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
exports.isCombineBuiltInCss = exports.buildInCssSet = void 0;
__exportStar(require("./ast"), exports);
__exportStar(require("./url"), exports);
__exportStar(require("./plugin"), exports);
__exportStar(require("./utils"), exports);
// 内置组件css列表，h5平台需要合并进去首页css中
exports.buildInCssSet = new Set();
function isCombineBuiltInCss(config) {
    if (process.env.UNI_COMPILE_TARGET === 'uni_modules') {
        return false;
    }
    return config.command === 'build' && config.build.cssCodeSplit;
}
exports.isCombineBuiltInCss = isCombineBuiltInCss;
