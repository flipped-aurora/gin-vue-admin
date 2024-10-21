"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildDirectiveArgs = exports.buildProps = exports.resolveComponentType = exports.transformElement = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const errors_1 = require("../errors");
const compiler_core_2 = require("@vue/compiler-core");
const compiler_core_3 = require("@vue/compiler-core");
const vSlot_1 = require("./vSlot");
const compiler_core_4 = require("@vue/compiler-core");
const compiler_core_5 = require("@vue/compiler-core");
const compiler_core_6 = require("@vue/compiler-core");
const utils_1 = require("../utils");
// some directive transforms (e.g. v-model) may return a symbol for runtime
// import, which should be used instead of a resolveDirective call.
const directiveImportMap = new WeakMap();
// generate a JavaScript AST for this element's codegen
const transformElement = (node, context) => {
    // perform the work on exit, after all child expressions have been
    // processed and merged.
    return function postTransformElement() {
        node = context.currentNode;
        if (!(node.type === compiler_core_1.NodeTypes.ELEMENT &&
            (node.tagType === compiler_core_1.ElementTypes.ELEMENT ||
                node.tagType === compiler_core_1.ElementTypes.COMPONENT))) {
            return;
        }
        const { tag, props } = node;
        const isComponent = node.tagType === compiler_core_1.ElementTypes.COMPONENT;
        // The goal of the transform is to create a codegenNode implementing the
        // VNodeCall interface.
        let vnodeTag = isComponent
            ? resolveComponentType(node, context)
            : `"${tag}"`;
        const isDynamicComponent = (0, shared_1.isObject)(vnodeTag) && vnodeTag.callee === compiler_core_2.RESOLVE_DYNAMIC_COMPONENT;
        let vnodeProps;
        let vnodeChildren;
        let vnodePatchFlag;
        let patchFlag = 0;
        let vnodeDynamicProps;
        let dynamicPropNames;
        let vnodeDirectives;
        let shouldUseBlock = 
        // dynamic component may resolve to plain elements
        isDynamicComponent ||
            vnodeTag === compiler_core_2.TELEPORT ||
            vnodeTag === compiler_core_2.SUSPENSE ||
            (!isComponent &&
                // <svg> and <foreignObject> must be forced into blocks so that block
                // updates inside get proper isSVG flag at runtime. (#639, #643)
                // This is technically web-specific, but splitting the logic out of core
                // leads to too much unnecessary complexity.
                (tag === 'svg' || tag === 'foreignObject'));
        // props
        if (props.length > 0) {
            const propsBuildResult = buildProps(node, context, undefined, isComponent, isDynamicComponent);
            vnodeProps = propsBuildResult.props;
            patchFlag = propsBuildResult.patchFlag;
            dynamicPropNames = propsBuildResult.dynamicPropNames;
            const directives = propsBuildResult.directives;
            vnodeDirectives =
                directives && directives.length
                    ? (0, compiler_core_1.createArrayExpression)(directives.map((dir) => buildDirectiveArgs(dir, context)))
                    : undefined;
            if (propsBuildResult.shouldUseBlock) {
                shouldUseBlock = true;
            }
        }
        // children
        if (node.children.length > 0) {
            if (vnodeTag === compiler_core_2.KEEP_ALIVE) {
                // Although a built-in component, we compile KeepAlive with raw children
                // instead of slot functions so that it can be used inside Transition
                // or other Transition-wrapping HOCs.
                // To ensure correct updates with block optimizations, we need to:
                // 1. Force keep-alive into a block. This avoids its children being
                //    collected by a parent block.
                shouldUseBlock = true;
                // 2. Force keep-alive to always be updated, since it uses raw children.
                patchFlag |= shared_1.PatchFlags.DYNAMIC_SLOTS;
                if (utils_1.__DEV__ && node.children.length > 1) {
                    context.onError((0, errors_1.createCompilerError)(46 /* ErrorCodes.X_KEEP_ALIVE_INVALID_CHILDREN */, {
                        start: node.children[0].loc.start,
                        end: node.children[node.children.length - 1].loc.end,
                        source: '',
                    }));
                }
            }
            const shouldBuildAsSlots = isComponent &&
                // Teleport is not a real component and has dedicated runtime handling
                vnodeTag !== compiler_core_2.TELEPORT &&
                // explained above.
                vnodeTag !== compiler_core_2.KEEP_ALIVE;
            if (shouldBuildAsSlots) {
                const { slots, hasDynamicSlots } = (0, vSlot_1.buildSlots)(node, context);
                vnodeChildren = slots;
                if (hasDynamicSlots) {
                    patchFlag |= shared_1.PatchFlags.DYNAMIC_SLOTS;
                }
            }
            else if (node.children.length === 1 && vnodeTag !== compiler_core_2.TELEPORT) {
                const child = node.children[0];
                const type = child.type;
                // check for dynamic text children
                const hasDynamicTextChild = type === compiler_core_1.NodeTypes.INTERPOLATION ||
                    type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION;
                if (hasDynamicTextChild &&
                    (0, compiler_core_4.getConstantType)(child, context) === compiler_core_1.ConstantTypes.NOT_CONSTANT) {
                    patchFlag |= shared_1.PatchFlags.TEXT;
                }
                // pass directly if the only child is a text node
                // (plain / interpolation / expression)
                if (hasDynamicTextChild || type === compiler_core_1.NodeTypes.TEXT) {
                    vnodeChildren = child;
                }
                else {
                    vnodeChildren = node.children;
                }
            }
            else {
                vnodeChildren = node.children;
            }
        }
        // patchFlag & dynamicPropNames
        if (patchFlag !== 0) {
            if (utils_1.__DEV__) {
                if (patchFlag < 0) {
                    // special flags (negative and mutually exclusive)
                    vnodePatchFlag =
                        patchFlag + ` /* ${shared_1.PatchFlagNames[patchFlag]} */`;
                }
                else {
                    // bitwise flags
                    const flagNames = Object.keys(shared_1.PatchFlagNames)
                        .map(Number)
                        .filter((n) => n > 0 && patchFlag & n)
                        .map((n) => shared_1.PatchFlagNames[n])
                        .join(`, `);
                    vnodePatchFlag = patchFlag + ` /* ${flagNames} */`;
                }
            }
            else {
                vnodePatchFlag = String(patchFlag);
            }
            if (dynamicPropNames && dynamicPropNames.length) {
                vnodeDynamicProps = stringifyDynamicPropNames(dynamicPropNames);
            }
        }
        node.codegenNode = (0, compiler_core_1.createVNodeCall)(context, vnodeTag, vnodeProps, vnodeChildren, vnodePatchFlag, vnodeDynamicProps, vnodeDirectives, !!shouldUseBlock, false /* disableTracking */, isComponent, node.loc);
    };
};
exports.transformElement = transformElement;
function resolveComponentType(node, context, ssr = false) {
    let { tag } = node;
    // 1. dynamic component
    const isExplicitDynamic = isComponentTag(tag);
    const isProp = (0, compiler_core_3.findProp)(node, 'is');
    if (isProp) {
        if (isExplicitDynamic ||
            (utils_1.__COMPAT__ &&
                (0, utils_1.isCompatEnabled)(compiler_core_6.CompilerDeprecationTypes.COMPILER_IS_ON_ELEMENT, context))) {
            const exp = isProp.type === compiler_core_1.NodeTypes.ATTRIBUTE
                ? isProp.value && (0, compiler_core_1.createSimpleExpression)(isProp.value.content, true)
                : isProp.exp;
            if (exp) {
                return (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.RESOLVE_DYNAMIC_COMPONENT), [
                    exp,
                ]);
            }
        }
        else if (isProp.type === compiler_core_1.NodeTypes.ATTRIBUTE &&
            isProp.value.content.startsWith('vue:')) {
            // <button is="vue:xxx">
            // if not <component>, only is value that starts with "vue:" will be
            // treated as component by the parse phase and reach here, unless it's
            // compat mode where all is values are considered components
            tag = isProp.value.content.slice(4);
        }
    }
    // 2. built-in components (Teleport, Transition, KeepAlive, Suspense...)
    const builtIn = (0, compiler_core_3.isCoreComponent)(tag) || context.isBuiltInComponent(tag);
    if (builtIn) {
        // built-ins are simply fallthroughs / have special handling during ssr
        // so we don't need to import their runtime equivalents
        if (!ssr)
            context.helper(builtIn);
        return builtIn;
    }
    // 3. user component (from setup bindings)
    // this is skipped in browser build since browser builds do not perform
    // binding analysis.
    if (!utils_1.__BROWSER__) {
        const fromSetup = resolveSetupReference(tag, context);
        if (fromSetup) {
            return fromSetup;
        }
        const dotIndex = tag.indexOf('.');
        if (dotIndex > 0) {
            const ns = resolveSetupReference(tag.slice(0, dotIndex), context);
            if (ns) {
                return ns + tag.slice(dotIndex);
            }
        }
    }
    // 4. Self referencing component (inferred from filename)
    if (!utils_1.__BROWSER__ &&
        context.selfName &&
        (0, shared_1.capitalize)((0, shared_1.camelize)(tag)) === context.selfName) {
        context.helper(compiler_core_2.RESOLVE_COMPONENT);
        // codegen.ts has special check for __self postfix when generating
        // component imports, which will pass additional `maybeSelfReference` flag
        // to `resolveComponent`.
        context.components.add(tag + `__self`);
        return (0, compiler_core_3.toValidAssetId)(tag, `component`);
    }
    // 5. user component (resolve)
    context.helper(compiler_core_2.RESOLVE_COMPONENT);
    context.components.add(tag);
    return (0, compiler_core_3.toValidAssetId)(tag, `component`);
}
exports.resolveComponentType = resolveComponentType;
function resolveSetupReference(name, context) {
    const bindings = context.bindingMetadata;
    if (!bindings || bindings.__isScriptSetup === false) {
        return;
    }
    const camelName = (0, shared_1.camelize)(name);
    const PascalName = (0, shared_1.capitalize)(camelName);
    const checkType = (type) => {
        if (bindings[name] === type) {
            return name;
        }
        if (bindings[camelName] === type) {
            return camelName;
        }
        if (bindings[PascalName] === type) {
            return PascalName;
        }
    };
    const fromConst = checkType(compiler_core_5.BindingTypes.SETUP_CONST) ||
        checkType(compiler_core_5.BindingTypes.SETUP_REACTIVE_CONST) ||
        checkType(compiler_core_5.BindingTypes.LITERAL_CONST);
    if (fromConst) {
        return context.inline
            ? // in inline mode, const setup bindings (e.g. imports) can be used as-is
                fromConst
            : `$setup[${JSON.stringify(fromConst)}]`;
    }
    const fromMaybeRef = checkType(compiler_core_5.BindingTypes.SETUP_LET) ||
        checkType(compiler_core_5.BindingTypes.SETUP_REF) ||
        checkType(compiler_core_5.BindingTypes.SETUP_MAYBE_REF);
    if (fromMaybeRef) {
        return context.inline
            ? // setup scope bindings that may be refs need to be unrefed
                `${context.helperString(compiler_core_2.UNREF)}(${fromMaybeRef})`
            : `$setup[${JSON.stringify(fromMaybeRef)}]`;
    }
    const fromProps = checkType(compiler_core_5.BindingTypes.PROPS);
    if (fromProps) {
        return `${context.helperString(compiler_core_2.UNREF)}(${context.inline ? '__props' : '$props'}[${JSON.stringify(fromProps)}])`;
    }
}
function buildProps(node, context, props = node.props, isComponent, isDynamicComponent, ssr = false) {
    const { tag, loc: elementLoc, children } = node;
    let properties = [];
    const mergeArgs = [];
    const runtimeDirectives = [];
    const hasChildren = children.length > 0;
    let shouldUseBlock = false;
    // patchFlag analysis
    let patchFlag = 0;
    let hasRef = false;
    let hasClassBinding = false;
    let hasStyleBinding = false;
    let hasHydrationEventBinding = false;
    let hasDynamicKeys = false;
    let hasVnodeHook = false;
    const dynamicPropNames = [];
    const pushMergeArg = (arg) => {
        if (properties.length) {
            mergeArgs.push((0, compiler_core_1.createObjectExpression)(dedupeProperties(properties), elementLoc));
            properties = [];
        }
        if (arg)
            mergeArgs.push(arg);
    };
    const analyzePatchFlag = ({ key, value }) => {
        if ((0, compiler_core_3.isStaticExp)(key)) {
            const name = key.content;
            const isEventHandler = (0, shared_1.isOn)(name);
            if (isEventHandler &&
                (!isComponent || isDynamicComponent) &&
                // omit the flag for click handlers because hydration gives click
                // dedicated fast path.
                name.toLowerCase() !== 'onclick' &&
                // omit v-model handlers
                name !== 'onUpdate:modelValue' &&
                // omit onVnodeXXX hooks
                !(0, shared_1.isReservedProp)(name)) {
                hasHydrationEventBinding = true;
            }
            if (isEventHandler && (0, shared_1.isReservedProp)(name)) {
                hasVnodeHook = true;
            }
            if (isEventHandler && value.type === compiler_core_1.NodeTypes.JS_CALL_EXPRESSION) {
                // handler wrapped with internal helper e.g. withModifiers(fn)
                // extract the actual expression
                value = value.arguments[0];
            }
            if (value.type === compiler_core_1.NodeTypes.JS_CACHE_EXPRESSION ||
                ((value.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION ||
                    value.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) &&
                    (0, compiler_core_4.getConstantType)(value, context) > 0)) {
                // skip if the prop is a cached handler or has constant value
                return;
            }
            if (name === 'ref') {
                hasRef = true;
            }
            else if (name === 'class') {
                hasClassBinding = true;
            }
            else if (name === 'style') {
                hasStyleBinding = true;
            }
            else if (name !== 'key' && !dynamicPropNames.includes(name)) {
                dynamicPropNames.push(name);
            }
            // treat the dynamic class and style binding of the component as dynamic props
            if (isComponent &&
                (name === 'class' || name === 'style') &&
                !dynamicPropNames.includes(name)) {
                dynamicPropNames.push(name);
            }
        }
        else {
            hasDynamicKeys = true;
        }
    };
    for (let i = 0; i < props.length; i++) {
        // static attribute
        const prop = props[i];
        if (prop.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
            const { loc, name, nameLoc, value } = prop;
            let isStatic = true;
            if (name === 'ref') {
                hasRef = true;
                if (context.scopes.vFor > 0) {
                    properties.push((0, compiler_core_1.createObjectProperty)((0, compiler_core_1.createSimpleExpression)('ref_for', true), (0, compiler_core_1.createSimpleExpression)('true')));
                }
                // in inline mode there is no setupState object, so we can't use string
                // keys to set the ref. Instead, we need to transform it to pass the
                // actual ref instead.
                if (!utils_1.__BROWSER__ && value && context.inline) {
                    const binding = context.bindingMetadata[value.content];
                    if (binding === compiler_core_5.BindingTypes.SETUP_LET ||
                        binding === compiler_core_5.BindingTypes.SETUP_REF ||
                        binding === compiler_core_5.BindingTypes.SETUP_MAYBE_REF) {
                        isStatic = false;
                        properties.push((0, compiler_core_1.createObjectProperty)((0, compiler_core_1.createSimpleExpression)('ref_key', true), (0, compiler_core_1.createSimpleExpression)(value.content, true, value.loc)));
                    }
                }
            }
            // skip is on <component>, or is="vue:xxx"
            if (name === 'is' &&
                (isComponentTag(tag) ||
                    (value && value.content.startsWith('vue:')) ||
                    (utils_1.__COMPAT__ &&
                        (0, utils_1.isCompatEnabled)(compiler_core_6.CompilerDeprecationTypes.COMPILER_IS_ON_ELEMENT, context)))) {
                continue;
            }
            properties.push((0, compiler_core_1.createObjectProperty)((0, compiler_core_1.createSimpleExpression)(name, true, nameLoc), (0, compiler_core_1.createSimpleExpression)(value ? value.content : '', isStatic, value ? value.loc : loc)));
        }
        else {
            // directives
            const { name, arg, exp, loc, modifiers } = prop;
            const isVBind = name === 'bind';
            const isVOn = name === 'on';
            // skip v-slot - it is handled by its dedicated transform.
            if (name === 'slot') {
                if (!isComponent) {
                    context.onError((0, errors_1.createCompilerError)(40 /* ErrorCodes.X_V_SLOT_MISPLACED */, loc));
                }
                continue;
            }
            // skip v-once/v-memo - they are handled by dedicated transforms.
            if (name === 'once' || name === 'memo') {
                continue;
            }
            // skip v-is and :is on <component>
            if (name === 'is' ||
                (isVBind &&
                    (0, compiler_core_3.isStaticArgOf)(arg, 'is') &&
                    (isComponentTag(tag) ||
                        (utils_1.__COMPAT__ &&
                            (0, utils_1.isCompatEnabled)(compiler_core_6.CompilerDeprecationTypes.COMPILER_IS_ON_ELEMENT, context))))) {
                continue;
            }
            // skip v-on in SSR compilation
            if (isVOn && ssr) {
                continue;
            }
            if (
            // #938: elements with dynamic keys should be forced into blocks
            (isVBind && (0, compiler_core_3.isStaticArgOf)(arg, 'key')) ||
                // inline before-update hooks need to force block so that it is invoked
                // before children
                (isVOn && hasChildren && (0, compiler_core_3.isStaticArgOf)(arg, 'vue:before-update'))) {
                shouldUseBlock = true;
            }
            if (isVBind && (0, compiler_core_3.isStaticArgOf)(arg, 'ref') && context.scopes.vFor > 0) {
                properties.push((0, compiler_core_1.createObjectProperty)((0, compiler_core_1.createSimpleExpression)('ref_for', true), (0, compiler_core_1.createSimpleExpression)('true')));
            }
            // special case for v-bind and v-on with no argument
            if (!arg && (isVBind || isVOn)) {
                hasDynamicKeys = true;
                if (exp) {
                    if (isVBind) {
                        // have to merge early for compat build check
                        pushMergeArg();
                        if (utils_1.__COMPAT__) {
                            // 2.x v-bind object order compat
                            if (utils_1.__DEV__) {
                                const hasOverridableKeys = mergeArgs.some((arg) => {
                                    if (arg.type === compiler_core_1.NodeTypes.JS_OBJECT_EXPRESSION) {
                                        return arg.properties.some(({ key }) => {
                                            if (key.type !== compiler_core_1.NodeTypes.SIMPLE_EXPRESSION ||
                                                !key.isStatic) {
                                                return true;
                                            }
                                            return (key.content !== 'class' &&
                                                key.content !== 'style' &&
                                                !(0, shared_1.isOn)(key.content));
                                        });
                                    }
                                    else {
                                        // dynamic expression
                                        return true;
                                    }
                                });
                                if (hasOverridableKeys) {
                                    (0, compiler_core_6.checkCompatEnabled)(compiler_core_6.CompilerDeprecationTypes.COMPILER_V_BIND_OBJECT_ORDER, context, loc);
                                }
                            }
                            if ((0, utils_1.isCompatEnabled)(compiler_core_6.CompilerDeprecationTypes.COMPILER_V_BIND_OBJECT_ORDER, context)) {
                                mergeArgs.unshift(exp);
                                continue;
                            }
                        }
                        mergeArgs.push(exp);
                    }
                    else {
                        // v-on="obj" -> toHandlers(obj)
                        pushMergeArg({
                            type: compiler_core_1.NodeTypes.JS_CALL_EXPRESSION,
                            loc,
                            callee: context.helper(compiler_core_2.TO_HANDLERS),
                            arguments: isComponent ? [exp] : [exp, `true`],
                        });
                    }
                }
                else {
                    context.onError((0, errors_1.createCompilerError)(isVBind
                        ? 34 /* ErrorCodes.X_V_BIND_NO_EXPRESSION */
                        : 35 /* ErrorCodes.X_V_ON_NO_EXPRESSION */, loc));
                }
                continue;
            }
            // force hydration for v-bind with .prop modifier
            if (isVBind && modifiers.includes('prop')) {
                patchFlag |= shared_1.PatchFlags.NEED_HYDRATION;
            }
            const directiveTransform = context.directiveTransforms[name];
            if (directiveTransform) {
                // has built-in directive transform.
                const { props, needRuntime } = directiveTransform(prop, node, context);
                !ssr && props.forEach(analyzePatchFlag);
                if (isVOn && arg && !(0, compiler_core_3.isStaticExp)(arg)) {
                    pushMergeArg((0, compiler_core_1.createObjectExpression)(props, elementLoc));
                }
                else {
                    properties.push(...props);
                }
                if (needRuntime) {
                    runtimeDirectives.push(prop);
                    if ((0, shared_1.isSymbol)(needRuntime)) {
                        directiveImportMap.set(prop, needRuntime);
                    }
                }
            }
            else if (!(0, shared_1.isBuiltInDirective)(name)) {
                // no built-in transform, this is a user custom directive.
                runtimeDirectives.push(prop);
                // custom dirs may use beforeUpdate so they need to force blocks
                // to ensure before-update gets called before children update
                if (hasChildren) {
                    shouldUseBlock = true;
                }
            }
        }
    }
    let propsExpression = undefined;
    // has v-bind="object" or v-on="object", wrap with mergeProps
    if (mergeArgs.length) {
        // close up any not-yet-merged props
        pushMergeArg();
        if (mergeArgs.length > 1) {
            propsExpression = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.MERGE_PROPS), mergeArgs, elementLoc);
        }
        else {
            // single v-bind with nothing else - no need for a mergeProps call
            propsExpression = mergeArgs[0];
        }
    }
    else if (properties.length) {
        propsExpression = (0, compiler_core_1.createObjectExpression)(dedupeProperties(properties), elementLoc);
    }
    // patchFlag analysis
    if (hasDynamicKeys) {
        patchFlag |= shared_1.PatchFlags.FULL_PROPS;
    }
    else {
        if (hasClassBinding && !isComponent) {
            patchFlag |= shared_1.PatchFlags.CLASS;
        }
        if (hasStyleBinding && !isComponent) {
            patchFlag |= shared_1.PatchFlags.STYLE;
        }
        if (dynamicPropNames.length) {
            patchFlag |= shared_1.PatchFlags.PROPS;
        }
        if (hasHydrationEventBinding) {
            patchFlag |= shared_1.PatchFlags.NEED_HYDRATION;
        }
    }
    if (!shouldUseBlock &&
        (patchFlag === 0 || patchFlag === shared_1.PatchFlags.NEED_HYDRATION) &&
        (hasRef || hasVnodeHook || runtimeDirectives.length > 0)) {
        patchFlag |= shared_1.PatchFlags.NEED_PATCH;
    }
    // pre-normalize props, SSR is skipped for now
    if (!context.inSSR && propsExpression) {
        switch (propsExpression.type) {
            case compiler_core_1.NodeTypes.JS_OBJECT_EXPRESSION:
                // means that there is no v-bind,
                // but still need to deal with dynamic key binding
                let classKeyIndex = -1;
                let styleKeyIndex = -1;
                let hasDynamicKey = false;
                for (let i = 0; i < propsExpression.properties.length; i++) {
                    const key = propsExpression.properties[i].key;
                    if ((0, compiler_core_3.isStaticExp)(key)) {
                        if (key.content === 'class') {
                            classKeyIndex = i;
                        }
                        else if (key.content === 'style') {
                            styleKeyIndex = i;
                        }
                    }
                    else if (!key.isHandlerKey) {
                        hasDynamicKey = true;
                    }
                }
                const classProp = propsExpression.properties[classKeyIndex];
                const styleProp = propsExpression.properties[styleKeyIndex];
                // no dynamic key
                if (!hasDynamicKey) {
                    if (classProp && !(0, compiler_core_3.isStaticExp)(classProp.value)) {
                        classProp.value = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.NORMALIZE_CLASS), [classProp.value]);
                    }
                    if (styleProp &&
                        // the static style is compiled into an object,
                        // so use `hasStyleBinding` to ensure that it is a dynamic style binding
                        (hasStyleBinding ||
                            (styleProp.value.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
                                styleProp.value.content.trim()[0] === `[`) ||
                            // v-bind:style and style both exist,
                            // v-bind:style with static literal object
                            styleProp.value.type === compiler_core_1.NodeTypes.JS_ARRAY_EXPRESSION)) {
                        styleProp.value = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.NORMALIZE_STYLE), [styleProp.value]);
                    }
                }
                else {
                    // dynamic key binding, wrap with `normalizeProps`
                    propsExpression = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.NORMALIZE_PROPS), [propsExpression]);
                }
                break;
            case compiler_core_1.NodeTypes.JS_CALL_EXPRESSION:
                // mergeProps call, do nothing
                break;
            default:
                // single v-bind
                propsExpression = (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.NORMALIZE_PROPS), [
                    (0, compiler_core_1.createCallExpression)(context.helper(compiler_core_2.GUARD_REACTIVE_PROPS), [
                        propsExpression,
                    ]),
                ]);
                break;
        }
    }
    return {
        props: propsExpression,
        directives: runtimeDirectives,
        patchFlag,
        dynamicPropNames,
        shouldUseBlock,
    };
}
exports.buildProps = buildProps;
// Dedupe props in an object literal.
// Literal duplicated attributes would have been warned during the parse phase,
// however, it's possible to encounter duplicated `onXXX` handlers with different
// modifiers. We also need to merge static and dynamic class / style attributes.
// - onXXX handlers / style: merge into array
// - class: merge into single expression with concatenation
function dedupeProperties(properties) {
    const knownProps = new Map();
    const deduped = [];
    for (let i = 0; i < properties.length; i++) {
        const prop = properties[i];
        // dynamic keys are always allowed
        if (prop.key.type === compiler_core_1.NodeTypes.COMPOUND_EXPRESSION || !prop.key.isStatic) {
            deduped.push(prop);
            continue;
        }
        const name = prop.key.content;
        const existing = knownProps.get(name);
        if (existing) {
            if (name === 'style' || name === 'class' || (0, shared_1.isOn)(name)) {
                mergeAsArray(existing, prop);
            }
            // unexpected duplicate, should have emitted error during parse
        }
        else {
            knownProps.set(name, prop);
            deduped.push(prop);
        }
    }
    return deduped;
}
function mergeAsArray(existing, incoming) {
    if (existing.value.type === compiler_core_1.NodeTypes.JS_ARRAY_EXPRESSION) {
        existing.value.elements.push(incoming.value);
    }
    else {
        existing.value = (0, compiler_core_1.createArrayExpression)([existing.value, incoming.value], existing.loc);
    }
}
function buildDirectiveArgs(dir, context) {
    const dirArgs = [];
    const runtime = directiveImportMap.get(dir);
    if (runtime) {
        // built-in directive with runtime
        dirArgs.push(context.helperString(runtime));
    }
    else {
        // user directive.
        // see if we have directives exposed via <script setup>
        const fromSetup = !utils_1.__BROWSER__ && resolveSetupReference('v-' + dir.name, context);
        if (fromSetup) {
            dirArgs.push(fromSetup);
        }
        else {
            // inject statement for resolving directive
            context.helper(compiler_core_2.RESOLVE_DIRECTIVE);
            context.directives.add(dir.name);
            dirArgs.push((0, compiler_core_3.toValidAssetId)(dir.name, `directive`));
        }
    }
    const { loc } = dir;
    if (dir.exp)
        dirArgs.push(dir.exp);
    if (dir.arg) {
        if (!dir.exp) {
            dirArgs.push(`void 0`);
        }
        dirArgs.push(dir.arg);
    }
    if (Object.keys(dir.modifiers).length) {
        if (!dir.arg) {
            if (!dir.exp) {
                dirArgs.push(`void 0`);
            }
            dirArgs.push(`void 0`);
        }
        const trueExpression = (0, compiler_core_1.createSimpleExpression)(`true`, false, loc);
        dirArgs.push((0, compiler_core_1.createObjectExpression)(dir.modifiers.map((modifier) => (0, compiler_core_1.createObjectProperty)(modifier, trueExpression)), loc));
    }
    return (0, compiler_core_1.createArrayExpression)(dirArgs, dir.loc);
}
exports.buildDirectiveArgs = buildDirectiveArgs;
function stringifyDynamicPropNames(props) {
    let propsNamesString = `[`;
    for (let i = 0, l = props.length; i < l; i++) {
        propsNamesString += JSON.stringify(props[i]);
        if (i < l - 1)
            propsNamesString += ', ';
    }
    return propsNamesString + `]`;
}
function isComponentTag(tag) {
    return tag === 'component' || tag === 'Component';
}
