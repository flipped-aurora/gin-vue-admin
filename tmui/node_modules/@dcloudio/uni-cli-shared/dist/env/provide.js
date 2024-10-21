"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initH5Provide = exports.initAppProvide = void 0;
const path_1 = __importDefault(require("path"));
const libDir = path_1.default.resolve(__dirname, '../../lib');
function initAppProvide() {
    const cryptoDefine = [path_1.default.join(libDir, 'crypto.js'), 'default'];
    return {
        __f__: ['@dcloudio/uni-app', 'formatAppLog'],
        crypto: cryptoDefine,
        'window.crypto': cryptoDefine,
        'global.crypto': cryptoDefine,
        'uni.getCurrentSubNVue': ['@dcloudio/uni-app', 'getCurrentSubNVue'],
        'uni.requireNativePlugin': ['@dcloudio/uni-app', 'requireNativePlugin'],
    };
}
exports.initAppProvide = initAppProvide;
function initH5Provide() {
    return {
        __f__: ['@dcloudio/uni-app', 'formatH5Log'],
    };
}
exports.initH5Provide = initH5Provide;
