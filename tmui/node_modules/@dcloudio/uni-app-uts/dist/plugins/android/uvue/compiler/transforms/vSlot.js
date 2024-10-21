"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildSlots = exports.trackVForSlotScopes = exports.trackSlotScopes = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const errors_1 = require("../errors");
const vFor_1 = require("./vFor");
const runtimeHelpers_1 = require("../runtimeHelpers");
const utils_1 = require("../utils");
const defaultFallback = (0, compiler_core_1.createSimpleExpression)(`null`, false);
// A NodeTransform that:
// 1. Tracks scope identifiers for scoped slots so that they don't get prefixed
//    by transformExpression. This is only applied in non-browser builds with
//    { prefixIdentifiers: true }.
// 2. Track v-slot depths so that we know a slot is inside another slot.
//    Note the exit callback is executed before buildSlots() on the same node,
//    so only nested slots see positive numbers.
const trackSlotScopes = (node, context) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT &&
        (node.tagType === compiler_core_1.ElementTypes.COMPONENT ||
            node.tagType === compiler_core_1.ElementTypes.TEMPLATE)) {
        // We are only checking non-empty v-slot here
        // since we only care about slots that introduce scope variables.
        const vSlot = (0, compiler_core_1.findDir)(node, 'slot');
        if (vSlot) {
            const slotProps = vSlot.exp;
            // if (!__BROWSER__ && context.prefixIdentifiers) {
            if (context.prefixIdentifiers) {
                slotProps && context.addIdentifiers(slotProps);
            }
            context.scopes.vSlot++;
            return () => {
                // if (!__BROWSER__ && context.prefixIdentifiers) {
                if (context.prefixIdentifiers) {
                    slotProps && context.removeIdentifiers(slotProps);
                }
                context.scopes.vSlot--;
            };
        }
    }
};
exports.trackSlotScopes = trackSlotScopes;
// A NodeTransform that tracks scope identifiers for scoped slots with v-for.
// This transform is only applied in non-browser builds with { prefixIdentifiers: true }
const trackVForSlotScopes = (node, context) => {
    let vFor;
    if ((0, compiler_core_1.isTemplateNode)(node) &&
        node.props.some(compiler_core_1.isVSlot) &&
        (vFor = (0, compiler_core_1.findDir)(node, 'for'))) {
        const result = (vFor.forParseResult = (0, vFor_1.parseForExpression)(vFor.exp, context));
        if (result) {
            const { value, key, index } = result;
            const { addIdentifiers, removeIdentifiers } = context;
            value && addIdentifiers(value);
            key && addIdentifiers(key);
            index && addIdentifiers(index);
            return () => {
                value && removeIdentifiers(value);
                key && removeIdentifiers(key);
                index && removeIdentifiers(index);
            };
        }
    }
};
exports.trackVForSlotScopes = trackVForSlotScopes;
const buildClientSlotFn = (props, children, loc) => (0, compiler_core_1.createFunctionExpression)(props, children, false /* newline */, true /* isSlot */, children.length ? children[0].loc : loc);
// Instead of being a DirectiveTransform, v-slot processing is called during
// transformElement to build the slots object for a component.
function buildSlots(node, context, buildSlotFn = buildClientSlotFn) {
    context.helper(compiler_core_1.WITH_CTX);
    const { children, loc } = node;
    const slotsProperties = [];
    const dynamicSlots = [];
    // If the slot is inside a v-for or another v-slot, force it to be dynamic
    // since it likely uses a scope variable.
    let hasDynamicSlots = context.scopes.vSlot > 0 || context.scopes.vFor > 0;
    // with `prefixIdentifiers: true`, this can be further optimized to make
    // it dynamic only when the slot actually uses the scope variables.
    // if (!__BROWSER__ && !context.ssr && context.prefixIdentifiers) {
    if (context.prefixIdentifiers) {
        hasDynamicSlots = (0, compiler_core_1.hasScopeRef)(node, context.identifiers);
    }
    // 1. Check for slot with slotProps on component itself.
    //    <Comp v-slot="{ prop }"/>
    const onComponentSlot = (0, compiler_core_1.findDir)(node, 'slot', true);
    if (onComponentSlot) {
        const { arg, exp } = onComponentSlot;
        if (arg && !(0, compiler_core_1.isStaticExp)(arg)) {
            hasDynamicSlots = true;
        }
        slotsProperties.push((0, compiler_core_1.createObjectProperty)(arg || (0, compiler_core_1.createSimpleExpression)('default', true), buildSlotFn(exp, children, loc)));
    }
    // 2. Iterate through children and check for template slots
    //    <template v-slot:foo="{ prop }">
    let hasTemplateSlots = false;
    let hasNamedDefaultSlot = false;
    const implicitDefaultChildren = [];
    const seenSlotNames = new Set();
    let conditionalBranchIndex = 0;
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
            context.onError((0, errors_1.createCompilerError)(37 /* ErrorCodes.X_V_SLOT_MIXED_SLOT_USAGE */, slotDir.loc));
            break;
        }
        hasTemplateSlots = true;
        const { children: slotChildren, loc: slotLoc } = slotElement;
        const { arg: slotName = (0, compiler_core_1.createSimpleExpression)(`default`, true), exp: slotProps, loc: dirLoc, } = slotDir;
        // check if name is dynamic.
        let staticSlotName;
        if ((0, compiler_core_1.isStaticExp)(slotName)) {
            staticSlotName = slotName ? slotName.content : `default`;
        }
        else {
            hasDynamicSlots = true;
        }
        const slotFunction = buildSlotFn(slotProps, slotChildren, slotLoc);
        // check if this slot is conditional (v-if/v-for)
        let vIf;
        let vElse;
        let vFor;
        if ((vIf = (0, compiler_core_1.findDir)(slotElement, 'if'))) {
            hasDynamicSlots = true;
            dynamicSlots.push((0, compiler_core_1.createConditionalExpression)(vIf.exp, buildDynamicSlot(slotName, slotFunction, conditionalBranchIndex++), defaultFallback));
        }
        else if ((vElse = (0, compiler_core_1.findDir)(slotElement, /^else(-if)?$/, true /* allowEmpty */))) {
            // find adjacent v-if
            let j = i;
            let prev;
            while (j--) {
                prev = children[j];
                if (prev.type !== compiler_core_1.NodeTypes.COMMENT) {
                    break;
                }
            }
            if (prev && (0, compiler_core_1.isTemplateNode)(prev) && (0, compiler_core_1.findDir)(prev, 'if')) {
                // remove node
                children.splice(i, 1);
                i--;
                // __TEST__ && assert(dynamicSlots.length > 0)
                // attach this slot to previous conditional
                let conditional = dynamicSlots[dynamicSlots.length - 1];
                while (conditional.alternate.type === compiler_core_1.NodeTypes.JS_CONDITIONAL_EXPRESSION) {
                    conditional = conditional.alternate;
                }
                conditional.alternate = vElse.exp
                    ? (0, compiler_core_1.createConditionalExpression)(vElse.exp, buildDynamicSlot(slotName, slotFunction, conditionalBranchIndex++), defaultFallback)
                    : buildDynamicSlot(slotName, slotFunction, conditionalBranchIndex++);
            }
            else {
                context.onError((0, errors_1.createCompilerError)(30 /* ErrorCodes.X_V_ELSE_NO_ADJACENT_IF */, vElse.loc));
            }
        }
        else if ((vFor = (0, compiler_core_1.findDir)(slotElement, 'for'))) {
            hasDynamicSlots = true;
            const parseResult = vFor.forParseResult ||
                (0, vFor_1.parseForExpression)(vFor.exp, context);
            if (parseResult) {
                // Render the dynamic slots as an array and add it to the createSlot()
                // args. The runtime knows how to handle it appropriately.
                dynamicSlots.push((0, compiler_core_1.createCallExpression)(context.helper(runtimeHelpers_1.RENDER_LIST), [
                    parseResult.source,
                    (0, compiler_core_1.createFunctionExpression)((0, vFor_1.createForLoopParams)(parseResult, (0, compiler_core_1.createSimpleExpression)(`_cached`)), buildDynamicSlot(slotName, slotFunction), true /* force newline */),
                ]));
            }
            else {
                context.onError((0, errors_1.createCompilerError)(32 /* ErrorCodes.X_V_FOR_MALFORMED_EXPRESSION */, vFor.loc));
            }
        }
        else {
            // check duplicate static names
            if (staticSlotName) {
                if (seenSlotNames.has(staticSlotName)) {
                    context.onError((0, errors_1.createCompilerError)(38 /* ErrorCodes.X_V_SLOT_DUPLICATE_SLOT_NAMES */, dirLoc));
                    continue;
                }
                seenSlotNames.add(staticSlotName);
                if (staticSlotName === 'default') {
                    hasNamedDefaultSlot = true;
                }
            }
            slotsProperties.push((0, compiler_core_1.createObjectProperty)(slotName, slotFunction));
        }
    }
    if (!onComponentSlot) {
        const buildDefaultSlotProperty = (props, children) => {
            const fn = buildSlotFn(props, children, loc);
            // if (__COMPAT__ && context.compatConfig) {
            //   fn.isNonScopedSlot = true
            // }
            return (0, compiler_core_1.createObjectProperty)(`default`, fn);
        };
        if (!hasTemplateSlots) {
            // implicit default slot (on component)
            slotsProperties.push(buildDefaultSlotProperty(undefined, children));
        }
        else if (implicitDefaultChildren.length &&
            // #3766
            // with whitespace: 'preserve', whitespaces between slots will end up in
            // implicitDefaultChildren. Ignore if all implicit children are whitespaces.
            implicitDefaultChildren.some((node) => isNonWhitespaceContent(node))) {
            // implicit default slot (mixed with named slots)
            if (hasNamedDefaultSlot) {
                context.onError((0, errors_1.createCompilerError)(39 /* ErrorCodes.X_V_SLOT_EXTRANEOUS_DEFAULT_SLOT_CHILDREN */, implicitDefaultChildren[0].loc));
            }
            else {
                slotsProperties.push(buildDefaultSlotProperty(undefined, implicitDefaultChildren));
            }
        }
    }
    const slotFlag = hasDynamicSlots
        ? shared_1.SlotFlags.DYNAMIC
        : hasForwardedSlots(node.children)
            ? shared_1.SlotFlags.FORWARDED
            : shared_1.SlotFlags.STABLE;
    let slots = (0, compiler_core_1.createObjectExpression)(slotsProperties.concat((0, compiler_core_1.createObjectProperty)(`_`, 
    // 2 = compiled but dynamic = can skip normalization, but must run diff
    // 1 = compiled and static = can skip normalization AND diff as optimized
    (0, compiler_core_1.createSimpleExpression)(slotFlag + (utils_1.__DEV__ ? ` /* ${shared_1.slotFlagsText[slotFlag]} */` : ``), false))), loc);
    if (dynamicSlots.length) {
        slots = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_1.CREATE_SLOTS), [
            slots,
            (0, compiler_core_1.createArrayExpression)(dynamicSlots),
        ]);
    }
    return {
        slots,
        hasDynamicSlots,
    };
}
exports.buildSlots = buildSlots;
function buildDynamicSlot(name, fn, index) {
    const props = [
        (0, compiler_core_1.createObjectProperty)(`name`, name),
        (0, compiler_core_1.createObjectProperty)(`fn`, fn),
    ];
    if (index != null) {
        props.push((0, compiler_core_1.createObjectProperty)(`key`, (0, compiler_core_1.createSimpleExpression)(String(index), true)));
    }
    return (0, compiler_core_1.createObjectExpression)(props);
}
function hasForwardedSlots(children) {
    for (let i = 0; i < children.length; i++) {
        const child = children[i];
        switch (child.type) {
            case compiler_core_1.NodeTypes.ELEMENT:
                if (child.tagType === compiler_core_1.ElementTypes.SLOT ||
                    hasForwardedSlots(child.children)) {
                    return true;
                }
                break;
            case compiler_core_1.NodeTypes.IF:
                if (hasForwardedSlots(child.branches))
                    return true;
                break;
            case compiler_core_1.NodeTypes.IF_BRANCH:
            case compiler_core_1.NodeTypes.FOR:
                if (hasForwardedSlots(child.children))
                    return true;
                break;
            default:
                break;
        }
    }
    return false;
}
function isNonWhitespaceContent(node) {
    if (node.type !== compiler_core_1.NodeTypes.TEXT && node.type !== compiler_core_1.NodeTypes.TEXT_CALL)
        return true;
    return node.type === compiler_core_1.NodeTypes.TEXT
        ? !!node.content.trim()
        : isNonWhitespaceContent(node.content);
}
