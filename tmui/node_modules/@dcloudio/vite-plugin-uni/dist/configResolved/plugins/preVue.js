"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPreVuePlugin = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniPreVuePlugin() {
    let isNVue = false;
    return {
        name: 'uni:pre-vue',
        config(config) {
            isNVue = config.nvue;
        },
        async transform(code, id) {
            const { filename, query } = (0, uni_cli_shared_1.parseVueRequest)(id);
            if (query.vue) {
                return;
            }
            if (!(process.env.UNI_APP_X === 'true' ? uni_cli_shared_1.X_EXTNAME_VUE : uni_cli_shared_1.EXTNAME_VUE).includes(path_1.default.extname(filename))) {
                return;
            }
            // 清空当前页面已缓存的 filter 信息
            (0, uni_cli_shared_1.clearMiniProgramTemplateFilter)((0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(id, process.env.UNI_INPUT_DIR)));
            return {
                code: (0, uni_cli_shared_1.parseVueCode)(code, isNVue).code, // 暂不提供sourcemap,意义不大
                map: null,
            };
        },
    };
}
exports.uniPreVuePlugin = uniPreVuePlugin;
