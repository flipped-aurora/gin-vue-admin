"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppVuePlugin = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const code_1 = require("@dcloudio/uni-cli-shared/dist/json/app/pages/code");
const renderjs_1 = require("../plugins/renderjs");
const configResolved_1 = require("../../plugin/configResolved");
const utils_1 = require("../../utils");
function uniAppVuePlugin() {
    const inputDir = process.env.UNI_INPUT_DIR;
    const mainPath = (0, uni_cli_shared_1.resolveMainPathOnce)(inputDir);
    let appCss = '';
    function normalizeCssChunkFilename(id) {
        return (0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizePath)(path_1.default.relative(inputDir, id))) + '.css';
    }
    return {
        name: 'uni:app-vue',
        config() {
            return {
                css: {
                    postcss: {
                        plugins: (0, uni_cli_shared_1.initPostcssPlugin)({
                            uniApp: (0, uni_cli_shared_1.parseRpx2UnitOnce)(inputDir, process.env.UNI_PLATFORM),
                        }),
                    },
                },
                build: {
                    rollupOptions: {
                        external: ['vue', '@vue/shared'],
                        output: {
                            name: 'AppService',
                            banner: uni_cli_shared_1.polyfillCode + code_1.restoreGlobalCode,
                            format: process.env.UNI_APP_CODE_SPLITING ? 'amd' : 'iife',
                            amd: {
                                autoId: true,
                            },
                            entryFileNames: uni_cli_shared_1.APP_SERVICE_FILENAME,
                            globals: {
                                vue: 'Vue',
                                '@vue/shared': 'uni.VueShared',
                            },
                        },
                    },
                },
            };
        },
        configResolved: (0, configResolved_1.createConfigResolved)({
            createCssPostPlugin(config) {
                return (0, uni_cli_shared_1.cssPostPlugin)(config, {
                    platform: 'app',
                    chunkCssFilename(id) {
                        if (id === mainPath) {
                            return 'app.css';
                        }
                        else if ((0, uni_cli_shared_1.isUniPageSfcFile)(id, inputDir)) {
                            return normalizeCssChunkFilename(id);
                        }
                        else if (process.env.UNI_COMPILE_TARGET === 'uni_modules' &&
                            (0, uni_cli_shared_1.isVueSfcFile)(id)) {
                            return normalizeCssChunkFilename(id);
                        }
                    },
                    chunkCssCode(filename, cssCode) {
                        if (filename === 'app.css') {
                            if (process.env.UNI_PLATFORM === 'app-harmony') {
                                appCss = fs_extra_1.default.readFileSync((0, uni_cli_shared_1.resolveBuiltIn)('@dcloudio/uni-app-harmony/dist/style.css'), 'utf8');
                            }
                            if (!appCss) {
                                appCss = fs_extra_1.default.readFileSync((0, uni_cli_shared_1.resolveBuiltIn)('@dcloudio/uni-app-plus/dist/style.css'), 'utf8');
                            }
                            return appCss + '\n' + cssCode;
                        }
                        return cssCode;
                    },
                });
            },
        }),
        generateBundle(_, bundle) {
            if (process.env.UNI_COMPILE_TARGET !== 'uni_modules') {
                this.emitFile({
                    fileName: '__uniappview.html',
                    source: genViewHtml(bundle),
                    type: 'asset',
                });
            }
        },
    };
}
exports.uniAppVuePlugin = uniAppVuePlugin;
function genViewHtml(bundle) {
    const viewHtmlStr = fs_extra_1.default.readFileSync(path_1.default.join(utils_1.templateDir, '__uniappview.html'), 'utf8');
    const { globalStyle } = (0, uni_cli_shared_1.parsePagesJsonOnce)(process.env.UNI_INPUT_DIR, process.env.UNI_PLATFORM);
    const { darkmode = false } = (0, uni_cli_shared_1.getPlatformManifestJsonOnce)();
    const __uniConfig = {
        globalStyle: {
            rpxCalcMaxDeviceWidth: globalStyle.rpxCalcMaxDeviceWidth,
            rpxCalcBaseDeviceWidth: globalStyle.rpxCalcBaseDeviceWidth,
        },
        darkmode,
    };
    const wxsCode = bundle[renderjs_1.APP_WXS_JS]
        ? `<script src="${renderjs_1.APP_WXS_JS}"></script>`
        : '';
    const renderjsCode = bundle[renderjs_1.APP_RENDERJS_JS]
        ? `<script src="${renderjs_1.APP_RENDERJS_JS}"></script>`
        : '';
    const automatorCode = process.env.UNI_AUTOMATOR_WS_ENDPOINT
        ? `<script src="__uniappautomator.js"></script>`
        : '';
    return viewHtmlStr
        .toString()
        .replace('<!--wxsCode-->', wxsCode)
        .replace('<!--renderjsCode-->', renderjsCode)
        .replace('<!--automatorCode-->', automatorCode)
        .replace('/*__uniConfig*/', `var __uniConfig = ${JSON.stringify(__uniConfig)}`);
}
