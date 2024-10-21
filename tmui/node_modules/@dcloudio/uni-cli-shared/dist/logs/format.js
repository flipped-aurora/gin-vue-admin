"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.formatWarnMsg = exports.formatInfoMsg = exports.removeNVueInfoFormatter = exports.formatErrMsg = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
const shared_1 = require("@vue/shared");
const env_1 = require("../hbx/env");
const alias_1 = require("../hbx/alias");
const log_1 = require("../hbx/log");
const errFormatters = [];
const infoFormatters = [];
const warnFormatters = [];
const initErrFormattersOnce = (0, uni_shared_1.once)(() => {
    if ((0, env_1.isInHBuilderX)()) {
        errFormatters.push(alias_1.moduleAliasFormatter);
    }
    errFormatters.push(log_1.errorFormatter);
});
const initInfoFormattersOnce = (0, uni_shared_1.once)(() => {
    if ((0, env_1.runByHBuilderX)()) {
        if (
        // 开发模式下
        process.env.UNI_PLATFORM === 'h5' &&
            process.env.NODE_ENV !== 'production') {
            infoFormatters.push(log_1.h5ServeFormatter);
        }
    }
    infoFormatters.push(log_1.removeInfoFormatter);
});
const initWarnFormattersOnce = (0, uni_shared_1.once)(() => {
    warnFormatters.push(log_1.removeWarnFormatter);
});
function formatErrMsg(msg, options) {
    if (options && (0, env_1.isInHBuilderX)()) {
        options.timestamp = false;
    }
    initErrFormattersOnce();
    const formatter = errFormatters.find(({ test }) => test(msg, options));
    if (formatter) {
        return formatter.format(msg, options);
    }
    return msg;
}
exports.formatErrMsg = formatErrMsg;
const REMOVED_NVUE_MSGS = [
    (msg) => {
        // vite v2.7.10 building for development... (x2)
        return msg.includes('vite v') && msg.includes('building ');
    },
];
exports.removeNVueInfoFormatter = {
    test(msg) {
        return !!REMOVED_NVUE_MSGS.find((m) => (0, shared_1.isString)(m) ? msg.includes(m) : m(msg));
    },
    format() {
        return '';
    },
};
const nvueInfoFormatters = [];
const initNVueInfoFormattersOnce = (0, uni_shared_1.once)(() => {
    nvueInfoFormatters.push(exports.removeNVueInfoFormatter);
});
function formatInfoMsg(msg, options) {
    if (options && (0, env_1.isInHBuilderX)()) {
        options.timestamp = false;
    }
    initInfoFormattersOnce();
    const formatter = infoFormatters.find(({ test }) => test(msg, options));
    if (formatter) {
        return formatter.format(msg, options);
    }
    if (options?.nvue) {
        initNVueInfoFormattersOnce();
        const formatter = nvueInfoFormatters.find(({ test }) => test(msg, options));
        if (formatter) {
            return formatter.format(msg, options);
        }
    }
    return msg;
}
exports.formatInfoMsg = formatInfoMsg;
function formatWarnMsg(msg, options) {
    if (options && (0, env_1.isInHBuilderX)()) {
        options.timestamp = false;
    }
    initWarnFormattersOnce();
    const formatter = warnFormatters.find(({ test }) => test(msg, options));
    if (formatter) {
        return formatter.format(msg, options);
    }
    return msg;
}
exports.formatWarnMsg = formatWarnMsg;
