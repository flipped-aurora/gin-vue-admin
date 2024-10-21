"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createTransformTag = void 0;
const ast_1 = require("../../vite/utils/ast");
function createTransformTag(opts) {
    return function transformTag(node, context) {
        if (!(0, ast_1.isElementNode)(node)) {
            return;
        }
        const oldTag = node.tag;
        const newTag = opts[oldTag];
        if (!newTag) {
            return;
        }
        node.tag = newTag;
    };
}
exports.createTransformTag = createTransformTag;
