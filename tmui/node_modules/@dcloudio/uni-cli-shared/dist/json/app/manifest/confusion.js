"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initConfusion = exports.hasConfusionFile = exports.isConfusionFile = exports.APP_CONFUSION_FILENAME = void 0;
const path_1 = __importDefault(require("path"));
const utils_1 = require("../../../utils");
const constants_1 = require("../../../constants");
const manifest_1 = require("../../manifest");
function isJsFile(filename) {
    return constants_1.EXTNAME_JS_RE.test(filename);
}
function isStaticJsFile(filename) {
    return (filename.indexOf('hybrid/html') === 0 ||
        filename.indexOf('static/') === 0 ||
        filename.indexOf('/static/') !== -1); // subpackages, uni_modules 中的 static 目录
}
const dynamicConfusionJsFiles = [];
exports.APP_CONFUSION_FILENAME = 'app-confusion.js';
function isConfusionFile(filename) {
    return dynamicConfusionJsFiles.includes((0, utils_1.normalizePath)(filename));
}
exports.isConfusionFile = isConfusionFile;
function hasConfusionFile(inputDir) {
    if (inputDir) {
        const manifestJson = (0, manifest_1.parseManifestJsonOnce)(inputDir);
        const resources = manifestJson['app-plus']?.confusion?.resources;
        if (resources && parseConfusion(resources)[1].length) {
            return true;
        }
    }
    return !!dynamicConfusionJsFiles.length;
}
exports.hasConfusionFile = hasConfusionFile;
function parseConfusion(resources) {
    const res = {};
    const dynamicJsFiles = [];
    Object.keys(resources).reduce((res, name) => {
        const extname = path_1.default.extname(name);
        if (extname === '.nvue') {
            res[name.replace('.nvue', '.js')] = resources[name];
        }
        else if (isJsFile(name)) {
            // 静态 js 加密
            if (isStaticJsFile(name)) {
                res[name] = resources[name];
            }
            else {
                // 非静态 js 将被合并进 app-confusion.js
                dynamicJsFiles.push(name);
            }
        }
        else {
            throw new Error(`原生混淆仅支持 nvue 页面，错误的页面路径：${name}`);
        }
        // TODO 旧编译器会检查要加密的 nvue 页面（包括subnvue）是否被使用？后续有时间再考虑支持吧，意义不太大
        return res;
    }, res);
    if (dynamicJsFiles.length) {
        res[exports.APP_CONFUSION_FILENAME] = {};
    }
    return [res, dynamicJsFiles];
}
function initConfusion(manifestJson) {
    dynamicConfusionJsFiles.length = 0;
    if (!manifestJson.plus.confusion?.resources) {
        return;
    }
    const resources = manifestJson.plus.confusion.resources;
    const [res, dynamicJsFiles] = parseConfusion(resources);
    manifestJson.plus.confusion.resources = res;
    dynamicConfusionJsFiles.push(...dynamicJsFiles);
}
exports.initConfusion = initConfusion;
