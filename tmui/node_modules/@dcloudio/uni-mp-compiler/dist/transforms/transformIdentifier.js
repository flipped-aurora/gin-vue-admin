"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformIdentifier = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const utils_1 = require("./utils");
const transformClass_1 = require("./transformClass");
const transformStyle_1 = require("./transformStyle");
const runtimeHelpers_1 = require("../runtimeHelpers");
const transformSlot_1 = require("./transformSlot");
const vSlot_1 = require("./vSlot");
const transformRef_1 = require("./transformRef");
const transformComponent_1 = require("./transformComponent");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const shared_1 = require("@vue/shared");
const transformIdentifier = (node, context) => {
    return function transformIdentifier() {
        if (node.type === compiler_core_1.NodeTypes.INTERPOLATION) {
            const content = node.content;
            let isFilter = false;
            if (content.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
                const firstChild = content.children[0];
                isFilter =
                    !(0, shared_1.isString)(firstChild) &&
                        !(0, shared_1.isSymbol)(firstChild) &&
                        firstChild.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
                        context.filters.includes(firstChild.content);
            }
            if (!isFilter) {
                node.content = (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
                    `${context.helperString(runtimeHelpers_1.TO_DISPLAY_STRING)}(`,
                    content,
                    `)`,
                ]), context);
            }
        }
        else if ((0, compiler_core_1.isSlotOutlet)(node)) {
            (0, transformSlot_1.rewriteSlot)(node, context);
        }
        else if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
            let hasClassBinding = false;
            let hasStyleBinding = false;
            (0, transformRef_1.rewriteRef)(node, context);
            if ((0, uni_cli_shared_1.isUserComponent)(node, context)) {
                (0, transformComponent_1.rewriteBinding)(node, context);
            }
            const { props } = node;
            const virtualHost = !!(context.miniProgram.component?.mergeVirtualHostAttributes &&
                context.rootNode === node);
            for (let i = 0; i < props.length; i++) {
                const dir = props[i];
                if (dir.type === compiler_core_1.NodeTypes.DIRECTIVE) {
                    const arg = dir.arg;
                    if (arg) {
                        // TODO 指令暂不不支持动态参数,v-bind:[arg] v-on:[event]
                        if (!(arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && arg.isStatic)) {
                            // v-slot:[slotName] 支持
                            if (dir.name === 'slot') {
                                (0, vSlot_1.rewriteVSlot)(dir, context);
                            }
                            else {
                                props.splice(i, 1);
                                i--;
                                continue;
                            }
                        }
                    }
                    const exp = dir.exp;
                    if (exp) {
                        if (isBuiltIn(dir)) {
                            // noop
                        }
                        else if ((0, transformClass_1.isClassBinding)(dir)) {
                            hasClassBinding = true;
                            (0, transformClass_1.rewriteClass)(i, dir, props, virtualHost, context);
                        }
                        else if ((0, transformStyle_1.isStyleBinding)(dir)) {
                            hasStyleBinding = true;
                            (0, transformStyle_1.rewriteStyle)(i, dir, props, virtualHost, context);
                        }
                        else if ((0, transformComponent_1.isPropsBinding)(dir)) {
                            (0, transformComponent_1.rewritePropsBinding)(dir, node, context);
                        }
                        else {
                            dir.exp = (0, utils_1.rewriteExpression)(exp, context);
                        }
                    }
                }
            }
            if (virtualHost) {
                if (!hasClassBinding) {
                    hasClassBinding = true;
                    props.push((0, transformClass_1.createVirtualHostClass)(props, context));
                }
                if (!hasStyleBinding) {
                    hasStyleBinding = true;
                    props.push((0, transformStyle_1.createVirtualHostStyle)(props, context));
                }
            }
            if (hasClassBinding) {
                const staticClassIndex = (0, transformClass_1.findStaticClassIndex)(props);
                if (staticClassIndex > -1) {
                    props.splice(staticClassIndex, 1);
                }
            }
            if (hasStyleBinding) {
                const staticStyleIndex = (0, transformStyle_1.findStaticStyleIndex)(props);
                if (staticStyleIndex > -1) {
                    props.splice(staticStyleIndex, 1);
                }
            }
        }
    };
};
exports.transformIdentifier = transformIdentifier;
const builtInProps = [utils_1.ATTR_VUE_SLOTS];
function isBuiltIn({ arg, exp }) {
    return (arg?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
        builtInProps.includes(arg.content) &&
        exp?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION);
}
