"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genRuntimeSlots = exports.processDefineSlots = exports.DEFINE_SLOTS = void 0;
const types_1 = require("@babel/types");
const utils_1 = require("./utils");
exports.DEFINE_SLOTS = 'defineSlots';
function processDefineSlots(ctx, node, declId) {
    if (!(0, utils_1.isCallOf)(node, exports.DEFINE_SLOTS)) {
        return false;
    }
    if (ctx.hasDefineSlotsCall) {
        ctx.error(`duplicate ${exports.DEFINE_SLOTS}() call`, node);
    }
    ctx.hasDefineSlotsCall = true;
    if (node.arguments.length > 0) {
        ctx.error(`${exports.DEFINE_SLOTS}() cannot accept arguments`, node);
    }
    if (node.typeParameters &&
        node.typeParameters.params.length === 1 &&
        (0, types_1.isTSTypeLiteral)(node.typeParameters.params[0])) {
        ctx.slotsRuntimeDecl = node.typeParameters.params[0];
    }
    if (declId) {
        ctx.s.overwrite(ctx.startOffset + node.start, ctx.startOffset + node.end, `${ctx.helper('useSlots')}()`);
    }
    return true;
}
exports.processDefineSlots = processDefineSlots;
function genRuntimeSlots(ctx) {
    if (!ctx.slotsRuntimeDecl) {
        return;
    }
    const slots = [];
    ctx.slotsRuntimeDecl.members.forEach((member) => {
        if ((0, types_1.isTSMethodSignature)(member) &&
            member.parameters.length === 1 &&
            (0, types_1.isIdentifier)(member.key)) {
            const param = member.parameters[0];
            if ((0, types_1.isIdentifier)(param) && param.typeAnnotation) {
                const typeAnn = param.typeAnnotation;
                slots.push(member.key.name +
                    ': ' +
                    ctx.source.slice(ctx.startOffset + typeAnn.start + 1, ctx.startOffset + typeAnn.end));
            }
        }
    });
    if (!slots.length) {
        return;
    }
    return `Object as SlotsType<{${slots.join(';')}}>`;
}
exports.genRuntimeSlots = genRuntimeSlots;
