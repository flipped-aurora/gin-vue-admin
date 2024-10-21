"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const vue_1 = require("./vue");
const nvue_1 = require("./nvue");
const plugin_1 = require("./plugin");
exports.default = () => {
    return [
        (0, plugin_1.uniAppPlugin)({
            renderer: process.env.UNI_RENDERER,
            appService: process.env.UNI_RENDERER_NATIVE === 'appService',
        }),
        (0, uni_cli_shared_1.uniUTSAppUniModulesPlugin)({
            x: false,
            isSingleThread: false,
            extApis: (0, uni_cli_shared_1.parseUniExtApiNamespacesOnce)(process.env.UNI_UTS_PLATFORM, process.env.UNI_UTS_TARGET_LANGUAGE),
        }),
        ...(process.env.UNI_COMPILER === 'nvue'
            ? (0, nvue_1.initNVuePlugins)()
            : (0, vue_1.initVuePlugins)()),
    ];
};
