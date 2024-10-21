"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.initLogger = exports.createConfigResolved = void 0;
const shared_1 = require("@vue/shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const env_1 = require("./env");
const options_1 = require("./options");
const plugins_1 = require("./plugins");
const resolve_1 = require("../config/resolve");
function createConfigResolved(options) {
    return (config) => {
        // 如果是混合编译且是 nvue 时，部分逻辑无需执行
        if (!(0, uni_cli_shared_1.isInHybridNVue)(config)) {
            (0, env_1.initEnv)(config);
        }
        initLogger(config);
        (0, options_1.initOptions)(options, config);
        (0, plugins_1.initPlugins)(config, options);
        if (!(0, uni_cli_shared_1.isInHybridNVue)(config)) {
            initCheckUpdate();
        }
        if (uni_cli_shared_1.isWindows) {
            // TODO 等 https://github.com/vitejs/vite/issues/3331 修复后，可以移除下列代码
            // 2.8.0 已修复，但为了兼容旧版本，先不移除
            const item = config.resolve.alias.find((item) => !(0, shared_1.isString)(item.find) ? item.find.test('@/') : false);
            if (item) {
                item.customResolver = resolve_1.customResolver;
            }
        }
    };
}
exports.createConfigResolved = createConfigResolved;
function initCheckUpdate() {
    (0, uni_cli_shared_1.checkUpdate)({
        inputDir: process.env.UNI_INPUT_DIR,
        compilerVersion: process.env.UNI_COMPILER_VERSION,
        versionType: process.env.UNI_COMPILER_VERSION_TYPE,
    });
}
function initLogger({ logger, nvue, }) {
    const { info, warn, error } = logger;
    logger.info = (msg, opts) => {
        msg = (0, uni_cli_shared_1.formatInfoMsg)(msg, (0, shared_1.extend)(opts || {}, { nvue }));
        if (msg) {
            return info(msg, opts);
        }
    };
    logger.warn = (msg, opts) => {
        msg = (0, uni_cli_shared_1.formatWarnMsg)(msg, opts);
        if (msg) {
            return warn(msg, opts);
        }
    };
    logger.error = (msg, opts) => {
        msg = (0, uni_cli_shared_1.formatErrMsg)(msg, opts);
        if (msg) {
            return error(msg, opts);
        }
    };
}
exports.initLogger = initLogger;
