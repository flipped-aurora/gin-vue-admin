"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniAppCssPlugin = exports.APP_CSS_JS = void 0;
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
exports.APP_CSS_JS = './app.css.js';
function uniAppCssPlugin() {
    const inputDir = process.env.UNI_INPUT_DIR;
    const appVueFilename = (0, uni_cli_shared_1.resolveAppVue)(inputDir);
    return {
        name: 'uni:app-nvue-app-style',
        // 提前到 @vite/plugin-vue 之前执行，因为在 nvue 编译时，仅 import 了 App.vue 的 styles，这样导致 descriptor
        // 一直使用的是上一次的（plugin-vue 会在 transformMain 中生成新的 descriptor），故不再交由 plugin-vue 来 load
        // 而是当前插件直接处理
        enforce: 'pre',
        resolveId(id) {
            if (id === exports.APP_CSS_JS) {
                return exports.APP_CSS_JS;
            }
        },
        load(id) {
            if (id === exports.APP_CSS_JS) {
                return genAppStylesCode(appVueFilename, this);
            }
            const { filename, query } = (0, uni_cli_shared_1.parseVueRequest)(id);
            if (query.vue && query.type === 'style' && appVueFilename === filename) {
                const descriptor = createAppDescriptor(filename, this);
                const block = descriptor.styles[query.index];
                if (block) {
                    return {
                        code: block.content,
                        map: '',
                    };
                }
            }
        },
    };
}
exports.uniAppCssPlugin = uniAppCssPlugin;
const defaultAppStylesCode = `exports.styles = []`;
async function genAppStylesCode(filename, pluginContext) {
    pluginContext.addWatchFile(filename);
    const descriptor = createAppDescriptor(filename, pluginContext);
    if (!descriptor.styles.length) {
        return defaultAppStylesCode;
    }
    let stylesCode = ``;
    const styleVars = [];
    for (let i = 0; i < descriptor.styles.length; i++) {
        const style = descriptor.styles[i];
        const src = style.src || descriptor.filename;
        const attrsQuery = attrsToQuery(style.attrs, 'css');
        const srcQuery = style.src ? `&src=${descriptor.id}` : ``;
        const query = `?vue&type=style&index=${i}${srcQuery}&inline`;
        const styleRequest = src + query + attrsQuery;
        stylesCode += `\nimport _style_${i} from ${JSON.stringify(styleRequest)}`;
        styleVars.push(`_style_${i}`);
    }
    return `
${stylesCode}
exports.styles = [${styleVars.join(',')}]
`;
}
function readAppCode(filename) {
    if (!fs_extra_1.default.existsSync(filename)) {
        return ``;
    }
    const source = fs_extra_1.default.readFileSync(filename, 'utf8');
    if (source.includes('#endif')) {
        return (0, uni_cli_shared_1.preNVueJs)((0, uni_cli_shared_1.preNVueHtml)(source));
    }
    return source;
}
let appDescriptor;
function createAppDescriptor(filename, pluginContext) {
    const source = readAppCode(filename);
    const id = (0, uni_cli_shared_1.hash)(source);
    if (!appDescriptor || appDescriptor.id !== id) {
        const { descriptor, errors } = require('@vue/compiler-sfc').parse(source, {
            filename,
        });
        descriptor.id = id;
        if (errors.length) {
            errors.forEach((error) => pluginContext.error((0, uni_cli_shared_1.createRollupError)('uni:app-nvue-app-style', filename, error)));
        }
        appDescriptor = descriptor;
    }
    return appDescriptor;
}
// these are built-in query parameters so should be ignored
// if the user happen to add them as attrs
const ignoreList = ['id', 'index', 'src', 'type', 'lang', 'module'];
function attrsToQuery(attrs, langFallback, forceLangFallback = false) {
    let query = ``;
    for (const name in attrs) {
        const value = attrs[name];
        if (!ignoreList.includes(name)) {
            query += `&${encodeURIComponent(name)}${value ? `=${encodeURIComponent(value)}` : ``}`;
        }
    }
    if (langFallback || attrs.lang) {
        query +=
            `lang` in attrs
                ? forceLangFallback
                    ? `&lang.${langFallback}`
                    : `&lang.${attrs.lang}`
                : `&lang.${langFallback}`;
    }
    return query;
}
