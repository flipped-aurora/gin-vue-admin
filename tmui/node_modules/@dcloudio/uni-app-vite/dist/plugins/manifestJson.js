"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniManifestJsonPlugin = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniManifestJsonPlugin() {
    return (0, uni_cli_shared_1.defineUniManifestJsonPlugin)((opts) => {
        const inputDir = process.env.UNI_INPUT_DIR;
        return {
            name: 'uni:app-manifest-json',
            enforce: 'pre',
            transform(code, id) {
                if (!opts.filter(id)) {
                    return;
                }
                this.addWatchFile(path_1.default.resolve(inputDir, 'manifest.json'));
                (0, uni_cli_shared_1.getLocaleFiles)(path_1.default.resolve(inputDir, 'locale')).forEach((filepath) => {
                    this.addWatchFile(filepath);
                });
                const manifestJson = (0, uni_cli_shared_1.normalizeAppManifestJson)((0, uni_cli_shared_1.parseJson)(code), (0, uni_cli_shared_1.parsePagesJsonOnce)(inputDir, process.env.UNI_PLATFORM));
                // 生成一个空的 app-config.js，兼容基座已有规范
                this.emitFile({
                    fileName: uni_cli_shared_1.APP_CONFIG,
                    type: 'asset',
                    source: '(function(){})();',
                });
                this.emitFile({
                    fileName: `manifest.json`,
                    type: 'asset',
                    source: JSON.stringify(manifestJson, null, 2),
                });
                return {
                    code: '',
                    map: null,
                };
            },
        };
    });
}
exports.uniManifestJsonPlugin = uniManifestJsonPlugin;
