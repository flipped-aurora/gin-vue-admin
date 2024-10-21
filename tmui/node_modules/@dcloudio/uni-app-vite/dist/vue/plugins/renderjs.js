"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniRenderjsPlugin = exports.APP_RENDERJS_JS = exports.APP_WXS_JS = void 0;
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const debugRenderjs = (0, debug_1.default)('uni:app-renderjs');
exports.APP_WXS_JS = 'app-wxs.js';
exports.APP_RENDERJS_JS = 'app-renderjs.js';
const wxsModulesCache = new WeakMap();
const renderjsModulesCache = new WeakMap();
function uniRenderjsPlugin() {
    let resolvedConfig;
    let userConfig;
    let changed = false;
    return {
        name: 'uni:app-vue-renderjs',
        config(config) {
            userConfig = config;
        },
        configResolved(config) {
            resolvedConfig = config;
            wxsModulesCache.set(resolvedConfig, new Map());
            renderjsModulesCache.set(resolvedConfig, new Map());
        },
        async transform(code, id) {
            const { type, name, filename } = (0, uni_cli_shared_1.parseRenderjs)(id);
            if (!type) {
                return;
            }
            if (type !== 'wxs' && type !== 'renderjs') {
                return;
            }
            debugRenderjs(id);
            this.addWatchFile((0, uni_cli_shared_1.cleanUrl)(id));
            if (!name) {
                this.error((0, uni_cli_shared_1.missingModuleName)(type, code));
            }
            const modulePath = (0, uni_cli_shared_1.normalizePath)(path_1.default.normalize(path_1.default.relative(process.env.UNI_INPUT_DIR, id)));
            const moduleHashId = (0, uni_cli_shared_1.hash)(modulePath);
            const globalName = type === 'wxs' ? uni_shared_1.WXS_MODULES : uni_shared_1.RENDERJS_MODULES;
            const { isProduction } = resolvedConfig;
            const resultCode = normalizeCode(type === 'wxs'
                ? await transformWxs(code, filename, `__${globalName}['${moduleHashId}']`, isProduction, userConfig)
                : await transformRenderjs(code, filename, `__${globalName}['${moduleHashId}']`, isProduction, userConfig), globalName, isProduction);
            if (type === 'wxs') {
                wxsModulesCache.get(resolvedConfig).set(moduleHashId, resultCode);
            }
            else {
                renderjsModulesCache.get(resolvedConfig).set(moduleHashId, resultCode);
            }
            changed = true;
            debugRenderjs(type, modulePath, moduleHashId);
            return {
                code: `export default Comp => {
          ;(Comp.$${type} || (Comp.$${type} = [])).push('${name}')
          ;(Comp.$${globalName} || (Comp.$${globalName} = {}))['${name}'] = '${moduleHashId}'
        }`,
                map: { mappings: '' },
            };
        },
        generateBundle() {
            if (!changed) {
                return;
            }
            const wxsCode = [...wxsModulesCache.get(resolvedConfig).values()].join('\n');
            if (wxsCode) {
                this.emitFile({
                    fileName: exports.APP_WXS_JS,
                    source: `var __${uni_shared_1.WXS_MODULES}={};\n` + wxsCode,
                    type: 'asset',
                });
            }
            const renderjsCode = [
                ...renderjsModulesCache.get(resolvedConfig).values(),
            ].join('\n');
            if (renderjsCode) {
                this.emitFile({
                    fileName: exports.APP_RENDERJS_JS,
                    source: `var __${uni_shared_1.RENDERJS_MODULES}={};\n` + renderjsCode,
                    type: 'asset',
                });
            }
        },
    };
}
exports.uniRenderjsPlugin = uniRenderjsPlugin;
function normalizeCode(code, globalName, isProduction) {
    return code.replace(isProduction
        ? `var __${globalName}=__${globalName}||{};`
        : `var __${globalName} = __${globalName} || {};`, '');
}
function transformWxs(code, filename, globalName, isProduction, config) {
    return (0, uni_cli_shared_1.transformWithEsbuild)(code, filename, {
        format: 'iife',
        globalName,
        target: config.build?.target || 'es6',
        minify: isProduction ? true : false,
        bundle: true,
        write: false,
    }).then((res) => {
        if (res.outputFiles) {
            return res.outputFiles[0].text;
        }
        return '';
    });
}
function transformRenderjs(code, filename, globalName, isProduction, config) {
    return (0, uni_cli_shared_1.transformWithEsbuild)(code, filename, {
        format: 'iife',
        globalName,
        target: config.build?.target || 'es6',
        minify: isProduction ? true : false,
        bundle: true,
        write: false,
    }).then((res) => {
        if (res.outputFiles) {
            return res.outputFiles[0].text;
        }
        return '';
    });
}
