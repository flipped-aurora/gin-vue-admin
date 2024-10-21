"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniUVuePlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const define_1 = require("../../config/define");
const resolve_1 = require("../../config/resolve");
const css_1 = require("../../config/css");
const configResolved_1 = require("../../configResolved");
function uniUVuePlugin(options) {
    return {
        name: 'uni:uvue',
        config(config, env) {
            options.command = env.command;
            let base = config.base;
            if (!base) {
                const manifestJson = (0, uni_cli_shared_1.parseManifestJsonOnce)(options.inputDir);
                const h5 = (0, uni_cli_shared_1.getPlatformManifestJson)(manifestJson, 'h5');
                base = (h5 && h5.router && h5.router.base) || '';
            }
            if (!base) {
                base = '/';
            }
            options.base = base;
            return {
                base: process.env.UNI_H5_BASE || base,
                root: process.env.VITE_ROOT_DIR,
                // TODO 临时设置为__static__,屏蔽警告：https://github.com/vitejs/vite/blob/824d042535033a5c3d7006978c0d05c201cd1c25/packages/vite/src/node/server/middlewares/transform.ts#L125
                publicDir: config.publicDir || '__static__',
                define: (0, define_1.createDefine)(options),
                resolve: (0, resolve_1.createResolve)(options, config),
                logLevel: config.logLevel || 'warn', // 默认使用 warn 等级，因为 info 等级vite:report 会输出文件列表等信息
                optimizeDeps: {
                    noDiscovery: true,
                    include: [],
                },
                css: (0, css_1.createCss)(options, config),
            };
        },
        configResolved(config) {
            (0, configResolved_1.initLogger)(config);
        },
    };
}
exports.uniUVuePlugin = uniUVuePlugin;
