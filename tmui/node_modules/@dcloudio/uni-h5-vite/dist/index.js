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
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const plugin_1 = require("./plugin");
const css_1 = require("./plugins/css");
const easycom_1 = require("./plugins/easycom");
const inject_1 = require("./plugins/inject");
const mainJs_1 = require("./plugins/mainJs");
const manifestJson_1 = require("./plugins/manifestJson");
const pagesJson_1 = require("./plugins/pagesJson");
const postVue_1 = require("./plugins/postVue");
const renderjs_1 = require("./plugins/renderjs");
const resolveId_1 = require("./plugins/resolveId");
const setup_1 = require("./plugins/setup");
const ssr_1 = require("./plugins/ssr");
const sourcemap_1 = require("./plugins/sourcemap");
const vueCompilerDom = __importStar(require("@vue/compiler-dom"));
const uniCliShared = __importStar(require("@dcloudio/uni-cli-shared"));
exports.default = [
    ...(process.env.UNI_APP_X === 'true'
        ? [
            (0, uni_cli_shared_1.uniDecryptUniModulesPlugin)(),
            (0, uni_cli_shared_1.uniUTSUVueJavaScriptPlugin)(),
            (0, uni_cli_shared_1.resolveUTSCompiler)().uts2js({
                inputDir: process.env.UNI_INPUT_DIR,
                version: process.env.UNI_COMPILER_VERSION,
                cacheRoot: path_1.default.resolve(process.env.UNI_APP_X_CACHE_DIR, '.uts2js/cache'),
                modules: {
                    vueCompilerDom,
                    uniCliShared,
                },
            }),
        ]
        : []),
    (0, easycom_1.uniEasycomPlugin)({ exclude: uni_cli_shared_1.UNI_EASYCOM_EXCLUDE }),
    (0, uni_cli_shared_1.uniCssScopedPlugin)({
        filter: (id) => (0, uni_cli_shared_1.isVueSfcFile)(id) && !(0, uni_cli_shared_1.isAppVue)(id),
    }),
    (0, resolveId_1.uniResolveIdPlugin)(),
    ...(process.env.UNI_COMPILE_TARGET === 'uni_modules'
        ? []
        : [(0, mainJs_1.uniMainJsPlugin)(), (0, manifestJson_1.uniManifestJsonPlugin)(), (0, pagesJson_1.uniPagesJsonPlugin)()]),
    (0, inject_1.uniInjectPlugin)(),
    (0, css_1.uniCssPlugin)(),
    (0, ssr_1.uniSSRPlugin)(),
    (0, setup_1.uniSetupPlugin)(),
    (0, renderjs_1.uniRenderjsPlugin)(),
    (0, plugin_1.uniH5Plugin)(),
    ...(process.env.UNI_COMPILE_TARGET === 'uni_modules'
        ? [(0, uni_cli_shared_1.uniEncryptUniModulesAssetsPlugin)(), (0, uni_cli_shared_1.uniEncryptUniModulesPlugin)()]
        : []),
    (0, postVue_1.uniPostVuePlugin)(),
    (0, sourcemap_1.uniPostSourceMapPlugin)(),
];
