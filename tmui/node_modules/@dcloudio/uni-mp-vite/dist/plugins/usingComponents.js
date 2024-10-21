"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.dynamicImport = exports.uniUsingComponentsPlugin = void 0;
const path_1 = __importDefault(require("path"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const entry_1 = require("./entry");
function uniUsingComponentsPlugin(options = {}) {
    const normalizeComponentName = options.normalizeComponentName || ((name) => name);
    const parseAst = (source, id) => {
        return (0, uni_cli_shared_1.parseProgram)(source, id, {
            babelParserPlugins: options.babelParserPlugins,
        });
    };
    const inputDir = process.env.UNI_INPUT_DIR;
    let resolvedConfig;
    return {
        name: 'uni:mp-using-component',
        enforce: 'post',
        configResolved(config) {
            resolvedConfig = config;
        },
        async transform(source, id) {
            const { filename, query } = (0, uni_cli_shared_1.parseVueRequest)(id);
            if ((0, uni_cli_shared_1.isAppVue)(filename)) {
                return null;
            }
            const sourceMap = (0, uni_cli_shared_1.withSourcemap)(resolvedConfig);
            const dynamicImportOptions = {
                id,
                sourceMap,
                dynamicImport,
            };
            const resolve = async (source, importer, options) => {
                const id = (0, uni_cli_shared_1.resolveUTSModule)(source, importer || process.env.UNI_INPUT_DIR);
                if (id) {
                    source = id;
                }
                return this.resolve(source, importer, options);
            };
            if (query.vue) {
                if (query.type === 'script') {
                    // 需要主动监听
                    this.addWatchFile(filename);
                    const descriptor = await (0, uni_cli_shared_1.parseScriptDescriptor)(filename, parseAst(source, id), {
                        resolve,
                        isExternal: true,
                    });
                    (0, uni_cli_shared_1.updateMiniProgramComponentsByScriptFilename)(filename, inputDir, normalizeComponentName);
                    return (0, uni_cli_shared_1.transformDynamicImports)(source, descriptor.imports, dynamicImportOptions);
                }
                else if (query.type === 'template') {
                    // 需要主动监听
                    this.addWatchFile(filename);
                    const descriptor = await (0, uni_cli_shared_1.parseTemplateDescriptor)(filename, parseAst(source, id), {
                        resolve,
                        isExternal: true,
                    });
                    (0, uni_cli_shared_1.updateMiniProgramComponentsByTemplateFilename)(filename, inputDir, normalizeComponentName);
                    return (0, uni_cli_shared_1.transformDynamicImports)(source, descriptor.imports, dynamicImportOptions);
                }
                return null;
            }
            if (!uni_cli_shared_1.EXTNAME_VUE.includes(path_1.default.extname(filename))) {
                return null;
            }
            const ast = parseAst(source, id);
            const descriptor = await (0, uni_cli_shared_1.parseMainDescriptor)(filename, ast, resolve);
            (0, uni_cli_shared_1.updateMiniProgramComponentsByMainFilename)(filename, inputDir, normalizeComponentName);
            return (0, uni_cli_shared_1.transformDynamicImports)(source, descriptor.imports, dynamicImportOptions);
        },
    };
}
exports.uniUsingComponentsPlugin = uniUsingComponentsPlugin;
function dynamicImport(name, value) {
    // 开发者可能将页面作为组件来引用
    if ((0, uni_cli_shared_1.isMiniProgramPageFile)(value, process.env.UNI_INPUT_DIR)) {
        return `const ${name} = ()=>import('${(0, entry_1.virtualPagePath)(value)}')`;
    }
    return `const ${name} = ()=>import('${(0, entry_1.virtualComponentPath)(value)}')`;
}
exports.dynamicImport = dynamicImport;
