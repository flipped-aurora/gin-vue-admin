"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformWxsProps = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const compiler_core_1 = require("@vue/compiler-core");
const runtimeHelpers_1 = require("./runtimeHelpers");
const transformWxsProps = (node, context) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    node.props.forEach((prop) => {
        if ((0, uni_cli_shared_1.isDirectiveNode)(prop) && prop.arg && (0, uni_cli_shared_1.isSimpleExpressionNode)(prop.arg)) {
            const { content } = prop.arg;
            if (content.startsWith(uni_shared_1.ATTR_CHANGE_PREFIX)) {
                const propName = content.substring(uni_shared_1.ATTR_CHANGE_PREFIX.length);
                const wxsProp = (0, compiler_core_1.findProp)(node, propName, true);
                if (wxsProp && (0, uni_cli_shared_1.isDirectiveNode)(wxsProp) && wxsProp.exp) {
                    wxsProp.exp = (0, compiler_core_1.createCompoundExpression)([
                        context.helperString(runtimeHelpers_1.WXS_PROP),
                        '(',
                        wxsProp.exp,
                        ')',
                    ]);
                }
            }
        }
    });
};
exports.transformWxsProps = transformWxsProps;
