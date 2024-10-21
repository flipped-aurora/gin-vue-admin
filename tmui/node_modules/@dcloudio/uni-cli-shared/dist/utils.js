"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createShadowImageUrl = exports.normalizeEmitAssetFileName = exports.requireUniHelpers = exports.enableSourceMap = exports.createResolveErrorMsg = exports.parseImporter = exports.resolveAppVue = exports.isAppVue = exports.installDepTips = exports.resolveSourceMapPath = exports.pathToGlob = exports.normalizeParsePlugins = exports.normalizeMiniProgramFilename = exports.normalizeNodeModules = exports.removeExt = exports.normalizePagePath = exports.normalizeIdentifier = exports.checkElementNodeTag = exports.normalizePath = exports.isWindows = exports.isArray = exports.capitalize = exports.camelize = exports.hash = void 0;
const fs_1 = __importDefault(require("fs"));
const os_1 = __importDefault(require("os"));
const path_1 = __importDefault(require("path"));
const picocolors_1 = __importDefault(require("picocolors"));
const shared_1 = require("@vue/shared");
var hash_sum_1 = require("hash-sum");
Object.defineProperty(exports, "hash", { enumerable: true, get: function () { return __importDefault(hash_sum_1).default; } });
const constants_1 = require("./constants");
const compiler_core_1 = require("@vue/compiler-core");
const platform_1 = require("./platform");
const hbx_1 = require("./hbx");
const json_1 = require("./json");
// 专为 uts.ts 服务
var shared_2 = require("@vue/shared");
Object.defineProperty(exports, "camelize", { enumerable: true, get: function () { return shared_2.camelize; } });
Object.defineProperty(exports, "capitalize", { enumerable: true, get: function () { return shared_2.capitalize; } });
Object.defineProperty(exports, "isArray", { enumerable: true, get: function () { return shared_2.isArray; } });
// export let isRunningWithYarnPnp: boolean
// try {
//   isRunningWithYarnPnp = Boolean(require('pnpapi'))
// } catch {}
exports.isWindows = os_1.default.platform() === 'win32';
function normalizePath(id) {
    return exports.isWindows ? id.replace(/\\/g, '/') : id;
}
exports.normalizePath = normalizePath;
function checkElementNodeTag(node, tag) {
    return !!node && node.type === compiler_core_1.NodeTypes.ELEMENT && node.tag === tag;
}
exports.checkElementNodeTag = checkElementNodeTag;
/**
 * 根据 path 返回合法 js 变量
 * @param str pages.json.page.path
 * @returns
 */
