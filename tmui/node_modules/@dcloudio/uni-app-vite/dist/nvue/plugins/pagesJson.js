"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPagesJsonPlugin = exports.parseNVuePageOptions = exports.nvuePagesCache = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
exports.nvuePagesCache = new Map();
// 在 @vue/compiler-sfc@3.2.47 执行前重写 @vue/compiler-dom compile 方法
const nvuePages = {};
function parseNVuePageOptions(filename) {
    return nvuePages[filename];
}
exports.parseNVuePageOptions = parseNVuePageOptions;
function uniPagesJsonPlugin({ renderer, appService, }) {
    return (0, uni_cli_shared_1.defineUniPagesJsonPlugin)((opts) => {
        return {
            name: 'uni:app-nvue-pages-json',
            enforce: 'pre',
            configResolved(config) {
                exports.nvuePagesCache.set(config, nvuePages);
            },
            transform(code, id) {
                if (!opts.filter(id)) {
                    return;
                }
                this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'pages.json'));
                (0, uni_cli_shared_1.getLocaleFiles)(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'locale')).forEach((filepath) => {
                    this.addWatchFile(filepath);
                });
                const pagesJson = (0, uni_cli_shared_1.normalizePagesJson)(code, process.env.UNI_PLATFORM);
                Object.keys(nvuePages).forEach((name) => {
                    delete nvuePages[name];
                });
                pagesJson.pages.forEach((page) => {
                    if (page.style.isNVue) {
                        const filename = (0, uni_cli_shared_1.normalizePath)(path_1.default.resolve(process.env.UNI_INPUT_DIR, page.path + '.nvue'));
                        nvuePages[filename] = {
                            disableScroll: page.style.disableScroll,
                            scrollIndicator: page.style.scrollIndicator,
                        };
                        // fix: question/164673
                        // this.addWatchFile(filename)
                    }
                });
                if (renderer === 'native' && appService) {
                    this.emitFile({
                        fileName: uni_cli_shared_1.APP_CONFIG_SERVICE,
                        type: 'asset',
                        source: (0, uni_cli_shared_1.normalizeAppConfigService)(pagesJson, (0, uni_cli_shared_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR)),
                    });
                    return {
                        code: '',
                        map: { mappings: '' },
                    };
                }
                return {
                    code: (0, uni_cli_shared_1.normalizeAppNVuePagesJson)(pagesJson),
                    map: { mappings: '' },
                };
            },
        };
    });
}
exports.uniPagesJsonPlugin = uniPagesJsonPlugin;
