"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniTemplatePlugin = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const utils_1 = require("../utils");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniTemplatePlugin({ renderer, } = {}) {
    let outputDir;
    return {
        name: 'uni:app-template',
        enforce: 'post',
        configResolved() {
            outputDir = process.env.UNI_OUTPUT_DIR;
            if (renderer !== 'native') {
                fs_extra_1.default.copySync((0, uni_cli_shared_1.resolveBuiltIn)('@dcloudio/uni-app-plus/dist/uni-app-view.umd.js'), path_1.default.resolve(outputDir, 'uni-app-view.umd.js'), {
                    overwrite: true,
                });
            }
            fs_extra_1.default.copySync(utils_1.templateDir, outputDir, {
                overwrite: true,
                filter(src) {
                    if (renderer === 'native') {
                        if (src.includes('__uniappquill') ||
                            src.includes('__uniappautomator')) {
                            return false;
                        }
                    }
                    return !src.includes('__uniappview.html');
                },
            });
        },
    };
}
exports.uniTemplatePlugin = uniTemplatePlugin;
