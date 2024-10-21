"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformDynamicImports = exports.parseScriptDescriptor = exports.parseTemplateDescriptor = exports.updateMiniProgramComponentsByMainFilename = exports.updateMiniProgramGlobalComponents = exports.updateMiniProgramComponentsByTemplateFilename = exports.updateMiniProgramComponentsByScriptFilename = exports.parseMainDescriptor = void 0;
const types_1 = require("@babel/types");
const estree_walker_1 = require("estree-walker");
const magic_string_1 = __importDefault(require("magic-string"));
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const messages_1 = require("../messages");
const constants_1 = require("../constants");
const utils_1 = require("../utils");
const utils_2 = require("../vite/utils");
const jsonFile_1 = require("../json/mp/jsonFile");
const mainDescriptors = new Map();
const scriptDescriptors = new Map();
const templateDescriptors = new Map();
function findImportTemplateSource(ast) {
    const importDeclaration = ast.body.find((node) => (0, types_1.isImportDeclaration)(node) &&
        node.source.value.includes('vue&type=template'));
    if (importDeclaration) {
        return importDeclaration.source.value;
    }
}
function findImportScriptSource(ast) {
    const importDeclaration = ast.body.find((node) => (0, types_1.isImportDeclaration)(node) && node.source.value.includes('vue&type=script'));
    if (importDeclaration) {
        return importDeclaration.source.value;
    }
}
async function resolveSource(filename, source, resolve) {
    if (!source) {
        return;
    }
    const resolveId = await resolve(source, filename);
    if (resolveId) {
        return resolveId.id;
    }
}
async function parseMainDescriptor(filename, ast, resolve) {
    const script = await resolveSource(filename, findImportScriptSource(ast), resolve);
    const template = await resolveSource(filename, findImportTemplateSource(ast), resolve);
    const imports = await parseVueComponentImports(filename, ast.body.filter((node) => (0, types_1.isImportDeclaration)(node)), resolve);
    if (!script) {
        // inline script
        await parseScriptDescriptor(filename, ast, { resolve, isExternal: false });
    }
    if (!template) {
        // inline template
        await parseTemplateDescriptor(filename, ast, { resolve, isExternal: false });
    }
    const descriptor = {
        imports,
        script: script ? (0, utils_2.parseVueRequest)(script).filename : filename,
        template: template ? (0, utils_2.parseVueRequest)(template).filename : filename,
    };
    mainDescriptors.set(filename, descriptor);
    return descriptor;
}
exports.parseMainDescriptor = parseMainDescriptor;
function updateMiniProgramComponentsByScriptFilename(scriptFilename, inputDir, normalizeComponentName) {
    const mainFilename = findMainFilenameByScriptFilename(scriptFilename);
    if (mainFilename) {
        updateMiniProgramComponentsByMainFilename(mainFilename, inputDir, normalizeComponentName);
    }
}
exports.updateMiniProgramComponentsByScriptFilename = updateMiniProgramComponentsByScriptFilename;
function updateMiniProgramComponentsByTemplateFilename(templateFilename, inputDir, normalizeComponentName) {
    const mainFilename = findMainFilenameByTemplateFilename(templateFilename);
    if (mainFilename) {
        updateMiniProgramComponentsByMainFilename(mainFilename, inputDir, normalizeComponentName);
    }
}
exports.updateMiniProgramComponentsByTemplateFilename = updateMiniProgramComponentsByTemplateFilename;
function findMainFilenameByScriptFilename(scriptFilename) {
    const keys = [...mainDescriptors.keys()];
    return keys.find((key) => mainDescriptors.get(key).script === scriptFilename);
}
function findMainFilenameByTemplateFilename(templateFilename) {
    const keys = [...mainDescriptors.keys()];
    return keys.find((key) => mainDescriptors.get(key).template === templateFilename);
}
async function updateMiniProgramGlobalComponents(filename, ast, { inputDir, resolve, normalizeComponentName, }) {
    const { bindingComponents, imports } = await parseGlobalDescriptor(filename, ast, resolve);
    (0, jsonFile_1.addMiniProgramUsingComponents)('app', createUsingComponents(bindingComponents, imports, inputDir, normalizeComponentName));
    return {
        imports,
    };
}
exports.updateMiniProgramGlobalComponents = updateMiniProgramGlobalComponents;
function createUsingComponents(bindingComponents, imports, inputDir, normalizeComponentName) {
    const usingComponents = {};
    imports.forEach(({ source: { value }, specifiers: [specifier] }) => {
        const { name } = specifier.local;
        if (!bindingComponents[name]) {
            return;
        }
        const componentName = normalizeComponentName((0, shared_1.hyphenate)(bindingComponents[name].tag));
        if (!usingComponents[componentName]) {
            usingComponents[componentName] = (0, uni_shared_1.addLeadingSlash)((0, utils_1.removeExt)((0, utils_1.normalizeMiniProgramFilename)(value, inputDir)));
        }
    });
    return usingComponents;
}
function updateMiniProgramComponentsByMainFilename(mainFilename, inputDir, normalizeComponentName) {
    const mainDescriptor = mainDescriptors.get(mainFilename);
    if (!mainDescriptor) {
        return;
    }
    const templateDescriptor = templateDescriptors.get(mainDescriptor.template);
    if (!templateDescriptor) {
        return;
    }
    const scriptDescriptor = scriptDescriptors.get(mainDescriptor.script);
    if (!scriptDescriptor) {
        return;
    }
    const bindingComponents = parseBindingComponents({
        ...templateDescriptor.bindingComponents,
        ...scriptDescriptor.setupBindingComponents,
    }, scriptDescriptor.bindingComponents);
    const imports = parseImports(mainDescriptor.imports, scriptDescriptor.imports, templateDescriptor.imports);
    (0, jsonFile_1.addMiniProgramUsingComponents)((0, utils_1.removeExt)((0, utils_1.normalizeMiniProgramFilename)(mainFilename, inputDir)), createUsingComponents(bindingComponents, imports, inputDir, normalizeComponentName));
}
exports.updateMiniProgramComponentsByMainFilename = updateMiniProgramComponentsByMainFilename;
function findBindingComponent(tag, bindingComponents) {
    return Object.keys(bindingComponents).find((name) => {
        const componentTag = bindingComponents[name].tag;
        const camelName = (0, shared_1.camelize)(componentTag);
        const PascalName = (0, shared_1.capitalize)(camelName);
        return tag === componentTag || tag === camelName || tag === PascalName;
    });
}
function normalizeComponentId(id) {
    // _unref(test) => test
    if (id.includes('_unref(')) {
        return id.replace('_unref(', '').replace(')', '');
    }
    // $setup["test"] => test
    if (id.includes('$setup[')) {
        return id.replace('$setup["', '').replace('"', '');
    }
    return id;
}
function parseBindingComponents(templateBindingComponents, scriptBindingComponents) {
    const bindingComponents = {};
    Object.keys(templateBindingComponents).forEach((id) => {
        bindingComponents[normalizeComponentId(id)] = templateBindingComponents[id];
    });
    Object.keys(scriptBindingComponents).forEach((id) => {
        const { tag } = scriptBindingComponents[id];
        const name = findBindingComponent(tag, templateBindingComponents);
        if (name) {
            bindingComponents[id] = bindingComponents[name];
        }
    });
    return bindingComponents;
}
function parseImports(mainImports, scriptImports, templateImports) {
    const imports = [...mainImports, ...templateImports, ...scriptImports];
    return imports;
}
/**
 * 解析 template
 * @param filename
 * @param code
 * @param ast
 * @param options
 * @returns
 */
