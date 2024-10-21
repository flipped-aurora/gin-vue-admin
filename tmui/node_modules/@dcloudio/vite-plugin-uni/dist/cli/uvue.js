"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildUVue = exports.runUVueAndroidBuild = exports.runUVueAndroidDev = exports.initUVueEnv = void 0;
const vite_1 = require("vite");
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const build_1 = require("./build");
const utils_1 = require("./utils");
const easycom_1 = require("../utils/easycom");
const action_1 = require("./action");
function initUVueEnv() {
    // 直接指定了
    if (process.env.UNI_APP_X === 'false') {
        return;
    }
    // 没有手动指定时，才需要动态读取 manifest.json
    if (process.env.UNI_APP_X !== 'true') {
        const manifestJson = (0, uni_cli_shared_1.parseManifestJsonOnce)(process.env.UNI_INPUT_DIR);
        const isNVueEnabled = (0, shared_1.hasOwn)(manifestJson, 'uni-app-x');
        if (!isNVueEnabled) {
            return;
        }
        process.env.UNI_APP_X = 'true';
        if (manifestJson['uni-app-x']?.singleThread === false) {
            process.env.UNI_APP_X_SINGLE_THREAD = 'false';
        }
    }
    if (process.env.UNI_PLATFORM === 'app') {
        process.env.UNI_APP_X_UVUE_SCRIPT_ENGINE = 'native';
        // 如果是app-ios，目前强制使用js引擎
        if (process.env.UNI_UTS_PLATFORM === 'app-ios') {
            process.env.UNI_APP_X_UVUE_SCRIPT_ENGINE = 'js';
        }
    }
}
exports.initUVueEnv = initUVueEnv;
async function runUVueAndroidDev(options) {
    if (options.platform !== 'app') {
        (0, uni_cli_shared_1.output)('error', uni_cli_shared_1.M['uvue.unsupported'].replace('{platform}', options.platform));
        return process.exit(0);
    }
    (0, easycom_1.initEasycom)();
    const watcher = (await buildUVue(options));
    let isFirstStart = true;
    let isFirstEnd = true;
    watcher.on('event', async (event) => {
        if (event.code === 'BUNDLE_START') {
            if (isFirstStart) {
                isFirstStart = false;
                return;
            }
            (0, uni_cli_shared_1.output)('log', uni_cli_shared_1.M['dev.watching.start']);
            // 重置一下，uts编译报错会导致下一次开始差量编译紧接着上一次的差量编译，导致无法正常输出
            (0, uni_cli_shared_1.resetOutput)('log');
        }
        else if (event.code === 'BUNDLE_END') {
            event.result.close();
            const dex = process.env.UNI_APP_UTS_CHANGED_FILES;
            process.env.UNI_APP_UTS_CHANGED_FILES = '';
            if (isFirstEnd) {
                // 首次全量同步
                isFirstEnd = false;
                (0, uni_cli_shared_1.output)('log', uni_cli_shared_1.M['dev.watching.end']);
                (0, utils_1.printStartupDuration)((0, vite_1.createLogger)(options.logLevel), false);
                await (0, action_1.stopProfiler)((message) => (0, vite_1.createLogger)(options.logLevel).info(message));
                return;
            }
            if (dex) {
                const files = JSON.parse(dex);
                if (!files.length) {
                    // 本次无变动
                    return (0, uni_cli_shared_1.output)('log', uni_cli_shared_1.M['uvue.dev.watching.end.empty']);
                }
                return (0, uni_cli_shared_1.output)('log', uni_cli_shared_1.M['dev.watching.end.files'].replace('{files}', JSON.stringify(files)));
            }
            return (0, uni_cli_shared_1.output)('log', uni_cli_shared_1.M['dev.watching.end']);
        }
    });
}
exports.runUVueAndroidDev = runUVueAndroidDev;
async function runUVueAndroidBuild(options) {
    try {
        (0, uni_cli_shared_1.initEasycomsOnce)(process.env.UNI_INPUT_DIR, {
            dirs: [(0, uni_cli_shared_1.resolveComponentsLibPath)()],
            platform: process.env.UNI_PLATFORM,
            isX: true,
        });
        await buildUVue(options);
        await (0, action_1.stopProfiler)((message) => (0, vite_1.createLogger)(options.logLevel).info(message));
        console.log(uni_cli_shared_1.M['build.done']);
        // 开发者可能用了三方插件，三方插件有可能阻止退出，导致HBuilderX打包状态识别不正确
        if ((0, uni_cli_shared_1.isInHBuilderX)()) {
            process.exit(0);
        }
    }
    catch (e) {
        console.error(e);
        console.error(`Build failed with errors.`);
        process.exit(1);
    }
}
exports.runUVueAndroidBuild = runUVueAndroidBuild;
/**
 * 目前的简易实现逻辑
 * node层：
 *  1. 监听项目，生成资源到临时目录 .uts/android
 *  2. uvue 文件，做解析，拆分生成 render.kt, css.kt, uts.uvue
 *  3. static 文件，copy 到最终目录
 *  4. uvue、vue、uts 文件发生变化，调用 uts 编译器
 * @param options
 */
async function buildUVue(options) {
    return (0, build_1.buildByVite)((0, utils_1.addConfigFile)((0, shared_1.extend)({ nvueAppService: true, uvue: true }, (0, build_1.initBuildOptions)(options, (0, utils_1.cleanOptions)(options)))));
}
exports.buildUVue = buildUVue;
