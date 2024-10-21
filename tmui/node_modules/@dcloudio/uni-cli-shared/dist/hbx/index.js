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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniHBuilderXConsolePlugin = exports.formatInstallHBuilderXPluginTips = exports.installHBuilderXPlugin = exports.initModuleAlias = exports.formatAtFilename = void 0;
const path_1 = __importDefault(require("path"));
const utils_1 = require("../utils");
const console_1 = require("../vite/plugins/console");
var log_1 = require("./log");
Object.defineProperty(exports, "formatAtFilename", { enumerable: true, get: function () { return log_1.formatAtFilename; } });
__exportStar(require("./env"), exports);
var alias_1 = require("./alias");
Object.defineProperty(exports, "initModuleAlias", { enumerable: true, get: function () { return alias_1.initModuleAlias; } });
Object.defineProperty(exports, "installHBuilderXPlugin", { enumerable: true, get: function () { return alias_1.installHBuilderXPlugin; } });
Object.defineProperty(exports, "formatInstallHBuilderXPluginTips", { enumerable: true, get: function () { return alias_1.formatInstallHBuilderXPluginTips; } });
function uniHBuilderXConsolePlugin(method = '__f__') {
    return (0, console_1.uniConsolePlugin)({
        method,
        filename(filename) {
            filename = path_1.default.relative(process.env.UNI_INPUT_DIR, filename);
            if (filename.startsWith('.') || path_1.default.isAbsolute(filename)) {
                return '';
            }
            return (0, utils_1.normalizePath)(filename);
        },
    });
}
exports.uniHBuilderXConsolePlugin = uniHBuilderXConsolePlugin;
