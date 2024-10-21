"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getExtApiComponents = exports.addExtApiComponents = exports.getGlobalPageOrientation = exports.setGlobalPageOrientation = exports.updateManifestModules = exports.normalizeManifestJson = exports.relativeInputDir = exports.configResolved = exports.isPages = exports.isManifest = exports.createUniOptions = void 0;
const node_path_1 = __importDefault(require("node:path"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_i18n_1 = require("@dcloudio/uni-i18n");
const compiler_core_1 = require("@vue/compiler-core");
function createUniOptions(platform) {
    return {
        copyOptions() {
            const platform = process.env.UNI_PLATFORM;
            const inputDir = process.env.UNI_INPUT_DIR;
            const outputDir = process.env.UNI_OUTPUT_DIR;
            const targets = [];
            // 自动化测试时，不启用隐私政策
            if (!process.env.UNI_AUTOMATOR_WS_ENDPOINT) {
                if (process.env.UNI_UTS_PLATFORM === 'app-android') {
                    targets.push({
                        src: 'androidPrivacy.json',
                        dest: outputDir,
                        transform(source) {
                            const options = (0, uni_cli_shared_1.initI18nOptions)(platform, inputDir, false, true);
                            if (!options) {
                                return;
                            }
                            return (0, uni_i18n_1.compileI18nJsonStr)(source.toString(), options);
                        },
                    });
                }
            }
            return {
                assets: ['hybrid/html/**/*', 'uni_modules/*/hybrid/html/**/*'],
                targets,
            };
        },
        compilerOptions: platform === 'ios'
            ? {
                isNativeTag(tag) {
                    return (0, uni_cli_shared_1.matchUTSComponent)(tag) || (0, uni_shared_1.isAppIOSUVueNativeTag)(tag);
                },
                nodeTransforms: [
                    uni_cli_shared_1.transformTapToClick,
                    uni_cli_shared_1.transformUTSComponent,
                    // TODO 合并复用安卓插件逻辑
                    function (node, context) {
                        if (node.type === 2) {
                            const parent = context.parent;
                            if (parent && parent.type === 1 && parent.tag === 'text') {
                                // 解析文本节点转义，暂时仅处理换行
                                node.content = node.content.replace(/[\\]+n/g, function (match) {
                                    return JSON.parse(`"${match}"`);
                                });
                            }
                        }
                    },
                    (node) => {
                        // 收集可能的 extApiComponents
                        if (node.type === compiler_core_1.NodeTypes.ELEMENT &&
                            node.tagType === compiler_core_1.ElementTypes.ELEMENT) {
                            if (!(0, uni_cli_shared_1.parseUTSComponent)(node.tag, 'swift')) {
                                addExtApiComponents([node.tag]);
                            }
                        }
                    },
                ],
            }
            : {},
    };
}
exports.createUniOptions = createUniOptions;
function isManifest(id) {
    return id.endsWith(uni_cli_shared_1.MANIFEST_JSON_UTS);
}
exports.isManifest = isManifest;
function isPages(id) {
    return id.endsWith(uni_cli_shared_1.PAGES_JSON_UTS);
}
exports.isPages = isPages;
// TODO vite 升级需要考虑调整以下列表
const REMOVED_PLUGINS = [
    'vite:build-metadata',
    'vite:modulepreload-polyfill',
    // 'vite:css', // iOS replace
    'vite:esbuild',
    'vite:wasm-helper',
    'vite:worker',
    // 'vite:json',
    // 'vite:asset', // replace
    'vite:wasm-fallback',
    // 'vite:define',
    // 'vite:css-post', // iOS replace
    'vite:build-html',
    'vite:html-inline-proxy',
    'vite:worker-import-meta-url',
    'vite:asset-import-meta-url',
    'vite:force-systemjs-wrap-complete',
    'vite:watch-package-data',
    'commonjs',
    'vite:data-uri',
    'vite:dynamic-import-vars',
    'vite:import-glob',
    'vite:build-import-analysis',
    'vite:esbuild-transpile',
    'vite:terser',
    'vite:reporter',
];
function configResolved(config, isAndroidX = false) {
    const plugins = config.plugins;
    const len = plugins.length;
    const removedPlugins = REMOVED_PLUGINS.slice(0);
    if (isAndroidX) {
        removedPlugins.push('vite:css');
        removedPlugins.push('vite:css-post');
    }
    for (let i = len - 1; i >= 0; i--) {
        const plugin = plugins[i];
        if (removedPlugins.includes(plugin.name)) {
            plugins.splice(i, 1);
        }
    }
    // console.log(plugins.map((p) => p.name))
    // 强制不inline
    config.build.assetsInlineLimit = 0;
    (0, uni_cli_shared_1.injectAssetPlugin)(config, { isAndroidX });
}
exports.configResolved = configResolved;
function relativeInputDir(filename) {
    const inputDir = process.env.UNI_INPUT_DIR;
    filename = (0, uni_cli_shared_1.normalizeNodeModules)(filename);
    if (filename.startsWith(inputDir)) {
        return (0, uni_cli_shared_1.normalizePath)(node_path_1.default.relative(inputDir, filename));
    }
    return filename;
}
exports.relativeInputDir = relativeInputDir;
function normalizeManifestJson(userManifestJson) {
    const app = userManifestJson.app || {};
    const x = userManifestJson['uni-app-x'] || {};
    x.compilerVersion = process.env.UNI_COMPILER_VERSION || '';
    const pageOrientation = getGlobalPageOrientation();
    if (pageOrientation) {
        if (!app.distribute) {
            app.distribute = {};
        }
        app.distribute['_uni-app-x_'] = {
            pageOrientation,
        };
    }
    return {
        id: userManifestJson.appid || '',
        name: userManifestJson.name || '',
        description: userManifestJson.description || '',
        version: {
            name: userManifestJson.versionName || '',
            code: userManifestJson.versionCode || '',
        },
        'uni-app-x': x,
        app,
    };
}
exports.normalizeManifestJson = normalizeManifestJson;
function updateManifestModules(manifest, modules) {
    // 执行了摇树逻辑，就需要设置 modules 节点
    const app = manifest.app;
    if (!app.distribute) {
        app.distribute = {};
    }
    if (!app.distribute.modules) {
        app.distribute.modules = {};
    }
    if (modules) {
        modules.forEach((name) => {
            const value = app.distribute.modules[name];
            app.distribute.modules[name] =
                typeof value === 'object' && value !== null ? value : {};
        });
    }
    return manifest;
}
exports.updateManifestModules = updateManifestModules;
let pageOrientation = '';
function setGlobalPageOrientation(value) {
    pageOrientation = value;
}
exports.setGlobalPageOrientation = setGlobalPageOrientation;
function getGlobalPageOrientation() {
    return pageOrientation;
}
exports.getGlobalPageOrientation = getGlobalPageOrientation;
const extApiComponents = new Set();
function addExtApiComponents(components) {
    components.forEach((component) => {
        extApiComponents.add(component);
    });
}
exports.addExtApiComponents = addExtApiComponents;
function getExtApiComponents() {
    return extApiComponents;
}
exports.getExtApiComponents = getExtApiComponents;
