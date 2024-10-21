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
exports.updateMiniProgramComponentExternalClasses = exports.findMiniProgramComponentExternalClasses = exports.parseExternalClasses = exports.hasExternalClasses = exports.updateMiniProgramComponentsByTemplateFilename = exports.updateMiniProgramComponentsByScriptFilename = exports.updateMiniProgramComponentsByMainFilename = exports.updateMiniProgramGlobalComponents = exports.transformDynamicImports = exports.parseTemplateDescriptor = exports.parseScriptDescriptor = exports.parseMainDescriptor = exports.copyMiniProgramThemeJson = exports.copyMiniProgramPluginJson = exports.HTML_TO_MINI_PROGRAM_TAGS = void 0;
__exportStar(require("./ast"), exports);
__exportStar(require("./wxs"), exports);
__exportStar(require("./nvue"), exports);
__exportStar(require("./event"), exports);
__exportStar(require("./style"), exports);
__exportStar(require("./assets"), exports);
__exportStar(require("./template"), exports);
__exportStar(require("./constants"), exports);
var tags_1 = require("./tags");
Object.defineProperty(exports, "HTML_TO_MINI_PROGRAM_TAGS", { enumerable: true, get: function () { return tags_1.HTML_TO_MINI_PROGRAM_TAGS; } });
var plugin_1 = require("./plugin");
Object.defineProperty(exports, "copyMiniProgramPluginJson", { enumerable: true, get: function () { return plugin_1.copyMiniProgramPluginJson; } });
Object.defineProperty(exports, "copyMiniProgramThemeJson", { enumerable: true, get: function () { return plugin_1.copyMiniProgramThemeJson; } });
var usingComponents_1 = require("./usingComponents");
Object.defineProperty(exports, "parseMainDescriptor", { enumerable: true, get: function () { return usingComponents_1.parseMainDescriptor; } });
Object.defineProperty(exports, "parseScriptDescriptor", { enumerable: true, get: function () { return usingComponents_1.parseScriptDescriptor; } });
Object.defineProperty(exports, "parseTemplateDescriptor", { enumerable: true, get: function () { return usingComponents_1.parseTemplateDescriptor; } });
Object.defineProperty(exports, "transformDynamicImports", { enumerable: true, get: function () { return usingComponents_1.transformDynamicImports; } });
Object.defineProperty(exports, "updateMiniProgramGlobalComponents", { enumerable: true, get: function () { return usingComponents_1.updateMiniProgramGlobalComponents; } });
Object.defineProperty(exports, "updateMiniProgramComponentsByMainFilename", { enumerable: true, get: function () { return usingComponents_1.updateMiniProgramComponentsByMainFilename; } });
Object.defineProperty(exports, "updateMiniProgramComponentsByScriptFilename", { enumerable: true, get: function () { return usingComponents_1.updateMiniProgramComponentsByScriptFilename; } });
Object.defineProperty(exports, "updateMiniProgramComponentsByTemplateFilename", { enumerable: true, get: function () { return usingComponents_1.updateMiniProgramComponentsByTemplateFilename; } });
var externalClasses_1 = require("./externalClasses");
Object.defineProperty(exports, "hasExternalClasses", { enumerable: true, get: function () { return externalClasses_1.hasExternalClasses; } });
Object.defineProperty(exports, "parseExternalClasses", { enumerable: true, get: function () { return externalClasses_1.parseExternalClasses; } });
Object.defineProperty(exports, "findMiniProgramComponentExternalClasses", { enumerable: true, get: function () { return externalClasses_1.findMiniProgramComponentExternalClasses; } });
Object.defineProperty(exports, "updateMiniProgramComponentExternalClasses", { enumerable: true, get: function () { return externalClasses_1.updateMiniProgramComponentExternalClasses; } });
