"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateCodeFrameColumns = exports.createRollupError = exports.isSsr = exports.isInHybridNVue = exports.withSourcemap = void 0;
const shared_1 = require("@vue/shared");
const code_frame_1 = require("@babel/code-frame");
const utils_1 = require("../plugins/vitejs/utils");
function withSourcemap(config) {
    if (process.env.UNI_COMPILE_TARGET === 'uni_modules') {
        return false;
    }
    if (config.command === 'serve') {
        return true;
    }
    return !!config.build.sourcemap;
}
exports.withSourcemap = withSourcemap;
function isInHybridNVue(config) {
    return config.nvue && process.env.UNI_RENDERER !== 'native';
}
exports.isInHybridNVue = isInHybridNVue;
function isSsr(command, config) {
    if (command === 'serve') {
        return !!(config.server && config.server.middlewareMode);
    }
    if (command === 'build') {
        return !!(config.build && config.build.ssr);
    }
    return false;
}
exports.isSsr = isSsr;
function createRollupError(plugin, id, error, source) {
    const { message, name, stack } = error;
    const rollupError = (0, shared_1.extend)(new Error(message), {
        id,
        plugin,
        name,
        stack,
    });
    if ('code' in error && error.loc) {
        rollupError.loc = {
            file: id,
            line: error.loc.start.line,
            column: error.loc.start.column,
        };
        if (source && source.length > 0) {
            if ('offsetStart' in error && 'offsetEnd' in error) {
                rollupError.frame = (0, code_frame_1.codeFrameColumns)(source, (0, utils_1.offsetToStartAndEnd)(source, error.offsetStart, error.offsetEnd));
            }
            else {
                rollupError.frame = (0, code_frame_1.codeFrameColumns)(source, error.loc);
            }
        }
    }
    if (id) {
        // 指定了id后，不让后续的rollup重写
        Object.defineProperty(rollupError, 'id', {
            get() {
                return id;
            },
            set(_v) { },
        });
    }
    return rollupError;
}
exports.createRollupError = createRollupError;
exports.generateCodeFrameColumns = code_frame_1.codeFrameColumns;
