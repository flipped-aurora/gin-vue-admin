"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseUniModulesArtifacts = exports.checkEncryptUniModules = exports.resolveEncryptUniModule = exports.initCheckEnv = exports.packUploadEncryptUniModules = exports.findUploadEncryptUniModulesFiles = exports.findEncryptUniModules = exports.parseEasyComComponents = exports.parseUniModulesWithComponents = exports.genEncryptEasyComModuleIndex = exports.parseUTSModuleDeps = exports.capitalize = exports.camelize = exports.parseInjects = exports.parseUniExtApi = exports.parseUniExtApis = exports.getUniExtApiProviderRegisters = exports.formatExtApiProviderName = exports.getUniExtApiPlugins = exports.getUniExtApiProviders = void 0;
// 重要：此文件编译后的js，需同步至 vue2 编译器中 uni-cli-shared/lib/uts/uni_modules.js
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const fast_glob_1 = require("fast-glob");
const utils_1 = require("./utils");
const easycom_1 = require("./easycom");
const messages_1 = require("./messages");
const extApiProviders = [];
const extApiPlugins = new Set();
function getUniExtApiProviders() {
    return extApiProviders;
}
exports.getUniExtApiProviders = getUniExtApiProviders;
function getUniExtApiPlugins() {
    return [...extApiPlugins].map((plugin) => {
        return { plugin };
    });
}
exports.getUniExtApiPlugins = getUniExtApiPlugins;
function formatExtApiProviderName(service, name) {
    if (service === 'oauth') {
        service = 'OAuth';
    }
    return `Uni${(0, exports.capitalize)((0, exports.camelize)(service))}${(0, exports.capitalize)((0, exports.camelize)(name))}ProviderImpl`;
}
exports.formatExtApiProviderName = formatExtApiProviderName;
function getUniExtApiProviderRegisters() {
    const result = [];
    extApiProviders.forEach((provider) => {
        if (provider.name && provider.service) {
            result.push({
                name: provider.name,
                plugin: provider.plugin,
                service: provider.service,
                class: `uts.sdk.modules.${(0, exports.camelize)(provider.plugin)}.${formatExtApiProviderName(provider.service, provider.name)}`,
            });
        }
    });
    return result;
}
exports.getUniExtApiProviderRegisters = getUniExtApiProviderRegisters;
function parseUniExtApis(vite = true, platform, language = 'javascript') {
    if (!process.env.UNI_INPUT_DIR) {
        return {};
    }
    const uniModulesDir = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'uni_modules');
    if (!fs_extra_1.default.existsSync(uniModulesDir)) {
        return {};
    }
    const injects = {};
    extApiProviders.length = 0;
    extApiPlugins.clear();
    fs_extra_1.default.readdirSync(uniModulesDir).forEach((uniModuleDir) => {
        // 必须以 uni- 开头
        if (!uniModuleDir.startsWith('uni-')) {
            return;
        }
        const uniModuleRootDir = path_1.default.resolve(uniModulesDir, uniModuleDir);
        const pkgPath = path_1.default.resolve(uniModuleRootDir, 'package.json');
        if (!fs_extra_1.default.existsSync(pkgPath)) {
            return;
        }
        try {
            let exports;
            const pkg = JSON.parse(fs_extra_1.default.readFileSync(pkgPath, 'utf8'));
            if (pkg && pkg.uni_modules && pkg.uni_modules['uni-ext-api']) {
                exports = pkg.uni_modules['uni-ext-api'];
            }
            if (exports) {
                const provider = exports.provider;
                if (provider && provider.service) {
                    provider.plugin = uniModuleDir;
                    extApiProviders.push(provider);
                }
                extApiPlugins.add(uniModuleDir);
                const curInjects = parseInjects(vite, platform, language, `@/uni_modules/${uniModuleDir}`, uniModuleRootDir, exports);
                Object.assign(injects, curInjects);
            }
        }
        catch (e) { }
    });
    return injects;
}
exports.parseUniExtApis = parseUniExtApis;
function parseUniExtApi(pluginDir, pluginId, vite = true, platform, language = 'javascript') {
    const pkgPath = path_1.default.resolve(pluginDir, 'package.json');
    if (!fs_extra_1.default.existsSync(pkgPath)) {
        return;
    }
    let exports;
    const pkg = JSON.parse(fs_extra_1.default.readFileSync(pkgPath, 'utf8'));
    if (pkg && pkg.uni_modules && pkg.uni_modules['uni-ext-api']) {
        exports = pkg.uni_modules['uni-ext-api'];
    }
    if (exports) {
        return parseInjects(vite, platform, language, `@/uni_modules/${pluginId}`, pluginDir, exports);
    }
}
exports.parseUniExtApi = parseUniExtApi;
/**
 *  uni:'getBatteryInfo'
 * import getBatteryInfo from '..'
 *
 * uni:['getBatteryInfo']
 * import { getBatteryInfo } from '..'
 *
 * uni:['openLocation','chooseLocation']
 * import { openLocation, chooseLocation } from '..'
 *
 * uni:{
 *  onUserCaptureScreen: "onCaptureScreen"
 *  offUserCaptureScreen: "offCaptureScreen"
 * }
 *
 * uni.getBatteryInfo = getBatteryInfo
 * @param source
 * @param globalObject
 * @param define
 * @returns
 */
