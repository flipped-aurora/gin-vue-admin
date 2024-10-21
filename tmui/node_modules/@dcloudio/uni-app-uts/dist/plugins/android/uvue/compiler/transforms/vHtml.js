"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformVHtml = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const errors_1 = require("../errors");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const transformVHtml = (node, context) => {
    if (node.tagType !== compiler_core_1.ElementTypes.ELEMENT) {
        return;
    }
    // check whether bind v-html
    if (node.props?.length) {
        ;
        node.props.forEach((prop, index) => {
            if (prop.name === 'html' && prop.loc.source.startsWith('v-html=')) {
                if (!prop.exp ||
                    !prop.exp?.content.trim()) {
                    context.onError((0, errors_1.createDOMCompilerError)(53 /* ErrorCodes.X_V_HTML_NO_EXPRESSION */, prop.loc));
                }
                if (node.children.length) {
                    context.onError((0, errors_1.createDOMCompilerError)(54 /* ErrorCodes.X_V_HTML_WITH_CHILDREN */, prop.loc));
                }
                ;
                node.children = [
                    createRichText(node, prop),
                ];
                node.props.splice(index, 1);
            }
        });
    }
};
exports.transformVHtml = transformVHtml;
function createRichText(node, prop) {
    return {
        tag: 'rich-text',
        type: compiler_core_1.NodeTypes.ELEMENT,
        tagType: compiler_core_1.ElementTypes.ELEMENT,
        props: [(0, uni_cli_shared_1.createBindDirectiveNode)('nodes', prop.exp || '')],
        isSelfClosing: true,
        children: [],
        codegenNode: undefined,
        ns: node.ns,
        loc: node.loc,
    };
}
