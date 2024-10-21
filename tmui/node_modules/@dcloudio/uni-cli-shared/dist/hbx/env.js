"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.fixBinaryPath = exports.initModulePaths = exports.runByHBuilderX = exports.isInHBuilderX = void 0;
const path_1 = __importDefault(require("path"));
const module_1 = __importDefault(require("module"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const resolve_1 = require("../resolve");
const utils_1 = require("../utils");
exports.isInHBuilderX = (0, uni_shared_1.once)(() => {
    // 自动化测试传入了 HX_APP_ROOT(其实就是UNI_HBUILDERX_PLUGINS)
    if (process.env.HX_APP_ROOT) {
        process.env.UNI_HBUILDERX_PLUGINS = process.env.HX_APP_ROOT + '/plugins';
        return true;
    }
    try {
        const { name } = require(path_1.default.resolve(process.cwd(), '../about/package.json'));
        if (name === 'about') {
            process.env.UNI_HBUILDERX_PLUGINS = path_1.default.resolve(process.cwd(), '..');
            return true;
        }
    }
    catch (e) {
        // console.error(e)
    }
    return false;
});
exports.runByHBuilderX = (0, uni_shared_1.once)(() => {
    return (!!process.env.UNI_HBUILDERX_PLUGINS &&
        (!!process.env.RUN_BY_HBUILDERX || !!process.env.HX_Version));
});
/**
 * 增加 node_modules
 */
function initModulePaths() {
    if (!(0, exports.isInHBuilderX)()) {
        return;
    }
    const Module = module.constructor.length > 1 ? module.constructor : module_1.default;
    const nodeModulesPath = path_1.default.resolve(process.env.UNI_CLI_CONTEXT, 'node_modules');
    const oldNodeModulePaths = Module._nodeModulePaths;
    Module._nodeModulePaths = function (from) {
        const paths = oldNodeModulePaths.call(this, from);
        if (!paths.includes(nodeModulesPath)) {
            paths.push(nodeModulesPath);
        }
        return paths;
    };
}
exports.initModulePaths = initModulePaths;
function resolveEsbuildModule(name) {
    try {
        return path_1.default.dirname(require.resolve(name + '/package.json', {
            paths: [path_1.default.dirname((0, resolve_1.resolveBuiltIn)('esbuild/package.json'))],
        }));
    }
    catch (e) { }
    return '';
}
function fixBinaryPath() {
    // cli 工程在 HBuilderX 中运行
    if (!(0, exports.isInHBuilderX)() && (0, exports.runByHBuilderX)()) {
        if (utils_1.isWindows) {
            const win64 = resolveEsbuildModule('esbuild-windows-64');
            if (win64) {
                process.env.ESBUILD_BINARY_PATH = path_1.default.join(win64, 'esbuild.exe');
            }
        }
        else {
            const arm64 = resolveEsbuildModule('esbuild-darwin-arm64');
            if (arm64) {
                process.env.ESBUILD_BINARY_PATH = path_1.default.join(arm64, 'bin/esbuild');
            }
        }
    }
}
exports.fixBinaryPath = fixBinaryPath;
