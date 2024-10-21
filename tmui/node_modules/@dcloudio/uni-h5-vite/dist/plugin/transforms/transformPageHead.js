"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformPageHead = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformPageHead = (node, context) => {
    if ((0, uni_cli_shared_1.checkElementNodeTag)(node, 'page-meta')) {
        const headNode = node.children.find((child) => (0, uni_cli_shared_1.checkElementNodeTag)(child, 'head'));
        if (headNode) {
            headNode.tag = 'page-meta-head';
            headNode.tagType = compiler_core_1.ElementTypes.COMPONENT;
        }
        return;
    }
    if ((0, uni_cli_shared_1.checkElementNodeTag)(node, 'head') &&
        (0, uni_cli_shared_1.checkElementNodeTag)(context.parent, 'page-meta')) {
        ;
        node.tag = 'page-meta-head';
        node.tagType = compiler_core_1.ElementTypes.COMPONENT;
    }
};
exports.transformPageHead = transformPageHead;
