"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genRuntimeProps = exports.processDefineProps = exports.WITH_DEFAULTS = exports.DEFINE_PROPS = void 0;
const shared_1 = require("@vue/shared");
const compiler_dom_1 = require("@vue/compiler-dom");
const resolveType_1 = require("./resolveType");
const utils_1 = require("./utils");
const defineModel_1 = require("./defineModel");
const analyzeScriptBindings_1 = require("./analyzeScriptBindings");
const definePropsDestructure_1 = require("./definePropsDestructure");
exports.DEFINE_PROPS = 'defineProps';
exports.WITH_DEFAULTS = 'withDefaults';
function processDefineProps(ctx, node, declId) {
    if (!(0, utils_1.isCallOf)(node, exports.DEFINE_PROPS)) {
        return processWithDefaults(ctx, node, declId);
    }
    if (node.arguments.length > 1) {
        ctx.error(`${exports.DEFINE_PROPS}() can only accept one argument.`, node);
    }
    if (node.arguments.length > 0) {
        const arg = node.arguments[0];
        if (arg.type !== 'ArrayExpression' && arg.type !== 'ObjectExpression') {
            ctx.error(`${exports.DEFINE_PROPS}() argument must be an object or array literal.`, arg);
        }
    }
    if (ctx.hasDefinePropsCall) {
        ctx.error(`duplicate ${exports.DEFINE_PROPS}() call`, node);
    }
    ctx.hasDefinePropsCall = true;
    ctx.propsRuntimeDecl = node.arguments[0];
    // register bindings
    if (ctx.propsRuntimeDecl) {
        for (const key of (0, analyzeScriptBindings_1.getObjectOrArrayExpressionKeys)(ctx.propsRuntimeDecl)) {
            if (!(key in ctx.bindingMetadata)) {
                ctx.bindingMetadata[key] = compiler_dom_1.BindingTypes.PROPS;
            }
        }
    }
    // call has type parameters - infer runtime types from it
    if (node.typeParameters) {
        if (ctx.propsRuntimeDecl) {
            ctx.error(`${exports.DEFINE_PROPS}() cannot accept both type and non-type arguments ` +
                `at the same time. Use one or the other.`, node);
        }
        ctx.propsTypeDecl = node.typeParameters.params[0];
    }
    // handle props destructure
    if (declId && declId.type === 'ObjectPattern') {
        (0, definePropsDestructure_1.processPropsDestructure)(ctx, declId);
    }
    ctx.propsCall = node;
    ctx.propsDecl = declId;
    return true;
}
exports.processDefineProps = processDefineProps;
function processWithDefaults(ctx, node, declId) {
    if (!(0, utils_1.isCallOf)(node, exports.WITH_DEFAULTS)) {
        return false;
    }
    if (!processDefineProps(ctx, node.arguments[0], declId)) {
        ctx.error(`${exports.WITH_DEFAULTS}' first argument must be a ${exports.DEFINE_PROPS} call.`, node.arguments[0] || node);
    }
    if (node.arguments.length < 2) {
        ctx.error(`${exports.WITH_DEFAULTS}' second argument is required.`, node.arguments[1] || node);
    }
    if (node.arguments[1].type !== 'ObjectExpression') {
        ctx.error(`${exports.WITH_DEFAULTS}' second argument must be an object literal.`, node.arguments[1] || node);
    }
    if (node.arguments[1].properties.find((p) => p.type === 'SpreadElement')) {
        ctx.error(`${exports.WITH_DEFAULTS} does not support spread properties in the second argument.`, node.arguments[1]);
    }
    if (ctx.propsRuntimeDecl) {
        ctx.error(`${exports.WITH_DEFAULTS} can only be used with type-based ` +
            `${exports.DEFINE_PROPS} declaration.`, node);
    }
    if (ctx.propsDestructureDecl) {
        ctx.error(`${exports.WITH_DEFAULTS}() is unnecessary when using destructure with ${exports.DEFINE_PROPS}().\n` +
            `Prefer using destructure default values, e.g. const { foo = 1 } = defineProps(...).`, node.callee);
    }
    ctx.propsRuntimeDefaults = node.arguments[1];
    if (!ctx.propsRuntimeDefaults) {
        ctx.error(`The 2nd argument of ${exports.WITH_DEFAULTS} is required.`, node);
    }
    ctx.propsCall = node;
    return true;
}
function genRuntimeProps(ctx) {
    let propsDecls;
    if (ctx.propsRuntimeDecl) {
        propsDecls = ctx.getString(ctx.propsRuntimeDecl).trim();
        if (ctx.propsDestructureDecl) {
            const defaults = [];
            for (const key in ctx.propsDestructuredBindings) {
                const d = genDestructuredDefaultValue(ctx, key);
                const finalKey = (0, utils_1.getEscapedPropName)(key);
                if (d)
                    defaults.push(`${finalKey}: ${d.valueString}${d.needSkipFactory ? `, __skip_${finalKey}: true` : ``}`);
            }
            if (defaults.length) {
                propsDecls = `/*#__PURE__*/${ctx.helper(`mergeDefaults`)}(${propsDecls}, {\n  ${defaults.join(',\n  ')}\n})`;
            }
        }
    }
    else if (ctx.propsTypeDecl) {
        propsDecls = genRuntimePropsFromTypes(ctx);
    }
    const modelsDecls = (0, defineModel_1.genModelProps)(ctx);
    if (propsDecls && modelsDecls) {
        return `/*#__PURE__*/${ctx.helper('mergeModels')}(${propsDecls}, ${modelsDecls})`;
    }
    else {
        return modelsDecls || propsDecls;
    }
}
exports.genRuntimeProps = genRuntimeProps;
function genRuntimePropsFromTypes(ctx) {
    // this is only called if propsTypeDecl exists
    const props = resolveRuntimePropsFromType(ctx, ctx.propsTypeDecl);
    if (!props.length) {
        return;
    }
    const propStrings = [];
    const hasStaticDefaults = hasStaticWithDefaults(ctx);
    for (const prop of props) {
        propStrings.push(genRuntimePropFromType(ctx, prop, hasStaticDefaults));
        // register bindings
        if (!(prop.key in ctx.bindingMetadata)) {
            ctx.bindingMetadata[prop.key] = compiler_dom_1.BindingTypes.PROPS;
        }
    }
    let propsDecls = `{
    ${propStrings.join(',\n    ')}\n  }`;
    if (ctx.propsRuntimeDefaults && !hasStaticDefaults) {
        propsDecls = `/*#__PURE__*/${ctx.helper('mergeDefaults')}(${propsDecls}, ${ctx.getString(ctx.propsRuntimeDefaults)})`;
    }
    return propsDecls;
}
function parsePropType(ctx, node) {
    if (node.type === 'TSPropertySignature') {
        const typeAnn = node.typeAnnotation;
        if (typeAnn) {
            const tsType = typeAnn?.typeAnnotation;
            if (tsType?.type === 'TSTypeReference') {
                if (tsType.typeName.type === 'Identifier' &&
                    tsType.typeName.name === 'PropType') {
                    return [
                        `Object as ${ctx.source.slice(tsType.start + ctx.startOffset, tsType.end + ctx.startOffset)}`,
                    ];
                }
            }
        }
    }
    return [];
}
function resolveRuntimePropsFromType(ctx, node) {
    const props = [];
    const elements = (0, resolveType_1.resolveTypeElements)(ctx, node);
    for (const key in elements.props) {
        const e = elements.props[key];
        let type = parsePropType(ctx, e);
        let skipCheck = false;
        if (type.length) {
            skipCheck = true;
        }
        else {
            type = (0, resolveType_1.inferRuntimeType)(ctx, e, 'defineProps');
            // skip check for result containing unknown types
            if (type.includes(utils_1.UNKNOWN_TYPE)) {
                if (type.includes('Boolean') || type.includes('Function')) {
                    type = type.filter((t) => t !== utils_1.UNKNOWN_TYPE);
                    skipCheck = true;
                }
                else {
                    type = ['Object'];
                }
            }
        }
        let hasNull = false;
        type = type || [];
        if ((0, shared_1.isArray)(type)) {
            if (type.find((t) => t === 'null')) {
                hasNull = true;
                type = type.filter((t) => t !== 'null');
            }
        }
        props.push({
            key,
            required: !e.optional && !hasNull,
            type,
            skipCheck,
        });
    }
    return props;
}
function genRuntimePropFromType(ctx, { key, required, type, skipCheck }, hasStaticDefaults) {
    let defaultString;
    const destructured = genDestructuredDefaultValue(ctx, key, type);
    if (destructured) {
        defaultString = `default: ${destructured.valueString}${destructured.needSkipFactory ? `, skipFactory: true` : ``}`;
    }
    else if (hasStaticDefaults) {
        const prop = ctx.propsRuntimeDefaults.properties.find((node) => {
            if (node.type === 'SpreadElement')
                return false;
            return (0, utils_1.resolveObjectKey)(node.key, node.computed) === key;
        });
        if (prop) {
            if (prop.type === 'ObjectProperty') {
                // prop has corresponding static default value
                defaultString = `default: ${ctx.getString(prop.value)}`;
            }
            else {
                defaultString = `${prop.async ? 'async ' : ''}${prop.kind !== 'method' ? `${prop.kind} ` : ''}default() ${ctx.getString(prop.body)}`;
            }
        }
    }
    const finalKey = (0, utils_1.getEscapedPropName)(key);
    return `${finalKey}: { ${(0, utils_1.concatStrings)([
        `type: ${(0, utils_1.toRuntimeTypeString)(type)}`,
        `required: ${required}`,
        skipCheck && 'skipCheck: true',
        defaultString,
    ])} }`;
    // if (!ctx.options.isProd) {
    //   return `${finalKey}: { ${concatStrings([
    //     `type: ${toRuntimeTypeString(type)}`,
    //     `required: ${required}`,
    //     skipCheck && 'skipCheck: true',
    //     defaultString,
    //   ])} }`
    // } else if (
    //   type.some(
    //     (el) =>
    //       el === 'Boolean' ||
    //       ((!hasStaticDefaults || defaultString) && el === 'Function')
    //   )
    // ) {
    //   // #4783 for boolean, should keep the type
    //   // #7111 for function, if default value exists or it's not static, should keep it
    //   // in production
    //   return `${finalKey}: { ${concatStrings([
    //     `type: ${toRuntimeTypeString(type)}`,
    //     defaultString,
    //   ])} }`
    // } else {
    //   // production: checks are useless
    //   return `${finalKey}: ${defaultString ? `{ ${defaultString} }` : `{}`}`
    // }
}
/**
 * check defaults. If the default object is an object literal with only
 * static properties, we can directly generate more optimized default
 * declarations. Otherwise we will have to fallback to runtime merging.
 */
