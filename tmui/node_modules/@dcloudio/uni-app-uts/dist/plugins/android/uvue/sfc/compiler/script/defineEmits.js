"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.genRuntimeEmits = exports.processDefineEmits = exports.DEFINE_EMITS = void 0;
const utils_1 = require("./utils");
const resolveType_1 = require("./resolveType");
exports.DEFINE_EMITS = 'defineEmits';
function processDefineEmits(ctx, node, declId) {
    if (!(0, utils_1.isCallOf)(node, exports.DEFINE_EMITS)) {
        return false;
    }
    if (ctx.hasDefineEmitCall) {
        ctx.error(`duplicate ${exports.DEFINE_EMITS}() call`, node);
    }
    ctx.hasDefineEmitCall = true;
    ctx.emitsRuntimeDecl = node.arguments[0];
    if (node.typeParameters) {
        if (ctx.emitsRuntimeDecl) {
            ctx.error(`${exports.DEFINE_EMITS}() cannot accept both type and non-type arguments ` +
                `at the same time. Use one or the other.`, node);
        }
        ctx.emitsTypeDecl = node.typeParameters.params[0];
    }
    ctx.emitDecl = declId;
    return true;
}
exports.processDefineEmits = processDefineEmits;
function genRuntimeEmits(ctx) {
    let emitsDecl = '';
    if (ctx.emitsRuntimeDecl) {
        emitsDecl = ctx.getString(ctx.emitsRuntimeDecl).trim();
    }
    else if (ctx.emitsTypeDecl) {
        const typeDeclaredEmits = extractRuntimeEmits(ctx);
        emitsDecl = typeDeclaredEmits.size
            ? `[${Array.from(typeDeclaredEmits)
                .map((k) => JSON.stringify(k))
                .join(', ')}]`
            : ``;
    }
    if (ctx.hasDefineModelCall) {
        let modelEmitsDecl = `[${Object.keys(ctx.modelDecls)
            .map((n) => JSON.stringify(`update:${n}`))
            .join(', ')}]`;
        emitsDecl = emitsDecl
            ? `/*#__PURE__*/${ctx.helper('mergeModels')}(${emitsDecl}, ${modelEmitsDecl})`
            : modelEmitsDecl;
    }
    return emitsDecl;
}
exports.genRuntimeEmits = genRuntimeEmits;
function extractRuntimeEmits(ctx) {
    const emits = new Set();
    const node = ctx.emitsTypeDecl;
    if (node.type === 'TSFunctionType') {
        extractEventNames(ctx, node.parameters[0], emits);
        return emits;
    }
    const { props, calls } = (0, resolveType_1.resolveTypeElements)(ctx, node);
    let hasProperty = false;
    for (const key in props) {
        emits.add(key);
        hasProperty = true;
    }
    if (calls) {
        if (hasProperty) {
            ctx.error(`defineEmits() type cannot mixed call signature and property syntax.`, node);
        }
        for (const call of calls) {
            extractEventNames(ctx, call.parameters[0], emits);
        }
    }
    return emits;
}
function extractEventNames(ctx, eventName, emits) {
    if (eventName.type === 'Identifier' &&
        eventName.typeAnnotation &&
        eventName.typeAnnotation.type === 'TSTypeAnnotation') {
        const types = (0, resolveType_1.resolveUnionType)(ctx, eventName.typeAnnotation.typeAnnotation);
        for (const type of types) {
            if (type.type === 'TSLiteralType') {
                if (type.literal.type !== 'UnaryExpression' &&
                    type.literal.type !== 'TemplateLiteral') {
                    emits.add(String(type.literal.value));
                }
            }
        }
    }
}
