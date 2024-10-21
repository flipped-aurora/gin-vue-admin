"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPagesJsonPlugin = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const esbuild_1 = require("../../nvue/plugins/esbuild");
function uniPagesJsonPlugin() {
    return (0, uni_cli_shared_1.defineUniPagesJsonPlugin)((opts) => {
        return {
            name: 'uni:app-vue-pages-json',
            enforce: 'pre',
            transform(code, id) {
                if (!opts.filter(id)) {
                    return;
                }
                this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'pages.json'));
                (0, uni_cli_shared_1.getLocaleFiles)(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'locale')).forEach((filepath) => {
                    this.addWatchFile(filepath);
                });
                const pagesJson = (0, uni_cli_shared_1.normalizePagesJson)(code, process.env.UNI_PLATFORM);
                pagesJson.pages.forEach((page) => {
                    if (!page.style.isNVue) {
                        this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, page.path + '.vue'));
                    }
                });
                this.emitFile({
                    fileName: uni_cli_shared_1.APP_CONFIG_SERVICE,
                    type: 'asset',
                    source: (0, uni_cli_shared_1.normalizeAppConfigService)(pagesJson, (0, uni_cli_shared_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR)),
                });
                return {
                    code: `import './${uni_cli_shared_1.MANIFEST_JSON_JS}'\n` +
                        (0, uni_cli_shared_1.normalizeAppPagesJson)(pagesJson, process.env.UNI_PLATFORM),
                    map: { mappings: '' },
                };
            },
            generateBundle(_, bundle) {
                const outputFile = bundle[uni_cli_shared_1.APP_CONFIG_SERVICE];
                if (outputFile && outputFile.type === 'asset') {
                    // 补充 nvue styles
                    ;
                    outputFile.source = (0, esbuild_1.wrapperNVueAppStyles)(outputFile.source);
                }
            },
        };
    });
}
exports.uniPagesJsonPlugin = uniPagesJsonPlugin;
