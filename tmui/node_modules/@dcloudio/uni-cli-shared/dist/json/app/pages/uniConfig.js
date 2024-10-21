"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseEntryPagePath = exports.normalizeAppUniConfig = void 0;
const path_1 = __importDefault(require("path"));
const i18n_1 = require("../../../i18n");
const manifest_1 = require("../../manifest");
const manifest_2 = require("../manifest");
const arguments_1 = require("../manifest/arguments");
const splashscreen_1 = require("../manifest/splashscreen");
const theme_1 = require("../../theme");
function normalizeAppUniConfig(pagesJson, manifestJson) {
    const { autoclose, alwaysShowBeforeRender } = (0, splashscreen_1.getSplashscreen)(manifestJson);
    const platformConfig = (0, manifest_1.getPlatformManifestJsonOnce)();
    const config = {
        pages: [],
        globalStyle: pagesJson.globalStyle,
        nvue: {
            compiler: (0, manifest_2.getNVueCompiler)(manifestJson),
            styleCompiler: (0, manifest_2.getNVueStyleCompiler)(manifestJson),
            'flex-direction': (0, manifest_2.getNVueFlexDirection)(manifestJson),
        },
        renderer: manifestJson['app-plus']?.renderer === 'native' ? 'native' : 'auto',
        appname: manifestJson.name || '',
        splashscreen: {
            alwaysShowBeforeRender,
            autoclose,
        },
        compilerVersion: process.env.UNI_COMPILER_VERSION,
        ...parseEntryPagePath(pagesJson),
        networkTimeout: (0, manifest_1.normalizeNetworkTimeout)(manifestJson.networkTimeout),
        tabBar: pagesJson.tabBar,
        fallbackLocale: manifestJson.fallbackLocale,
        locales: (0, i18n_1.initLocales)(path_1.default.join(process.env.UNI_INPUT_DIR, 'locale')),
        darkmode: platformConfig.darkmode || false,
        themeConfig: (0, theme_1.normalizeThemeConfigOnce)(platformConfig),
    };
    // TODO 待支持分包
    return JSON.stringify(config);
}
exports.normalizeAppUniConfig = normalizeAppUniConfig;
function parseEntryPagePath(pagesJson) {
    const res = {
        entryPagePath: '',
        entryPageQuery: '',
        realEntryPagePath: '',
    };
    if (!pagesJson.pages.length) {
        return res;
    }
    res.entryPagePath = pagesJson.pages[0].path;
    const argsJsonStr = (0, arguments_1.parseArguments)(pagesJson);
    if (argsJsonStr) {
        try {
            const args = JSON.parse(argsJsonStr);
            const entryPagePath = args.path || args.pathName;
            const realEntryPagePath = res.entryPagePath;
            if (entryPagePath && realEntryPagePath !== entryPagePath) {
                res.entryPagePath = entryPagePath;
                res.entryPageQuery = args.query ? '?' + args.query : '';
                // non tabBar page
                if (!(pagesJson.tabBar?.list || []).find((page) => page.pagePath === entryPagePath)) {
                    res.realEntryPagePath = realEntryPagePath;
                }
            }
        }
        catch (e) { }
    }
    return res;
}
exports.parseEntryPagePath = parseEntryPagePath;
