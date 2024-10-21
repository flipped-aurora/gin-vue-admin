"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppPlugin = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_1 = require("./uni");
const build_1 = require("./build");
function uniAppPlugin({ renderer, appService, } = {
    appService: false,
}) {
    return {
        name: 'uni:app',
        uni: (0, uni_1.uniOptions)(),
        config(config, env) {
            return {
                base: '/', // app 平台强制 base
                build: (0, build_1.buildOptions)({ renderer, appService }, config, env),
                optimizeDeps: {
                    noDiscovery: true,
                    include: [],
                },
                resolve: {
                    alias: {
                        // vue-i18n 默认会启用 new Function 来构造翻译函数，导致在 Android 上可能报`TypeError: no access` 错误
                        // 故：启用 runtime 模式，内部定制了简易的 compileToFunction
                        'vue-i18n': (0, uni_cli_shared_1.resolveVueI18nRuntime)(),
                    },
                },
            };
        },
    };
}
exports.uniAppPlugin = uniAppPlugin;
