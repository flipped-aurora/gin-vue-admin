"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppCssPlugin = void 0;
const picocolors_1 = __importDefault(require("picocolors"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_nvue_styler_1 = require("@dcloudio/uni-nvue-styler");
function uniAppCssPlugin(resolvedConfig) {
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
            if (source.includes('#endif')) {
                source = (0, uni_cli_shared_1.preUVueCss)(source);
            }
            source = (0, uni_cli_shared_1.parseAssets)(resolvedConfig, source);
            // 仅做校验使用
            const { messages, code } = await (0, uni_nvue_styler_1.parse)(source, {
                filename,
                logLevel: 'WARNING',
                type: 'uvue',
                platform: process.env.UNI_UTS_PLATFORM,
            });
            messages.forEach((message) => {
                if (message.type === 'warning') {
                    // 拆分成多行，第一行输出信息（有颜色），后续输出错误代码+文件行号
                    resolvedConfig.logger.warn(picocolors_1.default.yellow(`[plugin:uni:app-uvue-css] ${message.text}`));
                    let msg = '';
                    if (message.line && message.column) {
                        msg += `\n${(0, uni_cli_shared_1.generateCodeFrame)(source, {
                            line: message.line,
                            column: message.column,
                        }).replace(/\t/g, ' ')}\n`;
                    }
                    msg += `${(0, uni_cli_shared_1.formatAtFilename)(filename)}`;
                    resolvedConfig.logger.warn(msg);
                }
            });
            return { code: `export default ${code}`, map: { mappings: '' } };
        },
        generateBundle() {
            // 用于覆盖原始插件方法
            // noop
        },
    };
}
exports.uniAppCssPlugin = uniAppCssPlugin;
