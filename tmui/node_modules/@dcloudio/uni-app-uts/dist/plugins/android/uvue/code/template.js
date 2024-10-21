"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.tryResolveTemplateSrc = exports.genTemplateCode = exports.genTemplate = void 0;
const fs_extra_1 = __importDefault(require("fs-extra"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_1 = require("../compiler");
const utils_1 = require("../compiler/utils");
const descriptorCache_1 = require("../descriptorCache");
function genTemplate({ template }, options) {
    if (!template || !template.content) {
        return {
            code: options.mode === 'module'
                ? (0, utils_1.genRenderFunctionDecl)(options) + ` { return null }`
                : `null`,
            easyComponentAutoImports: {},
            preamble: '',
            elements: [],
            imports: [],
        };
    }
    const { preprocessLang, preprocessOptions } = options;
    const preprocessor = preprocessLang
        ? require('@vue/consolidate')[preprocessLang]
        : false;
    return (0, compiler_1.compile)(preprocessor
        ? preprocess({ source: template.content, filename: '', preprocessOptions }, preprocessor)
        : template.content, options);
}
exports.genTemplate = genTemplate;
exports.genTemplateCode = genTemplate;
function preprocess({ source, filename, preprocessOptions, }, preprocessor) {
    // Consolidate exposes a callback based API, but the callback is in fact
    // called synchronously for most templating engines. In our case, we have to
    // expose a synchronous API so that it is usable in Jest transforms (which
    // have to be sync because they are applied via Node.js require hooks)
    let res = '';
    let err = null;
    preprocessor.render(source, { filename, ...preprocessOptions }, (_err, _res) => {
        if (_err)
            err = _err;
        res = _res;
    });
    if (err)
        throw err;
    return res;
}
async function tryResolveTemplateSrc(descriptor, pluginContext) {
    if (!pluginContext) {
        return;
    }
    if (!descriptor.template) {
        return;
    }
    if (descriptor.template.src) {
        const resolved = await pluginContext.resolve(descriptor.template.src, descriptor.filename);
        if (resolved) {
            const filename = resolved.id;
            // 如果引入的vue文件，读取对应的template
            if ((0, uni_cli_shared_1.isVueSfcFile)(filename)) {
                const srcDescriptor = (0, descriptorCache_1.getDescriptor)(filename, (0, descriptorCache_1.getResolvedOptions)());
                if (srcDescriptor && srcDescriptor.template?.content) {
                    descriptor.template.content = srcDescriptor.template.content;
                }
            }
            else {
                descriptor.template.content = (0, uni_cli_shared_1.preUVueHtml)(fs_extra_1.default.readFileSync(filename, 'utf-8'));
            }
        }
    }
}
exports.tryResolveTemplateSrc = tryResolveTemplateSrc;
