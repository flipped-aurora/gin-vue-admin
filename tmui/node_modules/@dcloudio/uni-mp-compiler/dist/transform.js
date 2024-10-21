"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createStructuralDirectiveTransform = exports.createTransformContext = exports.traverseChildren = exports.traverseNode = exports.transform = exports.isScopedSlotVFor = exports.isVForScope = exports.isVIfScope = exports.isRootScope = void 0;
const shared_1 = require("@vue/shared");
const types_1 = require("@babel/types");
const compiler_core_1 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const identifier_1 = __importDefault(require("./identifier"));
const runtimeHelpers_1 = require("./runtimeHelpers");
const ast_1 = require("./ast");
const utils_1 = require("./transforms/utils");
const codegen_1 = require("./codegen");
function isRootScope(scope) {
    return !isVIfScope(scope) && !isVForScope(scope);
}
exports.isRootScope = isRootScope;
function isVIfScope(scope) {
    return (!!scope.condition ||
        scope.name === 'else');
}
exports.isVIfScope = isVIfScope;
function isVForScope(scope) {
    return !!scope.source;
}
exports.isVForScope = isVForScope;
function isScopedSlotVFor({ source }) {
    if (source.type !== compiler_core_1.NodeTypes.COMPOUND_EXPRESSION) {
        return false;
    }
    const first = source.children[0];
    return (first.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
        first.content.includes(utils_1.SCOPED_SLOT_IDENTIFIER));
}
exports.isScopedSlotVFor = isScopedSlotVFor;
function transform(root, options) {
    const context = createTransformContext(root, options);
    findRootNode(root, context);
    traverseNode(root, context);
    root.renderData = createRenderDataExpr(context.scope.properties, context);
    // finalize meta information
    root.helpers = new Set([...context.helpers.keys()]);
    root.components = [...context.components];
    root.imports = context.imports;
    root.cached = context.cached;
    return context;
}
exports.transform = transform;
function findRootNode(root, context) {
    const children = root.children.filter((node) => node.type === compiler_core_1.NodeTypes.ELEMENT && node.tag !== 'template');
    if (children.length === 1) {
        context.rootNode = children[0];
    }
}
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
            // context.helper(CREATE_COMMENT)
            break;
        case compiler_core_1.NodeTypes.INTERPOLATION:
            context.helper(compiler_core_1.TO_DISPLAY_STRING);
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
function defaultOnError(error) {
    throw error;
}
function defaultOnWarn(msg) {
    console.warn(`[Vue warn] ${msg.message}`);
}
function createTransformContext(rootNode, { root = '', filename = '', isTS = false, inline = false, hashId = null, scopeId = null, filters = [], bindingCssVars = [], bindingMetadata = shared_1.EMPTY_OBJ, cacheHandlers = false, prefixIdentifiers = false, skipTransformIdentifier = false, renderDataSpread = false, nodeTransforms = [], directiveTransforms = {}, miniProgram = {
    class: {
        array: true,
    },
    slot: {
        fallbackContent: false,
        dynamicSlotNames: true,
    },
    directive: '',
}, isBuiltInComponent = shared_1.NOOP, isCustomElement = shared_1.NOOP, expressionPlugins = [], onError = defaultOnError, onWarn = defaultOnWarn, }) {
    const rootScope = {
        id: new identifier_1.default(),
        identifiers: [],
        properties: [],
        parent: null,
    };
    function findVIfParentScope() {
        for (let i = scopes.length - 1; i >= 0; i--) {
            const scope = scopes[i];
            if (isVForScope(scope) || isRootScope(scope)) {
                return scope;
            }
        }
        return rootScope;
    }
    function createScope(id, initScope) {
        return (0, shared_1.extend)({
            id,
            properties: [],
            parent: scopes[scopes.length - 1],
            get identifiers() {
                return Object.keys(identifiers);
            },
        }, initScope);
    }
    const vueIds = [];
    const identifiers = Object.create(null);
    const scopes = [rootScope];
    const miniProgramComponents = (0, uni_cli_shared_1.findMiniProgramUsingComponents)({
        filename,
        componentsDir: miniProgram.component?.dir,
        inputDir: root,
    });
    // const nameMatch = filename.replace(/\?.*$/, '').match(/([^/\\]+)\.\w+$/)
    const context = {
        // options
        // 暂不提供根据文件名生成递归组件
        selfName: '', //nameMatch && capitalize(camelize(nameMatch[1])),
        miniProgram,
        isTS,
        inline,
        hashId,
        scopeId,
        filters,
        bindingCssVars,
        bindingMetadata,
        cacheHandlers,
        prefixIdentifiers,
        nodeTransforms,
        directiveTransforms,
        expressionPlugins,
        skipTransformIdentifier,
        renderDataSpread,
        isBuiltInComponent,
        isCustomElement,
        onError,
        onWarn,
        // state
        parent: null,
        childIndex: 0,
        helpers: new Map(),
        components: new Set(),
        imports: [],
        bindingComponents: Object.create(null),
        cached: 0,
        identifiers,
        scope: rootScope,
        scopes: {
            vFor: 0,
            vueId: 0,
        },
        get currentScope() {
            return scopes[scopes.length - 1];
        },
        currentNode: rootNode,
        vueIds,
        get currentVueId() {
            return vueIds[vueIds.length - 1];
        },
        inVOnce: false,
        get inVFor() {
            let parent = scopes[scopes.length - 1];
            while (parent) {
                if (isVForScope(parent) && !isScopedSlotVFor(parent)) {
                    return true;
                }
                parent = parent.parent;
            }
            return false;
        },
        // methods
        getScopeIndex(scope) {
            return scopes.indexOf(scope);
        },
        popScope() {
            return scopes.pop();
        },
        addVIfScope(initScope) {
            const vIfScope = createScope(scopes[scopes.length - 1].id, (0, shared_1.extend)(initScope, { parentScope: findVIfParentScope() }));
            scopes.push(vIfScope);
            return vIfScope;
        },
        addVForScope(initScope) {
            const vForScope = createScope(new identifier_1.default(), initScope);
            scopes.push(vForScope);
            return vForScope;
        },
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
            return `_${compiler_core_1.helperNameMap[context.helper(name)]}`;
        },
        replaceNode(node) {
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
            return createCacheExpression(context.cached++, exp, isVNode);
        },
        isMiniProgramComponent(name) {
            return miniProgramComponents[name];
        },
        rootNode: null,
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
function createCacheExpression(index, value, isVNode = false) {
    return {
        type: compiler_core_1.NodeTypes.JS_CACHE_EXPRESSION,
        index,
        value,
        isVNode,
        loc: compiler_core_1.locStub,
    };
}
function createStructuralDirectiveTransform(name, fn) {
    const matches = (0, shared_1.isString)(name)
        ? (n) => n === name
        : (n) => name.test(n);
    return (node, context) => {
        if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
            const { props } = node;
            // structural directive transforms are not concerned with slots
            // as they are handled separately in vSlot.ts
            // if (node.tagType === ElementTypes.TEMPLATE && props.some(isVSlot)) {
            //   return
            // }
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
function createRenderDataExpr(properties, context) {
    const objExpr = (0, ast_1.createObjectExpression)(properties);
    if (!hasSpreadElement(objExpr)) {
        return objExpr;
    }
    // filters: ['test']
    // v-if="text.aa()"
    if (context.filters.length) {
        transformFilterObjectSpreadExpr(objExpr, context);
    }
    if (context.renderDataSpread) {
        return objExpr;
    }
    return transformObjectSpreadExpr(objExpr, context);
}
function hasSpreadElement(expr) {
    return expr.properties.some((prop) => {
        if ((0, types_1.isSpreadElement)(prop)) {
            return true;
        }
        else {
            const returnStatement = parseReturnStatement(prop);
            if (returnStatement) {
                return hasSpreadElement(returnStatement.argument);
            }
        }
    });
}
// 目前硬编码识别 _f,应该读取 context.helperString
const returnObjExprMap = {
    _f: 1, // _f(_ctx.items,()=>{return {}})
    _w: 0, // _w(()=>{return {}})
};
function parseReturnStatement(prop) {
    if ((0, types_1.isObjectProperty)(prop) &&
        (0, types_1.isCallExpression)(prop.value) &&
        (0, types_1.isIdentifier)(prop.value.callee)) {
        const { name } = prop.value.callee;
        if ((0, shared_1.hasOwn)(returnObjExprMap, name)) {
            return prop.value.arguments[returnObjExprMap[name]].body.body[0];
        }
    }
}
function transformObjectPropertyExpr(prop, context) {
    // vFor,withScopedSlot
    const returnStatement = parseReturnStatement(prop);
    if (returnStatement) {
        const objExpr = returnStatement.argument;
        if (hasSpreadElement(objExpr)) {
            returnStatement.argument = transformObjectSpreadExpr(objExpr, context);
        }
    }
    return prop;
}
function transformObjectSpreadExpr(objExpr, context) {
    const properties = objExpr.properties;
    const args = [];
    let objExprProperties = [];
    properties.forEach((prop) => {
        if ((0, types_1.isObjectProperty)(prop)) {
            objExprProperties.push(transformObjectPropertyExpr(prop, context));
        }
        else {
            if (objExprProperties.length) {
                args.push((0, types_1.objectExpression)(objExprProperties));
            }
            args.push(transformConditionalExpression(prop.argument, context));
            objExprProperties = [];
        }
    });
    if (objExprProperties.length) {
        args.push((0, types_1.objectExpression)(objExprProperties));
    }
    if (args.length === 1) {
        return args[0];
    }
    return (0, types_1.callExpression)((0, types_1.identifier)(context.helperString(runtimeHelpers_1.EXTEND)), args);
}
function transformConditionalExpression(expr, context) {
    const { consequent, alternate } = expr;
    if ((0, types_1.isObjectExpression)(consequent) && hasSpreadElement(consequent)) {
        expr.consequent = transformObjectSpreadExpr(consequent, context);
    }
    if ((0, types_1.isObjectExpression)(alternate)) {
        if (hasSpreadElement(alternate)) {
            expr.alternate = transformObjectSpreadExpr(alternate, context);
        }
    }
    else if ((0, types_1.isConditionalExpression)(alternate)) {
        transformConditionalExpression(alternate, context);
    }
    return expr;
}
function transformFilterObjectSpreadExpr(objExpr, context) {
    const properties = objExpr.properties;
    properties.forEach((prop) => {
        if ((0, types_1.isObjectProperty)(prop)) {
            transformFilterObjectPropertyExpr(prop, context);
        }
        else {
            prop.argument = transformFilterConditionalExpression(prop.argument, context);
        }
    });
}
function transformFilterObjectPropertyExpr(prop, context) {
    // vFor, withScopedSlot
    const returnStatement = parseReturnStatement(prop);
    if (returnStatement) {
        const objExpr = returnStatement.argument;
        if (hasSpreadElement(objExpr)) {
            transformFilterObjectSpreadExpr(objExpr, context);
        }
    }
}
function transformFilterConditionalExpression(expr, context) {
    const { test, consequent, alternate } = expr;
    if ((0, types_1.isObjectExpression)(consequent) && hasSpreadElement(consequent)) {
        transformFilterObjectSpreadExpr(consequent, context);
    }
    if ((0, types_1.isObjectExpression)(alternate)) {
        if (hasSpreadElement(alternate)) {
            transformFilterObjectSpreadExpr(alternate, context);
        }
    }
    else if ((0, types_1.isConditionalExpression)(alternate)) {
        expr.alternate = transformFilterConditionalExpression(alternate, context);
    }
    const testCode = (0, codegen_1.genBabelExpr)(test);
    // filter test
    if (context.filters.find((filter) => testCode.includes(filter + '.'))) {
        // test.aa() ? {a:1} : {b:2} => {...{a:1},...{b:2}}
        const properties = [];
        if (!(0, types_1.isObjectExpression)(consequent) || consequent.properties.length) {
            properties.push((0, types_1.spreadElement)(consequent));
        }
        if (!(0, types_1.isObjectExpression)(expr.alternate) ||
            expr.alternate.properties.length) {
            properties.push((0, types_1.spreadElement)(expr.alternate));
        }
        return (0, types_1.objectExpression)(properties);
    }
    return expr;
}
