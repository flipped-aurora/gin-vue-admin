"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformUTSComponent = void 0;
const ast_1 = require("../../vite/utils/ast");
const utsUtils_1 = require("../../utsUtils");
/**
 * 将uts组件保存到自定义组件列表中
 * @param node
 * @param context
 * @returns
 */
const transformUTSComponent = (node, context) => {
    if (!(0, ast_1.isElementNode)(node)) {
        return;
    }
    if ((0, utsUtils_1.matchUTSComponent)(node.tag)) {
        if (!context.root.components.includes(node.tag)) {
            context.components.add(node.tag);
        }
    }
};
exports.transformUTSComponent = transformUTSComponent;
