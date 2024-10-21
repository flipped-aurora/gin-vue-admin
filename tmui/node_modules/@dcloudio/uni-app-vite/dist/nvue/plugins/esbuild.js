"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.wrapperNVueAppStyles = exports.uniEsbuildPlugin = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const debug_1 = __importDefault(require("debug"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../../utils");
const utils_2 = require("../utils");
const appCss_1 = require("./appCss");
const debugEsbuild = (0, debug_1.default)('uni:app-nvue-esbuild');
const emittedHashMap = new WeakMap();
function uniEsbuildPlugin({ appService, }) {
    let resolvedConfig;
    let buildOptions;
    const nvueOutputDir = (0, utils_1.nvueOutDir)();
    const outputDir = process.env.UNI_OUTPUT_DIR;
    let isFirst = true;
    return {
        name: 'uni:app-nvue-esbuild',
        enforce: 'post',
        configResolved(config) {
            buildOptions = {
                format: 'iife',
                target: 'es6',
                minify: config.build.minify ? true : false,
                banner: {
                    js: `"use weex:vue";
${uni_cli_shared_1.polyfillCode}`,
                },
                bundle: true,
                write: false,
                plugins: [esbuildGlobalPlugin((0, utils_2.esbuildGlobals)(appService))],
            };
            resolvedConfig = config;
            emittedHashMap.set(resolvedConfig, new Map());
        },
        async writeBundle(_, bundle) {
            const entryPoints = [];
            const assets = [];
            Object.keys(bundle).forEach((name) => {
                const chunk = bundle[name];
                if (chunk.type === 'chunk' &&
                    chunk.facadeModuleId &&
                    chunk.facadeModuleId.endsWith('.nvue')) {
                    entryPoints.push(name);
                }
                else if (chunk.type === 'asset') {
                    assets.push(name);
                }
            });
            // 仅 nvueOutputDir 时 copy
            if (!appService) {
                assets.forEach((name) => {
                    fs_extra_1.default.copySync(path_1.default.resolve(nvueOutputDir, name), path_1.default.resolve(outputDir, name), { overwrite: false });
                });
            }
            if (!entryPoints.length) {
                return;
            }
            const emittedHash = emittedHashMap.get(resolvedConfig);
            const changedFiles = [];
            if (buildAppCss()) {
                changedFiles.push(uni_cli_shared_1.APP_CONFIG_SERVICE);
            }
            debugEsbuild('start', entryPoints.length, entryPoints);
            for (const filename of entryPoints) {
                await buildNVuePage(filename, buildOptions).then((code) => {
                    const outputFileHash = (0, uni_cli_shared_1.hash)(code);
                    if (emittedHash.get(filename) !== outputFileHash) {
                        changedFiles.push(filename);
                        emittedHash.set(filename, outputFileHash);
                        return fs_extra_1.default.outputFile(path_1.default.resolve(outputDir, filename), code);
                    }
                });
            }
            if (!isFirst && changedFiles.length) {
                process.env[changedFiles.includes(uni_cli_shared_1.APP_CONFIG_SERVICE)
                    ? 'UNI_APP_CHANGED_FILES'
                    : 'UNI_APP_CHANGED_PAGES'] = JSON.stringify(changedFiles);
            }
            debugEsbuild('end');
            isFirst = false;
        },
    };
}
exports.uniEsbuildPlugin = uniEsbuildPlugin;
/**
 * 将 nvue 全局 css 样式注入 app-config-service.js
 * @returns
 */
function buildAppCss() {
    const appCssJsFilename = path_1.default.join((0, utils_1.nvueOutDir)(), appCss_1.APP_CSS_JS);
    if (!fs_extra_1.default.existsSync(appCssJsFilename)) {
        return;
    }
    const appCssJsCode = fs_extra_1.default.readFileSync(appCssJsFilename, 'utf8');
    const appCssJsFn = new Function('module', 
    // vite build.target为esnext时, 生成的代码没有export default
    appCssJsCode.includes('export default')
        ? appCssJsCode.replace(`export default`, `module.exports=`)
        : appCssJsCode.replace(`exports`, `module.exports`));
    const module = { exports: { styles: [] } };
    appCssJsFn(module);
    const appCssJsonCode = JSON.stringify(module.exports.styles);
    if (process.env.UNI_NVUE_APP_STYLES === appCssJsonCode) {
        return;
    }
    process.env.UNI_NVUE_APP_STYLES = appCssJsonCode;
    // 首次 build 时，可能还没生成 app-config-service 的文件，故仅写入环境变量
    const appConfigServiceFilename = path_1.default.join(process.env.UNI_OUTPUT_DIR, uni_cli_shared_1.APP_CONFIG_SERVICE);
    if (!fs_extra_1.default.existsSync(appConfigServiceFilename)) {
        return;
    }
    const appConfigServiceCode = fs_extra_1.default.readFileSync(appConfigServiceFilename, 'utf8');
    fs_extra_1.default.writeFileSync(appConfigServiceFilename, wrapperNVueAppStyles(appConfigServiceCode));
    return true;
}
function buildNVuePage(filename, options) {
    return (0, uni_cli_shared_1.transformWithEsbuild)(`import App from './${filename}'
const webview = plus.webview.currentWebview()
if(webview){
  const __pageId = parseInt(webview.id)
  const __pagePath = '${(0, uni_cli_shared_1.removeExt)(filename)}'
  let __pageQuery = {}
  try{ __pageQuery = JSON.parse(webview.__query__) }catch(e){}
  App.mpType = 'page'
  const app = Vue.createPageApp(App,{$store:getApp({allowDefault:true}).$store,__pageId,__pagePath,__pageQuery})
  app.provide('__globalStyles', Vue.useCssStyles([...__uniConfig.styles, ...(App.styles||[])]))
  app.mount('#root')
}`, path_1.default.join((0, utils_1.nvueOutDir)(), 'main.js'), options).then((res) => {
        if (res.outputFiles) {
            return res.outputFiles[0].text;
        }
        return '';
    });
}
function esbuildGlobalPlugin(options) {
    const keys = Object.keys(options);
    return {
        name: 'global',
        setup(build) {
            keys.forEach((key) => {
                const namespace = key + '-ns';
                build.onResolve({ filter: new RegExp('^' + key + '$') }, ({ path }) => {
                    return {
                        path,
                        namespace,
                    };
                });
                build.onLoad({ filter: /.*/, namespace }, () => ({
                    contents: `module.exports = ${options[key]}`,
                    loader: 'js',
                }));
            });
        },
    };
}
function wrapperNVueAppStyles(code) {
    return code.replace(/__uniConfig.styles=(.*);\/\/styles/, `__uniConfig.styles=${process.env.UNI_NVUE_APP_STYLES || '[]'};//styles`);
}
exports.wrapperNVueAppStyles = wrapperNVueAppStyles;
