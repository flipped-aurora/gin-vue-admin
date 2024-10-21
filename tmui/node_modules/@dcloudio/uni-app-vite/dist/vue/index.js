"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initVuePlugins = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const template_1 = require("../plugins/template");
const manifestJson_1 = require("../plugins/manifestJson");
const stats_1 = require("../plugins/stats");
const confusion_1 = require("./plugins/confusion");
const mainJs_1 = require("./plugins/mainJs");
const pagesJson_1 = require("./plugins/pagesJson");
const renderjs_1 = require("./plugins/renderjs");
const plugin_1 = require("./plugin");
function initUniCssScopedPluginFilter(inputDir) {
    const styleIsolation = (0, uni_cli_shared_1.getAppStyleIsolation)((0, uni_cli_shared_1.parseManifestJsonOnce)(inputDir));
    if (styleIsolation === 'shared') {
        return;
    }
    if (styleIsolation === 'isolated') {
        // isolated: 对所有非 App.vue 增加 scoped
        return (id) => (0, uni_cli_shared_1.isVueSfcFile)(id) && !(0, uni_cli_shared_1.isAppVue)(id);
    }
    // apply-shared: 仅对非页面组件增加 scoped
    return (id) => (0, uni_cli_shared_1.isVueSfcFile)(id) && !(0, uni_cli_shared_1.isAppVue)(id) && !(0, uni_cli_shared_1.isUniPageFile)(id, inputDir);
}
function initVuePlugins() {
    const plugins = [
        (0, uni_cli_shared_1.uniEasycomPlugin)({ exclude: uni_cli_shared_1.UNI_EASYCOM_EXCLUDE }),
        (0, uni_cli_shared_1.uniHBuilderXConsolePlugin)(),
    ];
    if (process.env.UNI_COMPILE_TARGET === 'uni_modules') {
        plugins.push((0, uni_cli_shared_1.uniViteInjectPlugin)('uni:app-inject', (0, uni_cli_shared_1.initAppProvide)()), (0, renderjs_1.uniRenderjsPlugin)(), (0, plugin_1.uniAppVuePlugin)(), (0, uni_cli_shared_1.uniEncryptUniModulesPlugin)());
    }
    else {
        plugins.push((0, mainJs_1.uniMainJsPlugin)(), (0, manifestJson_1.uniManifestJsonPlugin)(), (0, pagesJson_1.uniPagesJsonPlugin)(), (0, uni_cli_shared_1.uniViteInjectPlugin)('uni:app-inject', (0, uni_cli_shared_1.initAppProvide)()), (0, renderjs_1.uniRenderjsPlugin)(), (0, template_1.uniTemplatePlugin)(), (0, stats_1.uniStatsPlugin)(), (0, plugin_1.uniAppVuePlugin)(), (0, confusion_1.uniConfusionPlugin)());
        const filter = initUniCssScopedPluginFilter(process.env.UNI_INPUT_DIR);
        if (filter) {
            plugins.unshift((0, uni_cli_shared_1.uniCssScopedPlugin)({ filter }));
        }
    }
    return plugins;
}
exports.initVuePlugins = initVuePlugins;
