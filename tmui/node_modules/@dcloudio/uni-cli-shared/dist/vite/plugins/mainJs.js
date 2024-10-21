"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineUniMainJsPlugin = void 0;
const path_1 = __importDefault(require("path"));
const utils_1 = require("../../utils");
function defineUniMainJsPlugin(createUniMainJsPlugin) {
    const opts = {
        resolvedConfig: {},
        filter(id) {
            return id === mainJsPath || id === mainTsPath || id === mainUTsPath;
        },
    };
    const plugin = createUniMainJsPlugin(opts);
    const origConfigResolved = plugin.configResolved;
    let mainJsPath = '';
    let mainTsPath = '';
    let mainUTsPath = '';
    plugin.configResolved = function (config) {
        opts.resolvedConfig = config;
        const mainPath = (0, utils_1.normalizePath)(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'main'));
        mainJsPath = mainPath + '.js';
        mainTsPath = mainPath + '.ts';
        mainUTsPath = mainPath + '.uts';
        return origConfigResolved && origConfigResolved(config);
    };
    return plugin;
}
exports.defineUniMainJsPlugin = defineUniMainJsPlugin;
