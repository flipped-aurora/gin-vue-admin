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
Object.defineProperty(exports, "__esModule", { value: true });
exports.init = void 0;
const path = __importStar(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const plugin_1 = require("./plugin");
const mainUTS_1 = require("./mainUTS");
const manifestJson_1 = require("./manifestJson");
const pagesJson_1 = require("./pagesJson");
const vueCompilerDom = __importStar(require("@vue/compiler-dom"));
const uniCliShared = __importStar(require("@dcloudio/uni-cli-shared"));
function init() {
    return [
        (0, uni_cli_shared_1.uniDecryptUniModulesPlugin)(),
        (0, uni_cli_shared_1.uniHBuilderXConsolePlugin)('uni.__log__'),
        (0, uni_cli_shared_1.uniUTSAppUniModulesPlugin)({
            x: true,
            isSingleThread: process.env.UNI_APP_X_SINGLE_THREAD !== 'false',
            extApis: (0, uni_cli_shared_1.parseUniExtApiNamespacesOnce)(process.env.UNI_UTS_PLATFORM, process.env.UNI_UTS_TARGET_LANGUAGE),
        }),
        (0, uni_cli_shared_1.uniEasycomPlugin)({ exclude: uni_cli_shared_1.UNI_EASYCOM_EXCLUDE }),
        (0, plugin_1.uniAppIOSPlugin)(),
        ...(process.env.UNI_COMPILE_TARGET === 'uni_modules'
            ? [(0, uni_cli_shared_1.uniEncryptUniModulesAssetsPlugin)(), (0, uni_cli_shared_1.uniEncryptUniModulesPlugin)()]
            : [(0, mainUTS_1.uniAppIOSMainPlugin)(), (0, manifestJson_1.uniAppManifestPlugin)(), (0, pagesJson_1.uniAppPagesPlugin)()]),
        (0, uni_cli_shared_1.uniUTSUVueJavaScriptPlugin)(),
        (0, uni_cli_shared_1.resolveUTSCompiler)().uts2js({
            inputDir: process.env.UNI_INPUT_DIR,
            version: process.env.UNI_COMPILER_VERSION,
            cacheRoot: path.resolve(process.env.UNI_APP_X_CACHE_DIR, '.uts2js/cache'),
            modules: {
                vueCompilerDom,
                uniCliShared,
            },
        }),
    ];
}
exports.init = init;
