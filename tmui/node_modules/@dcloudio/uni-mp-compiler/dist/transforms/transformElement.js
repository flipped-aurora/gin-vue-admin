"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.processProps = exports.transformElement = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const uni_shared_1 = require("@dcloudio/uni-shared");
const errors_1 = require("../errors");
const vModel_1 = require("./vModel");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const transformElement = (node, context) => {
    return function postTransformElement() {
        node = context.currentNode;
        if (!(node.type === compiler_core_1.NodeTypes.ELEMENT &&
            (node.tagType === compiler_core_1.ElementTypes.ELEMENT ||
                node.tagType === compiler_core_1.ElementTypes.COMPONENT))) {
            return;
        }
        if (node.tagType === compiler_core_1.ElementTypes.COMPONENT) {
            processComponent(node, context);
        }
        if (context.scopeId) {
            addScopeId(node, context.scopeId);
        }
        const { props } = node;
        if (props.length > 0) {
            processProps(node, context);
        }
    };
};
exports.transformElement = transformElement;
function addScopeId(node, scopeId) {
    return (0, uni_cli_shared_1.addStaticClass)(node, scopeId);
}
function processComponent(node, context) {
    const { tag } = node;
    if (context.bindingComponents[tag]) {
        return;
    }
    // 1. dynamic component
    if ((0, uni_shared_1.isComponentTag)(tag)) {
        return context.onError((0, errors_1.createMPCompilerError)(8 /* MPErrorCodes.X_DYNAMIC_COMPONENT_NOT_SUPPORTED */, node.loc));
    }
    if ((0, compiler_core_1.findDir)(node, 'is')) {
        return context.onError((0, errors_1.createMPCompilerError)(6 /* MPErrorCodes.X_V_IS_NOT_SUPPORTED */, node.loc));
    }
    // TODO not supported
    // const isProp = findProp(node, 'is')
    // if (isProp) {
    // }
    // 2. built-in components (Teleport, Transition, KeepAlive, Suspense...)
    const builtIn = (0, compiler_core_1.isCoreComponent)(tag) || context.isBuiltInComponent(tag);
    if (builtIn) {
        return context.onError((0, errors_1.createMPCompilerError)(7 /* MPErrorCodes.X_NOT_SUPPORTED */, node.loc, tag));
    }
    // 3. user component (from setup bindings)
    const fromSetup = resolveSetupReference(tag, context);
    if (fromSetup) {
        return (context.bindingComponents[tag] = {
            name: fromSetup,
            type: "setup" /* BindingComponentTypes.SETUP */,
        });
    }
    const dotIndex = tag.indexOf('.');
    if (dotIndex > 0) {
        return context.onError((0, errors_1.createMPCompilerError)(7 /* MPErrorCodes.X_NOT_SUPPORTED */, node.loc, tag));
    }
    // 4. Self referencing component (inferred from filename)
    if (context.selfName && (0, shared_1.capitalize)((0, shared_1.camelize)(tag)) === context.selfName) {
        return (context.bindingComponents[tag] = {
            name: (0, compiler_core_1.toValidAssetId)(tag, `component`),
            type: "self" /* BindingComponentTypes.SELF */,
        });
    }
    // 5. user component (resolve)
    context.bindingComponents[tag] = {
        name: (0, compiler_core_1.toValidAssetId)(tag, `component`),
        type: "unknown" /* BindingComponentTypes.UNKNOWN */,
    };
    context.helper(compiler_core_1.RESOLVE_COMPONENT);
}
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
    const fromConst = checkType(compiler_core_1.BindingTypes.SETUP_CONST) ||
        checkType(compiler_core_1.BindingTypes.SETUP_REACTIVE_CONST) ||
        checkType(compiler_core_1.BindingTypes.LITERAL_CONST);
    if (fromConst) {
        return context.inline
            ? // in inline mode, const setup bindings (e.g. imports) can be used as-is
                fromConst
            : `$setup[${JSON.stringify(fromConst)}]`;
    }
    const fromMaybeRef = checkType(compiler_core_1.BindingTypes.SETUP_LET) ||
        checkType(compiler_core_1.BindingTypes.SETUP_REF) ||
        checkType(compiler_core_1.BindingTypes.SETUP_MAYBE_REF);
    if (fromMaybeRef) {
        return context.inline
            ? // setup scope bindings that may be refs need to be unrefed
                `${context.helperString(compiler_core_1.UNREF)}(${fromMaybeRef})`
            : `$setup[${JSON.stringify(fromMaybeRef)}]`;
    }
}
function processProps(node, context, props = node.props) {
    const { tag } = node;
    const isComponent = node.tagType === compiler_core_1.ElementTypes.COMPONENT;
    const isPluginComponent = isComponent && context.isMiniProgramComponent(node.tag) === 'plugin';
    for (let i = 0; i < props.length; i++) {
        const prop = props[i];
        if (prop.type === compiler_core_1.NodeTypes.DIRECTIVE) {
            // directives
            const { name, arg, loc } = prop;
            const isVBind = name === 'bind';
            const isVOn = name === 'on';
            // skip v-slot - it is handled by its dedicated transform.
            if (name === 'slot') {
                if (!isComponent) {
                    context.onError((0, compiler_core_1.createCompilerError)(compiler_core_1.ErrorCodes.X_V_SLOT_MISPLACED, loc));
                }
                continue;
            }
            // skip v-once/v-memo - they are handled by dedicated transforms.
            if (name === 'once' || name === 'memo') {
                continue;
            }
            // skip v-is and :is on <component>
            if (name === 'is' ||
                (isVBind && (0, compiler_core_1.isStaticArgOf)(arg, 'is') && (0, uni_shared_1.isComponentTag)(tag))) {
                continue;
            }
            if (isVBind || isVOn) {
                // v-on=""
                // v-bind=""
                if (!arg) {
                    if (isVOn) {
                        context.onError((0, errors_1.createMPCompilerError)(0 /* MPErrorCodes.X_V_ON_NO_ARGUMENT */, loc));
                    }
                    if (isVBind && (!isComponent || isPluginComponent)) {
                        context.onError((0, errors_1.createMPCompilerError)(2 /* MPErrorCodes.X_V_BIND_NO_ARGUMENT */, loc));
                    }
                    continue;
                }
                // v-on:[a]=""
                // v-on:[a.b]=""
                // v-bind:[a]=""
                // v-bind:[a.b]=""
                if (!(arg.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && arg.isStatic)) {
                    if (isVOn) {
                        context.onError((0, errors_1.createMPCompilerError)(1 /* MPErrorCodes.X_V_ON_DYNAMIC_EVENT */, loc));
                    }
                    if (isVBind && (!isComponent || isPluginComponent)) {
                        context.onError((0, errors_1.createMPCompilerError)(3 /* MPErrorCodes.X_V_BIND_DYNAMIC_ARGUMENT */, loc));
                    }
                    continue;
                }
            }
            const directiveTransform = context.directiveTransforms[name];
            if (name !== 'model' && directiveTransform) {
                const { props } = directiveTransform(prop, node, context);
                if (props.length) {
                    prop.exp = props[0].value;
                }
            }
        }
    }
    const transformVModel = (context.directiveTransforms.model ||
        vModel_1.transformModel);
    processVModel(node, transformVModel, context);
}
exports.processProps = processProps;
function processVModel(node, transformVModel, context) {
    const { props } = node;
    const dirs = [];
    for (let i = 0; i < props.length; i++) {
        const prop = props[i];
        if (prop.type === compiler_core_1.NodeTypes.DIRECTIVE && prop.name === 'model') {
            dirs.push(...transformVModel(prop, node, context)
                .props);
            props.splice(i, 1);
            i--;
        }
    }
    if (dirs.length) {
        props.push(...dirs);
    }
}
