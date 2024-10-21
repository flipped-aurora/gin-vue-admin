"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformOn = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const vOn_1 = require("./vOn");
const runtimeHelpers_1 = require("../runtimeHelpers");
const errors_1 = require("../errors");
const isEventOptionModifier = /*#__PURE__*/ (0, shared_1.makeMap)(`passive,once,capture`);
const isNonKeyModifier = /*#__PURE__*/ (0, shared_1.makeMap)(
// event propagation management
`stop,prevent,self,` +
    // system modifiers + exact
    `ctrl,shift,alt,meta,exact,` +
    // mouse
    `middle`);
// left & right could be mouse or key modifiers based on event type
const maybeKeyModifier = /*#__PURE__*/ (0, shared_1.makeMap)('left,right');
const resolveModifiers = (key, modifiers, context, loc) => {
    const keyModifiers = [];
    const nonKeyModifiers = [];
    const eventOptionModifiers = [];
    for (let i = 0; i < modifiers.length; i++) {
        const modifier = modifiers[i];
        if (isEventOptionModifier(modifier)) {
            // eventOptionModifiers: modifiers for addEventListener() options,
            // e.g. .passive & .capture
            if (modifier !== 'once') {
                onModifierWarn(modifier, context, loc);
            }
            else {
                eventOptionModifiers.push(modifier);
            }
        }
        else {
            // runtimeModifiers: modifiers that needs runtime guards
            if (maybeKeyModifier(modifier)) {
                onModifierWarn(modifier, context, loc);
            }
            else {
                if (isNonKeyModifier(modifier)) {
                    if (modifier !== 'stop') {
                        onModifierWarn(modifier, context, loc);
                    }
                    else {
                        nonKeyModifiers.push(modifier);
                    }
                }
                else {
                    onModifierWarn(modifier, context, loc);
                }
            }
        }
    }
    return {
        keyModifiers,
        nonKeyModifiers,
        eventOptionModifiers,
    };
};
const transformOn = (dir, node, context) => {
    return (0, vOn_1.transformOn)(dir, node, context, (baseResult) => {
        const { modifiers } = dir;
        if (!modifiers.length)
            return baseResult;
        let { key, value: handlerExp } = baseResult.props[0];
        const { nonKeyModifiers, eventOptionModifiers } = resolveModifiers(key, modifiers, context, dir.loc);
        if (nonKeyModifiers.length) {
            handlerExp = (0, compiler_core_1.createCallExpression)(context.helper(runtimeHelpers_1.V_ON_WITH_MODIFIERS), [
                handlerExp,
                JSON.stringify(nonKeyModifiers),
            ]);
        }
        if (eventOptionModifiers.length) {
            const modifierPostfix = eventOptionModifiers.map(shared_1.capitalize).join('');
            key = (0, compiler_core_1.isStaticExp)(key)
                ? (0, compiler_core_1.createSimpleExpression)(`${key.content}${modifierPostfix}`, true)
                : (0, compiler_core_1.createCompoundExpression)([`(`, key, `) + "${modifierPostfix}"`]);
        }
        return {
            props: [(0, compiler_core_1.createObjectProperty)(key, handlerExp)],
        };
    });
};
exports.transformOn = transformOn;
function onModifierWarn(modifier, context, loc) {
    context.onWarn((0, errors_1.createCompilerError)(100, loc, {
        100: '.' + modifier + ' is not supported',
    }));
}
