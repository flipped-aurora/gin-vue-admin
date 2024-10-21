"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformDestructuredProps = exports.processPropsDestructure = void 0;
const estree_walker_1 = require("estree-walker");
const compiler_dom_1 = require("@vue/compiler-dom");
const shared_1 = require("@vue/shared");
const utils_1 = require("./utils");
const defineProps_1 = require("./defineProps");
const warn_1 = require("../warn");
function processPropsDestructure(ctx, declId) {
    if (!ctx.options.propsDestructure && !ctx.options.reactivityTransform) {
        return;
    }
    (0, warn_1.warnOnce)(`This project is using reactive props destructure, which is an experimental ` +
        `feature. It may receive breaking changes or be removed in the future, so ` +
        `use at your own risk.\n` +
        `To stay updated, follow the RFC at https://github.com/vuejs/rfcs/discussions/502.`);
    ctx.propsDestructureDecl = declId;
    const registerBinding = (key, local, defaultValue) => {
        ctx.propsDestructuredBindings[key] = { local, default: defaultValue };
        if (local !== key) {
            ctx.bindingMetadata[local] = compiler_dom_1.BindingTypes.PROPS_ALIASED;
            (ctx.bindingMetadata.__propsAliases ||
                (ctx.bindingMetadata.__propsAliases = {}))[local] = key;
        }
    };
    for (const prop of declId.properties) {
        if (prop.type === 'ObjectProperty') {
            const propKey = (0, utils_1.resolveObjectKey)(prop.key, prop.computed);
            if (!propKey) {
                ctx.error(`${defineProps_1.DEFINE_PROPS}() destructure cannot use computed key.`, prop.key);
            }
            if (prop.value.type === 'AssignmentPattern') {
                // default value { foo = 123 }
                const { left, right } = prop.value;
                if (left.type !== 'Identifier') {
                    ctx.error(`${defineProps_1.DEFINE_PROPS}() destructure does not support nested patterns.`, left);
                }
                registerBinding(propKey, left.name, right);
            }
            else if (prop.value.type === 'Identifier') {
                // simple destructure
                registerBinding(propKey, prop.value.name);
            }
            else {
                ctx.error(`${defineProps_1.DEFINE_PROPS}() destructure does not support nested patterns.`, prop.value);
            }
        }
        else {
            // rest spread
            ctx.propsDestructureRestId = prop.argument.name;
            // register binding
            ctx.bindingMetadata[ctx.propsDestructureRestId] =
                compiler_dom_1.BindingTypes.SETUP_REACTIVE_CONST;
        }
    }
}
exports.processPropsDestructure = processPropsDestructure;
function transformDestructuredProps(ctx, vueImportAliases) {
    if (!ctx.options.propsDestructure && !ctx.options.reactivityTransform) {
        return;
    }
    const rootScope = {};
    const scopeStack = [rootScope];
    let currentScope = rootScope;
    const excludedIds = new WeakSet();
    const parentStack = [];
    const propsLocalToPublicMap = Object.create(null);
    for (const key in ctx.propsDestructuredBindings) {
        const { local } = ctx.propsDestructuredBindings[key];
        rootScope[local] = true;
        propsLocalToPublicMap[local] = key;
    }
    function pushScope() {
        scopeStack.push((currentScope = Object.create(currentScope)));
    }
    function popScope() {
        scopeStack.pop();
        currentScope = scopeStack[scopeStack.length - 1] || null;
    }
    function registerLocalBinding(id) {
        excludedIds.add(id);
        if (currentScope) {
            currentScope[id.name] = false;
        }
        else {
            ctx.error('registerBinding called without active scope, something is wrong.', id);
        }
    }
    function walkScope(node, isRoot = false) {
        for (const stmt of node.body) {
            if (stmt.type === 'VariableDeclaration') {
                walkVariableDeclaration(stmt, isRoot);
            }
            else if (stmt.type === 'FunctionDeclaration' ||
                stmt.type === 'ClassDeclaration') {
                if (stmt.declare || !stmt.id)
                    continue;
                registerLocalBinding(stmt.id);
            }
            else if ((stmt.type === 'ForOfStatement' || stmt.type === 'ForInStatement') &&
                stmt.left.type === 'VariableDeclaration') {
                walkVariableDeclaration(stmt.left);
            }
            else if (stmt.type === 'ExportNamedDeclaration' &&
                stmt.declaration &&
                stmt.declaration.type === 'VariableDeclaration') {
                walkVariableDeclaration(stmt.declaration, isRoot);
            }
            else if (stmt.type === 'LabeledStatement' &&
                stmt.body.type === 'VariableDeclaration') {
                walkVariableDeclaration(stmt.body, isRoot);
            }
        }
    }
    function walkVariableDeclaration(stmt, isRoot = false) {
        if (stmt.declare) {
            return;
        }
        for (const decl of stmt.declarations) {
            const isDefineProps = isRoot && decl.init && (0, utils_1.isCallOf)((0, utils_1.unwrapTSNode)(decl.init), 'defineProps');
            for (const id of (0, compiler_dom_1.extractIdentifiers)(decl.id)) {
                if (isDefineProps) {
                    // for defineProps destructure, only exclude them since they
                    // are already passed in as knownProps
                    excludedIds.add(id);
                }
                else {
                    registerLocalBinding(id);
                }
            }
        }
    }
    function rewriteId(id, parent, parentStack) {
        if ((parent.type === 'AssignmentExpression' && id === parent.left) ||
            parent.type === 'UpdateExpression') {
            ctx.error(`Cannot assign to destructured props as they are readonly.`, id);
        }
        if ((0, compiler_dom_1.isStaticProperty)(parent) && parent.shorthand) {
            // let binding used in a property shorthand
            // skip for destructure patterns
            if (!parent.inPattern ||
                (0, compiler_dom_1.isInDestructureAssignment)(parent, parentStack)) {
                // { prop } -> { prop: __props.prop }
                ctx.s.appendLeft(id.end + ctx.startOffset, `: ${(0, shared_1.genPropsAccessExp)(propsLocalToPublicMap[id.name])}`);
            }
        }
        else {
            // x --> __props.x
            ctx.s.overwrite(id.start + ctx.startOffset, id.end + ctx.startOffset, (0, shared_1.genPropsAccessExp)(propsLocalToPublicMap[id.name]));
        }
    }
    function checkUsage(node, method, alias = method) {
        if ((0, utils_1.isCallOf)(node, alias)) {
            const arg = (0, utils_1.unwrapTSNode)(node.arguments[0]);
            if (arg.type === 'Identifier' && currentScope[arg.name]) {
                ctx.error(`"${arg.name}" is a destructured prop and should not be passed directly to ${method}(). ` +
                    `Pass a getter () => ${arg.name} instead.`, arg);
            }
        }
    }
    // check root scope first
    const ast = ctx.scriptSetupAst;
    walkScope(ast, true);
    (0, estree_walker_1.walk)(ast, {
        enter(node, parent) {
            parent && parentStack.push(parent);
            // skip type nodes
            if (parent &&
                parent.type.startsWith('TS') &&
                parent.type !== 'TSAsExpression' &&
                parent.type !== 'TSNonNullExpression' &&
                parent.type !== 'TSTypeAssertion') {
                return this.skip();
            }
            checkUsage(node, 'watch', vueImportAliases.watch);
            checkUsage(node, 'toRef', vueImportAliases.toRef);
            // function scopes
            if ((0, compiler_dom_1.isFunctionType)(node)) {
                pushScope();
                (0, compiler_dom_1.walkFunctionParams)(node, registerLocalBinding);
                if (node.body.type === 'BlockStatement') {
                    walkScope(node.body);
                }
                return;
            }
            // catch param
            if (node.type === 'CatchClause') {
                pushScope();
                if (node.param && node.param.type === 'Identifier') {
                    registerLocalBinding(node.param);
                }
                walkScope(node.body);
                return;
            }
            // non-function block scopes
            if (node.type === 'BlockStatement' && !(0, compiler_dom_1.isFunctionType)(parent)) {
                pushScope();
                walkScope(node);
                return;
            }
            if (node.type === 'Identifier') {
                if ((0, compiler_dom_1.isReferencedIdentifier)(node, parent, parentStack) &&
                    !excludedIds.has(node)) {
                    if (currentScope[node.name]) {
                        rewriteId(node, parent, parentStack);
                    }
                }
            }
        },
        leave(node, parent) {
            parent && parentStack.pop();
            if ((node.type === 'BlockStatement' && !(0, compiler_dom_1.isFunctionType)(parent)) ||
                (0, compiler_dom_1.isFunctionType)(node)) {
                popScope();
            }
        },
    });
}
exports.transformDestructuredProps = transformDestructuredProps;
