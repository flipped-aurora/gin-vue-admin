"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniConfusionPlugin = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function uniConfusionPlugin() {
    const inputDir = process.env.UNI_INPUT_DIR;
    const hasConfusion = process.env.NODE_ENV === 'production' && (0, uni_cli_shared_1.hasConfusionFile)(inputDir);
    return {
        name: 'uni:app-vue-confusion',
        enforce: 'post',
        apply: 'build',
        config() {
            if (!hasConfusion) {
                return;
            }
            return {
                build: {
                    rollupOptions: {
                        output: {
                            format: process.env.UNI_APP_CODE_SPLITING ? 'amd' : 'cjs',
                            manualChunks(id) {
                                if ((0, uni_cli_shared_1.isConfusionFile)(path_1.default.relative(inputDir, id))) {
                                    return (0, uni_cli_shared_1.removeExt)(uni_cli_shared_1.APP_CONFUSION_FILENAME);
                                }
                            },
                        },
                    },
                },
            };
        },
        generateBundle(_, bundle) {
            if (!hasConfusion) {
                return;
            }
            const appConfusionChunk = bundle[uni_cli_shared_1.APP_CONFUSION_FILENAME];
            if (!appConfusionChunk) {
                return;
            }
            appConfusionChunk.code = wrapperAppConfusionCode(appConfusionChunk.code);
            const appServiceChunk = bundle[uni_cli_shared_1.APP_SERVICE_FILENAME];
            if (!appServiceChunk) {
                return;
            }
            appServiceChunk.code = wrapperAppServiceCode(appServiceChunk.code);
        },
    };
}
exports.uniConfusionPlugin = uniConfusionPlugin;
function replaceRequireVueCode(code) {
    // 目前会生成require("@vue/shared"); 理论上摇树之后，是不应该有的，后续排查为什么
    return code
        .replace(/require\(['"]vue['"]\)/gi, `$cjs_require$('vue')`)
        .replace(/require\(['"]@vue\/shared['"]\)/gi, `$cjs_require$('@vue/shared')`);
}
function replaceRequireAppConfusionCode(code) {
    return code.replace(new RegExp(`require\\(['"].\\/${uni_cli_shared_1.APP_CONFUSION_FILENAME}['"]\\)`, 'gi'), `$cjs_require$('./${uni_cli_shared_1.APP_CONFUSION_FILENAME}')`);
}
function wrapperAppServiceCode(code) {
    return replaceRequireAppConfusionCode(replaceRequireVueCode(code));
}
function wrapperAppConfusionCode(code) {
    return `function $cjs_require$(name){if(name==='vue'){return Vue;}if(name==='./${uni_cli_shared_1.APP_CONFUSION_FILENAME}'){return $appConfusion$;}};const $appConfusion$ = {};(function(exports){${replaceRequireVueCode(code)}})($appConfusion$);
`;
}
