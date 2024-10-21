"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.processDefineOptions = exports.DEFINE_OPTIONS = void 0;
const utils_1 = require("./utils");
const defineProps_1 = require("./defineProps");
const defineEmits_1 = require("./defineEmits");
const defineExpose_1 = require("./defineExpose");
const defineSlots_1 = require("./defineSlots");
exports.DEFINE_OPTIONS = 'defineOptions';
function processDefineOptions(ctx, node) {
    if (!(0, utils_1.isCallOf)(node, exports.DEFINE_OPTIONS)) {
        return false;
    }
    if (ctx.hasDefineOptionsCall) {
        ctx.error(`duplicate ${exports.DEFINE_OPTIONS}() call`, node);
    }
    if (node.typeParameters) {
        ctx.error(`${exports.DEFINE_OPTIONS}() cannot accept type arguments`, node);
    }
    if (!node.arguments[0])
        return true;
    ctx.hasDefineOptionsCall = true;
    ctx.optionsRuntimeDecl = (0, utils_1.unwrapTSNode)(node.arguments[0]);
    let propsOption = undefined;
    let emitsOption = undefined;
    let exposeOption = undefined;
    let slotsOption = undefined;
    if (ctx.optionsRuntimeDecl.type === 'ObjectExpression') {
        for (const prop of ctx.optionsRuntimeDecl.properties) {
            if ((prop.type === 'ObjectProperty' || prop.type === 'ObjectMethod') &&
                prop.key.type === 'Identifier') {
                if (prop.key.name === 'props')
                    propsOption = prop;
                if (prop.key.name === 'emits')
                    emitsOption = prop;
                if (prop.key.name === 'expose')
                    exposeOption = prop;
                if (prop.key.name === 'slots')
                    slotsOption = prop;
            }
        }
    }
    else {
        ctx.error(`${exports.DEFINE_OPTIONS}() options must be an object expression`, ctx.optionsRuntimeDecl);
    }
    if (ctx.optionsRuntimeDecl.properties.find((p) => p.type === 'SpreadElement')) {
        ctx.error(`options does not support spread properties.`, ctx.optionsRuntimeDecl);
    }
    if (propsOption) {
        ctx.error(`${exports.DEFINE_OPTIONS}() cannot be used to declare props. Use ${defineProps_1.DEFINE_PROPS}() instead.`, propsOption);
    }
    if (emitsOption) {
        ctx.error(`${exports.DEFINE_OPTIONS}() cannot be used to declare emits. Use ${defineEmits_1.DEFINE_EMITS}() instead.`, emitsOption);
    }
    if (exposeOption) {
        ctx.error(`${exports.DEFINE_OPTIONS}() cannot be used to declare expose. Use ${defineExpose_1.DEFINE_EXPOSE}() instead.`, exposeOption);
    }
    if (slotsOption) {
        ctx.error(`${exports.DEFINE_OPTIONS}() cannot be used to declare slots. Use ${defineSlots_1.DEFINE_SLOTS}() instead.`, slotsOption);
    }
    return true;
}
exports.processDefineOptions = processDefineOptions;
