"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniStaticPlugin = void 0;
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../../utils");
const debugStatic = (0, debug_1.default)('uni:static');
/**
 * 提供static等目录静态资源加载
 * @param _options
 * @param config
 * @returns
 */
function uniStaticPlugin(_options, config) {
    const filter = (0, utils_1.createPublicFileFilter)();
    return {
        name: 'uni:static',
        resolveId(id) {
            if (!config.assetsInclude((0, uni_cli_shared_1.cleanUrl)(id))) {
                return;
            }
            const publicFile = filter(id);
            if (publicFile) {
                debugStatic(id);
                return id;
            }
        },
        async load(id) {
            if (!config.assetsInclude((0, uni_cli_shared_1.cleanUrl)(id))) {
                return;
            }
            if (filter(id)) {
                return `export default ${JSON.stringify(fileToUrl(id, config))}`;
            }
        },
    };
}
exports.uniStaticPlugin = uniStaticPlugin;
function fileToUrl(id, config) {
    return config.base + id.slice(1);
}
