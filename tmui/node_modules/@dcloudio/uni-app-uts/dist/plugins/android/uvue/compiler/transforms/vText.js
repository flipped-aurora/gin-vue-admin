"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformVText = void 0;
const compiler_core_1 = require("@vue/compiler-core");
// import { createDOMCompilerError, DOMErrorCodes } from '../errors'
const transformVText = (dir, node, context) => {
    const { exp, loc } = dir;
    // if (!exp) {
    //   context.onError(
    //     createDOMCompilerError(DOMErrorCodes.X_V_TEXT_NO_EXPRESSION, loc)
    //   )
    // }
    // if (node.children.length) {
    //   context.onError(
    //     createDOMCompilerError(DOMErrorCodes.X_V_TEXT_WITH_CHILDREN, loc)
    //   )
    //   node.children.length = 0
    // }
    return {
        props: [
            (0, compiler_core_1.createObjectProperty)((0, compiler_core_1.createSimpleExpression)(`value`, true), exp
                ? (0, compiler_core_1.getConstantType)(exp, context) > 0
                    ? exp
                    : (0, compiler_core_1.createCallExpression)(context.helperString(compiler_core_1.TO_DISPLAY_STRING), [exp], loc)
                : (0, compiler_core_1.createSimpleExpression)('', true)),
        ],
    };
};
exports.transformVText = transformVText;
