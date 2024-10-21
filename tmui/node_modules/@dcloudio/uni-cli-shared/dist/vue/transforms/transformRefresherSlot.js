"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformRefresherSlot = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const ast_1 = require("../../vite/utils/ast");
/**
 * 将scroll-view、list-view内的<view slot="refresher">转为vue支持的用法，此transform需要再较早时机执行，暂时放在transformTag前。此时node.tag还没有加上v-uni-前缀
 * @param node
 * @param context
 * @returns
 */
function transformRefresherSlot(node) {
    if (!(0, ast_1.isElementNode)(node)) {
        return;
    }
    if (node.tag !== 'scroll-view' && node.tag !== 'list-view') {
        return;
    }
    let refresher = null, refresherIndex = -1;
    for (let i = 0; i < node.children.length; i++) {
        const child = node.children[i];
        if (!(0, ast_1.isElementNode)(child)) {
            continue;
        }
        if (child.props.find((prop) => (0, ast_1.isAttributeNode)(prop) &&
            prop.name === 'slot' &&
            prop.value?.content === 'refresher')) {
            refresher = child;
            refresherIndex = i;
            break;
        }
    }
    if (!refresher) {
        return;
    }
    node.children.splice(refresherIndex, 1, {
        type: compiler_core_1.NodeTypes.ELEMENT,
        tag: 'template',
        tagType: compiler_core_1.ElementTypes.TEMPLATE,
        props: [
            {
                type: compiler_core_1.NodeTypes.DIRECTIVE,
                name: 'slot',
                loc: compiler_core_1.locStub,
                arg: {
                    loc: compiler_core_1.locStub,
                    type: compiler_core_1.NodeTypes.SIMPLE_EXPRESSION,
                    content: 'refresher',
                    isStatic: true,
                    constType: compiler_core_1.ConstantTypes.CAN_STRINGIFY,
                },
            },
        ],
        children: [refresher],
        loc: compiler_core_1.locStub,
    });
}
exports.transformRefresherSlot = transformRefresherSlot;
