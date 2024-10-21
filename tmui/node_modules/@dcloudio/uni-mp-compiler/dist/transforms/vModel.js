"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformModel = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const compiler_dom_1 = require("@vue/compiler-dom");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const runtimeHelpers_1 = require("../runtimeHelpers");
const codegen_1 = require("../codegen");
const vOn_1 = require("./vOn");
const transformModel = (dir, node, _context) => {
    const context = _context;
    const baseResult = (0, compiler_core_1.transformModel)(dir, node, _context);
    // base transform has errors OR component v-model (only need props)
    if (!baseResult.props.length || node.tagType === compiler_core_1.ElementTypes.COMPONENT) {
        return transformComponentVModel(baseResult.props, node, context);
    }
    if (dir.arg) {
        context.onError((0, compiler_dom_1.createDOMCompilerError)(compiler_dom_1.DOMErrorCodes.X_V_MODEL_ARG_ON_ELEMENT, dir.arg.loc));
    }
    function checkDuplicatedValue() {
        const value = (0, compiler_core_1.findProp)(node, 'value');
        if (value) {
            context.onError((0, compiler_dom_1.createDOMCompilerError)(compiler_dom_1.DOMErrorCodes.X_V_MODEL_UNNECESSARY_VALUE, value.loc));
        }
    }
    const { tag } = node;
    if (tag === 'input' || tag === 'textarea') {
        checkDuplicatedValue();
    }
    else {
        context.onError((0, compiler_dom_1.createDOMCompilerError)(compiler_dom_1.DOMErrorCodes.X_V_MODEL_ON_INVALID_ELEMENT, dir.loc));
    }
    if (dir.modifiers.length) {
        const arg = dir.arg;
        const modifiers = dir.modifiers
            .map((m) => ((0, compiler_core_1.isSimpleIdentifier)(m) ? m : JSON.stringify(m)) + `: true`)
            .join(`, `);
        const modifiersKey = arg
            ? (0, compiler_core_1.isStaticExp)(arg)
                ? `${arg.content}Modifiers`
                : (0, compiler_core_1.createCompoundExpression)([arg, ' + "Modifiers"'])
            : `modelModifiers`;
        baseResult.props.push((0, compiler_core_1.createObjectProperty)(modifiersKey, (0, compiler_core_1.createSimpleExpression)(`{ ${modifiers} }`, false, dir.loc, compiler_core_1.ConstantTypes.CAN_HOIST)));
    }
    return transformElementVModel(baseResult.props, node, context);
};
exports.transformModel = transformModel;
function findInputDirectiveNode(props) {
    return props.find((prop) => prop.type === compiler_core_1.NodeTypes.DIRECTIVE &&
        prop.name === 'on' &&
        prop.arg?.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
        prop.arg.content === 'input');
}
function transformElementVModel(props, node, context) {
    const dirs = transformVModel(props, node, context, {
        isComponent: false,
        binding: 'value',
        event: 'input',
        formatEventCode(code) {
            return code.replace(/=\s\$event/g, `= $event.detail.value`);
        },
    });
    if (dirs.length === 2) {
        // 快手小程序的 input v-model 被转换到 data-e-o 中，补充到 data-e-o 中
        const inputExp = findDatasetEventOpts(node);
        if (inputExp) {
            inputExp.children[2] = combineVOn(dirs[1].exp, inputExp.children[2], node, context);
            dirs.length = 1;
        }
        else {
            const inputDir = findInputDirectiveNode(node.props);
            if (inputDir && inputDir.exp) {
                // 合并到已有的 input 事件中
                inputDir.exp = combineVOn(dirs[1].exp, inputDir.exp, node, context);
                dirs.length = 1;
            }
        }
    }
    return { props: dirs };
}
/**
 * {
 *  "type": 7,
 *  "name": "bind",
 *  "loc": {},
 *  "modifiers": [],
 *  "arg": {
 *    "type": 4,
 *    "loc": {},
 *    "content": "data-e-o",
 *    "isStatic": true,
 *    "constType": 3
 *  },
 * "exp": {
 *  "type": 8,
 *  "loc": {},
 *  "children": ["{", {
 *   "type": 8,
 *   "loc": {},
 *   "children": ["'input'", ": ", {
 *    "type": 8,
 *    "loc": {},
 *    "children": ["_o(", {
 *     "type": 4,
 *     "content": "_ctx.input",
 *     "isStatic": false,
 *     "constType": 0,
 *     "loc": {}
 *    }, ")"]
 *   }, ","]
 *  }, "}"]
 * }
 * }
 * @param node
 * @returns
 */
