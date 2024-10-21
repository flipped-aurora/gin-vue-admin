"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initNVueEnv = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function initNVueEnv() {
    if (process.env.UNI_APP_X === 'true') {
        return;
    }
    const manifestJson = (0, uni_cli_shared_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR);
    if ((0, uni_cli_shared_1.getAppRenderer)(manifestJson) === 'native') {
        process.env.UNI_RENDERER = 'native';
    }
    const nvueCompiler = (0, uni_cli_shared_1.getNVueCompiler)(manifestJson);
    if (nvueCompiler === 'uni-app') {
        process.env.UNI_NVUE_COMPILER = 'uni-app';
    }
    else if (nvueCompiler === 'vue') {
        process.env.UNI_NVUE_COMPILER = 'vue';
    }
    if ((0, uni_cli_shared_1.getNVueStyleCompiler)(manifestJson) === 'uni-app') {
        process.env.UNI_NVUE_STYLE_COMPILER = 'uni-app';
    }
    if ((0, uni_cli_shared_1.getAppCodeSpliting)(manifestJson)) {
        process.env.UNI_APP_CODE_SPLITING = 'true';
    }
}
exports.initNVueEnv = initNVueEnv;
