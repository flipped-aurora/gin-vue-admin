"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.uniViteInjectPlugin = void 0;
const path_1 = require("path");
const debug_1 = __importDefault(require("debug"));
const pluginutils_1 = require("@rollup/pluginutils");
const estree_walker_1 = require("estree-walker");
const shared_1 = require("@vue/shared");
const magic_string_1 = __importDefault(require("magic-string"));
const utils_1 = require("../utils");
const debugInject = (0, debug_1.default)('uni:inject');
// const debugInjectTry = debug('uni:inject-try')
function uniViteInjectPlugin(name, options) {
    if (!options)
        throw new Error('Missing options');
    const filter = (0, pluginutils_1.createFilter)(options.include, options.exclude);
    const modules = (0, shared_1.extend)({}, options);
    delete modules.include;
    delete modules.exclude;
    delete modules.sourceMap;
    delete modules.callback;
    const reassignments = new Set();
    const modulesMap = new Map();
    const namespaceModulesMap = new Map();
    Object.keys(modules).forEach((name) => {
        if (name.endsWith('.')) {
            namespaceModulesMap.set(name, modules[name]);
        }
        modulesMap.set(name, modules[name]);
    });
    const hasNamespace = namespaceModulesMap.size > 0;
    // Fix paths on Windows
    if (path_1.sep !== '/') {
        normalizeModulesMap(modulesMap);
        normalizeModulesMap(namespaceModulesMap);
    }
    const firstpass = new RegExp(`(?:${Array.from(modulesMap.keys()).map(escape).join('|')})`, 'g');
    const sourceMap = options.sourceMap !== false;
    const callback = options.callback;
    return {
        name,
        // 确保在 commonjs 之后，否则会混合 es6 module 与 cjs 的代码，导致 commonjs 失效
        enforce: options.enforce ?? 'post',
        transform(code, id) {
            if (!filter(id))
                return null;
            if (!(0, utils_1.isJsFile)(id))
                return null;
            // debugInjectTry(id)
            if (code.search(firstpass) === -1)
                return null;
            if (path_1.sep !== '/')
                id = id.split(path_1.sep).join('/');
            const ast = this.parse(code);
            const imports = new Set();
            ast.body.forEach((node) => {
                if (node.type === 'ImportDeclaration') {
                    node.specifiers.forEach((specifier) => {
                        imports.add(specifier.local.name);
                    });
                }
            });
            // analyse scopes
            let scope = (0, pluginutils_1.attachScopes)(ast, 'scope');
            const magicString = new magic_string_1.default(code);
            const newImports = new Map();
            function handleReference(node, name, keypath, parent) {
                let mod = modulesMap.get(keypath);
                if (!mod && hasNamespace) {
                    const mods = keypath.split('.');
                    if (mods.length === 2) {
                        mod = namespaceModulesMap.get(mods[0] + '.');
                        if (mod) {
                            if ((0, shared_1.isArray)(mod)) {
                                const testFn = mod[1];
                                if (testFn(mods[1])) {
                                    mod = [mod[0], mods[1]];
                                }
                                else {
                                    mod = undefined;
                                }
                            }
                            else {
                                mod = [mod, mods[1]];
                            }
                        }
                    }
                }
                if (mod && !imports.has(name) && !scope.contains(name)) {
                    if ((0, shared_1.isString)(mod))
                        mod = [mod, 'default'];
                    if (mod[0] === id)
                        return false;
                    const hash = `${keypath}:${mod[0]}:${mod[1]}`;
                    // 当 API 被覆盖定义后，不再摇树
                    if (reassignments.has(hash)) {
                        return false;
                    }
                    if (parent &&
                        (0, utils_1.isAssignmentExpression)(parent) &&
                        parent.left === node) {
                        reassignments.add(hash);
                        return false;
                    }
                    const importLocalName = name === keypath ? name : (0, pluginutils_1.makeLegalIdentifier)(`$inject_${keypath}`);
                    if (!newImports.has(hash)) {
                        if (mod[1] === '*') {
                            newImports.set(hash, `import * as ${importLocalName} from '${mod[0]}';`);
                        }
                        else {
                            newImports.set(hash, `import { ${mod[1]} as ${importLocalName} } from '${mod[0]}';`);
                            callback && callback(newImports, mod);
                        }
                    }
                    if (name !== keypath) {
                        magicString.overwrite(node.start, node.end, importLocalName, {
                            storeName: true,
                        });
                    }
                    return true;
                }
                return false;
            }
            (0, estree_walker_1.walk)(ast, {
                enter(node, parent) {
                    if (sourceMap) {
                        magicString.addSourcemapLocation(node.start);
                        magicString.addSourcemapLocation(node.end);
                    }
                    if (node.scope) {
                        scope = node.scope;
                    }
                    if ((0, utils_1.isProperty)(node) && node.shorthand) {
                        const { name } = node.key;
                        handleReference(node, name, name);
                        this.skip();
                        return;
                    }
                    if ((0, utils_1.isReference)(node, parent)) {
                        const { name, keypath } = flatten(node);
                        const handled = handleReference(node, name, keypath, parent);
                        if (handled) {
                            this.skip();
                        }
                    }
                },
                leave(node) {
                    if (node.scope) {
                        scope = scope.parent;
                    }
                },
            });
            debugInject(id, newImports.size);
            if (newImports.size === 0) {
                return {
                    code,
                    // 不能返回 ast ，否则会导致代码不能被再次修改
                    // 比如 App.vue 中，console.log('uniCloud') 触发了 inject 检测，检测完，发现不需要
                    // 此时返回 ast，会导致 import { setupApp } from '@dcloudio/uni-h5' 不会被编译
                    // ast
                    map: null,
                };
            }
            const importBlock = Array.from(newImports.values()).join('\n\n');
            magicString.prepend(`${importBlock}\n\n`);
            return {
                code: magicString.toString(),
                map: sourceMap ? magicString.generateMap({ hires: true }) : null,
            };
        },
    };
}
exports.uniViteInjectPlugin = uniViteInjectPlugin;
const escape = (str) => str.replace(/[-[\]/{}()*+?.\\^$|]/g, '\\$&');
const flatten = (startNode) => {
    const parts = [];
    let node = startNode;
    while ((0, utils_1.isMemberExpression)(node)) {
        parts.unshift(node.property.name);
        node = node.object;
    }
    const { name } = node;
    parts.unshift(name);
    return { name, keypath: parts.join('.') };
};
function normalizeModulesMap(modulesMap) {
    modulesMap.forEach((mod, key) => {
        modulesMap.set(key, (0, shared_1.isArray)(mod)
            ? [mod[0].split(path_1.sep).join('/'), mod[1]]
            : mod.split(path_1.sep).join('/'));
    });
}