function parseInjects(vite = true, platform, language, source, uniModuleRootDir, exports = {}) {
    if (platform === 'app-plus') {
        platform = 'app';
    }
    let rootDefines = {};
    Object.keys(exports).forEach((name) => {
        if (name.startsWith('uni')) {
            rootDefines[name] = exports[name];
        }
    });
    const injects = {};
    if (Object.keys(rootDefines).length) {
        const platformIndexFileName = path_1.default.resolve(uniModuleRootDir, 'utssdk', platform);
        const rootIndexFileName = path_1.default.resolve(uniModuleRootDir, 'utssdk', 'index.uts');
        let hasPlatformFile = uniModuleRootDir
            ? fs_extra_1.default.existsSync(rootIndexFileName) || fs_extra_1.default.existsSync(platformIndexFileName)
            : true;
        if (!hasPlatformFile) {
            if (platform === 'app') {
                hasPlatformFile =
                    fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'utssdk', 'app-android')) ||
                        fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'utssdk', 'app-ios')) ||
                        fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'utssdk', 'app-harmony'));
            }
        }
        // 其他平台修改source，直接指向目标文件，否则 uts2js 找不到类型信息
        if (platform !== 'app' &&
            platform !== 'app-android' &&
            platform !== 'app-ios' &&
            platform !== 'app-harmony') {
            if (fs_extra_1.default.existsSync(platformIndexFileName)) {
                source = `${source}/utssdk/${platform}/index.uts`;
            }
            else if (fs_extra_1.default.existsSync(rootIndexFileName)) {
                source = `${source}/utssdk/index.uts`;
            }
        }
        else if (process.env.UNI_APP_X_UVUE_SCRIPT_ENGINE === 'js') {
            if (fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'utssdk', 'app-js', 'index.uts'))) {
                source = `${source}/utssdk/app-js/index.uts`;
            }
        }
        for (const key in rootDefines) {
            Object.assign(injects, parseInject(vite, platform, language, source, 'uni', rootDefines[key], hasPlatformFile));
        }
    }
    return injects;
}
exports.parseInjects = parseInjects;
function parseInject(vite = true, platform, language, source, globalObject, define, hasPlatformFile) {
    const injects = {};
    if (define === false) {
    }
    else if (typeof define === 'string') {
        // {'uni.getBatteryInfo' : '@dcloudio/uni-getbatteryinfo'}
        if (hasPlatformFile) {
            injects[globalObject + '.' + define] = vite ? source : [source, 'default'];
        }
    }
    else if (Array.isArray(define)) {
        // {'uni.getBatteryInfo' : ['@dcloudio/uni-getbatteryinfo','getBatteryInfo]}
        if (hasPlatformFile) {
            define.forEach((d) => {
                injects[globalObject + '.' + d] = [source, d];
            });
        }
    }
    else {
        const keys = Object.keys(define);
        keys.forEach((d) => {
            if (typeof define[d] === 'string') {
                if (hasPlatformFile) {
                    injects[globalObject + '.' + d] = [source, define[d]];
                }
            }
            else {
                const defineOptions = define[d];
                const p = platform === 'app-android' ||
                    platform === 'app-ios' ||
                    platform === 'app-harmony'
                    ? 'app'
                    : platform;
                if (!(p in defineOptions)) {
                    if (hasPlatformFile) {
                        injects[globalObject + '.' + d] = [source, defineOptions.name || d];
                    }
                }
                else {
                    if (defineOptions[p] !== false) {
                        if (p === 'app') {
                            const appOptions = defineOptions.app;
                            if (isPlainObject(appOptions)) {
                                // js engine 下且存在 app-js，不检查
                                const skipCheck = process.env.UNI_APP_X_UVUE_SCRIPT_ENGINE === 'js' &&
                                    source.includes('app-js');
                                if (!skipCheck) {
                                    const targetLanguage = language === 'javascript' ? 'js' : language;
                                    if (targetLanguage && appOptions[targetLanguage] === false) {
                                        return;
                                    }
                                }
                            }
                            injects[globalObject + '.' + d] = [
                                source,
                                defineOptions.name || d,
                                defineOptions.app,
                            ];
                        }
                        else {
                            injects[globalObject + '.' + d] = [
                                source,
                                defineOptions.name || d,
                            ];
                        }
                    }
                }
            }
        });
    }
    return injects;
}
const objectToString = Object.prototype.toString;
const toTypeString = (value) => objectToString.call(value);
function isPlainObject(val) {
    return toTypeString(val) === '[object Object]';
}
const cacheStringFunction = (fn) => {
    const cache = Object.create(null);
    return ((str) => {
        const hit = cache[str];
        return hit || (cache[str] = fn(str));
    });
};
const camelizeRE = /-(\w)/g;
/**
 * @private
 */
