"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.mergeMiniProgramAppJson = exports.parseMiniProgramPagesJson = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const json_1 = require("../json");
const pages_1 = require("../pages");
const utils_1 = require("./utils");
const utils_2 = require("../../utils");
const project_1 = require("./project");
const manifest_1 = require("../manifest");
const theme_1 = require("../theme");
function parseMiniProgramPagesJson(jsonStr, platform, options = { subpackages: false }) {
    return parsePagesJson(jsonStr, platform, options);
}
exports.parseMiniProgramPagesJson = parseMiniProgramPagesJson;
const NON_APP_JSON_KEYS = [
    'unipush',
    'secureNetwork',
    'usingComponents',
    'optimization',
    'scopedSlotsCompiler',
    'uniStatistics',
    'mergeVirtualHostAttributes',
    'styleIsolation',
];
function mergeMiniProgramAppJson(appJson, platformJson = {}) {
    Object.keys(platformJson).forEach((name) => {
        if (!(0, project_1.isMiniProgramProjectJsonKey)(name) &&
            !NON_APP_JSON_KEYS.includes(name)) {
            appJson[name] = platformJson[name];
        }
    });
}
exports.mergeMiniProgramAppJson = mergeMiniProgramAppJson;
function parsePagesJson(jsonStr, platform, { debug, darkmode, networkTimeout, subpackages, windowOptionsMap, tabBarOptionsMap, tabBarItemOptionsMap, } = {
    subpackages: false,
}) {
    let appJson = {
        pages: [],
    };
    let pageJsons = {};
    let nvuePages = [];
    // preprocess
    const pagesJson = (0, json_1.parseJson)(jsonStr, true);
    if (!pagesJson) {
        throw new Error(`[vite] Error: pages.json parse failed.\n`);
    }
    function addPageJson(pagePath, style) {
        const filename = path_1.default.join(process.env.UNI_INPUT_DIR, pagePath);
        if (fs_1.default.existsSync(filename + '.nvue') &&
            !fs_1.default.existsSync(filename + '.vue')) {
            nvuePages.push(pagePath);
        }
        const windowOptions = {};
        if (platform === 'mp-baidu') {
            // 仅百度小程序需要页面配置 component:true
            // 快手小程序反而不能配置 component:true，故不能统一添加，目前硬编码处理
            windowOptions.component = true;
        }
        pageJsons[pagePath] = (0, shared_1.extend)(windowOptions, (0, utils_1.parseWindowOptions)(style, platform, windowOptionsMap));
    }
    // pages
    (0, pages_1.validatePages)(pagesJson, jsonStr);
    pagesJson.pages.forEach((page) => {
        appJson.pages.push(page.path);
        addPageJson(page.path, page.style);
    });
    // subpackages
    pagesJson.subPackages = pagesJson.subPackages || pagesJson.subpackages;
    if (pagesJson.subPackages) {
        if (subpackages) {
            appJson.subPackages = pagesJson.subPackages.map(({ root, pages, ...rest }) => {
                return (0, shared_1.extend)({
                    root,
                    pages: pages.map((page) => {
                        addPageJson((0, utils_2.normalizePath)(path_1.default.join(root, page.path)), page.style);
                        return page.path;
                    }),
                }, rest);
            });
        }
        else {
            pagesJson.subPackages.forEach(({ root, pages }) => {
                pages.forEach((page) => {
                    const pagePath = (0, utils_2.normalizePath)(path_1.default.join(root, page.path));
                    appJson.pages.push(pagePath);
                    addPageJson(pagePath, page.style);
                });
            });
        }
    }
    // window
    if (pagesJson.globalStyle) {
        const windowOptions = (0, utils_1.parseWindowOptions)(pagesJson.globalStyle, platform, windowOptionsMap);
        const { usingComponents } = windowOptions;
        if (usingComponents) {
            delete windowOptions.usingComponents;
            appJson.usingComponents = usingComponents;
        }
        else {
            delete appJson.usingComponents;
        }
        appJson.window = windowOptions;
    }
    // tabBar
    if (pagesJson.tabBar) {
        const tabBar = (0, utils_1.parseTabBar)(pagesJson.tabBar, platform, tabBarOptionsMap, tabBarItemOptionsMap);
        if (tabBar) {
            appJson.tabBar = tabBar;
        }
    }
    (0, pages_1.filterPlatformPages)(platform, pagesJson);
    ['preloadRule', 'workers', 'plugins', 'entryPagePath'].forEach((name) => {
        if ((0, shared_1.hasOwn)(pagesJson, name)) {
            appJson[name] = pagesJson[name];
        }
    });
    if (debug) {
        appJson.debug = debug;
    }
    if (networkTimeout) {
        appJson.networkTimeout = networkTimeout;
    }
    const manifestJson = (0, manifest_1.getPlatformManifestJsonOnce)();
    if (!darkmode) {
        const { pages, window, tabBar } = (0, theme_1.initTheme)(manifestJson, appJson);
        (0, shared_1.extend)(appJson, JSON.parse(JSON.stringify({ pages, window, tabBar })));
        delete appJson.darkmode;
        delete appJson.themeLocation;
        pageJsons = (0, theme_1.initTheme)(manifestJson, pageJsons);
    }
    else {
        const themeLocation = manifestJson.themeLocation || 'theme.json';
        if ((0, theme_1.hasThemeJson)(path_1.default.join(process.env.UNI_INPUT_DIR, themeLocation)))
            appJson.themeLocation = themeLocation;
    }
    return {
        appJson,
        pageJsons,
        nvuePages,
    };
}
