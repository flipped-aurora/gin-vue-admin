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
exports.parseUniXSplashScreen = exports.parseUniXFlexDirection = exports.checkPagesJson = exports.normalizeUniAppXAppConfig = exports.normalizeUniAppXAppPagesJson = void 0;
__exportStar(require("./mp"), exports);
__exportStar(require("./app"), exports);
__exportStar(require("./json"), exports);
__exportStar(require("./pages"), exports);
__exportStar(require("./manifest"), exports);
__exportStar(require("./theme"), exports);
var uni_x_1 = require("./uni-x");
Object.defineProperty(exports, "normalizeUniAppXAppPagesJson", { enumerable: true, get: function () { return uni_x_1.normalizeUniAppXAppPagesJson; } });
Object.defineProperty(exports, "normalizeUniAppXAppConfig", { enumerable: true, get: function () { return uni_x_1.normalizeUniAppXAppConfig; } });
Object.defineProperty(exports, "checkPagesJson", { enumerable: true, get: function () { return uni_x_1.checkPagesJson; } });
Object.defineProperty(exports, "parseUniXFlexDirection", { enumerable: true, get: function () { return uni_x_1.parseUniXFlexDirection; } });
Object.defineProperty(exports, "parseUniXSplashScreen", { enumerable: true, get: function () { return uni_x_1.parseUniXSplashScreen; } });
