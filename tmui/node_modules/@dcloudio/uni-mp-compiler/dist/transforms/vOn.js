"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.wrapperVOn = exports.transformOn = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const __1 = require("..");
const runtimeHelpers_1 = require("../runtimeHelpers");
const transformExpression_1 = require("./transformExpression");
const vFor_1 = require("./vFor");
const fnExpRE = /^\s*([\w$_]+|(async\s*)?\([^)]*?\))\s*=>|^\s*(async\s+)?function(?:\s+[\w$]+)?\s*\(/;
const transformOn = (dir, node, _context, augmentor) => {
    const context = _context;
    const { loc, modifiers, arg } = dir;
    if (!dir.exp && !modifiers.length) {
        context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_ON_NO_EXPRESSION, loc));
    }
    let eventName;
    if (arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
        if (arg.isStatic) {
            const rawName = arg.content;
            // for all event listeners, auto convert it to camelCase. See issue #2249
            eventName = (0, compiler_core_1.createSimpleExpression)((0, shared_1.toHandlerKey)((0, shared_1.camelize)(rawName)), true, arg.loc);
        }
        else {
            // #2388
            eventName = (0, compiler_core_1.createCompoundExpression)([
                // `${context.helperString(TO_HANDLER_KEY)}(`,
                arg,
                // `)`,
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
                    !(0, compiler_core_1.hasScopeRef)(exp, context.identifiers) &&
                    // wxs event
                    !isFilterExpr(exp, context);
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
                    ? context.isTS
                        ? `($event: any)`
                        : `$event`
                    : `${context.isTS ? `\n//@ts-ignore\n` : ``}(...args)`} => ${hasMultipleStatements ? `{` : `(`}`,
                exp,
                hasMultipleStatements ? `}` : `)`,
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
    // TODO
    if (shouldCache) {
        // cache handlers so that it's always the same handler being passed down.
        // this avoids unnecessary re-renders when users use inline handlers on
        // components.
        // ret.props[0].value = wrapper(
        //   context.cache(ret.props[0].value) as ExpressionNode,
        //   context
        // )
        ret.props[0].value = wrapperVOn(ret.props[0].value, node, context);
    }
    else {
        ret.props[0].value = wrapperVOn(ret.props[0].value, node, context);
    }
    // mark the key as handler for props normalization check
    ret.props.forEach((p) => (p.key.isHandlerKey = true));
    return ret;
};
exports.transformOn = transformOn;
function isFilterExpr(value, context) {
    if (context.filters.length && value.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
        const firstChild = value.children[0];
        if (firstChild.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
            context.filters.includes(firstChild.content)) {
            return true;
        }
    }
    return false;
}
function wrapperVOn(value, node, context) {
    if ((0, transformExpression_1.isBuiltInIdentifier)(value)) {
        return value;
    }
    // wxs event
    if (isFilterExpr(value, context)) {
        return value;
    }
    const keys = [];
    if (context.miniProgram.event?.key && context.inVFor) {
        let keyProp = (0, compiler_core_1.findProp)(node, 'key');
        if (!keyProp) {
            const vForScope = (0, vFor_1.parseVForScope)(context.currentScope);
            if (vForScope) {
                keyProp = (0, compiler_core_1.findProp)(vForScope.node, 'key');
            }
        }
        // 对 for 中的所有事件增加 key 标记，避免微信小程序不更新事件对象
        if (keyProp && (0, uni_cli_shared_1.isDirectiveNode)(keyProp) && keyProp.exp) {
            const keyCode = (0, __1.genExpr)(keyProp.exp);
            if (keyCode) {
                keys.push(',');
                keys.push((0, __1.genExpr)(keyProp.exp));
            }
        }
    }
    return (0, compiler_core_1.createCompoundExpression)([
        `${context.helperString(runtimeHelpers_1.V_ON)}(`,
        value,
        ...keys,
        `)`,
    ]);
}
exports.wrapperVOn = wrapperVOn;
