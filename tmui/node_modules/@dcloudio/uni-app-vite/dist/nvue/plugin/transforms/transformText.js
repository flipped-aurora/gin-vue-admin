"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformText = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
function isTextNode({ tag }) {
    return tag === 'text' || tag === 'u-text' || tag === 'button';
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
        let currentContainer = undefined;
        if (isText(child)) {
            if (!currentContainer) {
                currentContainer = children[i] = createText(node, child);
            }
            for (let j = i + 1; j < children.length; j++) {
                const next = children[j];
                if (isText(next)) {
                    // 合并相邻的文本节点
                    currentContainer.children.push(next);
                    children.splice(j, 1);
                    j--;
                }
                else {
                    currentContainer = undefined;
                    break;
                }
            }
        }
    }
};
exports.transformText = transformText;
/*
  1. 转换 \\n 为 \n
  2. u-text 下只能有一个文本节点（不支持 children），需要移除子组件并合并文本
*/
function parseText(node) {
    if (node.children.length) {
        let firstTextChild;
        for (let i = 0; i < node.children.length; i++) {
            const child = node.children[i];
            if (isText(child) && typeof child.content === 'string') {
                if (!firstTextChild) {
                    firstTextChild = child;
                    firstTextChild.content = firstTextChild.content.replace(/\\n/g, '\n');
                }
                else {
                    ;
                    firstTextChild.content += child.content.replace(/\\n/g, '\n');
                    node.children.splice(i, 1);
                    i--;
                }
            }
            else if (child.type === 1 || child.type === 3) {
                node.children.splice(i, 1);
                i--;
            }
            else {
                firstTextChild = null;
            }
        }
    }
}
function createText(parent, node) {
    return {
        tag: 'u-text',
        type: compiler_core_1.NodeTypes.ELEMENT,
        tagType: compiler_core_1.ElementTypes.ELEMENT,
        props: [],
        isSelfClosing: false,
        children: [node],
        codegenNode: undefined,
        ns: parent.ns,
        loc: node.loc,
    };
}
