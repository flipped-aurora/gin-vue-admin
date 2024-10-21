"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createVSlotCallExpression = exports.rewriteScopedSlotVForScope = exports.findCurrentPath = exports.findSlotName = exports.rewriteVSlot = exports.transformSlot = void 0;
const types_1 = require("@babel/types");
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const runtimeHelpers_1 = require("../runtimeHelpers");
const ast_1 = require("../ast");
const codegen_1 = require("../codegen");
const transform_1 = require("../transform");
const utils_1 = require("./utils");
const vFor_1 = require("./vFor");
const runtimeHelpers_2 = require("../runtimeHelpers");
const transformSlot = (node, context) => {
    if (!(0, uni_cli_shared_1.isUserComponent)(node, context)) {
        return;
    }
    const { tag, children } = node;
    const slots = new Set();
    const onComponentSlot = (0, compiler_core_1.findDir)(node, 'slot', true);
    const implicitDefaultChildren = [];
    const isMiniProgramComponent = context.isMiniProgramComponent(tag);
    for (let i = 0; i < children.length; i++) {
        const slotElement = children[i];
        let slotDir;
        if (!(0, compiler_core_1.isTemplateNode)(slotElement) ||
            !(slotDir = (0, compiler_core_1.findDir)(slotElement, 'slot', true))) {
            // not a <template v-slot>, skip.
            if (slotElement.type !== compiler_core_1.NodeTypes.COMMENT) {
                implicitDefaultChildren.push(slotElement);
            }
            continue;
        }
        if (onComponentSlot) {
            // already has on-component slot - this is incorrect usage.
            context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_SLOT_MIXED_SLOT_USAGE, slotDir.loc));
            break;
        }
        if (!slotDir.arg) {
            // v-slot => v-slot:default
            slotDir.arg = (0, compiler_core_1.createSimpleExpression)('default', true);
        }
        const slotName = transformTemplateSlotElement(slotDir, slotElement, node, context);
        // 小程序组件默认插槽，直接移除<template #default>节点
        if (isMiniProgramComponent) {
            if (slotName === 'default' && slotElement.children.length === 1) {
                children.splice(i, 1, slotElement.children[0]);
            }
            continue;
        }
        if (slotName) {
            slots.add(slotName);
        }
    }
    if (isMiniProgramComponent) {
        return;
    }
    if (implicitDefaultChildren.length) {
        slots.add(uni_shared_1.SLOT_DEFAULT_NAME);
    }
    if (onComponentSlot) {
        // <unicloud-db v-slot:default="{data, loading, error, options}"/>
        // => <unicloud-db collection=""><template v-slot:default="{data, loading, error, options}"/></unicloud-db>
        slots.add(uni_shared_1.SLOT_DEFAULT_NAME);
        const templateNode = createTemplateNode(onComponentSlot, implicitDefaultChildren);
        transformTemplateSlotElement(onComponentSlot, templateNode, node, context);
        node.children = [templateNode];
    }
    // 不支持 $slots, 则自动补充 props
    if (slots.size && !context.miniProgram.slot.$slots) {
        const slotsArr = [...slots];
        const hasDynamic = slotsArr.find((name) => !(0, shared_1.isString)(name));
        let value;
        if (hasDynamic) {
            const children = [];
            const len = slotsArr.length - 1;
            slotsArr.forEach((name, index) => {
                if ((0, shared_1.isString)(name)) {
                    children.push(`'${(0, uni_shared_1.dynamicSlotName)(name)}'`);
                }
                else {
                    children.push(name);
                }
                if (index < len) {
                    children.push(',');
                }
            });
            value = (0, compiler_core_1.createCompoundExpression)([
                context.helperString(runtimeHelpers_2.DYNAMIC_SLOT) + '([',
                ...children,
                '])',
            ]);
        }
        else {
            value = `[${slotsArr
                .map((name) => `'${(0, uni_shared_1.dynamicSlotName)(name)}'`)
                .join(',')}]`;
        }
        node.props.unshift((0, uni_cli_shared_1.createBindDirectiveNode)(utils_1.ATTR_VUE_SLOTS, value));
    }
};
exports.transformSlot = transformSlot;
function rewriteVSlot(dir, context) {
    dir.arg = (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
        context.helperString(runtimeHelpers_2.DYNAMIC_SLOT) + '(',
        dir.arg,
        ')',
    ]), context);
}
exports.rewriteVSlot = rewriteVSlot;
function transformTemplateSlotElement(slotDir, slotTemplate, slotComponent, context) {
    const slotName = findSlotName(slotDir);
    if (!slotName) {
        return;
    }
    const { exp } = slotDir;
    // non scoped slots
    if (!exp) {
        return slotName;
    }
    // empty
    if (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && !exp.content.trim()) {
        return slotName;
    }
    // 使用vFor来简单处理scoped slot作用域问题
    slotTemplate.children = [
        createVForTemplate(slotTemplate, { name: slotName, value: (0, codegen_1.genExpr)(exp), slotComponent }, context),
    ];
    if (context.miniProgram.slot.dynamicSlotNames) {
        // 已经在 vFor 中补充 slot，故需要移除 slotTemplate 中的
        const index = slotTemplate.props.indexOf(slotDir);
        if (index > -1) {
            slotTemplate.props.splice(index, 1);
        }
    }
    // v-slot="slotProps" => v-slot 避免 transformIdentifier 生成 slotProps 的变量声明
    slotDir.exp = undefined;
    return slotName;
}
function createTemplateNode(slotDir, children) {
    return {
        type: compiler_core_1.NodeTypes.ELEMENT,
        tag: 'template',
        tagType: compiler_core_1.ElementTypes.TEMPLATE,
        loc: compiler_core_1.locStub,
        isSelfClosing: false,
        codegenNode: undefined,
        ns: 0,
        props: [slotDir],
        children,
    };
}
function findSlotName(slotDir) {
    if (!slotDir.arg) {
        return uni_shared_1.SLOT_DEFAULT_NAME;
    }
    if ((0, compiler_core_1.isStaticExp)(slotDir.arg)) {
        return slotDir.arg.content;
    }
    return slotDir.arg;
}
exports.findSlotName = findSlotName;
function findCurrentVForValueAlias(context) {
    let scope = context.currentScope;
    while (scope) {
        if ((0, transform_1.isVForScope)(scope)) {
            return scope.valueAlias;
        }
        scope = scope.parent;
    }
    return '';
}
function createVForTemplate(slotElement, { name, value, slotComponent, }, context) {
    const slotName = 's' + context.scopes.vFor;
    const keyProp = (0, uni_cli_shared_1.createBindDirectiveNode)('key', 'i' + context.scopes.vFor);
    const source = (0, shared_1.isString)(name) ? `'${name}'` : (0, codegen_1.genExpr)(name);
    const vForProp = {
        type: compiler_core_1.NodeTypes.DIRECTIVE,
        name: 'for',
        loc: compiler_core_1.locStub,
        modifiers: [],
        arg: undefined,
        exp: (0, compiler_core_1.createSimpleExpression)(`(${value}, ${slotName}) in ${utils_1.SCOPED_SLOT_IDENTIFIER}(${source}, ${findCurrentVForValueAlias(context) || `''`})`),
    };
    const props = [vForProp, keyProp];
    if (context.miniProgram.slot.dynamicSlotNames) {
        props.push((0, uni_cli_shared_1.createBindDirectiveNode)('slot', slotName));
    }
    return {
        loc: slotElement.loc,
        ns: 0,
        tag: 'template',
        type: compiler_core_1.NodeTypes.ELEMENT,
        tagType: compiler_core_1.ElementTypes.TEMPLATE,
        props,
        isSelfClosing: false,
        codegenNode: undefined,
        children: slotElement.children,
        slotComponent,
    };
}
const slotNameRE = /\('(.*)',/;
/**
 * ('default','') => default
 * @param source
 * @returns
 */
function findCurrentSlotName(source) {
    return (0, types_1.stringLiteral)((0, uni_shared_1.dynamicSlotName)(source.children[1].match(slotNameRE)[1]));
}
function createPathBinaryExpr(scope, computed = true) {
    return (0, types_1.binaryExpression)('+', (0, types_1.binaryExpression)('+', (0, types_1.stringLiteral)(parseVForPath(scope.sourceAlias) + (computed ? '[' : '.')), (0, types_1.identifier)(scope.indexAlias)), (0, types_1.stringLiteral)(computed ? '].' : '.'));
}
function findCurrentPath(id, scope) {
    let parent = scope.parent;
    let binaryExpr = null;
    while (parent) {
        if ((0, transform_1.isVForScope)(parent)) {
            // const computed = !isScopedSlotVFor(parent)
            if (!binaryExpr) {
                binaryExpr = createPathBinaryExpr(parent);
            }
            else {
                binaryExpr = (0, types_1.binaryExpression)('+', createPathBinaryExpr(parent), binaryExpr);
            }
        }
        parent = parent.parent;
    }
    return ((binaryExpr && (0, types_1.binaryExpression)('+', binaryExpr, (0, types_1.stringLiteral)(id))) ||
        (0, types_1.stringLiteral)(id));
}
exports.findCurrentPath = findCurrentPath;
function findCurrentVueIdExpr(node, context) {
    if (!node) {
        return (0, types_1.stringLiteral)('');
    }
    const vueIdProp = (0, compiler_core_1.findProp)(node, utils_1.ATTR_VUE_ID);
    if (vueIdProp.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
        return (0, types_1.stringLiteral)(vueIdProp.value.content);
    }
    return (0, ast_1.parseExpr)((0, codegen_1.genExpr)(vueIdProp.exp), context) || (0, types_1.stringLiteral)('');
}
/**
 * 目前无用
 * @param vForScope
 * @param parentScope
 * @param context
 */
function rewriteScopedSlotVForScope(vForScope, parentScope, context) {
    // 生成一个新的sourceAlias，用于scopedSlots
    const { source, sourceExpr } = vForScope;
    vForScope.sourceAlias = (0, utils_1.rewriteExpressionWithoutProperty)(source, context, sourceExpr, parentScope).content;
}
exports.rewriteScopedSlotVForScope = rewriteScopedSlotVForScope;
function parseVForPath(id) {
    return id.includes('.') ? id.split('.')[1] : id;
}
function createVSlotCallExpression(slotComponent, vForScope, context) {
    const { source, sourceAlias } = vForScope;
    const id = parseVForPath(sourceAlias);
    return (0, types_1.callExpression)((0, types_1.identifier)(context.helperString(runtimeHelpers_1.WITH_SCOPED_SLOT)), [
        (0, vFor_1.createVForArrowFunctionExpression)(vForScope),
        (0, types_1.objectExpression)([
            // 插槽名称，数据更新 path，vueId
            (0, types_1.objectProperty)((0, types_1.identifier)('name'), findCurrentSlotName(source)),
            // 暂不生成 path
            (0, types_1.objectProperty)((0, types_1.identifier)('path'), findCurrentPath(id, vForScope)),
            (0, types_1.objectProperty)((0, types_1.identifier)('vueId'), findCurrentVueIdExpr(slotComponent, context)),
        ]),
    ]);
}
exports.createVSlotCallExpression = createVSlotCallExpression;
