"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniRenderjsPlugin = exports.getFiltersCache = void 0;
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugRenderjs = (0, debug_1.default)('uni:mp-renderjs');
const filtersCache = new Map();
function getFiltersCache(resolvedConfig) {
    return filtersCache.get(resolvedConfig) || [];
}
exports.getFiltersCache = getFiltersCache;
function uniRenderjsPlugin({ lang }) {
    let resolvedConfig;
    return {
        name: 'uni:mp-renderjs',
        configResolved(config) {
            resolvedConfig = config;
            filtersCache.set(resolvedConfig, []);
        },
        transform(code, id) {
            const { type, name } = (0, uni_cli_shared_1.parseRenderjs)(id);
            if (!type) {
                return null;
            }
            debugRenderjs(id);
            if (type !== lang) {
                return {
                    code: 'export default {}',
                    map: { mappings: '' },
                };
            }
            this.addWatchFile((0, uni_cli_shared_1.cleanUrl)(id));
            if (!name) {
                this.error((0, uni_cli_shared_1.missingModuleName)(type, code));
            }
            else {
                let cache = filtersCache.get(resolvedConfig);
                if (cache) {
                    const index = cache.findIndex((item) => item.id === id);
                    if (index > -1) {
                        cache.splice(index, 1);
                    }
                    cache.push({
                        id,
                        type,
                        name,
                        code,
                    });
                }
            }
            return {
                code: (0, uni_cli_shared_1.genWxsCallMethodsCode)(code),
                map: { mappings: '' },
            };
        },
    };
}
exports.uniRenderjsPlugin = uniRenderjsPlugin;
