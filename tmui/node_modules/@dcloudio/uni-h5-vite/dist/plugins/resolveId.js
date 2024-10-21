"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniResolveIdPlugin = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../utils");
const debugResolve = (0, debug_1.default)('uni:resolve');
function uniResolveIdPlugin() {
    const resolveCache = {};
    const mainPath = (0, uni_cli_shared_1.resolveMainPathOnce)(process.env.UNI_INPUT_DIR);
    return {
        name: 'uni:h5-resolve-id',
        enforce: 'pre',
        configResolved(config) {
            resolveCache[utils_1.ownerModuleName] = (0, uni_cli_shared_1.resolveBuiltIn)(path_1.default.join(utils_1.ownerModuleName, (process.env.UNI_APP_X === 'true' ? 'dist-x' : 'dist') +
                '/uni-h5.es.js'));
            resolveCache['@dcloudio/uni-h5-vue'] = (0, uni_cli_shared_1.resolveBuiltIn)(path_1.default.join('@dcloudio/uni-h5-vue', (process.env.UNI_APP_X === 'true' ? 'dist-x' : 'dist') +
                `/vue.runtime.${process.env.VITEST ? 'cjs' : 'esm'}.js`));
        },
        resolveId(id, importer, options) {
            if (id === '/main' && importer && importer.endsWith('index.html')) {
                return mainPath;
            }
            if (id === 'vue') {
                id = '@dcloudio/uni-h5-vue';
            }
            if ((0, utils_1.isSSR)(options)) {
                if (id === '@dcloudio/uni-h5-vue') {
                    return (0, uni_cli_shared_1.resolveBuiltIn)(path_1.default.join('@dcloudio/uni-h5-vue', (process.env.UNI_APP_X === 'true' ? 'dist-x' : 'dist') +
                        `/vue.runtime.cjs.js`));
                }
            }
            const cache = resolveCache[id];
            if (cache) {
                debugResolve('cache', id, cache);
                return cache;
            }
            if (id.startsWith('@dcloudio/uni-h5/style')) {
                return (resolveCache[id] = (0, uni_cli_shared_1.resolveBuiltIn)(id));
            }
            if (id.startsWith('@dcloudio/uni-components/style')) {
                return (resolveCache[id] = (0, uni_cli_shared_1.resolveBuiltIn)(id));
            }
        },
    };
}
exports.uniResolveIdPlugin = uniResolveIdPlugin;
