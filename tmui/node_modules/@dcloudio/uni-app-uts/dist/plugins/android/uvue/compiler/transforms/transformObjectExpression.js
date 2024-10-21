"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformObjectExpression = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const utils_1 = require("../utils");
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const transformObjectExpression = (node, context) => {
    // 因为 v-bind without arg 是被 transformElements.ts 直接处理的，没办法在 vBind 中解析处理objectExpression
    // 所以 统一在这里拦截处理
    const needTransformAttributes = ['style', 'id', 'class', 'nodes'];
    return function postTransformObjectExpression() {
        node = context.currentNode;
        if (!(0, uni_cli_shared_1.isElementNode)(node)) {
            return;
        }
        node.props.forEach((p) => {
            if (!(0, uni_cli_shared_1.isDirectiveNode)(p) || !p.exp) {
                return;
            }
            if (p.name === 'bind' && !p.arg) {
                // v-bind="{key: value}"
                if (p.exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
                    if (p.exp.content.startsWith('{') && p.exp.content.endsWith('}')) {
                        p.exp.content = `utsMapOf(${p.exp.content})`;
                    }
                    // v-bind="x"
                }
                else if (p.exp.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
                    const children = p.exp.children;
                    if (children.length > 1 &&
                        (0, shared_1.isString)(children[0]) &&
                        (0, shared_1.isString)(children[children.length - 1]) &&
                        children[0].startsWith('{') &&
                        children[children.length - 1].endsWith('}')) {
                        children[0] = `utsMapOf(${children[0]}`;
                        children[children.length - 1] = `${children[children.length - 1]})`;
                    }
                }
                return;
            }
            if ((p.name === 'bind' &&
                needTransformAttributes.includes(p.arg?.content)) ||
                p.name === 'on') {
                const newExp = (0, utils_1.rewriteObjectExpression)(p.exp, context);
                if (newExp) {
                    p.exp = newExp;
                }
            }
        });
    };
};
exports.transformObjectExpression = transformObjectExpression;
