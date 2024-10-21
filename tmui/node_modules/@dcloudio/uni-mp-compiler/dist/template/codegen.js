"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genElementProps = exports.genNode = exports.generate = void 0;
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const codegen_1 = require("../codegen");
const vFor_1 = require("../transforms/vFor");
const vIf_1 = require("../transforms/vIf");
const vSlot_1 = require("../transforms/vSlot");
const utils_1 = require("../transforms/utils");
function generate({ children }, { slot, event, scopeId, emitFile, filename, directive, lazyElement, isBuiltInComponent, isMiniProgramComponent, component, }) {
    const context = {
        slot,
        event,
        code: '',
        scopeId,
        directive,
        lazyElement,
        component,
        isBuiltInComponent,
        isMiniProgramComponent,
        push(code) {
            context.code += code;
        },
    };
    children.forEach((node) => {
        genNode(node, context);
    });
    emitFile({ type: 'asset', fileName: filename, source: context.code });
}
exports.generate = generate;
function genNode(node, context) {
    switch (node.type) {
        case compiler_core_1.NodeTypes.IF:
            return node.branches.forEach((node) => {
                genNode(node, context);
            });
        case compiler_core_1.NodeTypes.TEXT:
            return genText(node, context);
        case compiler_core_1.NodeTypes.INTERPOLATION:
            return genExpression(node.content, context);
        case compiler_core_1.NodeTypes.ELEMENT:
            if (node.tagType === compiler_core_1.ElementTypes.SLOT) {
                return genSlot(node, context);
            }
            else if (node.tagType === compiler_core_1.ElementTypes.COMPONENT) {
                return genComponent(node, context);
            }
            else if (node.tagType === compiler_core_1.ElementTypes.TEMPLATE) {
                return genTemplate(node, context);
            }
            else if (isLazyElement(node, context)) {
                return genLazyElement(node, context);
            }
            return genElement(node, context);
    }
}
exports.genNode = genNode;
function genText(node, { push }) {
    push(node.content);
}
function genExpression(node, { push }) {
    push(`{{${(0, codegen_1.genExpr)(node)}}}`);
}
function genVIf(exp, { push, directive }) {
    push(` ${directive}if="{{${exp}}}"`);
}
function genVElseIf(exp, { push, directive }) {
    push(` ${directive}elif="{{${exp}}}"`);
}
function genVElse({ push, directive }) {
    push(` ${directive}else`);
}
function genVFor(node, { push, directive }) {
    const { sourceCode, valueAlias, indexAlias } = node.vFor;
    push(` ${directive}for="${sourceCode}"`);
    if (valueAlias) {
        push(` ${directive}for-item="${valueAlias}"`);
    }
    if (valueAlias === 'index') {
        push(` ${directive}for-index="${indexAlias}"`);
    }
    const keyProp = (0, compiler_core_1.findProp)(node, 'key', true);
    if (keyProp) {
        const key = keyProp.exp.content;
        push(` ${directive}key="${key.includes('.') ? key.split('.')[1] : key}"`);
        node.props.splice(node.props.indexOf(keyProp), 1);
    }
}
function genSlot(node, context) {
    // 移除掉所有非name属性，即移除作用域插槽的绑定指令
    node.props = node.props.filter((prop) => {
        if (prop.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
            return prop.name === 'name';
        }
        else if (prop.arg?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
            return prop.arg.content === 'name';
        }
    });
    if (!node.children.length || context.slot.fallbackContent) {
        // 无后备内容或支持后备内容
        return genElement(node, context);
    }
    const { push } = context;
    const isVIfSlot = (0, vIf_1.isIfElementNode)(node);
    if (isVIfSlot) {
        push(`<block`);
        genVIfCode(node, context);
        push(`>`);
        delete node.vIf;
    }
    const children = node.children.slice();
    node.children.length = 0;
    push(`<block`);
    const nameProp = (0, compiler_core_1.findProp)(node, 'name');
    let name = uni_shared_1.SLOT_DEFAULT_NAME;
    if (nameProp) {
        if ((0, uni_cli_shared_1.isAttributeNode)(nameProp)) {
            if (nameProp.value?.content) {
                name = nameProp.value.content;
            }
        }
        else {
            if (nameProp.slotName) {
                name = nameProp.slotName;
            }
        }
    }
    if (name.includes('-')) {
        genVIf(`$slots['${name}']`, context);
    }
    else {
        genVIf(`$slots.${name}`, context);
    }
    push(`>`);
    genElement(node, context);
    push(`</block>`);
    push(`<block`);
    genVElse(context);
    push(`>`);
    children.forEach((node) => {
        genNode(node, context);
    });
    push(`</block>`);
    if (isVIfSlot) {
        push(`</block>`);
    }
}
function genTemplate(node, context) {
    const slotProp = node.props.find((prop) => prop.type === compiler_core_1.NodeTypes.DIRECTIVE &&
        (prop.name === 'slot' ||
            (prop.name === 'bind' &&
                prop.arg?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
                prop.arg.content === 'slot')));
    // 为 bind 时，通常是作用域插槽生成的 vSlot.ts:197 createBindDirectiveNode('slot',...)
    if (slotProp && (slotProp.name === 'bind' || (0, vSlot_1.findSlotName)(slotProp))) {
        /**
         * 仅百度、字节支持使用 block 作为命名插槽根节点
         * 此处为了统一仅默认替换为view
         * <template v-slot/> => <view slot="">
         */
        node.tag = 'view';
    }
    else {
        // <template/> => <block/>
        node.tag = 'block';
    }
    // @ts-expect-error
    node.tagType = compiler_core_1.ElementTypes.ELEMENT;
    // 仅单个子节点的命名插槽(非作用域)，直接使用子节点作为插槽使用，避免多增加的 view 节点影响 flex 排版
    if (slotProp &&
        node.tag === 'view' &&
        !(0, vFor_1.isForElementNode)(node) &&
        node.children.length === 1) {
        const child = node.children[0];
        if ((0, uni_cli_shared_1.isElementNode)(child) &&
            !(0, vFor_1.isForElementNode)(child) &&
            !(0, compiler_core_1.isSlotOutlet)(child)) {
            child.props.push(slotProp);
            if ((0, vIf_1.isIfElementNode)(node)) {
                ;
                child.vIf = node.vIf;
            }
            return genElement(child, context);
        }
    }
    return genElement(node, context);
}
function genComponent(node, context) {
    if (context.component?.getPropertySync) {
        return genElement(node, context);
    }
    if ((0, vIf_1.isIfElementNode)(node) || (0, vFor_1.isForElementNode)(node)) {
        return genElement(node, context);
    }
    // 小程序原生组件，补充 if(r0)
    if (context.isMiniProgramComponent(node.tag)) {
        ;
        node.vIf = {
            name: 'if',
            condition: 'r0',
        };
        return genElement(node, context);
    }
    const prop = (0, compiler_core_1.findProp)(node, utils_1.ATTR_VUE_PROPS);
    if (!prop) {
        return genElement(node, context);
    }
    ;
    node.vIf = {
        name: 'if',
        condition: prop.exp.content,
    };
    return genElement(node, context);
}
function isLazyElement(node, context) {
    if (!context.lazyElement) {
        return false;
    }
    let lazyProps;
    if ((0, shared_1.isFunction)(context.lazyElement)) {
        const res = context.lazyElement(node, context);
        if (!(0, shared_1.isPlainObject)(res)) {
            return res;
        }
        lazyProps = res[node.tag];
    }
    else {
        lazyProps = context.lazyElement[node.tag];
    }
    if (lazyProps === true) {
        return true;
    }
    if (!lazyProps) {
        return;
    }
    return node.props.some((prop) => prop.type === compiler_core_1.NodeTypes.DIRECTIVE &&
        lazyProps.find((lazyProp) => {
            return (prop.name === lazyProp.name &&
                prop.arg?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
                lazyProp.arg.includes(prop.arg.content));
        }));
}
/**
 * 部分内置组件的部分事件在初始化时会立刻触发，但标准事件需要等首次渲染才能确认事件函数，故增加wx:if="{{r0}}"
 * @param node
 * @param context
 */
