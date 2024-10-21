"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformBind = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const compiler_core_2 = require("@vue/compiler-core");
const errors_1 = require("../errors");
// v-bind without arg is handled directly in ./transformElements.ts due to it affecting
// codegen for the entire props object. This transform here is only for v-bind
// *with* args.
const transformBind = (dir, _node, context) => {
    const { exp, modifiers, loc } = dir;
    const arg = dir.arg;
    if (arg.type !== compiler_core_2.NodeTypes.SIMPLE_EXPRESSION) {
        arg.children.unshift(`(`);
        arg.children.push(`) || ""`);
    }
    else if (!arg.isStatic) {
        arg.content = `${arg.content} !== null ? ${arg.content} : ""`;
    }
    // .sync is replaced by v-model:arg
    if (modifiers.includes('camel')) {
        if (arg.type === compiler_core_2.NodeTypes.SIMPLE_EXPRESSION) {
            if (arg.isStatic) {
                arg.content = (0, shared_1.camelize)(arg.content);
            }
            else {
                arg.content = `${context.helperString(compiler_core_1.CAMELIZE)}(${arg.content})`;
            }
        }
        else {
            arg.children.unshift(`${context.helperString(compiler_core_1.CAMELIZE)}(`);
            arg.children.push(`)`);
        }
    }
    if (!false) {
        if (modifiers.includes('prop')) {
            injectPrefix(arg, '.');
        }
        if (modifiers.includes('attr')) {
            injectPrefix(arg, '^');
        }
    }
    if (!exp ||
        (exp.type === compiler_core_2.NodeTypes.SIMPLE_EXPRESSION && !exp.content.trim())) {
        context.onError((0, errors_1.createCompilerError)(34 /* ErrorCodes.X_V_BIND_NO_EXPRESSION */, loc));
        return {
            props: [(0, compiler_core_2.createObjectProperty)(arg, (0, compiler_core_2.createSimpleExpression)('', true, loc))],
        };
    }
    return {
        props: [(0, compiler_core_2.createObjectProperty)(arg, exp)],
    };
};
exports.transformBind = transformBind;
const injectPrefix = (arg, prefix) => {
    if (arg.type === compiler_core_2.NodeTypes.SIMPLE_EXPRESSION) {
        if (arg.isStatic) {
            arg.content = prefix + arg.content;
        }
        else {
            arg.content = `\`${prefix}\${${arg.content}}\``;
        }
    }
    else {
        arg.children.unshift(`'${prefix}' + (`);
        arg.children.push(`)`);
    }
};
