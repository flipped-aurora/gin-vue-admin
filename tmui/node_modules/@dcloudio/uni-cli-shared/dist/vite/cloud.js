"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.addUniModulesExtApiComponents = exports.uniEncryptUniModulesPlugin = exports.uniEncryptUniModulesAssetsPlugin = exports.createEncryptCssUrlReplacer = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_modules_1 = require("../uni_modules");
const utils_1 = require("./plugins/vitejs/utils");
const uts_1 = require("../uts");
const utils_2 = require("../utils");
const easycom_1 = require("../easycom");
function createEncryptCssUrlReplacer(resolve) {
    return async (url, importer) => {
        if (url.startsWith('/') && !url.startsWith('//')) {
            // /static/logo.png => @/static/logo.png
            url = '@' + url;
        }
        const resolved = await resolve(url, importer);
        if (resolved) {
            return ('@/' + (0, utils_2.normalizePath)(path_1.default.relative(process.env.UNI_INPUT_DIR, resolved)));
        }
        return url;
    };
}
exports.createEncryptCssUrlReplacer = createEncryptCssUrlReplacer;
// 处理静态资源加载（目前仅限非app-android）
function uniEncryptUniModulesAssetsPlugin() {
    let resolvedConfig;
    return {
        name: 'uni:encrypt-uni-modules-assets',
        enforce: 'pre',
        configResolved(config) {
            resolvedConfig = config;
        },
        resolveId(id, importer) {
            if (resolvedConfig.assetsInclude((0, utils_1.cleanUrl)(id))) {
                id = (0, utils_2.normalizePath)(id);
                if (importer && (id.startsWith('./') || id.startsWith('../'))) {
                    id = (0, utils_2.normalizePath)(path_1.default.resolve(path_1.default.dirname(importer), id));
                }
                if (path_1.default.isAbsolute(id)) {
                    id = '@/' + path_1.default.relative(process.env.UNI_INPUT_DIR, id);
                }
                return `\0${id}`;
            }
        },
        load(id) {
            if (resolvedConfig.assetsInclude((0, utils_1.cleanUrl)(id))) {
                return {
                    code: `export default ${JSON.stringify(id.replace(/\0/g, ''))}`,
                    moduleSideEffects: false,
                };
            }
        },
    };
}
exports.uniEncryptUniModulesAssetsPlugin = uniEncryptUniModulesAssetsPlugin;
function uniEncryptUniModulesPlugin() {
    let resolvedConfig;
    return {
        name: 'uni:encrypt-uni-modules',
        apply: 'build',
        config() {
            return {
                resolve: {
                    alias: initEncryptUniModulesAlias(),
                },
                build: initEncryptUniModulesBuildOptions(process.env.UNI_UTS_PLATFORM, process.env.UNI_INPUT_DIR),
            };
        },
        configResolved(config) {
            // 编译组件时，禁用内联资源
            config.build.assetsInlineLimit = 0;
            config.build.rollupOptions.external = createExternal(config);
            resolvedConfig = config;
        },
        resolveId(id, importer) {
            if (process.env.UNI_UTS_PLATFORM !== 'app-android') {
                if (resolvedConfig.assetsInclude((0, utils_1.cleanUrl)(id))) {
                    id = (0, utils_2.normalizePath)(id);
                    if (importer && (id.startsWith('./') || id.startsWith('../'))) {
                        id = (0, utils_2.normalizePath)(path_1.default.resolve(path_1.default.dirname(importer), id));
                    }
                    if (path_1.default.isAbsolute(id)) {
                        id = '@/' + path_1.default.relative(process.env.UNI_INPUT_DIR, id);
                    }
                    return `\0${id}`;
                }
            }
        },
        load(id) {
            if (process.env.UNI_UTS_PLATFORM !== 'app-android') {
                if (resolvedConfig.assetsInclude((0, utils_1.cleanUrl)(id))) {
                    return {
                        code: `export default ${JSON.stringify(id.replace(/\0/g, ''))}`,
                        moduleSideEffects: false,
                    };
                }
            }
        },
        generateBundle(_, bundle) {
            Object.keys(bundle).forEach((fileName) => {
                if (fileName.endsWith('.module.js')) {
                    const uniModuleId = path_1.default.basename(fileName).replace('.module.js', '');
                    // app-android 不需要 js
                    if (process.env.UNI_UTS_PLATFORM !== 'app-android') {
                        const newFileName = 'uni_modules/' +
                            fileName.replace('.module.js', '/index.module.js');
                        bundle[newFileName] = bundle[fileName];
                        bundle[newFileName].fileName = newFileName;
                    }
                    delete bundle[fileName];
                    const pkg = `uni_modules/${uniModuleId}/package.json`;
                    bundle[pkg] = {
                        type: 'asset',
                        fileName: pkg,
                        name: pkg,
                        needsCodeReference: false,
                        source: genUniModulesPackageJson(uniModuleId, process.env.UNI_INPUT_DIR, {
                            env: (0, uni_modules_1.initCheckEnv)(),
                        }),
                    };
                }
            });
        },
        async writeBundle() {
            if (process.env.UNI_UTS_PLATFORM !== 'app-android') {
                return;
            }
            // 编译所有 uni_modules 插件
            const tempOutputDir = uvueOutDir();
            const tempUniModulesDir = path_1.default.join(tempOutputDir, 'uni_modules');
            const tempUniModules = [];
            if (fs_extra_1.default.existsSync(tempUniModulesDir)) {
                fs_extra_1.default.readdirSync(tempUniModulesDir).forEach((uniModuleDir) => {
                    if (fs_extra_1.default.existsSync(path_1.default.join(tempUniModulesDir, uniModuleDir, 'index.module.uts'))) {
                        tempUniModules.push(uniModuleDir);
                    }
                });
            }
            const compiler = (0, uts_1.resolveUTSCompiler)();
            for (const uniModule of tempUniModules) {
                const pluginDir = path_1.default.resolve(tempUniModulesDir, uniModule);
                // TODO 待优化autoImports，目前 uni-app x 的编译，autoImport 是在js层处理过，rust层基本不再使用
                // 但uts插件目前还是使用的rust层的autoImports
                const autoImports = {};
                const allAutoImports = (0, easycom_1.getUTSEasyComAutoImports)();
                Object.keys(allAutoImports).forEach((source) => {
                    if (!source.startsWith(`@/uni_modules/${uniModule}/components/`)) {
                        autoImports[source] = allAutoImports[source];
                    }
                });
                const uni_modules = [];
                const pkgJson = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'uni_modules', uniModule, 'package.json');
                if (fs_extra_1.default.existsSync(pkgJson)) {
                    try {
                        const pkg = require(pkgJson);
                        if (pkg.uni_modules?.dependencies) {
                            uni_modules.push(...pkg.uni_modules.dependencies);
                        }
                    }
                    catch (e) {
                        console.error(e);
                    }
                }
                const result = await compiler.compile(pluginDir, {
                    isX: process.env.UNI_APP_X === 'true',
                    isSingleThread: true,
                    isPlugin: false,
                    sourceMap: false,
                    uni_modules,
                    transform: {
                        uvueClassNamePrefix: 'Gen',
                        autoImports,
                        uvueGenDefaultAs: '__sfc__',
                    },
                });
                if (result) {
                    const apis = result.inject_apis;
                    const scopedSlots = result.scoped_slots;
                    const components = getUniModulesExtApiComponents(uniModule);
                    const modules = (0, uts_1.resolveUTSCompiler)().parseInjectModules(apis, {}, components);
                    fs_extra_1.default.writeFileSync(path_1.default.resolve(process.env.UNI_OUTPUT_DIR, 'uni_modules', uniModule, 'package.json'), genUniModulesPackageJson(uniModule, tempOutputDir, {
                        env: (0, uni_modules_1.initCheckEnv)(),
                        apis,
                        components,
                        modules,
                        scopedSlots,
                    }));
                }
            }
        },
    };
}
exports.uniEncryptUniModulesPlugin = uniEncryptUniModulesPlugin;
function uvueOutDir() {
    return path_1.default.join(process.env.UNI_OUTPUT_DIR, '../.uvue');
}
function createExternal(config) {
    return function external(source) {
        if ([
            'vue',
            'plugin-vue:export-helper',
            'vue-router',
            'pinia',
            'vuex',
            'vue-i18n',
        ].includes(source)) {
            return true;
        }
        if (source.startsWith('@vue/')) {
            return true;
        }
        if (source.startsWith('@dcloudio/')) {
            return true;
        }
        if (source.startsWith('@/uni_modules/')) {
            return true;
        }
        // 相对目录
        if (source.startsWith('@/') || source.startsWith('.')) {
            return false;
        }
        if (path_1.default.isAbsolute(source)) {
            return false;
        }
        // android 系统库，三方库，iOS 的库呢？一般不包含.
        if (source.includes('.')) {
            return true;
        }
        return false;
    };
}
function initEncryptUniModulesAlias() {
    return [
        {
            find: '\0plugin-vue:export-helper',
            replacement: 'plugin-vue:export-helper',
        },
    ];
}
const indexFiles = ['index.uts', 'index.ts', 'index.js'];
function hasIndexFile(uniModuleDir) {
    return fs_extra_1.default.readdirSync(uniModuleDir).some((file) => indexFiles.includes(file));
}
function initEncryptUniModulesBuildOptions(platform, inputDir) {
    const modules = (0, uni_modules_1.parseUniModulesWithComponents)(inputDir);
    const moduleNames = Object.keys(modules);
    if (!moduleNames.length) {
        throw new Error('No encrypt uni_modules found');
    }
    // 生成入口文件
    const input = {};
    moduleNames.forEach((module) => {
        const moduleDir = path_1.default.resolve(inputDir, 'uni_modules', module);
        const indexEncryptFile = path_1.default.resolve(moduleDir, 'index.module.uts');
        const codes = [];
        if (hasIndexFile(moduleDir)) {
            codes.push(`export * from './index'`);
        }
        // easyCom
        if (modules[module] && Object.keys(modules[module]).length) {
            codes.push((0, uni_modules_1.genEncryptEasyComModuleIndex)(platform, modules[module]));
        }
        if (codes.length) {
            fs_extra_1.default.writeFileSync(indexEncryptFile, codes.join(`\n`));
            // 输出 xxx.module ，确保相对路径的准确性，因为真正引用的时候，是从 @/uni_modules/xxx 引入的
            input[module + '.module'] = indexEncryptFile;
        }
    });
    return {
        lib: false, // 不使用 lib 模式，lib模式会直接内联资源
        cssCodeSplit: false,
        // outDir: process.env.UNI_OUTPUT_DIR,
        rollupOptions: {
            preserveEntrySignatures: 'strict',
            input,
            output: {
                format: 'es',
                banner: ``,
                entryFileNames: '[name].js',
                assetFileNames(asset) {
                    if (asset.name && path_1.default.isAbsolute(asset.name)) {
                        const uniModuleId = parseUniModuleId(path_1.default.relative(inputDir, asset.name));
                        if (uniModuleId) {
                            return `uni_modules/${uniModuleId}/assets/[name]-[hash][extname]`;
                        }
                    }
                    return 'assets/[name]-[hash][extname]';
                },
            },
        },
    };
}
function genUniModulesPackageJson(uniModuleId, inputDir, artifacts) {
    const pkg = require(path_1.default.resolve(inputDir, path_1.default.join('uni_modules', uniModuleId, 'package.json')));
    return JSON.stringify({
        id: pkg.id,
        version: pkg.version,
        uni_modules: {
            dependencies: pkg.uni_modules?.dependencies || [],
            artifacts,
        },
    }, null, 2);
}
function parseUniModuleId(relativeFilename) {
    const parts = (0, utils_2.normalizePath)(relativeFilename).split('/', 2);
    if (parts[0] === 'uni_modules') {
        return parts[1];
    }
}
const uniModulesExtApiComponents = new Map();
function addUniModulesExtApiComponents(relativeFilename, components) {
    const uniModuleId = parseUniModuleId(relativeFilename);
    if (uniModuleId) {
        let extApiComponents = uniModulesExtApiComponents.get(uniModuleId);
        if (!extApiComponents) {
            extApiComponents = new Set();
            uniModulesExtApiComponents.set(uniModuleId, extApiComponents);
        }
        components.forEach((component) => extApiComponents.add(component));
    }
}
exports.addUniModulesExtApiComponents = addUniModulesExtApiComponents;
function getUniModulesExtApiComponents(uniModuleId) {
    return [...(uniModulesExtApiComponents.get(uniModuleId) || [])];
}