async function parseTemplateDescriptor(filename, ast, options) {
    // 外置时查找所有 vue component import
    const imports = options.isExternal
        ? await parseVueComponentImports(filename, ast.body.filter((node) => (0, types_1.isImportDeclaration)(node)), options.resolve)
        : [];
    const descriptor = {
        bindingComponents: findBindingComponents(ast.body),
        imports,
    };
    templateDescriptors.set(filename, descriptor);
    return descriptor;
}
exports.parseTemplateDescriptor = parseTemplateDescriptor;
async function parseGlobalDescriptor(filename, ast, resolve) {
    // 外置时查找所有 vue component import
    const imports = (await parseVueComponentImports(filename, ast.body.filter((node) => (0, types_1.isImportDeclaration)(node)), resolve)).filter((item) => !(0, utils_1.isAppVue)((0, utils_2.cleanUrl)(item.source.value)));
    return {
        bindingComponents: parseGlobalComponents(ast),
        imports,
    };
}
/**
 * 解析 script
 * @param filename
 * @param code
 * @param ast
 * @param options
 * @returns
 */
async function parseScriptDescriptor(filename, ast, options) {
    // 外置时查找所有 vue component import
    const imports = options.isExternal
        ? await parseVueComponentImports(filename, ast.body.filter((node) => (0, types_1.isImportDeclaration)(node)), options.resolve)
        : [];
    const descriptor = {
        bindingComponents: parseComponents(ast),
        setupBindingComponents: findBindingComponents(ast.body),
        imports,
    };
    scriptDescriptors.set(filename, descriptor);
    return descriptor;
}
exports.parseScriptDescriptor = parseScriptDescriptor;
/**
 * 解析编译器生成的 bindingComponents
 * @param ast
 * @returns
 */
function findBindingComponents(ast) {
    const mapping = findUnpluginComponents(ast);
    for (const node of ast) {
        if (!(0, types_1.isVariableDeclaration)(node)) {
            continue;
        }
        const declarator = node.declarations[0];
        if ((0, types_1.isIdentifier)(declarator.id) &&
            declarator.id.name === constants_1.BINDING_COMPONENTS) {
            const bindingComponents = JSON.parse(declarator.init.value);
            return Object.keys(bindingComponents).reduce((bindings, tag) => {
                const { name, type } = bindingComponents[tag];
                bindings[mapping[name] || name] = {
                    tag,
                    type: type,
                };
                return bindings;
            }, {});
        }
    }
    return {};
}
/**
 * 兼容：unplugin_components
 * https://github.com/dcloudio/uni-app/issues/3057
 * @param ast
 * @returns
 */
