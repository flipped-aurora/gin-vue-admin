"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initNVuePlugins = exports.initNVueNodeTransforms = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const manifestJson_1 = require("../plugins/manifestJson");
const stats_1 = require("../plugins/stats");
const template_1 = require("../plugins/template");
const plugin_1 = require("./plugin");
const appCss_1 = require("./plugins/appCss");
const esbuild_1 = require("./plugins/esbuild");
const mainJs_1 = require("./plugins/mainJs");
const pagesJson_1 = require("./plugins/pagesJson");
const renderjs_1 = require("./plugins/renderjs");
var plugin_2 = require("./plugin");
Object.defineProperty(exports, "initNVueNodeTransforms", { enumerable: true, get: function () { return plugin_2.initNVueNodeTransforms; } });
function initNVuePlugins() {
    const renderer = process.env.UNI_RENDERER;
    const appService = process.env.UNI_RENDERER_NATIVE === 'appService';
    return process.env.UNI_COMPILE_TARGET === 'uni_modules'
        ? [
            (0, uni_cli_shared_1.uniEasycomPlugin)({ exclude: uni_cli_shared_1.UNI_EASYCOM_EXCLUDE }),
            (0, uni_cli_shared_1.uniHBuilderXConsolePlugin)(),
            (0, uni_cli_shared_1.uniViteInjectPlugin)('uni:app-inject', (0, uni_cli_shared_1.initAppProvide)()),
            (0, renderjs_1.uniRenderjsPlugin)(),
            (0, plugin_1.uniAppNVuePlugin)({ appService }),
            (0, uni_cli_shared_1.uniEncryptUniModulesPlugin)(),
        ]
        : [
            (0, appCss_1.uniAppCssPlugin)(),
            (0, uni_cli_shared_1.uniEasycomPlugin)({ exclude: uni_cli_shared_1.UNI_EASYCOM_EXCLUDE }),
            (0, uni_cli_shared_1.uniHBuilderXConsolePlugin)(),
            (0, mainJs_1.uniMainJsPlugin)({ renderer, appService }),
            (0, manifestJson_1.uniManifestJsonPlugin)(),
            (0, pagesJson_1.uniPagesJsonPlugin)({ renderer, appService }),
            (0, uni_cli_shared_1.uniViteInjectPlugin)('uni:app-inject', (0, uni_cli_shared_1.initAppProvide)()),
            (0, renderjs_1.uniRenderjsPlugin)(),
            (0, plugin_1.uniAppNVuePlugin)({ appService }),
            (0, esbuild_1.uniEsbuildPlugin)({ renderer, appService }),
            ...(appService
                ? [(0, stats_1.uniStatsPlugin)(), (0, template_1.uniTemplatePlugin)({ renderer })]
                : []),
        ];
}
exports.initNVuePlugins = initNVuePlugins;
