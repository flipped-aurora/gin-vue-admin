"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createTransformComponentLink = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const utils_1 = require("../utils");
function createTransformComponentLink(name, type = compiler_core_1.NodeTypes.DIRECTIVE) {
    return function transformComponentLink(node, context) {
        if (!(0, utils_1.isUserComponent)(node, context)) {
            return;
        }
        // 新版本的 vue，识别 template 有差异，可能认为是自定义组件
        if (node.tag === 'template') {
            return;
        }
        if (type === compiler_core_1.NodeTypes.DIRECTIVE) {
            node.props.push({
                type: compiler_core_1.NodeTypes.DIRECTIVE,
                name: 'on',
                modifiers: [],
                loc: compiler_core_1.locStub,
                arg: (0, compiler_core_1.createSimpleExpression)(name, true),
                exp: (0, compiler_core_1.createSimpleExpression)('__l', true),
            });
        }
        else {
            node.props.push((0, utils_1.createAttributeNode)(name, '__l'));
        }
    };
}
exports.createTransformComponentLink = createTransformComponentLink;
