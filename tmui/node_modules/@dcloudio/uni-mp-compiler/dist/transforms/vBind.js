"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformBind = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const errors_1 = require("../errors");
const transformBind = (dir, _node, context) => {
    const { exp, modifiers, loc } = dir;
    const arg = dir.arg;
    if (arg.type !== compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
        arg.children.unshift(`(`);
        arg.children.push(`) || ""`);
    }
    else if (!arg.isStatic) {
        arg.content = `${arg.content} || ""`;
    }
    // .sync is replaced by v-model:arg
    if (modifiers.includes('camel')) {
        if (arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
            if (arg.isStatic) {
                arg.content = (0, shared_1.camelize)(arg.content);
            }
            else {
                // arg.content = `${context.helperString(CAMELIZE)}(${arg.content})`
            }
        }
        else {
            // arg.children.unshift(`${context.helperString(CAMELIZE)}(`)
            // arg.children.push(`)`)
        }
    }
    if (modifiers.includes('prop')) {
        context.onWarn((0, errors_1.createMPCompilerError)(4 /* MPErrorCodes.X_V_BIND_MODIFIER_PROP */, loc));
    }
    if (modifiers.includes('attr')) {
        context.onWarn((0, errors_1.createMPCompilerError)(5 /* MPErrorCodes.X_V_BIND_MODIFIER_ATTR */, loc));
    }
    if (!exp ||
        (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && !exp.content.trim())) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_BIND_NO_EXPRESSION, loc));
        return {
            props: [(0, compiler_core_1.createObjectProperty)(arg, (0, compiler_core_1.createSimpleExpression)('', true, loc))],
        };
    }
    return {
        props: [(0, compiler_core_1.createObjectProperty)(arg, exp)],
    };
};
exports.transformBind = transformBind;
