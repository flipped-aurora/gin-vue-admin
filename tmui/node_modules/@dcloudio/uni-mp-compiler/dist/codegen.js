"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.genExpr = exports.genBabelExpr = exports.generate = void 0;
const shared_1 = require("@vue/shared");
const compiler_core_1 = require("@vue/compiler-core");
const generator_1 = __importDefault(require("@babel/generator"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
function generate(ast, options) {
    const context = createCodegenContext(ast, options);
    const { mode, push, indent, deindent, newline, prefixIdentifiers } = context;
    const helpers = Array.from(ast.helpers);
    const hasHelpers = helpers.length > 0;
    const useWithBlock = !prefixIdentifiers && mode !== 'module';
    const isSetupInlined = !!options.inline;
    // preambles
    // in setup() inline mode, the preamble is generated in a sub context
    // and returned separately.
    const preambleContext = isSetupInlined
        ? createCodegenContext(ast, options)
        : context;
    if (mode === 'module') {
        genModulePreamble(ast, preambleContext, isSetupInlined);
    }
    else {
        genFunctionPreamble(ast, preambleContext);
    }
    // enter render function
    const functionName = `render`;
    const args = ['_ctx', '_cache'];
    if (options.bindingMetadata && !options.inline) {
        // binding optimization args
        args.push('$props', '$setup', '$data', '$options');
    }
    const signature = options.isTS
        ? args.map((arg) => `${arg}: any`).join(',')
        : args.join(', ');
    if (isSetupInlined) {
        push(`(${signature}) => {`);
    }
    else {
        push(`function ${functionName}(${signature}) {`);
    }
    indent();
    if (useWithBlock) {
        push(`with (_ctx) {`);
        indent();
        if (hasHelpers) {
            push(`const { ${helpers
                .map((s) => `${compiler_core_1.helperNameMap[s]}: _${compiler_core_1.helperNameMap[s]}`)
                .join(', ')} } = _Vue`);
            push(`\n`);
            newline();
        }
    }
    push(`return `);
    push(genBabelExpr(ast.renderData, options.generatorOpts));
    if (useWithBlock) {
        deindent();
        push(`}`);
    }
    deindent();
    push(`}`);
    return {
        code: context.code,
        preamble: isSetupInlined ? preambleContext.code : ``,
        // SourceMapGenerator does have toJSON() method but it's not in the types
        map: context.map ? context.map.toJSON() : undefined,
    };
}
exports.generate = generate;
function createCodegenContext(ast, { mode = 'function', prefixIdentifiers = mode === 'module', filename = `template.vue.html`, scopeId = null, runtimeGlobalName = `Vue`, runtimeModuleName = `vue`, isTS = false, sourceMap = false, }) {
    const context = {
        mode,
        prefixIdentifiers,
        filename,
        scopeId,
        runtimeGlobalName,
        runtimeModuleName,
        bindingComponents: ast.bindingComponents,
        isTS,
        source: ast.loc.source,
        code: ``,
        column: 1,
        line: 1,
        offset: 0,
        indentLevel: 0,
        push(code, node) {
            context.code += code;
            if (context.map) {
                if (node) {
                    let name;
                    if (node.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION && !node.isStatic) {
                        const content = node.content.replace(/^_ctx\./, '');
                        if (content !== node.content && (0, compiler_core_1.isSimpleIdentifier)(content)) {
                            name = content;
                        }
                    }
                    addMapping(node.loc.start, name);
                }
                (0, compiler_core_1.advancePositionWithMutation)(context, code);
                if (node && node.loc !== compiler_core_1.locStub) {
                    addMapping(node.loc.end);
                }
            }
        },
        indent() {
            newline(++context.indentLevel);
        },
        deindent(withoutNewLine = false) {
            if (withoutNewLine) {
                --context.indentLevel;
            }
            else {
                newline(--context.indentLevel);
            }
        },
        newline() {
            newline(context.indentLevel);
        },
    };
    function newline(n) {
        context.push('\n' + `  `.repeat(n));
    }
    function addMapping(loc, name) {
        context.map.addMapping({
            name,
            source: context.filename || '',
            original: {
                line: loc.line,
                column: loc.column - 1, // source-map column is 0 based
            },
            generated: {
                line: context.line,
                column: context.column - 1,
            },
        });
    }
    // 暂时无需提供 sourcemap 支持
    // if (sourceMap) {
    //   // lazy require source-map implementation
    //   context.map = new SourceMapGenerator()
    //   context.map!.setSourceContent(filename, context.source)
    // }
    return context;
}
function genComponentImports(bindingComponents, { push, newline }) {
    const tags = Object.keys(bindingComponents);
    const importDeclarations = [];
    // 仅记录easycom和setup组件
    const components = [];
    tags.forEach((tag) => {
        const { name, type } = bindingComponents[tag];
        if (type === "unknown" /* BindingComponentTypes.UNKNOWN */) {
            const source = (0, uni_cli_shared_1.matchEasycom)(tag);
            if (source) {
                // 调整为easycom命名
                const easycomName = name.replace('component', 'easycom');
                bindingComponents[tag].name = easycomName;
                components.push(easycomName);
                (0, uni_cli_shared_1.addImportDeclaration)(importDeclarations, easycomName, source);
            }
        }
        else if (type === "setup" /* BindingComponentTypes.SETUP */) {
            components.push(name);
        }
    });
    if (tags.length) {
        push(`const __BINDING_COMPONENTS__ = '` +
            JSON.stringify(bindingComponents) +
            `'`);
        const resolveComponents = [];
        const names = [];
        Object.keys(bindingComponents).forEach((id) => {
            const { type, name } = bindingComponents[id];
            if (type === "unknown" /* BindingComponentTypes.UNKNOWN */) {
                resolveComponents.push(`const ${name} = _${compiler_core_1.helperNameMap[compiler_core_1.RESOLVE_COMPONENT]}("${id}");`);
                names.push(name);
            }
        });
        if (resolveComponents.length) {
            newline();
            push(`if (!Array) {`);
            resolveComponents.forEach((code) => {
                push(code);
            });
            push(`(${names.join('+')})()`);
            push(`}`);
        }
        newline();
        importDeclarations.forEach((str) => push(str));
        if (importDeclarations.length) {
            newline();
        }
        if (components.length) {
            push(`if (!Math) {`);
            push(` (${components.map((name) => name).join('+')})() `);
            push(`}`);
            newline();
        }
    }
}
function genFunctionPreamble(ast, context) {
    const { prefixIdentifiers, push, newline, runtimeGlobalName, bindingComponents, } = context;
    const VueBinding = runtimeGlobalName;
    const aliasHelper = (s) => `${compiler_core_1.helperNameMap[s]}: _${compiler_core_1.helperNameMap[s]}`;
    const helpers = Array.from(ast.helpers);
    if (helpers.length > 0) {
        if (prefixIdentifiers) {
            push(`const { ${helpers.map(aliasHelper).join(', ')} } = ${VueBinding}\n`);
        }
        else {
            push(`const _Vue = ${VueBinding}\n`);
        }
    }
    genComponentImports(bindingComponents, context);
    newline();
    push(`return `);
}
function genModulePreamble(ast, context, inline) {
    const { push, newline, runtimeModuleName, bindingComponents } = context;
    const helpers = Array.from(ast.helpers);
    if (helpers.length) {
        push(`import { ${helpers
            .map((s) => `${compiler_core_1.helperNameMap[s]} as _${compiler_core_1.helperNameMap[s]}`)
            .join(', ')} } from ${JSON.stringify(runtimeModuleName)}\n`);
    }
    if (ast.imports.length) {
        genImports(ast.imports, context);
    }
    genComponentImports(bindingComponents, context);
    newline();
    if (!inline) {
        push(`export `);
    }
}
function genImports(importsOptions, { push, newline }) {
    if (!importsOptions.length) {
        return;
    }
    importsOptions.forEach((imports) => {
        push(`import `);
        push(genExpr(imports.exp));
        push(` from '${imports.path}'`);
        newline();
    });
}
function createGenNodeContext() {
    const context = {
        code: '',
        helper(key) {
            return `_${compiler_core_1.helperNameMap[key]}`;
        },
        push(code) {
            context.code += code;
        },
    };
    return context;
}
function genBabelExpr(expr, opts = {}) {
    if (!(0, shared_1.hasOwn)(opts, 'jsescOption')) {
        opts.jsescOption = {};
    }
    opts.jsescOption.quotes = 'single';
    return (0, generator_1.default)(expr, opts).code;
}
exports.genBabelExpr = genBabelExpr;
function genExpr(node, context) {
    return genNode(node, context).code;
}
exports.genExpr = genExpr;
function genNode(node, context) {
    if (!context) {
        context = createGenNodeContext();
    }
    if ((0, shared_1.isString)(node)) {
        context.push(node);
        return context;
    }
    if ((0, shared_1.isSymbol)(node)) {
        context.push(context.helper(node));
        return context;
    }
    switch (node.type) {
        case compiler_core_1.NodeTypes.TEXT:
            genText(node, context);
            break;
        case compiler_core_1.NodeTypes.SIMPLE_EXPRESSION:
            genExpression(node, context);
            break;
        case compiler_core_1.NodeTypes.INTERPOLATION:
            genInterpolation(node, context);
            break;
        case compiler_core_1.NodeTypes.COMPOUND_EXPRESSION:
            genCompoundExpression(node, context);
            break;
    }
    return context;
}
function genText(node, context) {
    context.push(JSON.stringify(node.content), node);
}
function genExpression(node, context) {
    const { content, isStatic } = node;
    context.push(isStatic ? JSON.stringify(content) : content, node);
}
function genInterpolation(node, context) {
    const { push, helper } = context;
    push(`${helper(compiler_core_1.TO_DISPLAY_STRING)}(`);
    genExpr(node.content, context);
    push(`)`);
}
function genCompoundExpression(node, context) {
    for (let i = 0; i < node.children.length; i++) {
        const child = node.children[i];
        if ((0, shared_1.isString)(child)) {
            context.push(child);
        }
        else {
            genExpr(child, context);
        }
    }
}