exports.camelize = cacheStringFunction((str) => {
    return str.replace(camelizeRE, (_, c) => (c ? c.toUpperCase() : ''));
});
/**
 * @private
 */
exports.capitalize = cacheStringFunction((str) => str.charAt(0).toUpperCase() + str.slice(1));
/**
 * 解析 UTS 类型的模块依赖列表
 * @param deps
 * @param inputDir
 * @returns
 */
function parseUTSModuleDeps(deps, inputDir) {
    const modulesDir = path_1.default.resolve(inputDir, 'uni_modules');
    return deps.filter((dep) => {
        return fs_extra_1.default.existsSync(path_1.default.resolve(modulesDir, dep, 'utssdk'));
    });
}
exports.parseUTSModuleDeps = parseUTSModuleDeps;
function genEncryptEasyComModuleIndex(platform, components) {
    const imports = [];
    const ids = [];
    Object.keys(components).forEach((component) => {
        const id = (0, exports.capitalize)((0, exports.camelize)(component));
        ids.push(id);
        let instance = '';
        if (platform === 'app-android') {
            instance = (0, easycom_1.genUTSComponentPublicInstanceIdent)(component);
            // 类型
            ids.push(instance);
        }
        imports.push(`import ${id}${instance ? `, { ${instance} }` : ''} from './components/${component}/${component}${components[component]}'`);
    });
    return `
${imports.join('\n')}
export { 
  ${ids.join(',\n  ')} 
}
`;
}
exports.genEncryptEasyComModuleIndex = genEncryptEasyComModuleIndex;
// 目前该函数仅在云端使用（目前仅限iOS/web），云端编译时，提交上来的uni_modules是过滤好的
function parseUniModulesWithComponents(inputDir) {
    const modulesDir = path_1.default.resolve(inputDir, 'uni_modules');
    const uniModules = {};
    if (fs_extra_1.default.existsSync(modulesDir)) {
        fs_extra_1.default.readdirSync(modulesDir).forEach((uniModuleDir) => {
            if (!fs_extra_1.default.existsSync(path_1.default.resolve(modulesDir, uniModuleDir, 'package.json'))) {
                return;
            }
            // 解析加密的 easyCom 插件列表
            const components = parseEasyComComponents(uniModuleDir, inputDir, false);
            uniModules[uniModuleDir] = components;
        });
    }
    return uniModules;
}
exports.parseUniModulesWithComponents = parseUniModulesWithComponents;
/**
 * 解析 easyCom 组件列表
 * @param pluginId
 * @param inputDir
 * @returns
 */
