"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformText = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
function isTextNode({ tag }) {
    // TODO 临时解决text节点嵌套的问题
    return tag === 'text' || tag === 'button';
}
function isTextElement(node) {
    return node.type === compiler_core_1.NodeTypes.ELEMENT && node.tag === 'text';
}
function isText(node) {
    const { type } = node;
    return (type === compiler_core_1.NodeTypes.TEXT ||
        type === compiler_core_1.NodeTypes.TEXT_CALL ||
        type === compiler_core_1.NodeTypes.INTERPOLATION ||
        type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION);
}
const transformText = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    if (isTextNode(node)) {
        return;
    }
    const { children } = node;
    if (!children.length) {
        return;
    }
    for (let i = 0; i < children.length; i++) {
        const child = children[i];
        if (isTextElement(child)) {
            parseText(child);
        }
    }
};
exports.transformText = transformText;
/*
  1. 转换 \\n 为 \n
  2. u-text 下仅支持 slot 及 文本节点
*/
function parseText(node) {
    if (node.children.length) {
        let firstTextChild;
        for (let i = 0; i < node.children.length; i++) {
            const child = node.children[i];
            const content = child.content;
            if (isText(child) && typeof content === 'string') {
                if (!firstTextChild) {
                    firstTextChild = child;
                    firstTextChild.content = translateObliqueLine(content);
                }
                else {
                    ;
                    firstTextChild.content += translateObliqueLine(content);
                    node.children.splice(i, 1);
                    i--;
                }
            }
            else if (child.type === 3) {
                node.children.splice(i, 1);
                i--;
            }
            else {
                firstTextChild = null;
            }
        }
    }
}
function translateObliqueLine(content) {
    const strFragments = content.split('\\n');
    return strFragments
        .map((str, index) => {
        if (index === strFragments.length - 1)
            return str;
        str += '\\n';
        if (!(str.split('\\').length % 2)) {
            str = str.replaceAll(/\\n/g, '\n');
        }
        return str.replaceAll(/\\\\/g, '\\');
    })
        .join('');
}
