"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniEasycomPlugin = void 0;
const path_1 = __importDefault(require("path"));
const shared_1 = require("@vue/shared");
const pluginutils_1 = require("@rollup/pluginutils");
const url_1 = require("../utils/url");
const constants_1 = require("../../constants");
const easycom_1 = require("../../easycom");
function uniEasycomPlugin(options) {
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    return {
        name: 'uni:app-easycom',
        transform(code, id) {
            if (!filter(id)) {
                return;
            }
            const { filename } = (0, url_1.parseVueRequest)(id);
            if (!constants_1.EXTNAME_VUE_TEMPLATE.includes(path_1.default.extname(filename))) {
                return;
            }
            if (!code.includes('_resolveComponent')) {
                return;
            }
            let i = 0;
            const importDeclarations = [];
            code = code.replace(/_resolveComponent\("(.+?)"(, true)?\)/g, (str, name) => {
                if (name && !name.startsWith('_')) {
                    const source = (0, easycom_1.matchEasycom)(name);
                    if (source) {
                        // 处理easycom组件优先级
                        return (0, easycom_1.genResolveEasycomCode)(importDeclarations, str, (0, easycom_1.addImportDeclaration)(importDeclarations, `__easycom_${i++}`, source, source.includes('uts-proxy')
                            ? (0, shared_1.capitalize)((0, shared_1.camelize)(name)) + 'Component'
                            : source.includes('uni_helpers')
                                ? (0, shared_1.capitalize)((0, shared_1.camelize)(name))
                                : ''));
                    }
                }
                return str;
            });
            if (importDeclarations.length) {
                code = importDeclarations.join('') + code;
            }
            return {
                code,
                map: null,
            };
        },
    };
}
exports.uniEasycomPlugin = uniEasycomPlugin;
