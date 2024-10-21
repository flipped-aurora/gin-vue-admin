"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.genUTSComponentPublicInstanceImported = exports.genUTSComponentPublicInstanceIdent = exports.addUTSEasyComAutoImports = exports.getUTSEasyComAutoImports = exports.UNI_EASYCOM_EXCLUDE = exports.genResolveEasycomCode = exports.addImportDeclaration = exports.matchEasycom = exports.initEasycomsOnce = exports.initEasycoms = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const shared_1 = require("@vue/shared");
const pluginutils_1 = require("@rollup/pluginutils");
const uni_shared_1 = require("@dcloudio/uni-shared");
const utils_1 = require("./utils");
const pages_1 = require("./json/pages");
const messages_1 = require("./messages");
const uts_1 = require("./uts");
const utsUtils_1 = require("./utsUtils");
const debugEasycom = (0, debug_1.default)('uni:easycom');
const easycoms = [];
const easycomsCache = new Map();
const easycomsInvalidCache = new Set();
let hasEasycom = false;
function clearEasycom() {
    easycoms.length = 0;
    easycomsCache.clear();
    easycomsInvalidCache.clear();
}
function initEasycoms(inputDir, { dirs, platform, isX, }) {
    const componentsDir = path_1.default.resolve(inputDir, 'components');
    const uniModulesDir = path_1.default.resolve(inputDir, 'uni_modules');
    const initEasycomOptions = (pagesJson) => {
        // 初始化时，从once中读取缓存，refresh时，实时读取
        const { easycom } = pagesJson || (0, pages_1.parsePagesJson)(inputDir, platform, false);
        const easycomOptions = {
            isX,
            dirs: easycom && easycom.autoscan === false
                ? [...dirs] // 禁止自动扫描
                : [
                    ...dirs,
                    componentsDir,
                    ...initUniModulesEasycomDirs(uniModulesDir),
                ],
            rootDir: inputDir,
            autoscan: !!(easycom && easycom.autoscan),
            custom: (easycom && easycom.custom) || {},
            extensions: [...(isX ? ['.uvue'] : []), ...['.vue', '.jsx', '.tsx']],
        };
        debugEasycom(easycomOptions);
        return easycomOptions;
    };
    const options = initEasycomOptions((0, pages_1.parsePagesJsonOnce)(inputDir, platform));
    const initUTSEasycom = () => {
        (0, uts_1.initUTSComponents)(inputDir, platform).forEach((item) => {
            const index = easycoms.findIndex((easycom) => item.name === easycom.name);
            if (index > -1) {
                easycoms.splice(index, 1, item);
            }
            else {
                easycoms.push(item);
            }
        });
        if (isX && globalThis.uts2jsSourceCodeMap) {
            ;
            globalThis.uts2jsSourceCodeMap.initUts2jsEasycom(easycoms);
        }
    };
    initEasycom(options);
    initUTSEasycom();
    const componentExtNames = isX ? 'uvue|vue' : 'vue';
    const res = {
        options,
        filter: (0, pluginutils_1.createFilter)([
            'components/*/*.(' + componentExtNames + '|jsx|tsx)',
            'uni_modules/*/components/*/*.(' + componentExtNames + '|jsx|tsx)',
            'utssdk/*/**/*.(' + componentExtNames + ')',
            'uni_modules/*/utssdk/*/*.(' + componentExtNames + ')',
        ], [], {
            resolve: inputDir,
        }),
        refresh() {
            res.options = initEasycomOptions();
            initEasycom(res.options);
            initUTSEasycom();
        },
        easycoms,
    };
    return res;
}
exports.initEasycoms = initEasycoms;
exports.initEasycomsOnce = (0, uni_shared_1.once)(initEasycoms);
function initUniModulesEasycomDirs(uniModulesDir) {
    if (!fs_1.default.existsSync(uniModulesDir)) {
        return [];
    }
    return fs_1.default
        .readdirSync(uniModulesDir)
        .map((uniModuleDir) => {
        const uniModuleComponentsDir = path_1.default.resolve(uniModulesDir, uniModuleDir, 'components');
        if (fs_1.default.existsSync(uniModuleComponentsDir)) {
            return uniModuleComponentsDir;
        }
    })
        .filter(Boolean);
}
function initEasycom({ isX, dirs, rootDir, custom, extensions, }) {
    clearEasycom();
    rootDir = (0, utils_1.normalizePath)(rootDir);
    const easycomsObj = Object.create(null);
    if (dirs && dirs.length && rootDir) {
        const autoEasyComObj = initAutoScanEasycoms(dirs, rootDir, extensions);
        if (isX) {
            Object.keys(autoEasyComObj).forEach((tagName) => {
                let source = autoEasyComObj[tagName];
                tagName = tagName.slice(1, -1);
                if (path_1.default.isAbsolute(source) && source.startsWith(rootDir)) {
                    source = '@/' + (0, utils_1.normalizePath)(path_1.default.relative(rootDir, source));
                }
                // 加密插件easycom类型导入
                if (source.includes('?uts-proxy')) {
                    const moduleId = path_1.default.basename(source.split('?uts-proxy')[0]);
                    source = `uts.sdk.modules.${(0, shared_1.camelize)(moduleId)}`;
                }
                const ident = genUTSComponentPublicInstanceIdent(tagName);
                addUTSEasyComAutoImports(source, [ident, ident]);
            });
        }
        (0, shared_1.extend)(easycomsObj, autoEasyComObj);
    }
    if (custom) {
        Object.keys(custom).forEach((name) => {
            const componentPath = custom[name];
            easycomsObj[name] = componentPath.startsWith('@/')
                ? (0, utils_1.normalizePath)(path_1.default.join(rootDir, componentPath.slice(2)))
                : componentPath;
        });
    }
    Object.keys(easycomsObj).forEach((name) => {
        easycoms.push({
            name: name.startsWith('^') && name.endsWith('$') ? name.slice(1, -1) : name,
            pattern: new RegExp(name),
            replacement: easycomsObj[name],
        });
    });
    debugEasycom(easycoms);
    hasEasycom = !!easycoms.length;
    return easycoms;
}
function matchEasycom(tag) {
    if (!hasEasycom) {
        return;
    }
    let source = easycomsCache.get(tag);
    if (source) {
        return source;
    }
    if (easycomsInvalidCache.has(tag)) {
        return false;
    }
    const matcher = easycoms.find((matcher) => matcher.pattern.test(tag));
    if (!matcher) {
        easycomsInvalidCache.add(tag);
        return false;
    }
    source = tag.replace(matcher.pattern, matcher.replacement);
    easycomsCache.set(tag, source);
    debugEasycom('matchEasycom', tag, source);
    return source;
}
exports.matchEasycom = matchEasycom;
const isDir = (path) => {
    const stat = fs_1.default.lstatSync(path);
    if (stat.isDirectory()) {
        return true;
    }
    else if (stat.isSymbolicLink()) {
        return fs_1.default.lstatSync(fs_1.default.realpathSync(path)).isDirectory();
    }
    return false;
};
function initAutoScanEasycom(dir, rootDir, extensions) {
    if (!path_1.default.isAbsolute(dir)) {
        dir = path_1.default.resolve(rootDir, dir);
    }
    const easycoms = Object.create(null);
    if (!fs_1.default.existsSync(dir)) {
        return easycoms;
    }
    const is_uni_modules = path_1.default.basename(path_1.default.resolve(dir, '../..')) === 'uni_modules';
    const is_encrypt_uni_modules = // uni_modules模式不需要此逻辑
     process.env.UNI_COMPILE_TARGET !== 'uni_modules' &&
        is_uni_modules &&
        fs_1.default.existsSync(path_1.default.resolve(dir, '../encrypt'));
    const uni_modules_plugin_id = is_encrypt_uni_modules && path_1.default.basename(path_1.default.resolve(dir, '..'));
    fs_1.default.readdirSync(dir).forEach((name) => {
        const folder = path_1.default.resolve(dir, name);
        if (!isDir(folder)) {
            return;
        }
        const importDir = (0, utils_1.normalizePath)(folder);
        const files = fs_1.default.readdirSync(folder);
        // 读取文件夹文件列表，比对文件名（fs.existsSync在大小写不敏感的系统会匹配不准确）
        for (let i = 0; i < extensions.length; i++) {
            const ext = extensions[i];
            if (files.includes(name + ext)) {
                easycoms[`^${name}$`] = is_encrypt_uni_modules
                    ? (0, utils_1.normalizePath)(path_1.default.join(rootDir, `uni_modules/${uni_modules_plugin_id}?${
                    // android 走 proxy
                    process.env.UNI_APP_X === 'true' &&
                        process.env.UNI_UTS_PLATFORM === 'app-android'
                        ? 'uts-proxy'
                        : 'uni_helpers'}`))
                    : `${importDir}/${name}${ext}`;
                break;
            }
        }
    });
    return easycoms;
}
function initAutoScanEasycoms(dirs, rootDir, extensions) {
    const conflict = {};
    const res = dirs.reduce((easycoms, dir) => {
        const curEasycoms = initAutoScanEasycom(dir, rootDir, extensions);
        Object.keys(curEasycoms).forEach((name) => {
            // Use the first component when name conflict
            const componentPath = easycoms[name];
            if (!componentPath) {
                easycoms[name] = curEasycoms[name];
            }
            else {
                ;
                (conflict[componentPath] || (conflict[componentPath] = [])).push(normalizeComponentPath(curEasycoms[name], rootDir));
            }
        });
        return easycoms;
    }, Object.create(null));
    const conflictComponents = Object.keys(conflict);
    if (conflictComponents.length) {
        console.warn(messages_1.M['easycom.conflict']);
        conflictComponents.forEach((com) => {
            console.warn([normalizeComponentPath(com, rootDir), conflict[com]].join(','));
        });
    }
    return res;
}
function normalizeComponentPath(componentPath, rootDir) {
    return (0, utils_1.normalizePath)(path_1.default.relative(rootDir, componentPath));
}
function addImportDeclaration(importDeclarations, local, source, imported) {
    importDeclarations.push(createImportDeclaration(local, source, imported));
    return local;
}
exports.addImportDeclaration = addImportDeclaration;
function createImportDeclaration(local, source, imported) {
    if (imported) {
        return `import { ${imported} as ${local} } from '${source}';`;
    }
    return `import ${local} from '${source}';`;
}
const RESOLVE_EASYCOM_IMPORT_CODE = `import { resolveDynamicComponent as __resolveDynamicComponent } from 'vue';import { resolveEasycom } from '@dcloudio/uni-app';`;
function genResolveEasycomCode(importDeclarations, code, name) {
    if (!importDeclarations.includes(RESOLVE_EASYCOM_IMPORT_CODE)) {
        importDeclarations.push(RESOLVE_EASYCOM_IMPORT_CODE);
    }
    return `resolveEasycom(${code.replace('_resolveComponent', '__resolveDynamicComponent')}, ${name})`;
}
exports.genResolveEasycomCode = genResolveEasycomCode;
exports.UNI_EASYCOM_EXCLUDE = [/@dcloudio\/uni-h5/];
const utsEasyComAutoImports = {};
function getUTSEasyComAutoImports() {
    return utsEasyComAutoImports;
}
exports.getUTSEasyComAutoImports = getUTSEasyComAutoImports;
function addUTSEasyComAutoImports(source, imports) {
    if (!utsEasyComAutoImports[source]) {
        utsEasyComAutoImports[source] = [imports];
    }
    else {
        if (!utsEasyComAutoImports[source].find((item) => item[0] === imports[0])) {
            utsEasyComAutoImports[source].push(imports);
        }
    }
}
exports.addUTSEasyComAutoImports = addUTSEasyComAutoImports;
function genUTSComponentPublicInstanceIdent(tagName) {
    return (0, shared_1.capitalize)((0, shared_1.camelize)(tagName)) + 'ComponentPublicInstance';
}
exports.genUTSComponentPublicInstanceIdent = genUTSComponentPublicInstanceIdent;
function genUTSComponentPublicInstanceImported(root, fileName) {
    root = (0, utils_1.normalizePath)(root);
    if (path_1.default.isAbsolute(fileName) && fileName.startsWith(root)) {
        fileName = (0, utils_1.normalizePath)(path_1.default.relative(root, fileName));
    }
    if (fileName.startsWith('@/')) {
        return ((0, utsUtils_1.genUTSClassName)(fileName.replace('@/', '')) + 'ComponentPublicInstance');
    }
    return (0, utsUtils_1.genUTSClassName)(fileName) + 'ComponentPublicInstance';
}
exports.genUTSComponentPublicInstanceImported = genUTSComponentPublicInstanceImported;
