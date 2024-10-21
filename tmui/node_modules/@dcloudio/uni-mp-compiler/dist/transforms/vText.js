"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformText = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformText = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    const dir = (0, compiler_core_1.findDir)(node, 'text');
    if (!dir) {
        return;
    }
    // remove v-text
    node.props.splice(node.props.indexOf(dir), 1);
    if (node.tagType !== compiler_core_1.ElementTypes.ELEMENT) {
        return;
    }
    node.isSelfClosing = false;
    node.children = [
        {
            type: compiler_core_1.NodeTypes.INTERPOLATION,
            loc: dir.exp.loc,
            content: dir.exp,
        },
    ];
};
exports.transformText = transformText;
