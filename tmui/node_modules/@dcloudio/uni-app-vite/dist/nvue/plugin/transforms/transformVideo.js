"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformVideo = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
function isVideo(node) {
    return node.tag === 'video' || node.tag === 'u-video';
}
const transformVideo = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    if (!isVideo(node)) {
        return;
    }
    if (!node.children.length) {
        return;
    }
    const firstChild = node.children[0];
    if ((0, uni_cli_shared_1.isElementNode)(firstChild) && firstChild.tag === 'u-scalable') {
        return;
    }
    node.children = [createScalable(node)];
};
exports.transformVideo = transformVideo;
function createScalable(node) {
    return {
        tag: 'u-scalable',
        type: compiler_core_1.NodeTypes.ELEMENT,
        tagType: compiler_core_1.ElementTypes.ELEMENT,
        props: [
            (0, uni_cli_shared_1.createBindDirectiveNode)('style', (0, compiler_core_1.createSimpleExpression)('{position:"absolute",left:"0",right:"0",top:"0",bottom:"0"}')),
        ],
        isSelfClosing: true,
        children: node.children,
        codegenNode: undefined,
        ns: node.ns,
        loc: node.loc,
    };
}
