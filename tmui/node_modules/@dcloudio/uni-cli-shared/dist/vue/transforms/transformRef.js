"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformRef = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const utils_1 = require("../utils");
function transformRef(node, context) {
    if (!(0, utils_1.isUserComponent)(node, context)) {
        return;
    }
    addVueRef(node, context);
}
exports.transformRef = transformRef;
function addVueRef(node, context) {
    // 仅配置了 ref 属性的，才需要增补 vue-ref
    const refProp = (0, compiler_core_1.findProp)(node, 'ref');
    if (!refProp) {
        return;
    }
    if (refProp.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
        refProp.name = 'u-' + utils_1.VUE_REF;
    }
    else {
        ;
        refProp.arg.content = 'u-' + utils_1.VUE_REF;
    }
    return (0, utils_1.addStaticClass)(node, 
    // ref-in-for
    // ref
    context.inVFor
        ? utils_1.VUE_REF_IN_FOR
        : utils_1.VUE_REF);
}
