"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.initFixedEsbuildInitTSConfck = exports.rewriteCompilerSfcParse = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const shared_1 = require("@vue/shared");
/**
 * TODO 临时重写，解决 @vitejs/plugin-vue 的 Bug
 */
function rewriteCompilerSfcParse() {
    const compilerSfc = require((0, uni_cli_shared_1.resolveBuiltIn)('@vue/compiler-sfc'));
    const { parse } = compilerSfc;
    let extname_vue = process.env.UNI_APP_X === 'true' ? uni_cli_shared_1.X_EXTNAME_VUE : uni_cli_shared_1.EXTNAME_VUE;
    compilerSfc.parse = (source, options) => {
        if (options?.filename) {
            const extname = path_1.default.extname(options.filename);
            // wxs、filter、renderjs
            if (extname && !extname_vue.includes(extname)) {
                const tag = extname.slice(1);
                source = `<${tag}>` + source + `</${tag}>`;
            }
        }
        return parse(source, options);
    };
}
exports.rewriteCompilerSfcParse = rewriteCompilerSfcParse;
/**
 * 已废弃，交由 rewriteExistsSyncHasRootFile 实现，因为新的 vite 版本在 configResolved 中重写已经晚了
 * 解决 HBuilderX 项目未包含 package.json 时，initTSConfck 可能导致查找过慢，或递归目录时权限不足报错
 * 即：未包含 package.json 时，直接移除 initTSConfck 相关逻辑
 * @param inputDir
 * @returns
 */
function initFixedEsbuildInitTSConfck(inputDir) {
    if (!(0, uni_cli_shared_1.isInHBuilderX)()) {
        return [];
    }
    if (fs_1.default.existsSync(path_1.default.resolve(inputDir, 'package.json'))) {
        return [];
    }
    // 'vite:esbuild', 'vite:esbuild-transpile' initTSConfck
    const existsSync = fs_1.default.existsSync;
    // 根目录 lerna.json 路径
    const lernaJsonPath = (0, uni_cli_shared_1.normalizePath)(path_1.default.resolve(inputDir, 'lerna.json'));
    return [
        {
            name: 'uni:fixed-HBuilderX-initTSConfck-before',
            enforce: 'pre',
            configResolved() {
                // 在 HBuilderX 项目中，缺少 package.json 时， 确保 searchForWorkspaceRoot 使用项目根目录， 否则 initTSConfck 会查找很费时，且可能访问目录权限报错
                // https://github.com/vitejs/vite/blob/43b7b78b1834a4c7128d8a5d987f66a4defcbd93/packages/vite/src/node/plugins/esbuild.ts#L407
                fs_1.default.existsSync = (p) => {
                    if ((0, shared_1.isString)(p)) {
                        // 访问根目录 lerna.json 时，直接返回 true
                        // https://github.com/vitejs/vite/blob/43b7b78b1834a4c7128d8a5d987f66a4defcbd93/packages/vite/src/node/server/searchRoot.ts#L35
                        if (p.endsWith('lerna.json') &&
                            lernaJsonPath === (0, uni_cli_shared_1.normalizePath)(p)) {
                            return true;
                        }
                    }
                    return existsSync(p);
                };
            },
        },
    ];
}
exports.initFixedEsbuildInitTSConfck = initFixedEsbuildInitTSConfck;