function parseEasyComComponents(pluginId, inputDir, detectBinary = true) {
    const componentsDir = path_1.default.resolve(inputDir, 'uni_modules', pluginId, 'components');
    const components = {};
    if (fs_extra_1.default.existsSync(componentsDir)) {
        fs_extra_1.default.readdirSync(componentsDir).forEach((componentDir) => {
            const componentFile = path_1.default.resolve(componentsDir, componentDir, componentDir);
            const extname = ['.vue', '.uvue'].find((extname) => {
                const filename = componentFile + extname;
                // 探测 filename 是否是二进制文件
                if (fs_extra_1.default.existsSync(filename)) {
                    if (detectBinary) {
                        // 延迟require，这个是新增的依赖，无法及时同步到内部测试版本HBuilderX中，导致报错，所以延迟require吧
                        if (require('isbinaryfile').isBinaryFileSync(filename)) {
                            return true;
                        }
                    }
                    else {
                        return true;
                    }
                }
            });
            if (extname) {
                components[componentDir] = extname;
            }
        });
    }
    return components;
}
exports.parseEasyComComponents = parseEasyComComponents;
// 查找所有普通加密插件 uni_modules
function findEncryptUniModules(inputDir, cacheDir = '') {
    const modulesDir = path_1.default.resolve(inputDir, 'uni_modules');
    const uniModules = {};
    if (fs_extra_1.default.existsSync(modulesDir)) {
        fs_extra_1.default.readdirSync(modulesDir).forEach((uniModuleDir) => {
            const uniModuleRootDir = path_1.default.resolve(modulesDir, uniModuleDir);
            if (!fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'encrypt'))) {
                return;
            }
            // 仅扫描普通加密插件，无需依赖
            if (fs_extra_1.default.existsSync(path_1.default.resolve(uniModuleRootDir, 'utssdk'))) {
                return;
            }
            const pkg = require(path_1.default.resolve(uniModuleRootDir, 'package.json'));
            uniModules[uniModuleDir] = findEncryptUniModuleCache(uniModuleDir, cacheDir, { version: pkg.version, env: initCheckEnv() });
        });
    }
    return uniModules;
}
exports.findEncryptUniModules = findEncryptUniModules;
function findUploadEncryptUniModulesFiles(uniModules, platform, inputDir) {
    const modules = {};
    Object.keys(uniModules).forEach((uniModuleId) => {
        if (!uniModules[uniModuleId]) {
            modules[uniModuleId] = findUniModuleFiles(platform, uniModuleId, inputDir);
        }
    });
    return modules;
}
exports.findUploadEncryptUniModulesFiles = findUploadEncryptUniModulesFiles;
function packUploadEncryptUniModules(uniModules, platform, inputDir, cacheDir) {
    const modules = findUploadEncryptUniModulesFiles(uniModules, platform, inputDir);
    const uploadModuleIds = Object.keys(modules);
    if (uploadModuleIds.length) {
        // 延迟 require，避免 vue2 编译器需要安装此依赖，目前该方法仅在 vite 编译器中使用
        const AdmZip = require('adm-zip');
        const zip = new AdmZip();
        uploadModuleIds.forEach((moduleId) => {
            modules[moduleId].forEach((file) => {
                zip.addLocalFile(file, path_1.default.dirname(path_1.default.relative(inputDir, file)));
            });
        });
        const zipFile = path_1.default.resolve(cacheDir, 'cloud-compile-plugins.zip');
        zip.writeZip(zipFile);
        return {
            zipFile,
            modules: uploadModuleIds,
        };
    }
    return {
        zipFile: '',
        modules: [],
    };
}
exports.packUploadEncryptUniModules = packUploadEncryptUniModules;
function isEnvExpired(value, other) {
    const valueKeys = Object.keys(value);
    const otherKeys = Object.keys(other);
    if (valueKeys.length !== otherKeys.length) {
        return true;
    }
    if (valueKeys.find((name) => value[name] !== other[name])) {
        return true;
    }
    return false;
}
function findEncryptUniModuleCache(uniModuleId, cacheDir, options) {
    if (!cacheDir) {
        return;
    }
    const uniModuleCacheDir = path_1.default.resolve(cacheDir, 'uni_modules', uniModuleId);
    if (fs_extra_1.default.existsSync(uniModuleCacheDir)) {
        const pkg = require(path_1.default.resolve(uniModuleCacheDir, 'package.json'));
        // 插件版本以及各种环境一致
        if (pkg.version === options.version &&
            !isEnvExpired(pkg.uni_modules?.artifacts?.env || {}, options.env)) {
            const declaration = path_1.default.resolve(uniModuleCacheDir, 'utssdk/app-android/index.d.uts');
            pkg.uni_modules.artifacts.declaration = fs_extra_1.default.existsSync(declaration)
                ? declaration
                : '';
            return pkg;
        }
        console.log(`插件${uniModuleId} 缓存已过期，需要重新云编译。`);
        // 已过期的插件，删除缓存
        fs_extra_1.default.rmSync(uniModuleCacheDir, { recursive: true });
    }
}
const KNOWN_ASSET_TYPES = [
    // images
    'png',
    'jpe?g',
    'gif',
    'svg',
    'ico',
    'webp',
    'avif',
    // media
    'mp4',
    'webm',
    'ogg',
    'mp3',
    'wav',
    'flac',
    'aac',
    // fonts
    'woff2?',
    'eot',
    'ttf',
    'otf',
    // other
    'pdf',
    'txt',
];
function findUniModuleFiles(platform, id, inputDir) {
    return (0, fast_glob_1.sync)(`uni_modules/${id}/**/*`, {
        cwd: inputDir,
        absolute: true,
        ignore: [
            '**/*.md',
            ...(platform !== 'app-android' // 非 android 平台不需要扫描 assets
                ? [`**/*.{${KNOWN_ASSET_TYPES.join(',')}}`]
                : []),
        ],
    });
}
function initCheckEnv() {
    return {
        // 云端编译的版本号不带日期及小版本
        compilerVersion: process.env.UNI_COMPILER_VERSION,
    };
}
exports.initCheckEnv = initCheckEnv;
function findLastIndex(array, predicate) {
    for (let i = array.length - 1; i >= 0; i--) {
        if (predicate(array[i], i, array)) {
            return i;
        }
    }
    return -1;
}
let encryptUniModules = {};
function resolveEncryptUniModule(id, platform, isX = true) {
    const parts = id.split('?', 2)[0].split('/');
    const index = findLastIndex(parts, (part) => part === 'uni_modules');
    if (index !== -1) {
        const uniModuleId = parts[index + 1];
        if (uniModuleId in encryptUniModules) {
            if (parts[index + 2]) {
                console.warn(messages_1.M['uni_modules.import']
                    .replace('{0}', uniModuleId)
                    .replace('{1}', uniModuleId)
                    .replace('{2}', parts.slice(index + 2).join('/')));
            }
            // 原生平台走旧的uts-proxy
            return (0, utils_1.normalizePath)(path_1.default.join(process.env.UNI_INPUT_DIR, `uni_modules/${uniModuleId}?${isX && platform === 'app-android' ? 'uts-proxy' : 'uni_helpers'}`));
        }
    }
}
exports.resolveEncryptUniModule = resolveEncryptUniModule;
async function checkEncryptUniModules(inputDir, params) {
    // 扫描加密插件云编译
    encryptUniModules = findEncryptUniModules(inputDir, process.env.UNI_MODULES_ENCRYPT_CACHE_DIR);
    if (!Object.keys(encryptUniModules).length) {
        return {};
    }
    if (!process.env.UNI_HBUILDERX_PLUGINS) {
        return {};
    }
    const cacheDir = process.env.UNI_MODULES_ENCRYPT_CACHE_DIR;
    const { zipFile, modules } = packUploadEncryptUniModules(encryptUniModules, process.env.UNI_UTS_PLATFORM, inputDir, cacheDir);
    if (zipFile) {
        const downloadFile = path_1.default.resolve(cacheDir, 'uni_modules.download.zip');
        const { C, D, R, U } = (0, utils_1.requireUniHelpers)();
        try {
            const isLogin = await C();
            const tips = process.env.UNI_UTS_PLATFORM !== 'app-android'
                ? '（此过程耗时较长）'
                : '';
            console.log(`正在云编译插件${isLogin ? '' : '（请先登录）'}${tips}：${modules.join(',')}...`);
            let downloadUrl = '';
            try {
                downloadUrl = await U({
                    params,
                    attachment: zipFile,
                });
            }
            catch (e) {
                if (e.message && e.message === '{"error":"UserNotLogin"}') {
                    console.log('当前项目包含需要云编译的付费插件，需要您先登录HBuilderX账号。');
                }
                else {
                    console.error(e);
                }
                process.exit(0);
            }
            await D(downloadUrl, downloadFile);
            // unzip
            const AdmZip = require('adm-zip');
            const zip = new AdmZip(downloadFile);
            zip.extractAllTo(cacheDir, true);
            fs_extra_1.default.unlinkSync(zipFile);
            fs_extra_1.default.unlinkSync(downloadFile);
            R({
                dir: process.env.UNI_INPUT_DIR,
                cacheDir: process.env.UNI_MODULES_ENCRYPT_CACHE_DIR,
            });
            console.log(`云编译已完成`);
            console.log(`正在编译中...`);
        }
        catch (e) {
            fs_extra_1.default.existsSync(zipFile) && fs_extra_1.default.unlinkSync(zipFile);
            fs_extra_1.default.existsSync(downloadFile) && fs_extra_1.default.unlinkSync(downloadFile);
            console.error(e);
            process.exit(0);
        }
    }
    else {
        // android 平台需要在这里初始化
        if (params.platform === 'app-android') {
            const { R } = (0, utils_1.requireUniHelpers)();
            R({
                dir: process.env.UNI_INPUT_DIR,
                cacheDir: process.env.UNI_MODULES_ENCRYPT_CACHE_DIR,
            });
        }
    }
    encryptUniModules = findEncryptUniModules(inputDir, process.env.UNI_MODULES_ENCRYPT_CACHE_DIR);
}
exports.checkEncryptUniModules = checkEncryptUniModules;
function parseUniModulesArtifacts() {
    const res = [];
    Object.keys(encryptUniModules).forEach((uniModuleId) => {
        const pkg = encryptUniModules[uniModuleId];
        if (pkg?.uni_modules?.artifacts) {
            res.push({
                name: uniModuleId,
                package: `uts.sdk.modules.${(0, exports.camelize)(uniModuleId)}`,
                scopedSlots: pkg.uni_modules.artifacts.scopedSlots || [],
                declaration: pkg.uni_modules.artifacts.declaration,
            });
        }
    });
    return res;
}
exports.parseUniModulesArtifacts = parseUniModulesArtifacts;
