"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const plugin_1 = require("./plugin");
const usingComponents_1 = require("./plugins/usingComponents");
const mainJs_1 = require("./plugins/mainJs");
const manifestJson_1 = require("./plugins/manifestJson");
const pagesJson_1 = require("./plugins/pagesJson");
const entry_1 = require("./plugins/entry");
const renderjs_1 = require("./plugins/renderjs");
const runtimeHooks_1 = require("./plugins/runtimeHooks");
const subpackage_1 = require("./plugins/subpackage");
const plugin_2 = require("./plugins/plugin");
exports.default = (options) => {
    if (!options.app.subpackages) {
        delete process.env.UNI_SUBPACKAGE;
    }
    if (!options.app.plugins) {
        delete process.env.UNI_MP_PLUGIN;
    }
    const normalizeComponentName = options.template.component?.normalizeName;
    return [
        (options) => {
            return (0, mainJs_1.uniMainJsPlugin)({
                normalizeComponentName,
                babelParserPlugins: options.vueOptions?.script?.babelParserPlugins,
            });
        },
        (0, manifestJson_1.uniManifestJsonPlugin)(options),
        (0, pagesJson_1.uniPagesJsonPlugin)(options),
        (0, entry_1.uniEntryPlugin)(options),
        (0, uni_cli_shared_1.uniViteInjectPlugin)('uni:mp-inject', (0, shared_1.extend)({ exclude: [/uni.api.esm/, /uni.mp.esm/] }, options.vite.inject)),
        (0, renderjs_1.uniRenderjsPlugin)({ lang: options.template.filter?.lang }),
        (0, runtimeHooks_1.uniRuntimeHooksPlugin)(),
        (0, plugin_1.uniMiniProgramPlugin)(options),
        (options) => {
            return (0, usingComponents_1.uniUsingComponentsPlugin)({
                normalizeComponentName,
                babelParserPlugins: options.vueOptions?.script?.babelParserPlugins,
            });
        },
        ...(process.env.UNI_SUBPACKAGE ? [(0, subpackage_1.uniSubpackagePlugin)(options)] : []),
        ...(process.env.UNI_MP_PLUGIN ? [(0, plugin_2.uniMiniProgramPluginPlugin)(options)] : []),
    ];
};
