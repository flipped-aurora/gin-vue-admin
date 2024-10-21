"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniConsolePlugin = void 0;
const debug_1 = __importDefault(require("debug"));
const pluginutils_1 = require("@rollup/pluginutils");
const utils_1 = require("../utils");
const console_1 = require("../../logs/console");
const utils_2 = require("../../vite/utils/utils");
const debugConsole = (0, debug_1.default)('uni:console');
function uniConsolePlugin(options) {
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    let resolvedConfig;
    return {
        name: 'uni:console',
        enforce: 'pre',
        configResolved(config) {
            resolvedConfig = config;
        },
        transform(code, id) {
            if (!filter(id))
                return null;
            if (!(0, utils_1.isJsFile)(id))
                return null;
            let { filename } = (0, utils_1.parseVueRequest)(id);
            if (options.filename) {
                filename = options.filename(filename);
            }
            if (!filename) {
                return null;
            }
            if (!code.includes('console.')) {
                return null;
            }
            debugConsole(id);
            return (0, console_1.rewriteConsoleExpr)(options.method, id, filename, code, (0, utils_2.withSourcemap)(resolvedConfig));
        },
    };
}
exports.uniConsolePlugin = uniConsolePlugin;
