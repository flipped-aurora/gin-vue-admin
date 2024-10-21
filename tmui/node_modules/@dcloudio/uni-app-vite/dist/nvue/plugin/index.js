"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppNVuePlugin = exports.initNVueDirectiveTransforms = exports.initNVueNodeTransforms = void 0;
const path_1 = __importDefault(require("path"));
const picocolors_1 = __importDefault(require("picocolors"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_nvue_styler_1 = require("@dcloudio/uni-nvue-styler");
const utils_1 = require("../../utils");
// import { transformRenderWhole } from './transforms/transformRenderWhole'
// import { transformAppendAsTree } from './transforms/transformAppendAsTree'
const transformVideo_1 = require("./transforms/transformVideo");
const transformText_1 = require("./transforms/transformText");
const configResolved_1 = require("../../plugin/configResolved");
const utils_2 = require("../utils");
const transformRootNode_1 = require("./transforms/transformRootNode");
const vModel_1 = require("./transforms/vModel");
const vShow_1 = require("./transforms/vShow");
const transformAttrs_1 = require("./transforms/transformAttrs");
const pagesJson_1 = require("../plugins/pagesJson");
const uTags = {
    text: 'u-text',
    image: 'u-image',
    input: 'u-input',
    textarea: 'u-textarea',
    video: 'u-video',
    'web-view': 'u-web-view',
    slider: 'u-slider',
};
function initNVueNodeTransforms() {
    // 优先级必须确保 renderWhole > appendAsTree
    return [
        transformRootNode_1.transformRootNode,
        (0, uni_cli_shared_1.createTransformTag)(uTags),
        transformAttrs_1.transformAttrs,
        transformText_1.transformText,
        transformVideo_1.transformVideo,
        uni_cli_shared_1.transformUTSComponent,
        // transformRenderWhole,
        // transformAppendAsTree,
    ];
}
exports.initNVueNodeTransforms = initNVueNodeTransforms;
function initNVueDirectiveTransforms() {
    return {
        model: vModel_1.transformModel,
        show: vShow_1.transformShow,
    };
}
exports.initNVueDirectiveTransforms = initNVueDirectiveTransforms;
function uniAppNVuePlugin({ appService, }) {
    const inputDir = process.env.UNI_INPUT_DIR;
    const mainPath = (0, uni_cli_shared_1.resolveMainPathOnce)(inputDir);
    return {
        name: 'uni:app-nvue',
        config() {
            return {
                css: {
                    postcss: {
                        plugins: (0, uni_cli_shared_1.initPostcssPlugin)({
                            uniApp: uni_shared_1.defaultNVueRpx2Unit,
                            autoprefixer: false,
                        }),
                    },
                },
                build: {
                    lib: {
                        name: 'AppService',
                        // 必须使用 lib 模式，否则会生成 preload 等代码
                        fileName: appService ? 'app-service' : 'app',
                        entry: mainPath,
                        formats: [appService ? 'iife' : 'es'],
                    },
                    outDir: appService ? process.env.UNI_OUTPUT_DIR : (0, utils_1.nvueOutDir)(),
                    rollupOptions: {
                        external: (0, utils_2.external)(appService),
                        output: {
                            entryFileNames(chunk) {
                                if (chunk.name === 'main' && chunk.isEntry) {
                                    return appService ? uni_cli_shared_1.APP_SERVICE_FILENAME : 'app.js';
                                }
                                return chunk.name + '.js';
                            },
                            chunkFileNames: createChunkFileNames(inputDir),
                            plugins: [(0, uni_cli_shared_1.dynamicImportPolyfill)(true)],
                            globals: (0, utils_2.globals)(appService),
                        },
                    },
                },
            };
        },
        configResolved: (0, configResolved_1.createConfigResolved)({
            createCssPostPlugin(config) {
                return {
                    name: 'vite:css-post',
                    buildStart() {
                        // 用于覆盖原始插件方法
                        // noop
                    },
                    async transform(source, filename) {
                        if (!uni_cli_shared_1.cssLangRE.test(filename) || uni_cli_shared_1.commonjsProxyRE.test(filename)) {
                            return;
                        }
                        const nvuePages = pagesJson_1.nvuePagesCache.get(config);
                        if (!nvuePages || !Object.keys(nvuePages).length) {
                            // 当前项目没有 nvue 文件
                            return { code: `export default {}`, map: { mappings: '' } };
                        }
                        const { code, messages } = await (0, uni_nvue_styler_1.parse)(source, {
                            filename,
                            logLevel: 'WARNING',
                        });
                        messages.forEach((message) => {
                            if (message.type === 'warning') {
                                config.logger.warn(picocolors_1.default.yellow(`[plugin:vite:nvue-css] ${message.text}`));
                                let msg = '';
                                if (message.line && message.column) {
                                    msg += `\n${(0, uni_cli_shared_1.generateCodeFrame)(source, {
                                        line: message.line,
                                        column: message.column,
                                    })}\n`;
                                }
                                msg += `${(0, uni_cli_shared_1.formatAtFilename)(filename)}`;
                                config.logger.warn(msg);
                            }
                        });
                        return { code: `export default ${code}`, map: { mappings: '' } };
                    },
                    generateBundle() {
                        // 用于覆盖原始插件方法
                        // noop
                    },
                };
            },
        }),
    };
}
exports.uniAppNVuePlugin = uniAppNVuePlugin;
function createChunkFileNames(inputDir) {
    return function chunkFileNames(chunk) {
        if (chunk.isDynamicEntry && chunk.facadeModuleId) {
            const { filename } = (0, uni_cli_shared_1.parseVueRequest)(chunk.facadeModuleId);
            if (filename.endsWith('.nvue')) {
                return ((0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizePath)(path_1.default.relative(inputDir, filename))) + '.js');
            }
        }
        return '[name].js';
    };
}
