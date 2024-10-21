"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformModel = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const tags = ['u-input', 'u-textarea'];
const transformModel = (dir, node, context) => {
    const result = (0, compiler_core_1.transformModel)(dir, node, context);
    // 将 u-input,u-textarea 的 onUpdate:modelValue 事件转换为 onInput
    if (tags.includes(node.tag) && result.props.length >= 2) {
        const key = result.props[1].key;
        let value = result.props[1].value;
        if (value.type === compiler_core_1.NodeTypes.JS_CACHE_EXPRESSION) {
            value = value.value;
        }
        if ((0, compiler_core_1.isStaticExp)(key) &&
            key.content === 'onUpdate:modelValue' &&
            value.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
            key.content = 'onInput';
            // 替换 $event 为 $event.detail.value
            value.children = value.children.map((child) => {
                if ((0, shared_1.isString)(child)) {
                    return child.replace(/=\s\$event/g, `= $event.detail.value`);
                }
                return child;
            });
        }
    }
    return result;
};
exports.transformModel = transformModel;