function genLazyElement(node, context) {
    const { push } = context;
    if (!(0, vIf_1.isIfElementNode)(node)) {
        push(`<block`);
        // r0 => ready 首次渲染
        genVIf(`r0`, context);
        push(`>`);
        genElement(node, context);
        push(`</block>`);
        return;
    }
    // v-if,v-else-if 无需处理
    if (node.vIf.name !== 'else') {
        return genElement(node, context);
    }
    push(`<block`);
    genVElse(context);
    push(`>`);
    node.vIf.name = 'if';
    node.vIf.condition = 'r0';
    genElement(node, context);
    push(`</block>`);
}
function genVIfCode(node, context) {
    const { name, condition } = node.vIf;
    if (name === 'if') {
        genVIf(condition, context);
    }
    else if (name === 'else-if') {
        genVElseIf(condition, context);
    }
    else if (name === 'else') {
        genVElse(context);
    }
}
function genElement(node, context) {
    const { children, isSelfClosing, props } = node;
    let tag = node.tag;
    // <template slot="left"/> => <block slot="left"/>
    if (tag === 'template') {
        if ((0, compiler_core_1.findProp)(node, 'slot')) {
            tag = 'view';
        }
        else {
            tag = 'block';
        }
    }
    // 无用的 block
    if (tag === 'block' &&
        props.length === 0 &&
        !(0, vIf_1.isIfElementNode)(node) &&
        !(0, vFor_1.isForElementNode)(node)) {
        return children.forEach((node) => {
            genNode(node, context);
        });
    }
    let virtualHost = false;
    if ((0, uni_cli_shared_1.isUserComponent)(node, context)) {
        tag = (0, shared_1.hyphenate)(tag);
        if (context.component?.normalizeName) {
            tag = context.component?.normalizeName(tag);
        }
        if (context.component?.mergeVirtualHostAttributes) {
            virtualHost = true;
        }
    }
    const { push } = context;
    const hasVIf = (0, vIf_1.isIfElementNode)(node);
    const hasVFor = (0, vFor_1.isForElementNode)(node);
    const hasVIfAndVFor = hasVIf && hasVFor;
    // 小程序中 wx:else wx:elif 不支持与 wx:for 同时使用
    // 故 if 需要补充一层 block
    if (hasVIfAndVFor) {
        push(`<block`);
        genVIfCode(node, context);
        push(`>`);
    }
    push(`<${tag}`);
    if (!hasVIfAndVFor && hasVIf) {
        genVIfCode(node, context);
    }
    if (hasVFor) {
        genVFor(node, context);
    }
    if (props.length) {
        genElementProps(node, virtualHost, context);
    }
    if (isSelfClosing) {
        push(`/>`);
    }
    else {
        push(`>`);
        children.forEach((node) => {
            genNode(node, context);
        });
        push(`</${tag}>`);
    }
    if (hasVIfAndVFor) {
        push(`</block>`);
    }
}
function checkVirtualHostProps(name, virtualHost) {
    const names = [name];
    if (virtualHost) {
        const obj = {
            style: utils_1.VIRTUAL_HOST_STYLE,
            class: utils_1.VIRTUAL_HOST_CLASS,
        };
        if (name in obj) {
            // TODO 支付宝平台移除原有属性（支付宝小程序自定义组件外部属性始终无效）
            names.push(obj[name]);
        }
        return names;
    }
    return names;
}
function genElementProps(node, virtualHost, context) {
    node.props.forEach((prop) => {
        if (prop.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
            const { value } = prop;
            if (value) {
                checkVirtualHostProps(prop.name, virtualHost).forEach((name) => {
                    context.push(` ${name}="${value.content}"`);
                });
            }
            else {
                context.push(` ${prop.name}`);
            }
        }
        else {
            const { name } = prop;
            if (name === 'on') {
                genOn(prop, node, context);
            }
            else {
                genDirectiveNode(prop, node, virtualHost, context);
            }
        }
    });
}
exports.genElementProps = genElementProps;
function genOn(prop, node, { push, event, isBuiltInComponent }) {
    const arg = prop.arg.content;
    const exp = prop.exp;
    const modifiers = prop.modifiers;
    const name = (event?.format || uni_cli_shared_1.formatMiniProgramEvent)(arg, {
        isCatch: modifiers.includes('stop') || modifiers.includes('prevent'),
        isCapture: modifiers.includes('capture'),
        isComponent: (0, uni_cli_shared_1.isUserComponent)(node, { isBuiltInComponent }),
    });
    if (exp.isStatic) {
        push(` ${name}="${exp.content}"`);
    }
    else {
        push(` ${name}="{{${exp.content}}}"`);
    }
}
function genDirectiveNode(prop, node, virtualHost, context) {
    const { push, component } = context;
    if (prop.name === 'slot') {
        if (prop.arg) {
            const arg = prop.arg;
            if (arg.isStatic) {
                const slotName = (0, uni_shared_1.dynamicSlotName)(arg.content);
                // 非作用域默认插槽不生成 slot 属性
                if (slotName !== uni_shared_1.SLOT_DEFAULT_NAME) {
                    push(` slot="${slotName}"`);
                }
            }
            else {
                push(` slot="{{${arg.content}}}"`);
            }
        }
    }
    else if (prop.name === 'show') {
        let hiddenPropName = 'hidden';
        if ((0, uni_cli_shared_1.isUserComponent)(node, context) && component && component.vShow) {
            hiddenPropName = component.vShow;
        }
        push(` ${hiddenPropName}="{{!${prop.exp.content}}}"`);
    }
    else if (prop.arg && prop.exp) {
        const arg = prop.arg.content;
        const exp = prop.exp.content;
        checkVirtualHostProps(arg, virtualHost).forEach((arg) => {
            push(` ${arg}="{{${exp}}}"`);
        });
    }
    else {
        if (prop.name !== 'bind') {
            throw new Error(`unknown directive ` + JSON.stringify(prop));
        }
    }
}