function findDatasetEventOpts(node) {
    const eventOptsProp = (0, compiler_core_1.findProp)(node, uni_cli_shared_1.ATTR_DATASET_EVENT_OPTS, true, false);
    if (!eventOptsProp) {
        return;
    }
    const { exp } = eventOptsProp;
    if (exp?.type !== compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
        return;
    }
    for (let i = 0; i < exp.children.length; i++) {
        const childExp = exp.children[i];
        if ((0, shared_1.isSymbol)(childExp) || (0, shared_1.isString)(childExp)) {
            continue;
        }
        if (childExp.type !== compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
            continue;
        }
        if (childExp.children[0] !== `'input'`) {
            continue;
        }
        const inputExp = childExp.children[2];
        if ((0, shared_1.isSymbol)(inputExp) ||
            (0, shared_1.isString)(inputExp) ||
            inputExp.type !== compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
            continue;
        }
        return childExp;
    }
}
function parseVOn(exp, context) {
    return (0, codegen_1.genExpr)(exp).slice(context.helperString(runtimeHelpers_1.V_ON).length + 1, -1);
}
function combineVOn(exp1, exp2, node, context) {
    return (0, vOn_1.wrapperVOn)((0, compiler_core_1.createCompoundExpression)([
        `[`,
        parseVOn(exp1, context),
        ',',
        parseVOn(exp2, context),
        `]`,
    ]), node, context);
}
function transformComponentVModel(props, node, context) {
    return {
        props: transformVModel(props, node, context, {
            isComponent: true,
            formatEventCode(code) {
                return code;
            },
        }),
    };
}
function transformVModel(props, node, context, { isComponent, binding, event, formatEventCode, }) {
    if (props.length < 2) {
        return [];
    }
    const { key: modelValueArg, value: modelValeExpr } = props[0];
    const { key: onUpdateArg, value: onUpdateExpr } = props[1];
    if (modelValueArg.type !== compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
        return [];
    }
    if (onUpdateArg.type !== compiler_core_1.NodeTypes.SIMPLE_EXPRESSION ||
        !onUpdateArg.content.startsWith('onUpdate:')) {
        return [];
    }
    const vBindModelValue = (0, uni_cli_shared_1.createBindDirectiveNode)(binding || modelValueArg.content, (0, codegen_1.genExpr)(modelValeExpr));
    const modifiers = parseVModelModifiers(props[2]);
    // onUpdateExpr 通常是 ExpressionNode 或者被 cache 的 ExpressionNode
    const vOnValue = (onUpdateExpr.type === compiler_core_1.NodeTypes.JS_CACHE_EXPRESSION
        ? onUpdateExpr.value
        : onUpdateExpr);
    const vOnUpdate = (0, uni_cli_shared_1.createOnDirectiveNode)(event || (0, shared_1.camelize)(onUpdateArg.content.replace('onUpdate:', 'update-')), formatEventCode((0, codegen_1.genExpr)((0, vOn_1.wrapperVOn)(modifiers
        ? wrapperVModelModifiers(vOnValue, modifiers, context, isComponent)
        : vOnValue, node, context))));
    return [vBindModelValue, vOnUpdate];
}
function parseVModelModifiers(property) {
    if (property &&
        (0, uni_cli_shared_1.isSimpleExpressionNode)(property.key) &&
        property.key.content.endsWith('Modifiers') &&
        (0, uni_cli_shared_1.isSimpleExpressionNode)(property.value)) {
        return property.value.content;
    }
}
function wrapperVModelModifiers(exp, modifiers, context, isComponent = false) {
    return (0, compiler_core_1.createCompoundExpression)([
        `${context.helperString(runtimeHelpers_1.WITH_MODEL_MODIFIERS)}(`,
        exp,
        ',',
        modifiers,
        `${isComponent ? `, true` : ``}`,
        `)`,
    ]);
}
