"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.hasDefaultExport = exports.rewriteDefaultAST = exports.rewriteDefault = void 0;
const parser_1 = require("@babel/parser");
const magic_string_1 = __importDefault(require("magic-string"));
function rewriteDefault(input, as, define, parserPlugins) {
    const ast = (0, parser_1.parse)(input, {
        sourceType: 'module',
        plugins: parserPlugins,
    }).program.body;
    const s = new magic_string_1.default(input);
    rewriteDefaultAST(ast, s, as, define);
    return s.toString();
}
exports.rewriteDefault = rewriteDefault;
/**
 * Utility for rewriting `export default` in a script block into a variable
 * declaration so that we can inject things into it
 */
function rewriteDefaultAST(ast, s, as, define) {
    if (!hasDefaultExport(ast)) {
        s.append(`\nconst ${as} = ${define}({})`);
        return;
    }
    // if the script somehow still contains `default export`, it probably has
    // multi-line comments or template strings. fallback to a full parse.
    ast.forEach((node) => {
        if (node.type === 'ExportDefaultDeclaration') {
            if (node.declaration.type === 'ClassDeclaration' && node.declaration.id) {
                let start = node.declaration.decorators && node.declaration.decorators.length > 0
                    ? node.declaration.decorators[node.declaration.decorators.length - 1].end
                    : node.start;
                s.overwrite(start, node.declaration.id.start, ` class `);
                s.append(`\nconst ${as} = ${define}(${node.declaration.id.name})`);
            }
            else {
                s.overwrite(node.start, node.declaration.start, `const ${as} = ${define}(`);
                s.appendRight(node.declaration.end, ')');
            }
        }
        else if (node.type === 'ExportNamedDeclaration') {
            for (const specifier of node.specifiers) {
                if (specifier.type === 'ExportSpecifier' &&
                    specifier.exported.type === 'Identifier' &&
                    specifier.exported.name === 'default') {
                    if (node.source) {
                        if (specifier.local.name === 'default') {
                            s.prepend(`import { default as __VUE_DEFAULT__ } from '${node.source.value}'\n`);
                            const end = specifierEnd(s, specifier.local.end, node.end);
                            s.remove(specifier.start, end);
                            s.append(`\nconst ${as} = ${define}(__VUE_DEFAULT___)`);
                            continue;
                        }
                        else {
                            s.prepend(`import { ${s.slice(specifier.local.start, specifier.local.end)} as __VUE_DEFAULT__ } from '${node.source.value}'\n`);
                            const end = specifierEnd(s, specifier.exported.end, node.end);
                            s.remove(specifier.start, end);
                            s.append(`\nconst ${as} = ${define}(__VUE_DEFAULT__)`);
                            continue;
                        }
                    }
                    const end = specifierEnd(s, specifier.end, node.end);
                    s.remove(specifier.start, end);
                    s.append(`\nconst ${as} = ${specifier.local.name}`);
                }
            }
        }
    });
}
exports.rewriteDefaultAST = rewriteDefaultAST;
function hasDefaultExport(ast) {
    for (const stmt of ast) {
        if (stmt.type === 'ExportDefaultDeclaration') {
            return true;
        }
        else if (stmt.type === 'ExportNamedDeclaration' &&
            stmt.specifiers.some((spec) => spec.exported.name === 'default')) {
            return true;
        }
    }
    return false;
}
exports.hasDefaultExport = hasDefaultExport;
function specifierEnd(s, end, nodeEnd) {
    // export { default   , foo } ...
    let hasCommas = false;
    let oldEnd = end;
    while (end < nodeEnd) {
        if (/\s/.test(s.slice(end, end + 1))) {
            end++;
        }
        else if (s.slice(end, end + 1) === ',') {
            end++;
            hasCommas = true;
            break;
        }
        else if (s.slice(end, end + 1) === '}') {
            break;
        }
    }
    return hasCommas ? end : oldEnd;
}
