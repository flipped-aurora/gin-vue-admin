"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformAttr = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const shared_1 = require("@vue/shared");
const transformAttr = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    node.props.forEach((prop) => {
        if ((0, uni_cli_shared_1.isAttributeNode)(prop) && prop.value) {
            switch (prop.name) {
                case 'style':
                    prop.value.content = (0, shared_1.stringifyStyle)((0, shared_1.parseStringStyle)(prop.value.content)).slice(0, -1); // 移除最后一个分号，省点大小吧
                    break;
            }
        }
    });
};
exports.transformAttr = transformAttr;
