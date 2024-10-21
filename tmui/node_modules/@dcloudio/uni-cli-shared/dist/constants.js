"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TEXT_STYLE = exports.DEFAULT_ASSETS_RE = exports.KNOWN_ASSET_TYPES = exports.COMMON_EXCLUDE = exports.X_BASE_COMPONENTS_STYLE_PATH = exports.BASE_COMPONENTS_STYLE_PATH = exports.H5_COMPONENTS_STYLE_PATH = exports.H5_FRAMEWORK_STYLE_PATH = exports.H5_API_STYLE_PATH = exports.X_PAGE_EXTNAME_APP = exports.X_PAGE_EXTNAME = exports.PAGE_EXTNAME = exports.PAGE_EXTNAME_APP = exports.BINDING_COMPONENTS = exports.APP_CONFIG_SERVICE = exports.APP_CONFIG = exports.APP_SERVICE_FILENAME = exports.ASSETS_INLINE_LIMIT = exports.JSON_JS_MAP = exports.MANIFEST_JSON_UTS = exports.MANIFEST_JSON_JS = exports.PAGES_JSON_UTS = exports.PAGES_JSON_JS = exports.uni_app_x_extensions = exports.extensions = exports.EXTNAME_TS_RE = exports.EXTNAME_JS_RE = exports.EXTNAME_VUE_RE = exports.EXTNAME_VUE_TEMPLATE = exports.X_EXTNAME_VUE = exports.EXTNAME_VUE = exports.EXTNAME_TS = exports.EXTNAME_JS = exports.PUBLIC_DIR = void 0;
exports.PUBLIC_DIR = 'static';
exports.EXTNAME_JS = ['.js', '.ts', '.jsx', '.tsx', '.uts'];
exports.EXTNAME_TS = ['.ts', '.tsx'];
exports.EXTNAME_VUE = ['.vue', '.nvue', '.uvue'];
exports.X_EXTNAME_VUE = ['.uvue', '.vue'];
exports.EXTNAME_VUE_TEMPLATE = ['.vue', '.nvue', '.uvue', '.jsx', '.tsx'];
exports.EXTNAME_VUE_RE = /\.(vue|nvue|uvue)$/;
exports.EXTNAME_JS_RE = /\.(js|jsx|ts|uts|tsx|mjs)$/;
exports.EXTNAME_TS_RE = /\.tsx?$/;
const COMMON_EXTENSIONS = [
    '.uts',
    '.mjs',
    '.js',
    '.ts',
    '.jsx',
    '.tsx',
    '.json',
];
exports.extensions = COMMON_EXTENSIONS.concat(exports.EXTNAME_VUE);
exports.uni_app_x_extensions = COMMON_EXTENSIONS.concat(['.uvue', '.vue']);
exports.PAGES_JSON_JS = 'pages-json-js';
exports.PAGES_JSON_UTS = 'pages-json-uts';
exports.MANIFEST_JSON_JS = 'manifest-json-js';
exports.MANIFEST_JSON_UTS = 'manifest-json-uts';
exports.JSON_JS_MAP = {
    'pages.json': exports.PAGES_JSON_JS,
    'manifest.json': exports.MANIFEST_JSON_JS,
};
exports.ASSETS_INLINE_LIMIT = 40 * 1024;
exports.APP_SERVICE_FILENAME = 'app-service.js';
exports.APP_CONFIG = 'app-config.js';
exports.APP_CONFIG_SERVICE = 'app-config-service.js';
exports.BINDING_COMPONENTS = '__BINDING_COMPONENTS__';
// APP 平台解析页面后缀的优先级
exports.PAGE_EXTNAME_APP = ['.nvue', '.vue', '.tsx', '.jsx', '.js'];
// 其他平台解析页面后缀的优先级
exports.PAGE_EXTNAME = ['.vue', '.nvue', '.tsx', '.jsx', '.js'];
exports.X_PAGE_EXTNAME = ['.uvue', '.vue', '.tsx', '.jsx', '.js'];
exports.X_PAGE_EXTNAME_APP = ['.uvue', '.tsx', '.jsx', '.js'];
exports.H5_API_STYLE_PATH = '@dcloudio/uni-h5/style/api/';
exports.H5_FRAMEWORK_STYLE_PATH = '@dcloudio/uni-h5/style/framework/';
exports.H5_COMPONENTS_STYLE_PATH = '@dcloudio/uni-h5/style/';
exports.BASE_COMPONENTS_STYLE_PATH = '@dcloudio/uni-components/style/';
exports.X_BASE_COMPONENTS_STYLE_PATH = '@dcloudio/uni-components/style-x/';
exports.COMMON_EXCLUDE = [
    /\/pages\.json\.js$/,
    /\/manifest\.json\.js$/,
    /\/vite\//,
    /\/@vue\//,
    /\/vue-router\//,
    /\/vuex\//,
    /\/vue-i18n\//,
    /\/@dcloudio\/uni-h5-vue/,
    /\/@dcloudio\/uni-shared/,
];
exports.KNOWN_ASSET_TYPES = [
    // images
    'png',
    'jpe?g',
    'gif',
    'svg',
    'ico',
    'webp',
    'avif',
    // media
    'mp4',
    'webm',
    'ogg',
    'mp3',
    'wav',
    'flac',
    'aac',
    // fonts
    'woff2?',
    'eot',
    'ttf',
    'otf',
    // other
    'pdf',
    'txt',
];
exports.DEFAULT_ASSETS_RE = new RegExp(`\\.(` + exports.KNOWN_ASSET_TYPES.join('|') + `)(\\?.*)?$`);
exports.TEXT_STYLE = ['black', 'white'];
