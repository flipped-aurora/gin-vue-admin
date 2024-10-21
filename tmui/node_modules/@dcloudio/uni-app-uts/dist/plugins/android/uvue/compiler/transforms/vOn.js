"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformOn = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const errors_1 = require("../errors");
const transformExpression_1 = require("./transformExpression");
const fnExpRE = /^\s*([\w$_]+|(async\s*)?\([^)]*?\))\s*(:[^=]+)?=>|^\s*(async\s+)?function(?:\s+[\w$]+)?\s*\(/;
// TODO: inline statement $event type complement
const transformOn = (dir, node, context, augmentor) => {
    const { loc, modifiers, arg } = dir;
    if (!dir.exp && !modifiers.length) {
        context.onError((0, errors_1.createCompilerError)(35 /* ErrorCodes.X_V_ON_NO_EXPRESSION */, loc));
    }
    let eventName;
    if (arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
        if (arg.isStatic) {
            let rawName = arg.content;
            // TODO deprecate @vnodeXXX usage
            if (rawName.startsWith('vue:')) {
                rawName = `vnode-${rawName.slice(4)}`;
            }
            const eventString = node.tagType !== compiler_core_1.ElementTypes.ELEMENT ||
                rawName.startsWith('vnode') ||
                !/[A-Z]/.test(rawName)
                ? // for non-element and vnode lifecycle event listeners, auto convert
                    // it to camelCase. See issue #2249
                    (0, shared_1.toHandlerKey)((0, shared_1.camelize)(rawName))
                : // preserve case for plain element listeners that have uppercase
                    // letters, as these may be custom elements' custom events
                    `on:${rawName}`;
            eventName = (0, compiler_core_1.createSimpleExpression)(eventString, true, arg.loc);
        }
        else {
            // #2388
            eventName = (0, compiler_core_1.createCompoundExpression)([
                `${context.helperString(compiler_core_1.TO_HANDLER_KEY)}(`,
                arg,
                `)`,
            ]);
        }
    }
    else {
        // already a compound expression.
        eventName = arg;
        eventName.children.unshift(`${context.helperString(compiler_core_1.TO_HANDLER_KEY)}(`);
        eventName.children.push(`)`);
    }
    // handler processing
    let exp = dir.exp;
    if (exp && !exp.content.trim()) {
        exp = undefined;
    }
    let shouldCache = context.cacheHandlers && !exp && !context.inVOnce;
    if (exp) {
        // @ts-expect-error
        const isMemberExp = (0, compiler_core_1.isMemberExpression)(exp.content, context);
        const isInlineStatement = !(isMemberExp || fnExpRE.test(exp.content));
        const hasMultipleStatements = exp.content.includes(`;`);
        // process the expression since it's been skipped
        if (context.prefixIdentifiers) {
            isInlineStatement && context.addIdentifiers(`$event`);
            exp = dir.exp = (0, transformExpression_1.processExpression)(exp, context, false, hasMultipleStatements);
            isInlineStatement && context.removeIdentifiers(`$event`);
            // with scope analysis, the function is hoistable if it has no reference
            // to scope variables.
            shouldCache =
                context.cacheHandlers &&
                    // unnecessary to cache inside v-once
                    !context.inVOnce &&
                    // runtime constants don't need to be cached
                    // (this is analyzed by compileScript in SFC <script setup>)
                    !(exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && exp.constType > 0) &&
                    // #1541 bail if this is a member exp handler passed to a component -
                    // we need to use the original function to preserve arity,
                    // e.g. <transition> relies on checking cb.length to determine
                    // transition end handling. Inline function is ok since its arity
                    // is preserved even when cached.
                    !(isMemberExp && node.tagType === compiler_core_1.ElementTypes.COMPONENT) &&
                    // bail if the function references closure variables (v-for, v-slot)
                    // it must be passed fresh to avoid stale values.
                    !(0, compiler_core_1.hasScopeRef)(exp, context.identifiers);
            // If the expression is optimizable and is a member expression pointing
            // to a function, turn it into invocation (and wrap in an arrow function
            // below) so that it always accesses the latest value when called - thus
            // avoiding the need to be patched.
            if (shouldCache && isMemberExp) {
                if (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
                    exp.content = `${exp.content} && ${exp.content}(...args)`;
                }
                else {
                    exp.children = [...exp.children, ` && `, ...exp.children, `(...args)`];
                }
            }
        }
        if (isInlineStatement || (shouldCache && isMemberExp)) {
            // wrap inline statement in a function expression
            exp = (0, compiler_core_1.createCompoundExpression)([
                `${isInlineStatement
                    ? `${loc.source.includes('$event') ? '($event: any)' : '()'}`
                    : `${``}(...args)`
                // } => ${hasMultipleStatements ? `{` : `(`}`,
                } => ${`{`}`,
                exp,
                // hasMultipleStatements ? `}` : `)`
                `}`,
            ]);
        }
    }
    let ret = {
        props: [
            (0, compiler_core_1.createObjectProperty)(eventName, exp || (0, compiler_core_1.createSimpleExpression)(`() => {}`, false, loc)),
        ],
    };
    // apply extended compiler augmentor
    if (augmentor) {
        ret = augmentor(ret);
    }
    if (shouldCache) {
        // cache handlers so that it's always the same handler being passed down.
        // this avoids unnecessary re-renders when users use inline handlers on
        // components.
        ret.props[0].value = context.cache(ret.props[0].value);
    }
    // mark the key as handler for props normalization check
    ret.props.forEach((p) => (p.key.isHandlerKey = true));
    return ret;
};
exports.transformOn = transformOn;
