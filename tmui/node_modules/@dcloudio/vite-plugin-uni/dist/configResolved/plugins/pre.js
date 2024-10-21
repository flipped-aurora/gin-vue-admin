"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPrePlugin = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const pluginutils_1 = require("@rollup/pluginutils");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugPreJs = (0, debug_1.default)('uni:pre-js');
const debugPreHtml = (0, debug_1.default)('uni:pre-html');
// const debugPreJsTry = debug('uni:pre-js-try')
function uniPrePlugin(config, options) {
    const isX = process.env.UNI_APP_X === 'true';
    const PRE_JS_EXTNAME = ['.json', '.css']
        .concat(isX ? uni_cli_shared_1.X_EXTNAME_VUE : uni_cli_shared_1.EXTNAME_VUE)
        .concat(uni_cli_shared_1.EXTNAME_JS); // 因为 1.0 也会使用 uts uni_modules，所以 EXTNAME_JS 直接包含了 .uts 后缀
    const PRE_HTML_EXTNAME = isX ? uni_cli_shared_1.X_EXTNAME_VUE : uni_cli_shared_1.EXTNAME_VUE;
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    const isNVue = config.nvue;
    const preJsFile = isNVue ? uni_cli_shared_1.preNVueJs : uni_cli_shared_1.preJs;
    const preHtmlFile = isNVue ? uni_cli_shared_1.preNVueHtml : uni_cli_shared_1.preHtml;
    return {
        name: 'uni:pre',
        transform(code, id) {
            if (!filter(id)) {
                return;
            }
            const { filename, query } = (0, uni_cli_shared_1.parseVueRequest)(id);
            const extname = path_1.default.extname(filename);
            const isHtml = query.type === 'template' || PRE_HTML_EXTNAME.includes(extname);
            const isJs = PRE_JS_EXTNAME.includes(extname);
            const isPre = isHtml || isJs;
            if (isPre) {
                // debugPreJsTry(id)
            }
            const hasEndif = isPre && code.includes('#endif');
            if (!hasEndif) {
                return;
            }
            if (isHtml) {
                code = preHtmlFile(code);
                debugPreHtml(id);
            }
            if (isJs) {
                code = preJsFile(code);
                debugPreJs(id);
            }
            return {
                code,
                map: (0, uni_cli_shared_1.withSourcemap)(config) ? this.getCombinedSourcemap() : null,
            };
        },
    };
}
exports.uniPrePlugin = uniPrePlugin;
