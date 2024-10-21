"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPagesJsonPlugin = exports.getNVueCssPaths = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const entry_1 = require("./entry");
const uni_i18n_1 = require("@dcloudio/uni-i18n");
const debugPagesJson = (0, debug_1.default)('uni:pages-json');
const nvueCssPathsCache = new Map();
function getNVueCssPaths(config) {
    return nvueCssPathsCache.get(config);
}
exports.getNVueCssPaths = getNVueCssPaths;
function uniPagesJsonPlugin(options) {
    let resolvedConfig;
    const platform = process.env.UNI_PLATFORM;
    const inputDir = process.env.UNI_INPUT_DIR;
    return (0, uni_cli_shared_1.defineUniPagesJsonPlugin)((opts) => {
        return {
            name: 'uni:mp-pages-json',
            enforce: 'pre',
            configResolved(config) {
                resolvedConfig = config;
            },
            transform(code, id) {
                if (!opts.filter(id)) {
                    return null;
                }
                this.addWatchFile(path_1.default.resolve(inputDir, 'pages.json'));
                (0, uni_cli_shared_1.getLocaleFiles)(path_1.default.resolve(inputDir, 'locale')).forEach((filepath) => {
                    this.addWatchFile(filepath);
                });
                const manifestJson = (0, uni_cli_shared_1.parseManifestJsonOnce)(inputDir);
                const { appJson, pageJsons, nvuePages } = (0, uni_cli_shared_1.parseMiniProgramPagesJson)(code, platform, {
                    debug: !!manifestJson.debug,
                    darkmode: options.app.darkmode,
                    networkTimeout: manifestJson.networkTimeout,
                    subpackages: !!options.app.subpackages,
                    ...options.json,
                });
                nvueCssPathsCache.set(resolvedConfig, nvuePages.map((pagePath) => pagePath + options.style.extname));
                (0, uni_cli_shared_1.mergeMiniProgramAppJson)(appJson, manifestJson[platform]);
                if (options.json?.formatAppJson) {
                    options.json.formatAppJson(appJson, manifestJson, pageJsons);
                }
                // 使用 once 获取的话，可以节省编译时间，但 i18n 内容发生变化时，pages.json 不会自动更新
                const i18nOptions = (0, uni_cli_shared_1.initI18nOptionsOnce)(platform, inputDir, false, true);
                if (i18nOptions) {
                    const { locale, locales, delimiters } = i18nOptions;
                    (0, uni_i18n_1.parseI18nJson)(appJson, locales[locale], delimiters);
                    (0, uni_i18n_1.parseI18nJson)(pageJsons, locales[locale], delimiters);
                }
                const { normalize } = options.app;
                (0, uni_cli_shared_1.addMiniProgramAppJson)(normalize ? normalize(appJson) : appJson);
                Object.keys(pageJsons).forEach((name) => {
                    if (isNormalPage(name)) {
                        (0, uni_cli_shared_1.addMiniProgramPageJson)(name, pageJsons[name]);
                    }
                });
                return {
                    code: `import './${uni_cli_shared_1.MANIFEST_JSON_JS}'\n` + importPagesCode(appJson),
                    map: { mappings: '' },
                };
            },
            generateBundle() {
                (0, uni_cli_shared_1.findChangedJsonFiles)(options.app.usingComponents).forEach((value, key) => {
                    debugPagesJson('json.changed', key);
                    this.emitFile({
                        type: 'asset',
                        fileName: key + '.json',
                        source: value,
                    });
                });
            },
        };
    });
}
exports.uniPagesJsonPlugin = uniPagesJsonPlugin;
/**
 * 字节跳动小程序可以配置 ext:// 开头的插件页面模板，如 ext://microapp-trade-plugin/order-confirm
 * @param pagePath
 * @returns
 */
function isNormalPage(pagePath) {
    return !pagePath.startsWith('ext://');
}
function importPagesCode(pagesJson) {
    const importPagesCode = [];
    function importPageCode(pagePath) {
        if (!isNormalPage(pagePath)) {
            return;
        }
        const pagePathWithExtname = (0, uni_cli_shared_1.normalizePagePath)(pagePath, process.env.UNI_PLATFORM);
        if (pagePathWithExtname) {
            importPagesCode.push(`import('${(0, entry_1.virtualPagePath)(pagePathWithExtname)}')`);
        }
    }
    pagesJson.pages.forEach((pagePath) => importPageCode(pagePath));
    if (pagesJson.subPackages) {
        pagesJson.subPackages.forEach(({ root, pages }) => {
            pages &&
                pages.forEach((pagePath) => importPageCode(path_1.default.join(root, pagePath)));
        });
    }
    return `if(!Math){
${importPagesCode.join('\n')}
}`;
}