function findUnpluginComponents(ast) {
    const res = Object.create(null);
    // if(!Array){}
    const ifStatement = ast.find((statement) => (0, types_1.isIfStatement)(statement) &&
        (0, types_1.isUnaryExpression)(statement.test) &&
        statement.test.operator === '!' &&
        (0, types_1.isIdentifier)(statement.test.argument) &&
        statement.test.argument.name === 'Array');
    if (!ifStatement) {
        return res;
    }
    if (!(0, types_1.isBlockStatement)(ifStatement.consequent)) {
        return res;
    }
    for (const node of ifStatement.consequent.body) {
        if (!(0, types_1.isVariableDeclaration)(node)) {
            continue;
        }
        const { id, init } = node.declarations[0];
        if ((0, types_1.isIdentifier)(id) &&
            (0, types_1.isIdentifier)(init) &&
            init.name.includes('unplugin_components')) {
            res[id.name] = init.name;
        }
    }
    return res;
}
/**
 * 查找全局组件定义：app.component('component-a',{})
 * @param ast
 * @returns
 */
function parseGlobalComponents(ast) {
    const bindingComponents = {};
    estree_walker_1.walk(ast, {
        enter(child) {
            if (!(0, types_1.isCallExpression)(child)) {
                return;
            }
            const { callee } = child;
            // .component
            if (!(0, types_1.isMemberExpression)(callee) ||
                !(0, types_1.isIdentifier)(callee.property) ||
                callee.property.name !== 'component') {
                return;
            }
            // .component('component-a',{})
            const args = child.arguments;
            if (args.length !== 2) {
                return;
            }
            const [name, value] = args;
            if (!(0, types_1.isStringLiteral)(name)) {
                return console.warn(messages_1.M['mp.component.args[0]']);
            }
            if (!(0, types_1.isIdentifier)(value)) {
                return console.warn(messages_1.M['mp.component.args[1]']);
            }
            bindingComponents[value.name] = {
                tag: name.value,
                type: 'unknown',
            };
        },
    });
    return bindingComponents;
}
/**
 * 从 components 中查找定义的组件
 * @param ast
 * @param bindingComponents
 */
function parseComponents(ast) {
    const bindingComponents = {};
    estree_walker_1.walk(ast, {
        enter(child) {
            if (!(0, types_1.isObjectExpression)(child)) {
                return;
            }
            const componentsProp = child.properties.find((prop) => (0, types_1.isObjectProperty)(prop) &&
                (0, types_1.isIdentifier)(prop.key) &&
                prop.key.name === 'components');
            if (!componentsProp) {
                return;
            }
            const componentsExpr = componentsProp.value;
            if (!(0, types_1.isObjectExpression)(componentsExpr)) {
                return;
            }
            componentsExpr.properties.forEach((prop) => {
                if (!(0, types_1.isObjectProperty)(prop)) {
                    return;
                }
                if (!(0, types_1.isIdentifier)(prop.key) && !(0, types_1.isStringLiteral)(prop.key)) {
                    return;
                }
                if (!(0, types_1.isIdentifier)(prop.value)) {
                    return;
                }
                bindingComponents[prop.value.name] = {
                    tag: (0, types_1.isIdentifier)(prop.key) ? prop.key.name : prop.key.value,
                    type: 'unknown',
                };
            });
        },
    });
    return bindingComponents;
}
/**
 * vue component imports
 * @param filename
 * @param imports
 * @param resolve
 * @returns
 */
async function parseVueComponentImports(importer, imports, resolve) {
    const vueComponentImports = [];
    for (let i = 0; i < imports.length; i++) {
        const { source } = imports[i];
        if ((0, utils_2.parseVueRequest)(source.value).query.vue) {
            continue;
        }
        const resolveId = await resolve(source.value, importer);
        if (!resolveId) {
            continue;
        }
        const { filename } = (0, utils_2.parseVueRequest)(resolveId.id);
        if (constants_1.EXTNAME_VUE_RE.test(filename)) {
            source.value = resolveId.id;
            vueComponentImports.push(imports[i]);
        }
    }
    return vueComponentImports;
}
/**
 * static import => dynamic import
 * @param code
 * @param imports
 * @param dynamicImport
 * @returns
 */
async function transformDynamicImports(code, imports, { id, sourceMap, dynamicImport, }) {
    if (!imports.length) {
        return {
            code,
            map: null,
        };
    }
    const s = new magic_string_1.default(code);
    for (let i = 0; i < imports.length; i++) {
        const { start, end, specifiers: [specifier], source, } = imports[i];
        s.overwrite(start, end, dynamicImport(specifier.local.name, source.value) + ';');
    }
    return {
        code: s.toString(),
        map: null,
    };
}
exports.transformDynamicImports = transformDynamicImports;