function normalizeIdentifier(str) {
    let _str = str.replace(/[^a-zA-Z0-9]+/g, '-');
    _str = (0, shared_1.capitalize)((0, shared_1.camelize)(_str));
    // 不允许数字开头，补充 _
    if (/^\d/.test(_str)) {
        _str = '_' + _str;
    }
    return _str;
}
exports.normalizeIdentifier = normalizeIdentifier;
function normalizePagePath(pagePath, platform) {
    const absolutePagePath = path_1.default.resolve(process.env.UNI_INPUT_DIR, pagePath);
    const isX = process.env.UNI_APP_X === 'true';
    let extensions = isX ? constants_1.X_PAGE_EXTNAME : constants_1.PAGE_EXTNAME;
    if (platform === 'app') {
        extensions = isX ? constants_1.X_PAGE_EXTNAME_APP : constants_1.PAGE_EXTNAME_APP;
    }
    for (let i = 0; i < extensions.length; i++) {
        const extname = extensions[i];
        if (fs_1.default.existsSync(absolutePagePath + extname)) {
            return pagePath + extname;
        }
    }
    console.error(`${pagePath} not found`);
}
exports.normalizePagePath = normalizePagePath;
function removeExt(str) {
    return str.split('?')[0].replace(/\.\w+$/g, '');
}
exports.removeExt = removeExt;
const NODE_MODULES_REGEX = /(\.\.\/)?node_modules/g;
function normalizeNodeModules(str) {
    str = normalizePath(str).replace(NODE_MODULES_REGEX, 'node-modules');
    // HBuilderX 内置模块路径转换
    str = str.replace(/.*\/plugins\/uniapp-cli-vite\/node[-_]modules/, 'node-modules');
    if (!(0, hbx_1.isInHBuilderX)()) {
        // 内部测试
        if (str.includes('uni-app-next/packages/')) {
            str = str.replace(/.*\/uni-app-next\/packages\//, 'node-modules/@dcloudio/');
        }
    }
    if (process.env.UNI_PLATFORM === 'mp-alipay') {
        str = str.replace('node-modules/@', 'node-modules/npm-scope-');
    }
    return str;
}
exports.normalizeNodeModules = normalizeNodeModules;
function normalizeMiniProgramFilename(filename, inputDir) {
    if (!inputDir || !path_1.default.isAbsolute(filename)) {
        return normalizeNodeModules(filename);
    }
    return normalizeNodeModules(path_1.default.relative(inputDir, filename));
}
exports.normalizeMiniProgramFilename = normalizeMiniProgramFilename;
function normalizeParsePlugins(importer, babelParserPlugins) {
    const isTS = constants_1.EXTNAME_TS_RE.test(importer.split('?')[0]);
    const plugins = [];
    if (isTS) {
        plugins.push('jsx');
    }
    if (babelParserPlugins)
        plugins.push(...babelParserPlugins);
    if (isTS)
        plugins.push('typescript', 'decorators-legacy');
    return plugins;
}
exports.normalizeParsePlugins = normalizeParsePlugins;
function pathToGlob(pathString, glob, options = {}) {
    const isWindows = 'windows' in options ? options.windows : /^win/.test(process.platform);
    const useEscape = options.escape;
    const str = isWindows ? pathString.replace(/\\/g, '/') : pathString;
    let safeStr = str.replace(/[\\*?[\]{}()!]/g, isWindows || !useEscape ? '[$&]' : '\\$&');
    return path_1.default.posix.join(safeStr, glob);
}
exports.pathToGlob = pathToGlob;
function resolveSourceMapPath(outputDir, platform) {
    return path_1.default.resolve(outputDir || process.env.UNI_OUTPUT_DIR, '../.sourcemap/' + (platform || (0, platform_1.getPlatformDir)()));
}
exports.resolveSourceMapPath = resolveSourceMapPath;
function hasProjectYarn(cwd) {
    return fs_1.default.existsSync(path_1.default.join(cwd, 'yarn.lock'));
}
function hasProjectPnpm(cwd) {
    return fs_1.default.existsSync(path_1.default.join(cwd, 'pnpm-lock.yaml'));
}
function getInstallCommand(cwd) {
    return hasProjectYarn(cwd)
        ? 'yarn add'
        : hasProjectPnpm(cwd)
            ? 'pnpm i'
            : 'npm i';
}
function installDepTips(type, module, version) {
    return `Cannot find module: ${module}
Please run \`${picocolors_1.default.cyan(`${getInstallCommand(process.cwd())} ${module + (version ? '@' + version : '')}${type === 'devDependencies' ? ' -D' : ''}`)}\` and try again.`;
}
exports.installDepTips = installDepTips;
/**
 * 根据路径判断是否为 App.(u?)vue
 * @param {string} filename 相对、绝对路径
 * @returns
 */
function isAppVue(filename) {
    const _filePath = normalizePath(filename);
    return /(\/|\\)app\.(u?)vue$/.test(_filePath.toLowerCase());
}
exports.isAppVue = isAppVue;
function resolveAppVue(inputDir) {
    if (process.env.UNI_APP_X === 'true') {
        const appUVue = path_1.default.resolve(inputDir, 'App.uvue');
        if (fs_1.default.existsSync(appUVue)) {
            return normalizePath(appUVue);
        }
    }
    return normalizePath(path_1.default.resolve(inputDir, 'App.vue'));
}
exports.resolveAppVue = resolveAppVue;
function parseImporter(importer) {
    importer = importer.split('?')[0];
    if (path_1.default.isAbsolute(importer)) {
        return normalizePath(path_1.default.relative(process.env.UNI_INPUT_DIR, importer));
    }
    return importer;
}
exports.parseImporter = parseImporter;
function createResolveErrorMsg(source, importer) {
    return `Could not resolve "${source}" from "${parseImporter(importer)}"`;
}
exports.createResolveErrorMsg = createResolveErrorMsg;
function enableSourceMap() {
    return (process.env.NODE_ENV === 'development' &&
        process.env.UNI_COMPILE_TARGET !== 'uni_modules');
}
exports.enableSourceMap = enableSourceMap;
function requireUniHelpers() {
    require(path_1.default.resolve(process.env.UNI_HBUILDERX_PLUGINS, 'uni_helpers/lib/bytenode'));
    return require(path_1.default.join(process.env.UNI_HBUILDERX_PLUGINS, 'uni_helpers'));
}
exports.requireUniHelpers = requireUniHelpers;
function normalizeEmitAssetFileName(fileName) {
    const extname = path_1.default.extname(fileName);
    if (process.env.UNI_APP_X_TSC === 'true') {
        if (extname !== '.ts') {
            return fileName + '.ts';
        }
    }
    else {
        // logo.png、pages.json 等
        if (!['.ts', '.uts', '.uvue', '.vue'].includes(extname)) {
            fileName = fileName + '.uts';
        }
    }
    return fileName;
}
exports.normalizeEmitAssetFileName = normalizeEmitAssetFileName;
function createIdent() {
    if (process.env.UNI_INPUT_DIR) {
        const manifestJson = (0, json_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR);
        const id = (manifestJson.appid || '').replace('__UNI__', '');
        if (id) {
            return Buffer.from(Buffer.from(id).toString('base64')).toString('hex');
        }
    }
    return '';
}
function createShadowImageUrl(cdn, type = 'grey') {
    let identStr = '';
    if (process.env.UNI_PLATFORM !== 'h5' && process.env.UNI_PLATFORM !== 'web') {
        const ident = createIdent();
        identStr = ident ? `${ident}/` : '';
    }
    return `https://cdn${(cdn || 0) + (process.env.UNI_APP_X === 'true' ? 1000 : 0) || ''}.dcloud.net.cn/${identStr}img/shadow-${type}.png`;
}
exports.createShadowImageUrl = createShadowImageUrl;
