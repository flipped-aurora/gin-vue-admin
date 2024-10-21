"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createConfigResolved = void 0;
const debug_1 = __importDefault(require("debug"));
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const pagesJson_1 = require("../plugins/pagesJson");
const entry_1 = require("../plugins/entry");
const debugNVueCss = (0, debug_1.default)('uni:nvue-css');
const cssVars = `page{--status-bar-height:25px;--top-window-height:0px;--window-top:0px;--window-bottom:0px;--window-left:0px;--window-right:0px;--window-magin:0px}`;
const genShadowCss = (cdn) => {
    const url = (0, uni_cli_shared_1.createShadowImageUrl)(cdn, 'grey');
    return `page::after{position:fixed;content:'';left:-1000px;top:-1000px;-webkit-animation:shadow-preload .1s;-webkit-animation-delay:3s;animation:shadow-preload .1s;animation-delay:3s}@-webkit-keyframes shadow-preload{0%{background-image:url(${url})}100%{background-image:url(${url})}}@keyframes shadow-preload{0%{background-image:url(${url})}100%{background-image:url(${url})}}`;
};
const genComponentCustomHiddenCss = (name) => `[${name.replace(':', '')}="true"]{display: none !important;}`;
function createConfigResolved({ cdn, style: { extname }, template: { component }, }) {
    function normalizeCssChunkFilename(id, extname) {
        return ((0, uni_cli_shared_1.removeExt)((0, uni_cli_shared_1.normalizeMiniProgramFilename)(id, process.env.UNI_INPUT_DIR)) +
            extname);
    }
    return (config) => {
        const mainPath = (0, uni_cli_shared_1.resolveMainPathOnce)(process.env.UNI_INPUT_DIR);
        fixUnocss(config);
        (0, uni_cli_shared_1.injectCssPlugin)(config, process.env.UNI_COMPILE_TARGET === 'uni_modules'
            ? {
                createUrlReplacer: uni_cli_shared_1.createEncryptCssUrlReplacer,
            }
            : {});
        let unocssGlobalBuildBundleIndex = config.plugins.findIndex((p) => p.name === 'unocss:global:build:bundle');
        if (unocssGlobalBuildBundleIndex === -1) {
            unocssGlobalBuildBundleIndex = config.plugins.findIndex((p) => p.name === 'unocss:global:build:generate');
        }
        const hasUnocssGlobalBuildBundle = unocssGlobalBuildBundleIndex > -1;
        // unocss 是根据 .css 后缀来编译文件，需要先保持 css 文件后缀为 .css，等 unocss 处理完后，再重置回正确的文件后缀
        const cssExtname = hasUnocssGlobalBuildBundle ? '.css' : extname;
        (0, uni_cli_shared_1.injectCssPostPlugin)(config, (0, uni_cli_shared_1.cssPostPlugin)(config, {
            platform: process.env.UNI_PLATFORM,
            chunkCssFilename(id) {
                if (id === mainPath) {
                    return 'app' + cssExtname;
                }
                else if ((0, entry_1.isUniPageUrl)(id)) {
                    return normalizeCssChunkFilename((0, entry_1.parseVirtualPagePath)(id), cssExtname);
                }
                else if ((0, entry_1.isUniComponentUrl)(id)) {
                    return normalizeCssChunkFilename((0, entry_1.parseVirtualComponentPath)(id), cssExtname);
                }
            },
            chunkCssCode(filename, cssCode) {
                cssCode = (0, uni_cli_shared_1.transformScopedCss)(cssCode);
                if (filename === 'app' + cssExtname) {
                    const componentCustomHiddenCss = (component &&
                        component.vShow &&
                        genComponentCustomHiddenCss(component.vShow)) ||
                        '';
                    if (config.isProduction) {
                        return (cssCode +
                            genShadowCss(cdn || 0) +
                            cssVars +
                            componentCustomHiddenCss);
                    }
                    else {
                        return cssCode + cssVars + componentCustomHiddenCss;
                    }
                }
                const nvueCssPaths = (0, pagesJson_1.getNVueCssPaths)(config);
                if (!nvueCssPaths || !nvueCssPaths.length) {
                    return cssCode;
                }
                const normalized = (0, uni_cli_shared_1.normalizePath)(filename);
                if (nvueCssPaths.find((pageCssPath) => pageCssPath === normalized)) {
                    debugNVueCss(normalized);
                    return (`@import "${(0, uni_cli_shared_1.relativeFile)(normalized, 'nvue' + extname)}";\n` +
                        cssCode);
                }
                return cssCode;
            },
        }));
        (0, uni_cli_shared_1.injectAssetPlugin)(config);
        if (hasUnocssGlobalBuildBundle && extname !== '.css') {
            ;
            config.plugins.splice(unocssGlobalBuildBundleIndex + 1, 0, adjustCssExtname(extname));
        }
    };
}
exports.createConfigResolved = createConfigResolved;
function adjustCssExtname(extname) {
    return {
        name: 'uni:adjust-css-extname',
        generateBundle(_, bundle) {
            const files = Object.keys(bundle);
            files.forEach((name) => {
                if (name.endsWith('.css')) {
                    const asset = bundle[name];
                    (0, shared_1.isString)(asset.source) &&
                        (asset.source = asset.source.replace(/\*\,/g, 'page,'));
                    this.emitFile({
                        fileName: name.replace('.css', extname),
                        type: 'asset',
                        source: asset.source,
                    });
                    delete bundle[name];
                }
            });
        },
    };
}
function fixUnocss(config) {
    const unocssGlobalBuildScan = config.plugins.find((p) => p.name === 'unocss:global:build:scan');
    // TODO 原始的 scan 的 buildStart 会清空 vfsLayerMap，导致 watch 时，load 阶段 /__uno.css 获取不到
    // https://github.com/antfu/unocss/blob/main/packages/vite/src/modes/global/build.ts#L25
    if (unocssGlobalBuildScan) {
        // 隐患: task 未被清空
        unocssGlobalBuildScan.buildStart = () => { };
    }
}
