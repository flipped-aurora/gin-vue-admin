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
exports.offsetToStartAndEnd = exports.locToStartAndEnd = exports.generateCodeFrame = exports.rewriteScssReadFileSync = exports.commonjsProxyRE = exports.cssLangRE = exports.minifyCSS = exports.cssPostPlugin = exports.cssPlugin = exports.isCSSRequest = exports.getAssetHash = exports.parseAssets = exports.assetPlugin = exports.uniViteSfcSrcImportPlugin = void 0;
__exportStar(require("./cssScoped"), exports);
__exportStar(require("./copy"), exports);
__exportStar(require("./inject"), exports);
__exportStar(require("./mainJs"), exports);
__exportStar(require("./jsonJs"), exports);
__exportStar(require("./console"), exports);
__exportStar(require("./dynamicImportPolyfill"), exports);
__exportStar(require("./uts/uni_modules"), exports);
__exportStar(require("./uts/uvue"), exports);
__exportStar(require("./uts/ext-api"), exports);
__exportStar(require("./easycom"), exports);
__exportStar(require("./json"), exports);
var sfc_1 = require("./sfc");
Object.defineProperty(exports, "uniViteSfcSrcImportPlugin", { enumerable: true, get: function () { return sfc_1.uniViteSfcSrcImportPlugin; } });
var asset_1 = require("./vitejs/plugins/asset");
Object.defineProperty(exports, "assetPlugin", { enumerable: true, get: function () { return asset_1.assetPlugin; } });
Object.defineProperty(exports, "parseAssets", { enumerable: true, get: function () { return asset_1.parseAssets; } });
Object.defineProperty(exports, "getAssetHash", { enumerable: true, get: function () { return asset_1.getAssetHash; } });
var css_1 = require("./vitejs/plugins/css");
Object.defineProperty(exports, "isCSSRequest", { enumerable: true, get: function () { return css_1.isCSSRequest; } });
Object.defineProperty(exports, "cssPlugin", { enumerable: true, get: function () { return css_1.cssPlugin; } });
Object.defineProperty(exports, "cssPostPlugin", { enumerable: true, get: function () { return css_1.cssPostPlugin; } });
Object.defineProperty(exports, "minifyCSS", { enumerable: true, get: function () { return css_1.minifyCSS; } });
Object.defineProperty(exports, "cssLangRE", { enumerable: true, get: function () { return css_1.cssLangRE; } });
Object.defineProperty(exports, "commonjsProxyRE", { enumerable: true, get: function () { return css_1.commonjsProxyRE; } });
Object.defineProperty(exports, "rewriteScssReadFileSync", { enumerable: true, get: function () { return css_1.rewriteScssReadFileSync; } });
var utils_1 = require("./vitejs/utils");
Object.defineProperty(exports, "generateCodeFrame", { enumerable: true, get: function () { return utils_1.generateCodeFrame; } });
Object.defineProperty(exports, "locToStartAndEnd", { enumerable: true, get: function () { return utils_1.locToStartAndEnd; } });
Object.defineProperty(exports, "offsetToStartAndEnd", { enumerable: true, get: function () { return utils_1.offsetToStartAndEnd; } });
