"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createTransformEvent = void 0;
const shared_1 = require("@vue/shared");
const ast_1 = require("../../vite/utils/ast");
function createTransformEvent(options) {
    return function transformEvent(node) {
        if (!(0, ast_1.isElementNode)(node)) {
            return;
        }
        node.props.forEach((prop) => {
            const { name, arg } = prop;
            if (name === 'on' && arg && (0, ast_1.isSimpleExpressionNode)(arg)) {
                const eventType = options[arg.content];
                if (eventType) {
                    // e.g tap => click
                    if ((0, shared_1.isFunction)(eventType)) {
                        arg.content = eventType(node, prop);
                    }
                    else {
                        arg.content = eventType;
                    }
                }
            }
        });
    };
}
exports.createTransformEvent = createTransformEvent;
