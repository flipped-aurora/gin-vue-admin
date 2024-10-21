"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initPostcssPlugin = exports.uniPostcssScopedPlugin = exports.uniPostcssPlugin = void 0;
const stylePluginScoped_1 = __importDefault(require("./plugins/stylePluginScoped"));
exports.uniPostcssScopedPlugin = stylePluginScoped_1.default;
const uniapp_1 = __importDefault(require("./plugins/uniapp"));
exports.uniPostcssPlugin = uniapp_1.default;
function initPostcssPlugin({ uniApp, autoprefixer, } = {}) {
    const plugins = [(0, uniapp_1.default)(uniApp)];
    // nvue 不需要 autoprefixer
    if (autoprefixer !== false) {
        plugins.push(require('autoprefixer')(autoprefixer));
    }
    return plugins;
}
exports.initPostcssPlugin = initPostcssPlugin;
