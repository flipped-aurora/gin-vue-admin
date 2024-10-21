"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppManifestPlugin = exports.getOutputManifestJson = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../utils");
let outputManifestJson = undefined;
function getOutputManifestJson() {
    return outputManifestJson;
}
exports.getOutputManifestJson = getOutputManifestJson;
function uniAppManifestPlugin() {
    const manifestJsonPath = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'manifest.json');
    const manifestJsonUTSPath = path_1.default.resolve(process.env.UNI_INPUT_DIR, uni_cli_shared_1.MANIFEST_JSON_UTS);
    let manifestJson = {};
    return {
        name: 'uni:app-manifest',
        apply: 'build',
        resolveId(id) {
            if ((0, utils_1.isManifest)(id)) {
                return manifestJsonUTSPath;
            }
        },
        load(id) {
            if ((0, utils_1.isManifest)(id)) {
                return fs_extra_1.default.readFileSync(manifestJsonPath, 'utf8');
            }
        },
        transform(code, id) {
            if ((0, utils_1.isManifest)(id)) {
                this.addWatchFile(path_1.default.resolve(process.env.UNI_INPUT_DIR, 'manifest.json'));
                manifestJson = (0, uni_cli_shared_1.parseJson)(code);
                return {
                    code: `export default ${JSON.stringify(manifestJson)}`,
                    map: {
                        mappings: '',
                    },
                };
            }
        },
        buildEnd() {
            outputManifestJson = (0, utils_1.normalizeManifestJson)(manifestJson);
            const manifest = outputManifestJson;
            if (process.env.NODE_ENV !== 'development') {
                // 生产模式，记录使用到的modules
                const ids = Array.from(this.getModuleIds());
                const uniExtApis = new Set();
                ids.forEach((id) => {
                    const moduleInfo = this.getModuleInfo(id);
                    if (moduleInfo &&
                        moduleInfo.meta &&
                        Array.isArray(moduleInfo.meta.uniExtApis)) {
                        moduleInfo.meta.uniExtApis.forEach((api) => {
                            uniExtApis.add(api);
                        });
                    }
                });
                if (uniExtApis.size) {
                    const modules = (0, uni_cli_shared_1.resolveUTSCompiler)().parseInjectModules([...uniExtApis], {}, [...(0, utils_1.getExtApiComponents)()]);
                    // 执行了摇树逻辑，就需要设置 modules 节点
                    (0, utils_1.updateManifestModules)(manifest, modules);
                }
            }
            fs_extra_1.default.outputFileSync(path_1.default.resolve(process.env.UNI_OUTPUT_DIR, 'manifest.json'), JSON.stringify(manifest, null, 2));
        },
    };
}
exports.uniAppManifestPlugin = uniAppManifestPlugin;
