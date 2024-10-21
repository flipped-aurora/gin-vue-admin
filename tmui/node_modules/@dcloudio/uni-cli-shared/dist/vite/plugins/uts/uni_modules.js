"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniDecryptUniModulesPlugin = exports.resolveExtApiProvider = exports.buildUniExtApis = exports.uniUTSAppUniModulesPlugin = exports.getCurrentCompiledUTSPlugins = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uts_1 = require("../../../uts");
const utils_1 = require("../../utils");
const uni_modules_1 = require("../../../uni_modules");
const utils_2 = require("../../../utils");
const json_1 = require("../../../json");
const UTSProxyRE = /\?uts-proxy$/;
const UniHelpersRE = /\?uni_helpers$/;
function isUTSProxy(id) {
    return UTSProxyRE.test(id);
}
function isUniHelpers(id) {
    return UniHelpersRE.test(id);
}
const utsModuleCaches = new Map();
const utsPlugins = new Set();
function getCurrentCompiledUTSPlugins() {
    return utsPlugins;
}
exports.getCurrentCompiledUTSPlugins = getCurrentCompiledUTSPlugins;
let uniExtApiCompiler = async () => { };
// 该插件仅限app-android、app-ios、app-harmony
function uniUTSAppUniModulesPlugin(options = {}) {
    const inputDir = process.env.UNI_INPUT_DIR;
    process.env.UNI_UTS_USING_ROLLUP = 'true';
    const compilePlugin = async (pluginDir) => {
        utsPlugins.add(path_1.default.basename(pluginDir));
        const pkgJson = require(path_1.default.join(pluginDir, 'package.json'));
        const isExtApi = !!pkgJson.uni_modules?.['uni-ext-api'];
        const extApiProvider = resolveExtApiProvider(pkgJson);
        // 如果是 provider 扩展，需要判断 provider 的宿主插件是否在本地，在的话，自动导入该宿主插件包名
        let uniExtApiProviderServicePlugin = '';
        if (extApiProvider?.servicePlugin) {
            if (fs_1.default.existsSync(path_1.default.resolve(inputDir, 'uni_modules', extApiProvider.servicePlugin))) {
                uniExtApiProviderServicePlugin = extApiProvider.servicePlugin;
            }
        }
        const compiler = (0, uts_1.resolveUTSCompiler)();
        // 处理依赖的 uts 插件
        const deps = (0, uni_modules_1.parseUTSModuleDeps)(pkgJson.uni_modules?.dependencies || [], inputDir);
        if (deps.length) {
            for (const dep of deps) {
                await compilePlugin(path_1.default.resolve(inputDir, 'uni_modules', dep));
            }
        }
        if (process.env.UNI_PLATFORM === 'app-harmony') {
            return compiler.compileArkTS(pluginDir, {
                isX: !!options.x,
                isExtApi,
                transform: {
                    uniExtApiProviderName: extApiProvider?.name,
                    uniExtApiProviderService: extApiProvider?.service,
                    uniExtApiProviderServicePlugin,
                },
            });
        }
        return compiler.compile(pluginDir, {
            isX: !!options.x,
            isSingleThread: !!options.isSingleThread,
            isPlugin: true,
            isExtApi,
            extApis: options.extApis,
            sourceMap: (0, utils_2.enableSourceMap)(),
            uni_modules: deps,
            transform: {
                uniExtApiProviderName: extApiProvider?.name,
                uniExtApiProviderService: extApiProvider?.service,
                uniExtApiProviderServicePlugin,
            },
        });
    };
    uniExtApiCompiler = async () => {
        // 获取 provider 扩展(编译所有uni)
        const plugins = (0, uni_modules_1.getUniExtApiPlugins)().filter((provider) => !utsPlugins.has(provider.plugin));
        for (const plugin of plugins) {
            const pluginDir = path_1.default.resolve(inputDir, 'uni_modules', plugin.plugin);
            // 如果是 app-js 环境
            if (process.env.UNI_APP_X_UVUE_SCRIPT_ENGINE === 'js') {
                if (fs_1.default.existsSync(path_1.default.resolve(pluginDir, 'utssdk', 'app-js', 'index.uts'))) {
                    continue;
                }
            }
            const result = await compilePlugin(pluginDir);
            if (result) {
                // 时机不对，不能addWatch
                // result.deps.forEach((dep) => {
                //   this.addWatchFile(dep)
                // })
            }
        }
    };
    return {
        name: 'uni:uts-uni_modules',
        apply: 'build',
        enforce: 'pre',
        resolveId(id, importer) {
            if (isUTSProxy(id) || isUniHelpers(id)) {
                return id;
            }
            // 加密插件缓存目录的css文件
            if (id.endsWith('.css')) {
                return;
            }
            const module = (0, uts_1.resolveUTSAppModule)(process.env.UNI_UTS_PLATFORM, id, importer ? path_1.default.dirname(importer) : inputDir, options.x !== true);
            if (module) {
                // app-js 会直接返回 index.uts 路径，不需要 uts-proxy
                if (module.endsWith('.uts')) {
                    return module;
                }
                // prefix the polyfill id with \0 to tell other plugins not to try to load or transform it
                return module + '?uts-proxy';
            }
        },
        load(id) {
            if (isUTSProxy(id)) {
                return '';
            }
        },
        buildEnd() {
            utsModuleCaches.clear();
        },
        async transform(_, id, opts) {
            if (opts && opts.ssr) {
                return;
            }
            if (!isUTSProxy(id)) {
                return;
            }
            const { filename: pluginDir } = (0, utils_1.parseVueRequest)(id.replace('\0', ''));
            // 当 vue 和 nvue 均引用了相同 uts 插件，解决两套编译器会编译两次 uts 插件的问题
            // 通过缓存，保证同一个 uts 插件只编译一次
            if (utsModuleCaches.get(pluginDir)) {
                return utsModuleCaches.get(pluginDir)().then((result) => {
                    if (result) {
                        result.deps.forEach((dep) => {
                            this.addWatchFile(dep);
                        });
                        return {
                            code: result.code,
                            map: null,
                            syntheticNamedExports: result.encrypt,
                            meta: result.meta,
                        };
                    }
                });
            }
            const compile = (0, uni_shared_1.once)(() => {
                return compilePlugin(pluginDir);
            });
            utsModuleCaches.set(pluginDir, compile);
            const result = await compile();
            if (result) {
                result.deps.forEach((dep) => {
                    this.addWatchFile(dep);
                });
                return {
                    code: result.code,
                    map: null,
                    syntheticNamedExports: result.encrypt,
                    meta: result.meta,
                };
            }
        },
    };
}
exports.uniUTSAppUniModulesPlugin = uniUTSAppUniModulesPlugin;
async function buildUniExtApis() {
    await uniExtApiCompiler();
}
exports.buildUniExtApis = buildUniExtApis;
function resolveExtApiProvider(pkg) {
    const provider = pkg.uni_modules?.['uni-ext-api']?.provider;
    if (provider?.service) {
        if (provider.name && !provider.servicePlugin) {
            provider.servicePlugin = 'uni-' + provider.service;
        }
        return provider;
    }
}
exports.resolveExtApiProvider = resolveExtApiProvider;
function uniDecryptUniModulesPlugin() {
    const inputDir = process.env.UNI_INPUT_DIR;
    const isX = process.env.UNI_APP_X === 'true';
    return {
        name: 'uni:uni_modules-d',
        enforce: 'pre',
        async configResolved() {
            if (isX && process.env.UNI_COMPILE_TARGET !== 'uni_modules') {
                const manifest = (0, json_1.parseManifestJsonOnce)(inputDir);
                await (0, uni_modules_1.checkEncryptUniModules)(inputDir, {
                    mode: process.env.NODE_ENV !== 'development'
                        ? 'production'
                        : 'development',
                    packType: process.env.UNI_APP_PACK_TYPE ||
                        (process.env.NODE_ENV !== 'development' ? 'release' : 'debug'),
                    compilerVersion: process.env.UNI_COMPILER_VERSION,
                    appid: manifest.appid,
                    appname: manifest.name,
                    platform: process.env.UNI_UTS_PLATFORM,
                    'uni-app-x': isX,
                });
            }
        },
        resolveId(id) {
            if (isUTSProxy(id) || isUniHelpers(id)) {
                return id;
            }
            if (isX &&
                process.env.UNI_COMPILE_TARGET !== 'uni_modules' &&
                !id.endsWith('.css')) {
                const resolvedId = (0, uni_modules_1.resolveEncryptUniModule)(id, process.env.UNI_UTS_PLATFORM, process.env.UNI_APP_X === 'true');
                if (resolvedId) {
                    return resolvedId;
                }
            }
        },
    };
}
exports.uniDecryptUniModulesPlugin = uniDecryptUniModulesPlugin;
