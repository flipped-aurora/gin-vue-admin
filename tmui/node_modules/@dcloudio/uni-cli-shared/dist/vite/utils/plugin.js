"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.insertBeforePlugin = exports.removePlugins = exports.replacePlugins = exports.injectCssPostPlugin = exports.injectCssPlugin = exports.injectAssetPlugin = void 0;
const shared_1 = require("@vue/shared");
const asset_1 = require("../plugins/vitejs/plugins/asset");
const css_1 = require("../plugins/vitejs/plugins/css");
function injectAssetPlugin(config, options) {
    replacePlugins([(0, asset_1.assetPlugin)(config, options)], config);
}
exports.injectAssetPlugin = injectAssetPlugin;
function injectCssPlugin(config, options) {
    replacePlugins([
        (0, css_1.cssPlugin)(config, {
            isAndroidX: false,
            ...options,
        }),
    ], config);
}
exports.injectCssPlugin = injectCssPlugin;
function injectCssPostPlugin(config, newCssPostPlugin) {
    const oldCssPostPlugin = config.plugins.find((p) => p.name === newCssPostPlugin.name);
    // 直接覆盖原有方法，不能删除，替换，因为 unocss 在 pre 阶段已经获取到了旧的 css-post 插件对象
    if (oldCssPostPlugin) {
        (0, shared_1.extend)(oldCssPostPlugin, newCssPostPlugin);
    }
}
exports.injectCssPostPlugin = injectCssPostPlugin;
function replacePlugins(plugins, config) {
    plugins.forEach((plugin) => {
        const index = config.plugins.findIndex((p) => p.name === plugin.name);
        if (index > -1) {
            ;
            config.plugins.splice(index, 1, plugin);
        }
    });
}
exports.replacePlugins = replacePlugins;
function removePlugins(plugins, config) {
    if (!(0, shared_1.isArray)(plugins)) {
        plugins = [plugins];
    }
    plugins.forEach((name) => {
        const index = config.plugins.findIndex((p) => p.name === name);
        if (index > -1) {
            ;
            config.plugins.splice(index, 1);
        }
    });
}
exports.removePlugins = removePlugins;
function insertBeforePlugin(plugin, before, config) {
    const index = config.plugins.findIndex((p) => p.name === before);
    if (index > -1) {
        ;
        config.plugins.splice(index, 0, plugin);
    }
}
exports.insertBeforePlugin = insertBeforePlugin;
