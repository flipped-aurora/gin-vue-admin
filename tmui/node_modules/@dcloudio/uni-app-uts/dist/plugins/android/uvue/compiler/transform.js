"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createStructuralDirectiveTransform = exports.traverseNode = exports.traverseChildren = exports.isSingleElementRoot = exports.transform = exports.createTransformContext = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
const errors_1 = require("./errors");
const runtimeHelpers_1 = require("./runtimeHelpers");
function createTransformContext(root, { mode = 'default', rootDir = '', targetLanguage = 'kotlin', filename = '', cacheHandlers = false, prefixIdentifiers = mode === 'module', nodeTransforms = [], directiveTransforms = {}, scopeId = null, slotted = true, bindingMetadata = shared_1.EMPTY_OBJ, inline = false, isBuiltInComponent = shared_1.NOOP, isCustomElement = shared_1.NOOP, onError = errors_1.defaultOnError, onWarn = errors_1.defaultOnWarn, parseUTSComponent = shared_1.NOOP, }) {
    const nameMatch = filename.replace(/\?.*$/, '').match(/([^/\\]+)\.\w+$/);
    const context = {
        // options
        mode,
        rootDir,
        targetLanguage,
        selfName: nameMatch && (0, shared_1.capitalize)((0, shared_1.camelize)(nameMatch[1])),
        cacheHandlers,
        prefixIdentifiers,
        bindingMetadata,
        inline,
        nodeTransforms,
        directiveTransforms,
        elements: new Set(),
        isBuiltInComponent,
        isCustomElement,
        scopeId,
        slotted,
        onError,
        onWarn,
        parseUTSComponent,
        // state
        root,
        helpers: new Map(),
        components: new Set(),
        directives: new Set(),
        imports: [],
        constantCache: new Map(),
        temps: 0,
        cached: 0,
        identifiers: Object.create(null),
        scopes: {
            vFor: 0,
            vSlot: 0,
            vPre: 0,
            vOnce: 0,
        },
        parent: null,
        currentNode: root,
        childIndex: 0,
        inVOnce: false,
        expressionPlugins: ['typescript'],
        // methods
        helper(name) {
            const count = context.helpers.get(name) || 0;
            context.helpers.set(name, count + 1);
            return name;
        },
        removeHelper(name) {
            const count = context.helpers.get(name);
            if (count) {
                const currentCount = count - 1;
                if (!currentCount) {
                    context.helpers.delete(name);
                }
                else {
                    context.helpers.set(name, currentCount);
                }
            }
        },
        helperString(name) {
            // return `_${helperNameMap[context.helper(name)]}`
            return `${compiler_core_1.helperNameMap[context.helper(name)]}`;
        },
        replaceNode(node) {
            if (!context.currentNode) {
                throw new Error(`Node being replaced is already removed.`);
            }
            if (!context.parent) {
                throw new Error(`Cannot replace root node.`);
            }
            context.parent.children[context.childIndex] = context.currentNode = node;
        },
        removeNode(node) {
            if (!context.parent) {
                throw new Error(`Cannot remove root node.`);
            }
            const list = context.parent.children;
            const removalIndex = node
                ? list.indexOf(node)
                : context.currentNode
                    ? context.childIndex
                    : -1;
            /* istanbul ignore if */
            if (removalIndex < 0) {
                throw new Error(`node being removed is not a child of current parent`);
            }
            if (!node || node === context.currentNode) {
                // current node removed
                context.currentNode = null;
                context.onNodeRemoved();
            }
            else {
                // sibling node removed
                if (context.childIndex > removalIndex) {
                    context.childIndex--;
                    context.onNodeRemoved();
                }
            }
            context.parent.children.splice(removalIndex, 1);
        },
        onNodeRemoved: () => { },
        addIdentifiers(exp) {
            // identifier tracking only happens in non-browser builds.
            if ((0, shared_1.isString)(exp)) {
                addId(exp);
            }
            else if (exp.identifiers) {
                exp.identifiers.forEach(addId);
            }
            else if (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
                addId(exp.content);
            }
        },
        removeIdentifiers(exp) {
            if ((0, shared_1.isString)(exp)) {
                removeId(exp);
            }
            else if (exp.identifiers) {
                exp.identifiers.forEach(removeId);
            }
            else if (exp.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION) {
                removeId(exp.content);
            }
        },
        cache(exp, isVNode = false) {
            return (0, compiler_core_1.createCacheExpression)(context.cached++, exp, isVNode);
        },
    };
    function addId(id) {
        const { identifiers } = context;
        if (identifiers[id] === undefined) {
            identifiers[id] = 0;
        }
        identifiers[id]++;
    }
    function removeId(id) {
        context.identifiers[id]--;
    }
    return context;
}
exports.createTransformContext = createTransformContext;
function transform(root, options) {
    const context = createTransformContext(root, options);
    traverseNode(root, context);
    createRootCodegen(root, context);
    // finalize meta information
    root.helpers = new Set([...context.helpers.keys()]);
    root.components = [...context.components];
    root.directives = [...context.directives];
    root.imports = context.imports;
    // root.hoists = context.hoists
    root.temps = context.temps;
    root.cached = context.cached;
    // @ts-expect-error
    root.elements = Array.from(context.elements);
}
exports.transform = transform;
function isSingleElementRoot(root, child) {
    const { children } = root;
    return (children.length === 1 &&
        child.type === compiler_core_1.NodeTypes.ELEMENT &&
        !(0, compiler_core_1.isSlotOutlet)(child));
}
exports.isSingleElementRoot = isSingleElementRoot;
function createRootCodegen(root, context) {
    const { helper } = context;
    const { children } = root;
    if (children.length === 1) {
        const child = children[0];
        // if the single child is an element, turn it into a block.
        if (isSingleElementRoot(root, child) && child.codegenNode) {
            // single element root is never hoisted so codegenNode will never be
            // SimpleExpressionNode
            const codegenNode = child.codegenNode;
            if (codegenNode.type === compiler_core_1.NodeTypes.VNODE_CALL) {
                (0, compiler_core_1.convertToBlock)(codegenNode, context);
            }
            root.codegenNode = codegenNode;
        }
        else {
            // - single <slot/>, IfNode, ForNode: already blocks.
            // - single text node: always patched.
            // root codegen falls through via genNode()
            root.codegenNode = child;
        }
    }
    else if (children.length > 1) {
        // root has multiple nodes - return a fragment block.
        let patchFlag = shared_1.PatchFlags.STABLE_FRAGMENT;
        let patchFlagText = shared_1.PatchFlagNames[shared_1.PatchFlags.STABLE_FRAGMENT];
        // check if the fragment actually contains a single valid child with
        // the rest being comments
        if (children.filter((c) => c.type !== compiler_core_1.NodeTypes.COMMENT).length === 1) {
            patchFlag |= shared_1.PatchFlags.DEV_ROOT_FRAGMENT;
            patchFlagText += `, ${shared_1.PatchFlagNames[shared_1.PatchFlags.DEV_ROOT_FRAGMENT]}`;
        }
        root.codegenNode = (0, compiler_core_1.createVNodeCall)(
        // @ts-expect-error
        context, helper(runtimeHelpers_1.FRAGMENT), undefined, root.children, patchFlag + ` /* ${patchFlagText} */`, undefined, undefined, true, undefined, false /* isComponent */);
    }
    else {
        // no children = noop. codegen will return null.
    }
}
function traverseChildren(parent, context) {
    let i = 0;
    const nodeRemoved = () => {
        i--;
    };
    for (; i < parent.children.length; i++) {
        const child = parent.children[i];
        if ((0, shared_1.isString)(child))
            continue;
        context.parent = parent;
        context.childIndex = i;
        context.onNodeRemoved = nodeRemoved;
        traverseNode(child, context);
    }
}
exports.traverseChildren = traverseChildren;
function traverseNode(node, context) {
    context.currentNode = node;
    // apply transform plugins
    const { nodeTransforms } = context;
    const exitFns = [];
    for (let i = 0; i < nodeTransforms.length; i++) {
        const onExit = nodeTransforms[i](node, context);
        if (onExit) {
            if ((0, shared_1.isArray)(onExit)) {
                exitFns.push(...onExit);
            }
            else {
                exitFns.push(onExit);
            }
        }
        if (!context.currentNode) {
            // node was removed
            return;
        }
        else {
            // node may have been replaced
            node = context.currentNode;
        }
    }
    switch (node.type) {
        case compiler_core_1.NodeTypes.COMMENT:
            break;
        case compiler_core_1.NodeTypes.INTERPOLATION:
            break;
        // for container types, further traverse downwards
        case compiler_core_1.NodeTypes.IF:
            for (let i = 0; i < node.branches.length; i++) {
                traverseNode(node.branches[i], context);
            }
            break;
        case compiler_core_1.NodeTypes.IF_BRANCH:
        case compiler_core_1.NodeTypes.FOR:
        case compiler_core_1.NodeTypes.ELEMENT:
        case compiler_core_1.NodeTypes.ROOT:
            traverseChildren(node, context);
            break;
    }
    // exit transforms
    context.currentNode = node;
    let i = exitFns.length;
    while (i--) {
        exitFns[i]();
    }
}
exports.traverseNode = traverseNode;
function createStructuralDirectiveTransform(name, fn) {
    const matches = (0, shared_1.isString)(name)
        ? (n) => n === name
        : (n) => name.test(n);
    return (node, context) => {
        if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
            const { props } = node;
            // structural directive transforms are not concerned with slots
            // as they are handled separately in vSlot.ts
            if (node.tagType === compiler_core_1.ElementTypes.TEMPLATE && props.some(compiler_core_1.isVSlot)) {
                return;
            }
            const exitFns = [];
            for (let i = 0; i < props.length; i++) {
                const prop = props[i];
                if (prop.type === compiler_core_1.NodeTypes.DIRECTIVE && matches(prop.name)) {
                    // structural directives are removed to avoid infinite recursion
                    // also we remove them *before* applying so that it can further
                    // traverse itself in case it moves the node around
                    props.splice(i, 1);
                    i--;
                    const onExit = fn(node, prop, context);
                    if (onExit)
                        exitFns.push(onExit);
                }
            }
            return exitFns;
        }
    };
}
exports.createStructuralDirectiveTransform = createStructuralDirectiveTransform;
