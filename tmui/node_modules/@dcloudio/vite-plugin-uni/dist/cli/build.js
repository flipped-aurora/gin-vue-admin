"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildApp = exports.initBuildOptions = exports.buildSSR = exports.build = exports.buildByVite = void 0;
const path_1 = __importDefault(require("path"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("./utils");
async function buildByVite(inlineConfig) {
    return Promise.resolve().then(() => __importStar(require('vite'))).then(({ build }) => build(inlineConfig));
}
exports.buildByVite = buildByVite;
async function build(options) {
    if (process.env.UNI_APP_X !== 'true' && options.platform === 'app') {
        return buildApp(options);
    }
    return buildByVite((0, utils_1.addConfigFile)(initBuildOptions(options, (0, utils_1.cleanOptions)(options))));
}
exports.build = build;
async function buildSSR(options) {
    const outputDir = process.env.UNI_OUTPUT_DIR;
    const ssrClientDir = path_1.default.resolve(outputDir, 'client');
    process.env.UNI_OUTPUT_DIR = ssrClientDir;
    const ssrBuildClientOptions = (0, utils_1.cleanOptions)(options);
    ssrBuildClientOptions.ssrManifest = 'ssr-manifest.json';
    ssrBuildClientOptions.outDir = process.env.UNI_OUTPUT_DIR;
    process.env.UNI_SSR_CLIENT = 'true';
    await buildByVite((0, utils_1.addConfigFile)(initBuildOptions(options, ssrBuildClientOptions)));
    const ssrServerDir = path_1.default.resolve(outputDir, 'server');
    process.env.UNI_OUTPUT_DIR = ssrServerDir;
    const ssrBuildServerOptions = (0, utils_1.cleanOptions)(options);
    ssrBuildServerOptions.ssr = path_1.default.resolve(process.env.UNI_INPUT_DIR, 'entry-server.js');
    // 强制 cjs 输出
    ssrBuildServerOptions.rollupOptions = {
        output: {
            format: 'cjs',
        },
    };
    ssrBuildServerOptions.outDir = process.env.UNI_OUTPUT_DIR;
    process.env.UNI_SSR_CLIENT = '';
    process.env.UNI_SSR_SERVER = 'true';
    await buildByVite((0, utils_1.addConfigFile)(initBuildOptions(options, ssrBuildServerOptions)));
    // copy ssr-manfiest.json to server
    const assets = ['ssr-manifest.json', 'index.html'];
    assets.forEach((asset) => {
        const ssrManifestFile = path_1.default.join(ssrClientDir, asset);
        if (fs_extra_1.default.existsSync(ssrManifestFile)) {
            fs_extra_1.default.copyFileSync(ssrManifestFile, path_1.default.join(ssrServerDir, asset));
        }
    });
}
exports.buildSSR = buildSSR;
function initBuildOptions(options, build) {
    return {
        root: process.env.VITE_ROOT_DIR,
        configFile: options.config,
        base: options.base,
        logLevel: options.logLevel,
        clearScreen: options.clearScreen,
        mode: options.mode,
        build,
    };
}
exports.initBuildOptions = initBuildOptions;
function buildManifestJson() {
    const platform = 'app';
    const inputDir = process.env.UNI_INPUT_DIR;
    const outputDir = process.env.UNI_OUTPUT_DIR;
    const pkg = require(path_1.default.resolve(__dirname, '../../package.json'));
    process.env.UNI_COMPILER_VERSION =
        process.env.UNI_COMPILER_VERSION ||
            pkg['uni-app']?.['compilerVersion'] ||
            '';
    const manifestJson = (0, uni_cli_shared_1.normalizeAppManifestJson)((0, uni_cli_shared_1.parseManifestJsonOnce)(inputDir), (0, uni_cli_shared_1.parsePagesJsonOnce)(inputDir, platform));
    fs_extra_1.default.outputFileSync(path_1.default.resolve(outputDir, 'manifest.json'), JSON.stringify(manifestJson, null, 2));
}
async function buildApp(options) {
    if (options.manifest) {
        return buildManifestJson();
    }
    let appWatcher;
    if (options.watch) {
        appWatcher = new AppWatcher();
    }
    if (process.env.UNI_RENDERER === 'native') {
        // 纯原生渲染时，main.js + App.vue 需要跟页面分开，独立编译（因为需要包含 Vuex 等共享内容）
        process.env.UNI_COMPILER = 'nvue';
        process.env.UNI_RENDERER_NATIVE = 'appService';
        const nvueAppBuilder = await buildByVite((0, utils_1.addConfigFile)((0, shared_1.extend)({ nvueAppService: true, nvue: true }, initBuildOptions(options, (0, utils_1.cleanOptions)(options)))));
        if (appWatcher) {
            appWatcher.setFirstWatcher(nvueAppBuilder);
        }
        process.env.UNI_RENDERER_NATIVE = 'pages';
        const nvueBuilder = await buildByVite((0, utils_1.addConfigFile)((0, shared_1.extend)({ nvue: true }, initBuildOptions(options, (0, utils_1.cleanOptions)(options)))));
        if (appWatcher) {
            appWatcher.setSecondWatcher(nvueBuilder);
            return appWatcher;
        }
        return;
    }
    // 指定为 vue 方便 App 插件初始化 vue 所需插件列表
    process.env.UNI_COMPILER = 'vue';
    const vueBuilder = await buildByVite((0, utils_1.addConfigFile)(initBuildOptions(options, (0, utils_1.cleanOptions)(options))));
    if (process.env.UNI_COMPILE_TARGET === 'uni_modules') {
        // 不需要 nvue 编译器
        return vueBuilder;
    }
    if (appWatcher) {
        appWatcher.setFirstWatcher(vueBuilder);
    }
    // 临时指定为 nvue 方便 App 插件初始化 nvue 所需插件列表
    process.env.UNI_COMPILER = 'nvue';
    const nvueBuilder = await buildByVite((0, utils_1.addConfigFile)((0, shared_1.extend)({ nvue: true }, initBuildOptions(options, (0, utils_1.cleanOptions)(options)))));
    // 还原为 vue
    process.env.UNI_COMPILER = 'vue';
    if (appWatcher) {
        appWatcher.setSecondWatcher(nvueBuilder);
        return appWatcher;
    }
}
exports.buildApp = buildApp;
class AppWatcher {
    constructor() {
        this._firstStart = false;
        this._firstEnd = false;
        this._secondStart = false;
        this._secondEnd = false;
    }
    on(_event, callback) {
        this._callback = callback;
    }
    setFirstWatcher(firstWatcher) {
        firstWatcher.on('event', (event) => {
            if (event.code === 'BUNDLE_START') {
                this._bundleFirstStart(event);
            }
            else if (event.code === 'BUNDLE_END') {
                this._bundleFirstEnd(event);
            }
        });
    }
    setSecondWatcher(secondWatcher) {
        secondWatcher.on('event', (event) => {
            if (event.code === 'BUNDLE_START') {
                this._bundleSecondStart(event);
            }
            else if (event.code === 'BUNDLE_END') {
                this._bundleSecondEnd(event);
            }
        });
    }
    _bundleFirstStart(event) {
        this._firstStart = true;
        this._bundleStart(event);
    }
    _bundleFirstEnd(event) {
        this._firstEnd = true;
        this._bundleEnd(event);
    }
    _bundleSecondStart(event) {
        this._secondStart = true;
        this._bundleStart(event);
    }
    _bundleSecondEnd(event) {
        this._secondEnd = true;
        this._bundleEnd(event);
    }
    _bundleStart(event) {
        if (this._firstStart && this._secondStart) {
            this._callback(event);
        }
    }
    _bundleEnd(event) {
        if (this._firstEnd && this._secondEnd) {
            this._callback(event);
        }
    }
}
