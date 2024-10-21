"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformHtml = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformHtml = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    const dir = (0, compiler_core_1.findDir)(node, 'html');
    if (!dir) {
        return;
    }
    // remove v-html
    node.props.splice(node.props.indexOf(dir), 1);
    if (node.tagType !== compiler_core_1.ElementTypes.ELEMENT) {
        return;
    }
    node.isSelfClosing = false;
    node.children = [createRichText(node, dir)];
};
exports.transformHtml = transformHtml;
function createRichText(node, dir) {
    return {
        tag: 'rich-text',
        type: compiler_core_1.NodeTypes.ELEMENT,
        tagType: compiler_core_1.ElementTypes.ELEMENT,
        props: [(0, uni_cli_shared_1.createBindDirectiveNode)('nodes', dir.exp || '')],
        isSelfClosing: true,
        children: [],
        codegenNode: undefined,
        ns: node.ns,
        loc: node.loc,
    };
}
