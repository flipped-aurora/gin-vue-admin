"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.init = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const pre_1 = require("./pre");
const plugin_1 = require("./plugin");
const css_1 = require("./css");
const mainUTS_1 = require("./mainUTS");
const manifestJson_1 = require("./manifestJson");
const pagesJson_1 = require("./pagesJson");
const uvue_1 = require("./uvue");
const unicloud_1 = require("./unicloud");
function init() {
    return [
        (0, uni_cli_shared_1.uniDecryptUniModulesPlugin)(),
        (0, pre_1.uniPrePlugin)(),
        ...(process.env.UNI_COMPILE_TARGET === 'uni_modules'
            ? []
            : [
                (0, uni_cli_shared_1.uniUTSAppUniModulesPlugin)({
                    x: true,
                    isSingleThread: process.env.UNI_APP_X_SINGLE_THREAD !== 'false',
                    extApis: (0, uni_cli_shared_1.parseUniExtApiNamespacesOnce)(process.env.UNI_UTS_PLATFORM, process.env.UNI_UTS_TARGET_LANGUAGE),
                }),
            ]),
        (0, plugin_1.uniAppPlugin)(),
        ...(process.env.UNI_COMPILE_TARGET === 'uni_modules'
            ? [(0, uni_cli_shared_1.uniEncryptUniModulesPlugin)()]
            : [
                // 需要放到 uniAppPlugin 之后(TSC模式无需)，让 uniAppPlugin 先 emit 出真实的 main.uts，然后 MainPlugin 再返回仅包含 import 的 js code
                (0, mainUTS_1.uniAppMainPlugin)(),
                (0, manifestJson_1.uniAppManifestPlugin)(),
                (0, pagesJson_1.uniAppPagesPlugin)(),
            ]),
        (0, css_1.uniAppCssPlugin)(),
        // 解决所有的src引入
        (0, uni_cli_shared_1.uniViteSfcSrcImportPlugin)({ onlyVue: false }),
        (0, uvue_1.uniAppUVuePlugin)(),
        (0, unicloud_1.uniCloudPlugin)(),
    ];
}
exports.init = init;
