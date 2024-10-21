"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initExtraPlugins = exports.initPluginUniOptions = void 0;
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function initPluginUniOptions(UniVitePlugins) {
    const assets = [];
    const targets = [];
    const transformEvent = Object.create(null);
    const compilerOptions = {};
    const jsxOptions = {};
    const styleOptions = {};
    let compiler;
    UniVitePlugins.forEach((plugin) => {
        const { compiler: pluginTemplateCompiler, copyOptions: pluginCopyOptions, compilerOptions: pluginCompilerOptions, jsxOptions: pluginJsxOptions, styleOptions: pluginStyleOpitons, } = plugin.uni || {};
        if (pluginTemplateCompiler) {
            compiler = pluginTemplateCompiler;
        }
        if (pluginCompilerOptions) {
            (0, shared_1.extend)(compilerOptions, pluginCompilerOptions);
        }
        if (pluginJsxOptions) {
            (0, shared_1.extend)(jsxOptions, pluginJsxOptions);
        }
        if (pluginStyleOpitons) {
            (0, shared_1.extend)(styleOptions, pluginStyleOpitons);
        }
        if (pluginCopyOptions) {
            let copyOptions = pluginCopyOptions;
            if ((0, shared_1.isFunction)(pluginCopyOptions)) {
                copyOptions = pluginCopyOptions();
            }
            if (copyOptions.assets) {
                assets.push(...copyOptions.assets);
            }
            if (copyOptions.targets) {
                targets.push(...copyOptions.targets);
            }
        }
    });
    return {
        compiler,
        copyOptions: {
            assets,
            targets,
        },
        transformEvent,
        compilerOptions,
        jsxOptions,
        styleOptions,
    };
}
exports.initPluginUniOptions = initPluginUniOptions;
function initExtraPlugins(cliRoot, platform, options) {
    return initPlugins(cliRoot, resolvePlugins(cliRoot, platform, options.uvue), options);
}
exports.initExtraPlugins = initExtraPlugins;
function initPlugin(cliRoot, { id, config: { main } }, options) {
    let plugin = require(require.resolve(path_1.default.join(id, main || '/lib/uni.plugin.js'), { paths: [cliRoot] }));
    plugin = plugin.default || plugin;
    if ((0, shared_1.isFunction)(plugin)) {
        plugin = plugin(options);
    }
    return plugin;
}
function initPlugins(cliRoot, plugins, options) {
    return plugins
        .map((plugin) => initPlugin(cliRoot, plugin, options))
        .flat()
        .filter(Boolean)
        .map((plugin) => {
        if ((0, shared_1.isFunction)(plugin)) {
            return plugin(options);
        }
        return plugin;
    })
        .flat();
}
function resolvePlugins(cliRoot, platform, uvue = false) {
    const pkg = require(path_1.default.join(cliRoot, 'package.json'));
    return Object.keys(pkg.devDependencies || {})
        .concat(Object.keys(pkg.dependencies || {}))
        .map((id) => {
        try {
            const pluginPkg = require(require.resolve(id + '/package.json', {
                paths: [cliRoot],
            }));
            const config = pluginPkg['uni-app'];
            if (!config || !config.name) {
                return;
            }
            const { apply } = config;
            if ((0, shared_1.isArray)(apply)) {
                // 注册所有平台
                apply.forEach((p) => (0, uni_cli_shared_1.registerPlatform)(p));
                if (!apply.includes(platform)) {
                    return;
                }
            }
            else if ((0, shared_1.isString)(apply)) {
                if (!new RegExp(apply).test(platform)) {
                    return;
                }
            }
            // 插件必须支持 uvue
            if (uvue && !config.uvue) {
                return;
            }
            return {
                id,
                name: config.name,
                config,
            };
        }
        catch (e) { }
    })
        .filter(Boolean);
}
