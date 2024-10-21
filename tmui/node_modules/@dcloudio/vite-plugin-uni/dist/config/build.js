"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createBuild = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const shared_1 = require("@vue/shared");
function createBuild(options, config) {
    (0, uni_cli_shared_1.initEasycomsOnce)(options.inputDir, {
        dirs: [(0, uni_cli_shared_1.resolveComponentsLibPath)()],
        platform: process.env.UNI_PLATFORM,
        isX: process.env.UNI_APP_X === 'true',
    });
    const rollupOutputOption = config.build?.rollupOptions?.output;
    const sourcemap = process.env.SOURCEMAP === 'true' ? 'hidden' : config.build?.sourcemap;
    return {
        sourcemap,
        cssTarget: uni_cli_shared_1.cssTarget,
        chunkSizeWarningLimit: 100000000,
        minify: config.build && (0, shared_1.hasOwn)(config.build, 'minify')
            ? config.build.minify
            : process.env.NODE_ENV === 'production'
                ? 'terser'
                : false,
        rollupOptions: {
            onwarn(warning, warn) {
                if (warning.code === 'UNUSED_EXTERNAL_IMPORT') {
                    const { message } = warning;
                    // ignore
                    if (message.includes('"vue"') ||
                        message.includes('"resolveComponent"') ||
                        message.includes('"@dcloudio/uni-h5"')) {
                        return;
                    }
                }
                warn(warning);
            },
            output: {
                sourcemapExcludeSources: !(0, shared_1.isArray)(rollupOutputOption) &&
                    rollupOutputOption?.sourcemapExcludeSources === false
                    ? false
                    : process.env.SOURCEMAP === 'true',
            },
        },
    };
}
exports.createBuild = createBuild;
