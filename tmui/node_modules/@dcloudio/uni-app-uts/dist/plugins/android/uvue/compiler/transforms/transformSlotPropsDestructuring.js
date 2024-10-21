"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createDestructuringSlotProps = exports.isDestructuringSlotProps = exports.SLOT_PROPS_NAME = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
exports.SLOT_PROPS_NAME = 'slotProps';
function isDestructuringSlotProps(isSlot, params) {
    if (isSlot && params?.children?.length > 2) {
        const firstParam = params.children[0];
        const lastParam = params.children[params.children.length - 1];
        return ((0, shared_1.isString)(firstParam) &&
            firstParam.startsWith('{') &&
            (0, shared_1.isString)(lastParam) &&
            lastParam.endsWith('}'));
    }
    return false;
}
exports.isDestructuringSlotProps = isDestructuringSlotProps;
function createDestructuringSlotProps(params, context) {
    params.children.forEach((child, index) => {
        if (child.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
            context.newline();
            const content = child.content;
            if (isRename(params, index)) {
                const originKey = getOriginKey(params.children[index - 1]);
                context.push(`const ${content} = ${exports.SLOT_PROPS_NAME}["${originKey}"]`);
            }
            else if (hasDefaultValue(params, index)) {
                const defaultValue = getDefaultValue(params.children[index + 1]);
                context.push(`const ${content} = ${exports.SLOT_PROPS_NAME}["${content}"] !== null ? ${exports.SLOT_PROPS_NAME}["${content}"] : ${defaultValue}`);
            }
            else {
                context.push(`const ${content} = ${exports.SLOT_PROPS_NAME}["${content}"]`);
            }
        }
    });
}
exports.createDestructuringSlotProps = createDestructuringSlotProps;
function isRename(params, index) {
    const prevChild = params.children[index - 1];
    return (0, shared_1.isString)(prevChild) && prevChild.trim().endsWith(':');
}
function getOriginKey(prevChild) {
    const originKey = prevChild.substring(0, prevChild.indexOf(':')).trim();
    return originKey.startsWith('{') ? originKey.substring(1).trim() : originKey;
}
function hasDefaultValue(params, index) {
    const nextChild = params.children[index + 1];
    return (0, shared_1.isString)(nextChild) && nextChild.trim().startsWith('=');
}
function getDefaultValue(nextChild) {
    const defaultValue = nextChild.trim().substring(1).trim();
    return defaultValue.endsWith('}')
        ? defaultValue.substring(0, defaultValue.length - 1).trim()
        : defaultValue;
}
