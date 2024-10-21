"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createCss = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function resolveAdditionalData(inputDir, config) {
    const uniScssFile = path_1.default.resolve(inputDir, 'uni.scss');
    const userAdditionalData = config.css?.preprocessorOptions?.scss?.additionalData || '';
    if (!fs_extra_1.default.existsSync(uniScssFile)) {
        return userAdditionalData;
    }
    let content = fs_extra_1.default.readFileSync(uniScssFile, 'utf8');
    if (content.includes('#endif')) {
        content = (0, uni_cli_shared_1.preCss)(content);
    }
    return content + '\n' + userAdditionalData;
}
function createCss(options, config) {
    return {
        preprocessorOptions: {
            scss: {
                charset: false,
                additionalData: resolveAdditionalData(options.inputDir, config),
            },
        },
    };
}
exports.createCss = createCss;
