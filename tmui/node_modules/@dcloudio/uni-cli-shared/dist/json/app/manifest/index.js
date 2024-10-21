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
exports.parseArguments = exports.getNVueFlexDirection = exports.getNVueStyleCompiler = exports.getNVueCompiler = exports.hasConfusionFile = exports.isConfusionFile = exports.APP_CONFUSION_FILENAME = exports.normalizeAppManifestJson = void 0;
const shared_1 = require("@vue/shared");
const merge_1 = require("./merge");
const defaultManifestJson_1 = require("./defaultManifestJson");
const statusbar_1 = require("./statusbar");
const plus_1 = require("./plus");
const nvue_1 = require("./nvue");
const arguments_1 = require("./arguments");
const safearea_1 = require("./safearea");
const splashscreen_1 = require("./splashscreen");
const confusion_1 = require("./confusion");
const uniApp_1 = require("./uniApp");
const launchwebview_1 = require("./launchwebview");
const checksystemwebview_1 = require("./checksystemwebview");
const tabBar_1 = require("./tabBar");
const i18n_1 = require("./i18n");
const theme_1 = require("../../theme");
function normalizeAppManifestJson(userManifestJson, pagesJson) {
    const manifestJson = (0, merge_1.initRecursiveMerge)((0, defaultManifestJson_1.initDefaultManifestJson)(), userManifestJson);
    const { pages, globalStyle, tabBar } = (0, theme_1.initTheme)(manifestJson, pagesJson);
    (0, shared_1.extend)(pagesJson, JSON.parse(JSON.stringify({ pages, globalStyle, tabBar })));
    (0, statusbar_1.initAppStatusbar)(manifestJson, pagesJson);
    (0, arguments_1.initArguments)(manifestJson, pagesJson);
    (0, plus_1.initPlus)(manifestJson, pagesJson);
    (0, nvue_1.initNVue)(manifestJson, pagesJson);
    (0, safearea_1.initSafearea)(manifestJson, pagesJson);
    (0, splashscreen_1.initSplashscreen)(manifestJson, userManifestJson);
    (0, confusion_1.initConfusion)(manifestJson);
    (0, uniApp_1.initUniApp)(manifestJson);
    // 依赖 initArguments 先执行
    (0, tabBar_1.initTabBar)((0, launchwebview_1.initLaunchwebview)(manifestJson, pagesJson), manifestJson, pagesJson);
    // 依赖 initUniApp 先执行
    (0, checksystemwebview_1.initCheckSystemWebview)(manifestJson);
    return (0, i18n_1.initI18n)(manifestJson);
}
exports.normalizeAppManifestJson = normalizeAppManifestJson;
__exportStar(require("./env"), exports);
var confusion_2 = require("./confusion");
Object.defineProperty(exports, "APP_CONFUSION_FILENAME", { enumerable: true, get: function () { return confusion_2.APP_CONFUSION_FILENAME; } });
Object.defineProperty(exports, "isConfusionFile", { enumerable: true, get: function () { return confusion_2.isConfusionFile; } });
Object.defineProperty(exports, "hasConfusionFile", { enumerable: true, get: function () { return confusion_2.hasConfusionFile; } });
var nvue_2 = require("./nvue");
Object.defineProperty(exports, "getNVueCompiler", { enumerable: true, get: function () { return nvue_2.getNVueCompiler; } });
Object.defineProperty(exports, "getNVueStyleCompiler", { enumerable: true, get: function () { return nvue_2.getNVueStyleCompiler; } });
Object.defineProperty(exports, "getNVueFlexDirection", { enumerable: true, get: function () { return nvue_2.getNVueFlexDirection; } });
var arguments_2 = require("./arguments");
Object.defineProperty(exports, "parseArguments", { enumerable: true, get: function () { return arguments_2.parseArguments; } });
