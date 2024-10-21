"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniOptions = void 0;
const fs_extra_1 = __importDefault(require("fs-extra"));
const path_1 = __importDefault(require("path"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_i18n_1 = require("@dcloudio/uni-i18n");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const nvue_1 = require("../../nvue");
const plugin_1 = require("../../nvue/plugin");
const transformWxsProps_1 = require("./transforms/transformWxsProps");
function isAppNVueNativeTag(tag) {
    return (0, uni_cli_shared_1.matchUTSComponent)(tag) || (0, uni_shared_1.isAppNVueNativeTag)(tag);
}
function uniOptions(compilerType = process.env.UNI_COMPILER) {
    const isNVueCompiler = compilerType === 'nvue';
    return {
        copyOptions() {
            const platform = process.env.UNI_PLATFORM;
            const inputDir = process.env.UNI_INPUT_DIR;
            const outputDir = process.env.UNI_OUTPUT_DIR;
            const targets = [];
            // 自动化测试时，不启用隐私政策
            if (!process.env.UNI_AUTOMATOR_WS_ENDPOINT) {
                targets.push({
                    src: 'androidPrivacy.json',
                    dest: outputDir,
                    transform(source) {
                        const options = (0, uni_cli_shared_1.initI18nOptions)(platform, inputDir, false, true);
                        if (!options) {
                            return;
                        }
                        return (0, uni_i18n_1.compileI18nJsonStr)(source.toString(), options);
                    },
                });
                const debugFilename = '__nvue_debug__';
                if (fs_extra_1.default.existsSync(path_1.default.resolve(inputDir, debugFilename))) {
                    targets.push({
                        src: debugFilename,
                        dest: outputDir,
                    });
                }
            }
            return {
                assets: ['hybrid/html/**/*', 'uni_modules/*/hybrid/html/**/*'],
                targets,
            };
        },
        compilerOptions: {
            isNativeTag: isNVueCompiler ? isAppNVueNativeTag : uni_shared_1.isAppNativeTag,
            nodeTransforms: [
                ...(isNVueCompiler ? (0, nvue_1.initNVueNodeTransforms)() : [transformWxsProps_1.transformWxsProps]),
                uni_cli_shared_1.transformTapToClick,
                uni_cli_shared_1.transformMatchMedia,
                uni_cli_shared_1.transformPageHead,
            ],
            directiveTransforms: {
                ...(isNVueCompiler ? (0, plugin_1.initNVueDirectiveTransforms)() : {}),
            },
        },
    };
}
exports.uniOptions = uniOptions;
