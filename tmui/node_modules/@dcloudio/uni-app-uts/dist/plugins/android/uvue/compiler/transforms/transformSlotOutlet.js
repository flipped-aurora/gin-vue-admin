"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.processSlotOutlet = exports.transformSlotOutlet = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const compiler_core_2 = require("@vue/compiler-core");
const runtimeHelpers_1 = require("../runtimeHelpers");
const shared_1 = require("@vue/shared");
const transformSlotOutlet = (node, context) => {
    if ((0, compiler_core_2.isSlotOutlet)(node)) {
        const { children, loc } = node;
        const { slotName, slotProps } = processSlotOutlet(node, context);
        const slotArgs = [
            context.prefixIdentifiers ? `_ctx.$slots` : `$slots`,
            slotName,
            '{}',
            'undefined',
            'true',
        ];
        let expectedLen = 2;
        if (slotProps) {
            slotArgs[2] = slotProps;
            expectedLen = 3;
        }
        if (children.length) {
            let fn = (0, compiler_core_1.createFunctionExpression)([], children, false, false, loc);
            // @ts-expect-error 补充returnType
            fn.returnType = `any[]`;
            slotArgs[3] = fn;
            expectedLen = 4;
        }
        if (context.scopeId && !context.slotted) {
            expectedLen = 5;
        }
        slotArgs.splice(expectedLen) // remove unused arguments
        ;
        node.codegenNode = (0, compiler_core_1.createCallExpression)(context.helper(runtimeHelpers_1.RENDER_SLOT), slotArgs, loc);
    }
};
exports.transformSlotOutlet = transformSlotOutlet;
function processSlotOutlet(node, context) {
    let slotName = `"default"`;
    let slotProps = undefined;
    const nonNameProps = [];
    for (let i = 0; i < node.props.length; i++) {
        const p = node.props[i];
        if (p.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
            if (p.value) {
                if (p.name === 'name') {
                    slotName = JSON.stringify(p.value.content);
                }
                else {
                    p.name = (0, shared_1.camelize)(p.name);
                    nonNameProps.push(p);
                }
            }
        }
        else {
            if (p.name === 'bind' && (0, compiler_core_2.isStaticArgOf)(p.arg, 'name')) {
                if (p.exp)
                    slotName = p.exp;
            }
            else {
                if (p.name === 'bind' && p.arg && (0, compiler_core_2.isStaticExp)(p.arg)) {
                    p.arg.content = (0, shared_1.camelize)(p.arg.content);
                }
                nonNameProps.push(p);
            }
        }
    }
    if (nonNameProps.length > 0) {
        const { props, directives } = (0, compiler_core_1.buildProps)(node, context, nonNameProps, false, false);
        slotProps = props;
        if (directives.length) {
            context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_SLOT_UNEXPECTED_DIRECTIVE_ON_SLOT_OUTLET, directives[0].loc));
        }
    }
    return {
        slotName,
        slotProps,
    };
}
exports.processSlotOutlet = processSlotOutlet;
