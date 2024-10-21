"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildOptions = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../utils");
function buildOptions({ appService, renderer, }, userConfig, _) {
    const inputDir = process.env.UNI_INPUT_DIR;
    const outputDir = process.env.UNI_OUTPUT_DIR;
    // 开始编译时，清空输出目录
    function emptyNVueDir() {
        const nvueOutputDir = (0, utils_1.nvueOutDir)();
        if (fs_1.default.existsSync(nvueOutputDir)) {
            (0, uni_cli_shared_1.emptyDir)(nvueOutputDir);
        }
    }
    function emptyOutDir() {
        if (fs_1.default.existsSync(outputDir)) {
            (0, uni_cli_shared_1.emptyDir)(outputDir);
        }
    }
    if (renderer === 'native') {
        if (appService) {
            // 仅编译 main.js+App.vue 的时候才清空
            emptyNVueDir();
            emptyOutDir();
        }
    }
    else {
        if ((0, uni_cli_shared_1.isInHybridNVue)(userConfig)) {
            emptyNVueDir();
        }
        else {
            emptyOutDir();
        }
    }
    const sourcemap = process.env.SOURCEMAP === 'true'
        ? 'hidden'
        : userConfig.build?.sourcemap
            ? 'inline'
            : false;
    return {
        // App 端目前仅提供 inline
        sourcemap,
        emptyOutDir: false, // 不清空输出目录，否则会影响 webpack 的输出
        assetsInlineLimit: 0,
        rollupOptions: {
            input: (0, uni_cli_shared_1.resolveMainPathOnce)(inputDir),
            output: {
                sourcemapPathTransform(relativeSourcePath, sourcemapPath) {
                    const sourcePath = (0, uni_cli_shared_1.normalizePath)(path_1.default.relative(inputDir, path_1.default.resolve(path_1.default.dirname(sourcemapPath), relativeSourcePath)));
                    if (sourcePath.startsWith('..')) {
                        return '';
                    }
                    return 'uni-app:///' + sourcePath;
                },
                manualChunks: {},
                inlineDynamicImports: false,
                chunkFileNames(chunk) {
                    if (chunk.isDynamicEntry && chunk.facadeModuleId) {
                        const filepath = path_1.default.relative(inputDir, chunk.facadeModuleId);
                        return (0, uni_cli_shared_1.normalizePath)(filepath.replace(path_1.default.extname(filepath), '.js'));
                    }
                    return '[name].js';
                },
            },
        },
    };
}
exports.buildOptions = buildOptions;
