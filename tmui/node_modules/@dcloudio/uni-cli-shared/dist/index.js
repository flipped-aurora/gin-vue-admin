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
exports.checkUpdate = exports.M = exports.formatExtApiProviderName = exports.resolveEncryptUniModule = exports.getUniExtApiProviderRegisters = exports.parseUniModulesArtifacts = exports.parseInjects = exports.parseUniExtApis = exports.parseUniExtApi = void 0;
__exportStar(require("./fs"), exports);
__exportStar(require("./mp"), exports);
__exportStar(require("./url"), exports);
__exportStar(require("./env"), exports);
__exportStar(require("./hbx"), exports);
__exportStar(require("./ssr"), exports);
__exportStar(require("./vue"), exports);
__exportStar(require("./uts"), exports);
__exportStar(require("./logs"), exports);
__exportStar(require("./i18n"), exports);
__exportStar(require("./deps"), exports);
__exportStar(require("./json"), exports);
__exportStar(require("./vite"), exports);
__exportStar(require("./utils"), exports);
__exportStar(require("./easycom"), exports);
__exportStar(require("./constants"), exports);
__exportStar(require("./preprocess"), exports);
__exportStar(require("./postcss"), exports);
__exportStar(require("./filter"), exports);
__exportStar(require("./esbuild"), exports);
__exportStar(require("./resolve"), exports);
__exportStar(require("./scripts"), exports);
__exportStar(require("./platform"), exports);
__exportStar(require("./utsUtils"), exports);
var uni_modules_1 = require("./uni_modules");
Object.defineProperty(exports, "parseUniExtApi", { enumerable: true, get: function () { return uni_modules_1.parseUniExtApi; } });
Object.defineProperty(exports, "parseUniExtApis", { enumerable: true, get: function () { return uni_modules_1.parseUniExtApis; } });
Object.defineProperty(exports, "parseInjects", { enumerable: true, get: function () { return uni_modules_1.parseInjects; } });
Object.defineProperty(exports, "parseUniModulesArtifacts", { enumerable: true, get: function () { return uni_modules_1.parseUniModulesArtifacts; } });
Object.defineProperty(exports, "getUniExtApiProviderRegisters", { enumerable: true, get: function () { return uni_modules_1.getUniExtApiProviderRegisters; } });
Object.defineProperty(exports, "resolveEncryptUniModule", { enumerable: true, get: function () { return uni_modules_1.resolveEncryptUniModule; } });
Object.defineProperty(exports, "formatExtApiProviderName", { enumerable: true, get: function () { return uni_modules_1.formatExtApiProviderName; } });
var messages_1 = require("./messages");
Object.defineProperty(exports, "M", { enumerable: true, get: function () { return messages_1.M; } });
__exportStar(require("./exports"), exports);
var checkUpdate_1 = require("./checkUpdate");
Object.defineProperty(exports, "checkUpdate", { enumerable: true, get: function () { return checkUpdate_1.checkUpdate; } });
