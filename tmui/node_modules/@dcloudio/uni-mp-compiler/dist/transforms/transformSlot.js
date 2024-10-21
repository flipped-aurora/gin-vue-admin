"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteSlot = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const runtimeHelpers_1 = require("../runtimeHelpers");
const codegen_1 = require("../codegen");
const transform_1 = require("../transform");
const transformElement_1 = require("./transformElement");
const utils_1 = require("./utils");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const __1 = require("..");
const ast_1 = require("../ast");
function rewriteSlot(node, context) {
    let slotName = `"${uni_shared_1.SLOT_DEFAULT_NAME}"`;
    let hasOtherDir = false;
    const nonNameProps = [];
    const { props } = node;
    // 默认插槽强制设置name
    if (!(0, compiler_core_1.findProp)(node, 'name')) {
        props.unshift((0, uni_cli_shared_1.createAttributeNode)('name', 'default'));
    }
    for (let i = 0; i < props.length; i++) {
        const p = props[i];
        if (p.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
            if (p.value) {
                if (p.name === 'name') {
                    p.value.content = (0, uni_shared_1.dynamicSlotName)(p.value.content);
                    slotName = JSON.stringify(p.value.content);
                }
                else {
                    p.name = (0, shared_1.camelize)(p.name);
                    nonNameProps.push(p);
                }
            }
        }
        else {
            if (p.name !== 'bind') {
                hasOtherDir = true;
            }
            if (p.name === 'bind' && (0, compiler_core_1.isStaticArgOf)(p.arg, 'name')) {
                if (p.exp) {
                    slotName = (0, compiler_core_1.createCompoundExpression)([
                        context.helperString(__1.DYNAMIC_SLOT) + '(',
                        p.exp,
                        ')',
                    ]);
                    let slotKey;
                    const keys = parseVForKeyAlias(context);
                    if (keys.length) {
                        const babelNode = (0, ast_1.parseExpr)(p.exp, context);
                        // 在 v-for 中，判断是否插槽名使用 v-for 的 key 变量
                        if (babelNode && (0, utils_1.isReferencedByIds)(babelNode, keys)) {
                            slotKey = parseScopedSlotKey(context);
                        }
                    }
                    p.exp = (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
                        context.helperString(__1.DYNAMIC_SLOT) + '(',
                        p.exp,
                        slotKey ? `+'-'+` + slotKey : '',
                        ')',
                    ]), context);
                }
            }
            else {
                if (p.name === 'bind' && p.arg && (0, compiler_core_1.isStaticExp)(p.arg)) {
                    p.arg.content = (0, shared_1.camelize)(p.arg.content);
                }
                nonNameProps.push(p);
            }
        }
    }
    if (hasOtherDir) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_SLOT_UNEXPECTED_DIRECTIVE_ON_SLOT_OUTLET, node.loc));
    }
    if (nonNameProps.length > 0) {
        (0, transformElement_1.processProps)(node, context, nonNameProps);
        const properties = [];
        nonNameProps.forEach((prop) => {
            const property = transformProperty(prop, context);
            property && properties.push(property);
        });
        if (properties.length) {
            transformScopedSlotName(node, context);
            const vForIndexAlias = parseVForIndexAlias(context);
            (0, utils_1.rewriteExpression)((0, compiler_core_1.createCompoundExpression)([
                context.helperString(runtimeHelpers_1.RENDER_SLOT) + '(',
                slotName,
                ',',
                `{${properties.join(',')}}`,
                `${vForIndexAlias ? ',' + parseScopedSlotKey(context) : ''}`,
                ')',
            ]), context);
        }
        else {
            // 非作用域默认插槽直接移除命名
            if (slotName === `"${uni_shared_1.SLOT_DEFAULT_NAME}"`) {
                (0, utils_1.removeAttribute)(node, 'name');
            }
        }
    }
    else {
        // 非作用域默认插槽直接移除命名
        if (slotName === `"${uni_shared_1.SLOT_DEFAULT_NAME}"`) {
            (0, utils_1.removeAttribute)(node, 'name');
        }
    }
}
exports.rewriteSlot = rewriteSlot;
function parseVForIndexAlias(context) {
    let { currentScope } = context;
    while (currentScope) {
        if ((0, transform_1.isVForScope)(currentScope) && !(0, transform_1.isScopedSlotVFor)(currentScope)) {
            return currentScope.indexAlias;
        }
        currentScope = currentScope.parent;
    }
}
function transformScopedSlotName(node, context) {
    if (!context.miniProgram.slot.dynamicSlotNames) {
        return;
    }
    const slotKey = parseScopedSlotKey(context);
    if (!slotKey) {
        return;
    }
    const nameProps = (0, compiler_core_1.findProp)(node, 'name');
    if (nameProps && nameProps.type === compiler_core_1.NodeTypes.ATTRIBUTE && nameProps.value) {
        const { props } = node;
        props.splice(props.indexOf(nameProps), 1, createScopedSlotDirectiveNode(nameProps.value.content, slotKey, context));
    }
}
function createScopedSlotDirectiveNode(name, slotKey, context) {
    const dir = (0, uni_cli_shared_1.createBindDirectiveNode)('name', (0, utils_1.rewriteExpression)((0, compiler_core_1.createSimpleExpression)(`"${name}-"+` + slotKey), context)
        .content);
    // 存储原始的 slot 名称
    dir.slotName = name;
    return dir;
}
function parseVForKeyAlias(context) {
    let { currentScope } = context;
    const keys = [];
    while (currentScope) {
        if ((0, transform_1.isVForScope)(currentScope) && !(0, transform_1.isScopedSlotVFor)(currentScope)) {
            keys.push(currentScope.keyAlias);
        }
        currentScope = currentScope.parent;
    }
    return keys;
}
function parseVForIndexes(context) {
    let { currentScope } = context;
    const indexes = [];
    while (currentScope) {
        if ((0, transform_1.isVForScope)(currentScope) && !(0, transform_1.isScopedSlotVFor)(currentScope)) {
            indexes.push(currentScope.indexAlias);
        }
        currentScope = currentScope.parent;
    }
    return indexes;
}
function parseSlotKeyByVForIndexes(indexes) {
    return indexes.reverse().join(`+'-'+`);
}
function parseScopedSlotKey(context) {
    const indexes = parseVForIndexes(context);
    const inFor = !!indexes.length;
    if (inFor) {
        return parseSlotKeyByVForIndexes(indexes);
    }
}
function transformProperty(dir, _) {
    if ((0, uni_cli_shared_1.isDirectiveNode)(dir)) {
        if (!dir.arg || !dir.exp) {
            return;
        }
        const isStaticArg = dir.arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && dir.arg.isStatic;
        if (isStaticArg) {
            return `${dir.arg.content}:${(0, codegen_1.genExpr)(dir.exp)}`;
        }
        return `[${(0, codegen_1.genExpr)(dir.arg)}||'']:${(0, codegen_1.genExpr)(dir.exp)}`;
    }
    if (dir.value) {
        return `${dir.name}:${(0, codegen_1.genExpr)(dir.value)}`;
    }
}
