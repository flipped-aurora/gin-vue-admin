"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniPreCssPlugin = void 0;
const debug_1 = __importDefault(require("debug"));
const pluginutils_1 = require("@rollup/pluginutils");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugPre = (0, debug_1.default)('uni:pre-css');
// const debugPreTry = debug('uni:pre-css-try')
const cssLangs = `\\.(css|less|sass|scss|styl|stylus|postcss)($|\\?)`;
const cssLangRE = new RegExp(cssLangs);
/**
 * preprocess css
 * @param options
 */
function uniPreCssPlugin(config, options) {
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    const isNVue = config.nvue;
    const preJsFile = isNVue ? uni_cli_shared_1.preNVueJs : uni_cli_shared_1.preJs;
    return {
        name: 'uni:pre-css',
        transform(code, id) {
            if (!cssLangRE.test(id)) {
                return;
            }
            if (!filter(id)) {
                return;
            }
            // debugPreTry(id)
            if (!code.includes('#endif')) {
                return;
            }
            debugPre(id);
            return {
                code: preJsFile(code),
                map: (0, uni_cli_shared_1.withSourcemap)(config) ? this.getCombinedSourcemap() : null,
            };
        },
    };
}
exports.uniPreCssPlugin = uniPreCssPlugin;
