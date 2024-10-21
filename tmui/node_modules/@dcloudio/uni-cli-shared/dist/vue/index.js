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
exports.isExternalUrl = exports.transformUniH5Jsx = void 0;
__exportStar(require("./transforms"), exports);
__exportStar(require("./utils"), exports);
__exportStar(require("./parse"), exports);
var babel_1 = require("./babel");
Object.defineProperty(exports, "transformUniH5Jsx", { enumerable: true, get: function () { return babel_1.transformUniH5Jsx; } });
var templateUtils_1 = require("./transforms/templateUtils");
Object.defineProperty(exports, "isExternalUrl", { enumerable: true, get: function () { return templateUtils_1.isExternalUrl; } });
