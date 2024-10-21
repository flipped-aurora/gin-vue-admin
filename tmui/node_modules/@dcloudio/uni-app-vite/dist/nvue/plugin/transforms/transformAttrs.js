"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformAttrs = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const shared_1 = require("@vue/shared");
/**
 * 将内置组件属性调整为驼峰
 * @param node
 * @param context
 * @returns
 */
const transformAttrs = (node, context) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    if (!(0, uni_shared_1.isAppNVueNativeTag)(node.tag)) {
        return;
    }
    node.props.forEach((prop) => {
        if ((0, uni_cli_shared_1.isDirectiveNode)(prop)) {
            const { arg } = prop;
            if (arg && (0, uni_cli_shared_1.isSimpleExpressionNode)(arg)) {
                arg.content = normalizePropName(arg.content);
            }
        }
        else {
            prop.name = normalizePropName(prop.name);
        }
    });
};
exports.transformAttrs = transformAttrs;
function normalizePropName(name) {
    if (name.startsWith('data-')) {
        return name;
    }
    return (0, shared_1.camelize)(name);
}
