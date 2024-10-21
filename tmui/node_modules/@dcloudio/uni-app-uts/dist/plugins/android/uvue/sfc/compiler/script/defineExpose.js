"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.processDefineExpose = exports.DEFINE_EXPOSE = void 0;
const utils_1 = require("./utils");
exports.DEFINE_EXPOSE = 'defineExpose';
function processDefineExpose(ctx, node) {
    if ((0, utils_1.isCallOf)(node, exports.DEFINE_EXPOSE)) {
        if (ctx.hasDefineExposeCall) {
            ctx.error(`duplicate ${exports.DEFINE_EXPOSE}() call`, node);
        }
        ctx.hasDefineExposeCall = true;
        return true;
    }
    return false;
}
exports.processDefineExpose = processDefineExpose;
