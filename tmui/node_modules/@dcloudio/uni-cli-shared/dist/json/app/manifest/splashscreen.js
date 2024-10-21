"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getSplashscreen = exports.initSplashscreen = void 0;
const shared_1 = require("@vue/shared");
function initSplashscreen(manifestJson, userManifestJson) {
    if (!manifestJson.plus.splashscreen) {
        return;
    }
    // 强制白屏检测
    const splashscreenOptions = userManifestJson['app-plus'] && userManifestJson['app-plus'].splashscreen;
    const hasAlwaysShowBeforeRender = splashscreenOptions && (0, shared_1.hasOwn)(splashscreenOptions, 'alwaysShowBeforeRender');
    if (!hasAlwaysShowBeforeRender &&
        manifestJson.plus.splashscreen.autoclose === false) {
        // 兼容旧版本仅配置了 autoclose 为 false
        manifestJson.plus.splashscreen.alwaysShowBeforeRender = false;
    }
    if (manifestJson.plus.splashscreen.alwaysShowBeforeRender) {
        // 白屏检测
        if (!manifestJson.plus.splashscreen.target) {
            manifestJson.plus.splashscreen.target = 'id:1';
        }
        manifestJson.plus.splashscreen.autoclose = true;
        manifestJson.plus.splashscreen.delay = 0;
    }
    else {
        // 不启用白屏检测
        delete manifestJson.plus.splashscreen.target;
        if (manifestJson.plus.splashscreen.autoclose) {
            // 启用 uni-app 框架关闭 splash
            manifestJson.plus.splashscreen.autoclose = false; // 原 5+ autoclose 改为 false
        }
    }
    delete manifestJson.plus.splashscreen.alwaysShowBeforeRender;
}
exports.initSplashscreen = initSplashscreen;
function getSplashscreen(manifestJson) {
    const splashscreenOptions = manifestJson['app-plus']?.splashscreen || {};
    return {
        autoclose: splashscreenOptions.autoclose !== false,
        alwaysShowBeforeRender: splashscreenOptions.alwaysShowBeforeRender !== false,
    };
}
exports.getSplashscreen = getSplashscreen;
