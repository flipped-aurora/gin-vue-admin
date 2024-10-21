"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformPageHead = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const utils_1 = require("../../utils");
const transformPageHead = (node, context) => {
    // 发现是page-meta下的head,直接remove该节点
    if ((0, utils_1.checkElementNodeTag)(node, 'head') &&
        (0, utils_1.checkElementNodeTag)(context.parent, 'page-meta')) {
        ;
        node.tag = 'page-meta-head';
        node.tagType = compiler_core_1.ElementTypes.COMPONENT;
    }
};
exports.transformPageHead = transformPageHead;
