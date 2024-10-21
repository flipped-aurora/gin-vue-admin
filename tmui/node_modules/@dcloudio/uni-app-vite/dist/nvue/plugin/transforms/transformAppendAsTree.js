"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformAppendAsTree = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const unitaryTags = [
    'cell',
    'header',
    'cell-slot',
    'recycle-list',
    'text',
    'u-text',
];
const transformAppendAsTree = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    // append => appendAsTree: true
    const appendProp = (0, compiler_core_1.findProp)(node, 'append');
    if (appendProp) {
        (0, uni_cli_shared_1.renameProp)('appendAsTree', appendProp);
        return;
    }
    if (!unitaryTags.includes(node.tag)) {
        return;
    }
    node.props.push((0, uni_cli_shared_1.createBindDirectiveNode)('appendAsTree', 'true'));
};
exports.transformAppendAsTree = transformAppendAsTree;