function hasStaticWithDefaults(ctx) {
    return !!(ctx.propsRuntimeDefaults &&
        ctx.propsRuntimeDefaults.type === 'ObjectExpression' &&
        ctx.propsRuntimeDefaults.properties.every((node) => node.type !== 'SpreadElement' &&
            (!node.computed || node.key.type.endsWith('Literal'))));
}
function genDestructuredDefaultValue(ctx, key, inferredType) {
    const destructured = ctx.propsDestructuredBindings[key];
    const defaultVal = destructured && destructured.default;
    if (defaultVal) {
        const value = ctx.getString(defaultVal);
        const unwrapped = (0, utils_1.unwrapTSNode)(defaultVal);
        if (inferredType && inferredType.length && !inferredType.includes('null')) {
            const valueType = inferValueType(unwrapped);
            if (valueType && !inferredType.includes(valueType)) {
                ctx.error(`Default value of prop "${key}" does not match declared type.`, unwrapped);
            }
        }
        // If the default value is a function or is an identifier referencing
        // external value, skip factory wrap. This is needed when using
        // destructure w/ runtime declaration since we cannot safely infer
        // whether the expected runtime prop type is `Function`.
        const needSkipFactory = !inferredType &&
            ((0, compiler_dom_1.isFunctionType)(unwrapped) || unwrapped.type === 'Identifier');
        const needFactoryWrap = !needSkipFactory &&
            !(0, utils_1.isLiteralNode)(unwrapped) &&
            !inferredType?.includes('Function');
        return {
            valueString: needFactoryWrap ? `() => (${value})` : value,
            needSkipFactory,
        };
    }
}
// non-comprehensive, best-effort type infernece for a runtime value
// this is used to catch default value / type declaration mismatches
// when using props destructure.
function inferValueType(node) {
    switch (node.type) {
        case 'StringLiteral':
            return 'String';
        case 'NumericLiteral':
            return 'Number';
        case 'BooleanLiteral':
            return 'Boolean';
        case 'ObjectExpression':
            return 'Object';
        case 'ArrayExpression':
            return 'Array';
        case 'FunctionExpression':
        case 'ArrowFunctionExpression':
            return 'Function';
    }
}
