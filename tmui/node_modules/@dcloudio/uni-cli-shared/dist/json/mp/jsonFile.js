"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.findUsingComponentsJson = exports.findMiniProgramUsingComponents = exports.isMiniProgramUsingComponent = exports.addMiniProgramUsingComponents = exports.addMiniProgramComponentJson = exports.addMiniProgramPageJson = exports.addMiniProgramAppJson = exports.findChangedJsonFiles = exports.normalizeJsonFilename = exports.findUsingComponents = exports.findJsonFile = exports.getComponentJsonFilenames = exports.hasJsonFile = exports.isMiniProgramPageSfcFile = exports.isMiniProgramPageFile = void 0;
const path_1 = __importDefault(require("path"));
const fs_1 = __importDefault(require("fs"));
const shared_1 = require("@vue/shared");
const utils_1 = require("../../utils");
const resolve_1 = require("../../resolve");
const utils_2 = require("../../vue/utils");
let appJsonCache = {};
const jsonFilesCache = new Map();
const jsonPagesCache = new Map();
const jsonComponentsCache = new Map();
const jsonUsingComponentsCache = new Map();
function isMiniProgramPageFile(file, inputDir) {
    if (inputDir && path_1.default.isAbsolute(file)) {
        file = (0, utils_1.normalizePath)(path_1.default.relative(inputDir, file));
    }
    return jsonPagesCache.has((0, utils_1.removeExt)(file));
}
exports.isMiniProgramPageFile = isMiniProgramPageFile;
function isMiniProgramPageSfcFile(file, inputDir) {
    return (0, utils_2.isVueSfcFile)(file) && isMiniProgramPageFile(file, inputDir);
}
exports.isMiniProgramPageSfcFile = isMiniProgramPageSfcFile;
function hasJsonFile(filename) {
    return (filename === 'app' ||
        jsonPagesCache.has(filename) ||
        jsonComponentsCache.has(filename));
}
exports.hasJsonFile = hasJsonFile;
function getComponentJsonFilenames() {
    return [...jsonComponentsCache.keys()];
}
exports.getComponentJsonFilenames = getComponentJsonFilenames;
function findJsonFile(filename) {
    if (filename === 'app') {
        return appJsonCache;
    }
    return jsonPagesCache.get(filename) || jsonComponentsCache.get(filename);
}
exports.findJsonFile = findJsonFile;
function findUsingComponents(filename) {
    return jsonUsingComponentsCache.get(filename);
}
exports.findUsingComponents = findUsingComponents;
function normalizeJsonFilename(filename) {
    return (0, utils_1.normalizeNodeModules)(filename);
}
exports.normalizeJsonFilename = normalizeJsonFilename;
function findChangedJsonFiles(supportGlobalUsingComponents = true) {
    const changedJsonFiles = new Map();
    function findChangedFile(filename, json) {
        const newJson = JSON.parse(JSON.stringify(json));
        if (!newJson.usingComponents) {
            newJson.usingComponents = {};
        }
        (0, shared_1.extend)(newJson.usingComponents, jsonUsingComponentsCache.get(filename));
        // 格式化为相对路径，这样作为分包也可以直接运行
        // app.json mp-baidu 在 win 不支持相对路径。所有平台改用绝对路径
        if (filename !== 'app') {
            let usingComponents = newJson.usingComponents;
            // 如果小程序不支持 global 的 usingComponents
            if (!supportGlobalUsingComponents) {
                // 从取全局的 usingComponents 并补充到子组件 usingComponents 中
                const globalUsingComponents = appJsonCache?.usingComponents || {};
                const globalComponents = findUsingComponents('app') || {};
                usingComponents = {
                    ...globalUsingComponents,
                    ...globalComponents,
                    ...newJson.usingComponents,
                };
            }
            Object.keys(usingComponents).forEach((name) => {
                const componentFilename = usingComponents[name];
                if (componentFilename.startsWith('/')) {
                    usingComponents[name] = (0, resolve_1.relativeFile)(filename, componentFilename.slice(1));
                }
            });
            newJson.usingComponents = usingComponents;
        }
        const jsonStr = JSON.stringify(newJson, null, 2);
        if (jsonFilesCache.get(filename) !== jsonStr) {
            changedJsonFiles.set(filename, jsonStr);
            jsonFilesCache.set(filename, jsonStr);
        }
    }
    function findChangedFiles(jsonsCache) {
        for (const name of jsonsCache.keys()) {
            findChangedFile(name, jsonsCache.get(name));
        }
    }
    findChangedFile('app', appJsonCache);
    findChangedFiles(jsonPagesCache);
    findChangedFiles(jsonComponentsCache);
    return changedJsonFiles;
}
exports.findChangedJsonFiles = findChangedJsonFiles;
function addMiniProgramAppJson(appJson) {
    appJsonCache = appJson;
}
exports.addMiniProgramAppJson = addMiniProgramAppJson;
function addMiniProgramPageJson(filename, json) {
    jsonPagesCache.set(filename, json);
}
exports.addMiniProgramPageJson = addMiniProgramPageJson;
function addMiniProgramComponentJson(filename, json) {
    jsonComponentsCache.set(filename, json);
}
exports.addMiniProgramComponentJson = addMiniProgramComponentJson;
function addMiniProgramUsingComponents(filename, json) {
    jsonUsingComponentsCache.set(filename, json);
}
exports.addMiniProgramUsingComponents = addMiniProgramUsingComponents;
function isMiniProgramUsingComponent(name, options) {
    return !!findMiniProgramUsingComponents(options)[name];
}
exports.isMiniProgramUsingComponent = isMiniProgramUsingComponent;
function findMiniProgramUsingComponents({ filename, inputDir, componentsDir, }) {
    const globalUsingComponents = appJsonCache && appJsonCache.usingComponents;
    const miniProgramComponents = {};
    if (globalUsingComponents) {
        (0, shared_1.extend)(miniProgramComponents, findMiniProgramUsingComponent(globalUsingComponents, componentsDir));
    }
    const jsonFile = findJsonFile((0, utils_1.removeExt)((0, utils_1.normalizeMiniProgramFilename)(filename, inputDir)));
    if (jsonFile) {
        if (jsonFile.usingComponents) {
            (0, shared_1.extend)(miniProgramComponents, findMiniProgramUsingComponent(jsonFile.usingComponents, componentsDir));
        }
        // mp-baidu 特有
        if (jsonFile.usingSwanComponents) {
            (0, shared_1.extend)(miniProgramComponents, findMiniProgramUsingComponent(jsonFile.usingSwanComponents, componentsDir));
        }
    }
    return miniProgramComponents;
}
exports.findMiniProgramUsingComponents = findMiniProgramUsingComponents;
function findMiniProgramUsingComponent(usingComponents, componentsDir) {
    return Object.keys(usingComponents).reduce((res, name) => {
        const path = usingComponents[name];
        if (path.includes('plugin://')) {
            // mp-weixin & mp-alipay
            res[name] = 'plugin';
        }
        else if (path.includes('dynamicLib://')) {
            // mp-baidu
            res[name] = 'dynamicLib';
        }
        else if (path.includes('ext://')) {
            // mp-toutiao
            res[name] = 'ext';
        }
        else if (componentsDir &&
            path.includes(componentsDir + '/') &&
            findUsingComponentsJson(path, componentsDir).renderer === 'xr-frame') {
            // mp-weixin & x-frame
            res[name] = 'xr-frame';
        }
        else if (componentsDir && path.includes(componentsDir + '/')) {
            res[name] = 'component';
        }
        return res;
    }, {});
}
function findUsingComponentsJson(pathInpages, componentsDir) {
    // 兼容test case
    if (!process.env.UNI_INPUT_DIR)
        return {};
    let [, dir] = pathInpages.split(componentsDir);
    if (dir === '') {
        console.warn(`${pathInpages} 路径里没有找到对应的 ${componentsDir} 目录`);
        return {};
    }
    dir = '.' + dir;
    const fulldir = path_1.default.resolve(process.env.UNI_INPUT_DIR, componentsDir, dir);
    let filename = path_1.default.parse(pathInpages).name + '.json';
    let jsonPath = path_1.default.resolve(fulldir, filename);
    if (fs_1.default.existsSync(jsonPath)) {
        return require(jsonPath);
    }
    jsonPath = path_1.default.resolve(fulldir, 'index.json');
    if (fs_1.default.existsSync(jsonPath)) {
        return require(jsonPath);
    }
    console.warn(`${pathInpages} 路径下没有找到对应的json文件`);
    return {};
}
exports.findUsingComponentsJson = findUsingComponentsJson;
