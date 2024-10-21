"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.filterPrefersColorScheme = void 0;
const shared_1 = require("@vue/shared");
const postcss_selector_parser_1 = __importDefault(require("postcss-selector-parser"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const defaultUniAppCssProcessorOptions = (0, shared_1.extend)({}, uni_shared_1.defaultRpx2Unit);
const BG_PROPS = [
    'background',
    'background-clip',
    'background-color',
    'background-image',
    'background-origin',
    'background-position',
    'background-repeat',
    'background-size',
    'background-attachment',
];
function transform(selector, state, { rewriteTag }) {
    if (selector.type !== 'tag') {
        return;
    }
    const { value } = selector;
    selector.value = rewriteTag(value);
    if (value === 'page' && selector.value === 'uni-page-body') {
        state.bg = true;
    }
}
function createBodyBackgroundRule(origRule) {
    const bgDecls = [];
    origRule.walkDecls((decl) => {
        if (BG_PROPS.indexOf(decl.prop) !== -1) {
            bgDecls.push(decl.clone());
        }
    });
    if (bgDecls.length) {
        const { rule } = require('postcss');
        origRule.after(rule({ selector: 'body' }).append(bgDecls));
    }
}
function walkRules(options) {
    return (rule) => {
        const state = { bg: false };
        rule.selector = (0, postcss_selector_parser_1.default)((selectors) => selectors.walk((selector) => transform(selector, state, options))).processSync(rule.selector);
        state.bg && createBodyBackgroundRule(rule);
    };
}
function walkDecls(rpx2unit) {
    return (decl) => {
        const { value } = decl;
        if (value.indexOf('rpx') === -1 && value.indexOf('upx') === -1) {
            return;
        }
        decl.value = rpx2unit(decl.value);
    };
}
function filterPrefersColorScheme(root, force = false) {
    if (process.env.VUE_APP_DARK_MODE !== 'true') {
        const filePath = root.source?.input.file;
        if (force || (filePath && filePath.includes('@dcloudio'))) {
            root.walkAtRules((rule) => {
                if (/prefers-color-scheme\s*:\s*dark/.test(rule.params)) {
                    rule.remove();
                }
            });
        }
    }
}
exports.filterPrefersColorScheme = filterPrefersColorScheme;
const baiduTags = {
    navigator: 'nav',
};
function rewriteBaiduTags(tag) {
    return baiduTags[tag] || tag;
}
function rewriteUniH5Tags(tag) {
    if (tag === 'page') {
        return 'uni-page-body';
    }
    if ((0, uni_shared_1.isBuiltInComponent)(tag)) {
        return uni_shared_1.COMPONENT_SELECTOR_PREFIX + tag;
    }
    return tag;
}
function rewriteUniAppTags(tag) {
    if (tag === 'page') {
        return 'body';
    }
    if ((0, uni_shared_1.isBuiltInComponent)(tag)) {
        return uni_shared_1.COMPONENT_SELECTOR_PREFIX + tag;
    }
    return tag;
}
const transforms = {
    h5: rewriteUniH5Tags,
    app: rewriteUniAppTags,
    'app-harmony': rewriteUniAppTags,
    'mp-baidu': rewriteBaiduTags,
};
const uniapp = (opts) => {
    const platform = process.env.UNI_PLATFORM;
    const { unit, unitRatio, unitPrecision } = (0, shared_1.extend)({}, defaultUniAppCssProcessorOptions, opts);
    const rpx2unit = (0, uni_shared_1.createRpx2Unit)(unit, unitRatio, unitPrecision);
    return {
        postcssPlugin: 'uni-app',
        prepare() {
            return {
                OnceExit(root) {
                    root.walkDecls(walkDecls(rpx2unit));
                    const rewriteTag = transforms[platform];
                    filterPrefersColorScheme(root);
                    if (rewriteTag) {
                        root.walkRules(walkRules({
                            rewriteTag,
                        }));
                    }
                },
            };
        },
    };
};
uniapp.postcss = true;
exports.default = uniapp;
