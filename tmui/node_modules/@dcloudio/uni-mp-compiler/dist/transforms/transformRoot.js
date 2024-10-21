"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformRoot = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const ast_1 = require("../ast");
const types_1 = require("@babel/types");
const codegen_1 = require("../codegen");
const transformRoot = (node, context) => {
    if (node.type !== compiler_core_1.NodeTypes.ROOT) {
        return;
    }
    if (context.bindingCssVars.length) {
        node.children.forEach((child) => {
            if (child.type !== compiler_core_1.NodeTypes.ELEMENT) {
                return;
            }
            addCssVars(child, context);
        });
    }
};
exports.transformRoot = transformRoot;
const CSS_VARS = '__cssVars()';
function addCssVars(node, context) {
    const styleProp = (0, compiler_core_1.findProp)(node, 'style', true);
    if (!styleProp) {
        node.props.push((0, uni_cli_shared_1.createBindDirectiveNode)('style', CSS_VARS));
    }
    else {
        if (styleProp.exp?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
            let expr = (0, ast_1.parseExpr)(styleProp.exp.content, context);
            if ((0, types_1.isArrayExpression)(expr)) {
                expr.elements.push((0, types_1.identifier)(CSS_VARS));
            }
            else {
                expr = (0, types_1.arrayExpression)([expr, (0, types_1.identifier)(CSS_VARS)]);
            }
            styleProp.exp.content = (0, codegen_1.genBabelExpr)(expr);
        }
    }
}
