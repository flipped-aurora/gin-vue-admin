"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformModel = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformExpression_1 = require("./transformExpression");
const estree_walker_1 = require("estree-walker");
const parser_1 = require("@babel/parser");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const errors_1 = require("../errors");
const runtimeHelpers_1 = require("../runtimeHelpers");
const INPUT_TAGS = ['input', 'textarea'];
const AS = ' as ';
const transformModel = (dir, node, context) => {
    // 组件 v-model 绑定了复杂表达式，且没有手动 as 类型
    if (node.tagType === compiler_core_1.ElementTypes.COMPONENT &&
        dir.exp?.children?.length > 1 &&
        !dir.loc.source.includes(AS)) {
        context.onError((0, errors_1.createCompilerError)(100, dir.loc, {
            100: `When custom components use "v-model" to bind complex expressions, you must specify the type using "as", 详见：https://uniapp.dcloud.net.cn/uni-app-x/component/#v-model-complex-expression`,
        }));
    }
    const { exp, arg } = dir;
    if (!exp) {
        context.onError((0, errors_1.createCompilerError)(41 /* ErrorCodes.X_V_MODEL_NO_EXPRESSION */, dir.loc));
        return createTransformProps();
    }
    let rawExp = exp.loc.source;
    let expType = '';
    if ((0, uni_cli_shared_1.isCompoundExpressionNode)(exp)) {
        if (rawExp.includes(AS)) {
            // 目前简单处理(a as string)
            if (rawExp.startsWith('(') && rawExp.endsWith(')')) {
                rawExp = rawExp.slice(1, -1);
            }
            const parts = rawExp.split(AS);
            rawExp = parts[0].trim();
            expType = parts[1].trim();
            const source = (0, transformExpression_1.stringifyExpression)(exp);
            const ast = (0, parser_1.parseExpression)(source, {
                plugins: context.expressionPlugins,
            });
            let str = '';
            (0, estree_walker_1.walk)(ast, {
                enter(node) {
                    if (node.type === 'TSAsExpression') {
                        str = source.slice(node.expression.start, node.expression.end);
                    }
                },
            });
            exp.children = [(0, compiler_core_1.createSimpleExpression)(str)];
        }
    }
    const expString = exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION ? exp.content : rawExp;
    // im SFC <script setup> inline mode, the exp may have been transformed into
    // _unref(exp)
    const bindingType = context.bindingMetadata[rawExp];
    // check props
    if (bindingType === compiler_core_1.BindingTypes.PROPS ||
        bindingType === compiler_core_1.BindingTypes.PROPS_ALIASED) {
        context.onError((0, errors_1.createCompilerError)(44 /* ErrorCodes.X_V_MODEL_ON_PROPS */, exp.loc));
        return createTransformProps();
    }
    const maybeRef = context.inline &&
        (bindingType === compiler_core_1.BindingTypes.SETUP_LET ||
            bindingType === compiler_core_1.BindingTypes.SETUP_REF ||
            bindingType === compiler_core_1.BindingTypes.SETUP_MAYBE_REF);
    if (!expString.trim() ||
        (!(0, compiler_core_1.isMemberExpression)(expString, context) && !maybeRef)) {
        context.onError((0, errors_1.createCompilerError)(42 /* ErrorCodes.X_V_MODEL_MALFORMED_EXPRESSION */, exp.loc));
        return createTransformProps();
    }
    if (context.prefixIdentifiers &&
        (0, compiler_core_1.isSimpleIdentifier)(expString) &&
        context.identifiers[expString]) {
        context.onError((0, errors_1.createCompilerError)(43 /* ErrorCodes.X_V_MODEL_ON_SCOPE_VARIABLE */, exp.loc));
        return createTransformProps();
    }
    const isInputElement = INPUT_TAGS.includes(node.tag);
    const propName = arg ? arg : (0, compiler_core_1.createSimpleExpression)('modelValue', true);
    let eventName = arg
        ? (0, compiler_core_1.isStaticExp)(arg)
            ? `onUpdate:${(0, shared_1.camelize)(arg.content)}`
            : (0, compiler_core_1.createCompoundExpression)(['"onUpdate:" + ', arg])
        : `onUpdate:modelValue`;
    if (isInputElement && eventName === 'onUpdate:modelValue') {
        eventName = getEventName(dir);
    }
    const eventType = isInputElement ? getEventParamsType(dir) : expType;
    let eventValue = isInputElement ? `$event.detail.value` : `$event`;
    if (withTrim(dir)) {
        eventValue = `${eventValue}.trim()`;
    }
    if (withNumber(dir)) {
        eventValue = `${context.helperString(runtimeHelpers_1.LOOSE_TO_NUMBER)}(${eventValue})`;
    }
    let assignmentExp;
    const eventArg = eventType ? `($event: ${eventType})` : `$event`;
    if (maybeRef) {
        if (bindingType === compiler_core_1.BindingTypes.SETUP_REF) {
            // v-model used on known ref.
            assignmentExp = (0, compiler_core_1.createCompoundExpression)([
                `${eventArg} => {(`,
                (0, compiler_core_1.createSimpleExpression)(rawExp, false, exp.loc),
                `).value = ${eventValue}}`,
            ]);
        }
        else {
            // v-model used on a potentially ref binding in <script setup> inline mode.
            // the assignment needs to check whether the binding is actually a ref.
            // 如果是 const 仅设置值：trySetRefValue(innerValue, `$event`.detail.value)
            // 如果是 let 需要执行赋值动作 innerValue = trySetRefValue(innerValue, `$event`.detail.value)
            assignmentExp = (0, compiler_core_1.createCompoundExpression)([
                `${eventArg} => {${bindingType === compiler_core_1.BindingTypes.SETUP_LET ? `${rawExp} = ` : ''}${context.helperString(runtimeHelpers_1.TRY_SET_REF_VALUE)}(${rawExp}, ${eventValue})}`,
            ]);
        }
    }
    else {
        assignmentExp = (0, compiler_core_1.createCompoundExpression)([
            `${eventArg} => {(`,
            exp,
            `) = ${eventValue}}`,
        ]);
    }
    const props = [
        // modelValue: foo
        (0, compiler_core_1.createObjectProperty)(propName, exp),
        // "onUpdate:modelValue": $event => (foo = $event)
        (0, compiler_core_1.createObjectProperty)(eventName, assignmentExp),
    ];
    // cache v-model handler if applicable (when it doesn't refer any scope vars)
    if (context.prefixIdentifiers &&
        !context.inVOnce &&
        context.cacheHandlers &&
        !(0, compiler_core_1.hasScopeRef)(exp, context.identifiers)) {
        props[1].value = context.cache(props[1].value);
    }
    // modelModifiers: { foo: true, "bar-baz": true }
    if (dir.modifiers.length && node.tagType === compiler_core_1.ElementTypes.COMPONENT) {
        const modifiers = dir.modifiers
            .map((m) => ((0, compiler_core_1.isSimpleIdentifier)(m) ? m : JSON.stringify(m)) + `: true`)
            .join(`, `);
        const modifiersKey = arg
            ? (0, compiler_core_1.isStaticExp)(arg)
                ? `${arg.content}Modifiers`
                : (0, compiler_core_1.createCompoundExpression)([arg, ' + "Modifiers"'])
            : `modelModifiers`;
        props.push((0, compiler_core_1.createObjectProperty)(modifiersKey, (0, compiler_core_1.createSimpleExpression)(`{ ${modifiers} }`, false, dir.loc, compiler_core_1.ConstantTypes.CAN_HOIST)));
    }
    return createTransformProps(props);
};
exports.transformModel = transformModel;
function createTransformProps(props = []) {
    return { props };
}
function getEventName(dir) {
    return withLazy(dir) ? 'onBlur' : 'onInput';
}
function getEventParamsType(dir) {
    return withLazy(dir) ? 'InputBlurEvent' : 'InputEvent';
}
function withLazy(dir) {
    return dir.modifiers.includes('lazy');
}
function withNumber(dir) {
    return dir.modifiers.includes('number');
}
function withTrim(dir) {
    return dir.modifiers.includes('trim');
}
