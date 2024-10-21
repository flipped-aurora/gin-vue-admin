"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformTag = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformTag = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node) ||
        (_ && _.isCustomElement && _.isCustomElement(node.tag))) {
        return;
    }
    const newTag = uni_cli_shared_1.HTML_TO_MINI_PROGRAM_TAGS[node.tag];
    if (newTag) {
        node.tag = newTag;
        node.tagType = compiler_core_1.ElementTypes.ELEMENT;
    }
};
exports.transformTag = transformTag;
