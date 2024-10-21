"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.addEasyComponentAutoImports = exports.parseSource = exports.onCompilerError = exports.rewriteObjectExpression = exports.genRenderFunctionDecl = exports.isCompatEnabled = exports.__COMPAT__ = exports.__BROWSER__ = exports.__DEV__ = void 0;
const path_1 = __importDefault(require("path"));
const compiler_core_1 = require("@vue/compiler-core");
const estree_walker_1 = require("estree-walker");
const parser_1 = require("@babel/parser");
const magic_string_1 = __importDefault(require("magic-string"));
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const transformExpression_1 = require("./transforms/transformExpression");
exports.__DEV__ = true;
exports.__BROWSER__ = false;
exports.__COMPAT__ = false;
function isCompatEnabled(...args) {
    return false;
}
exports.isCompatEnabled = isCompatEnabled;
function genRenderFunctionDecl({ className, genDefaultAs, inline = false, }) {
    if (inline) {
        return `(): any | null =>`;
    }
    // 调整返回值类型为 any | null, 支持 <template>some text</template>
    const thisCode = genDefaultAs
        ? `this: InstanceType<typeof ${genDefaultAs}>`
        : '';
    return `function ${className || ''}Render(${thisCode}): any | null`;
}
exports.genRenderFunctionDecl = genRenderFunctionDecl;
function rewriteObjectExpression(exp, context) {
    const source = (0, transformExpression_1.stringifyExpression)(exp);
    if (source.includes('{')) {
        const s = new magic_string_1.default(source);
        const ast = (0, parser_1.parseExpression)(source, {
            plugins: context.expressionPlugins,
        });
        (0, estree_walker_1.walk)(ast, {
            enter(node) {
                if (node.type === 'ObjectExpression') {
                    s.prependLeft(node.start, node.properties.length > 0
                        ? 'utsMapOf('
                        : 'utsMapOf<string, any | null>(');
                    s.prependRight(node.end, ')');
                }
            },
        });
        return (0, compiler_core_1.createSimpleExpression)(s.toString(), false, exp.loc);
    }
}
exports.rewriteObjectExpression = rewriteObjectExpression;
function onCompilerError(error) { }
exports.onCompilerError = onCompilerError;
function parseSource(fileName, rootDir) {
    if (fileName.includes('@dcloudio')) {
        return fileName;
    }
    rootDir = (0, uni_cli_shared_1.normalizePath)(rootDir);
    if (path_1.default.isAbsolute(fileName) && fileName.startsWith(rootDir)) {
        return '@/' + (0, uni_cli_shared_1.normalizePath)(path_1.default.relative(rootDir, fileName));
    }
    return fileName;
}
exports.parseSource = parseSource;
function addEasyComponentAutoImports(easyComponentAutoImports, rootDir, tagName, fileName) {
    if (easyComponentAutoImports[fileName]) {
        return;
    }
    // 内置easycom，如 unicloud-db
    if (fileName.includes('@dcloudio')) {
        return;
    }
    rootDir = (0, uni_cli_shared_1.normalizePath)(rootDir);
    if (path_1.default.isAbsolute(fileName) && fileName.startsWith(rootDir)) {
        fileName = '@/' + (0, uni_cli_shared_1.normalizePath)(path_1.default.relative(rootDir, fileName));
    }
    // 加密插件easycom类型导入
    if (fileName.includes('?uts-proxy')) {
        const moduleId = path_1.default.basename(fileName.split('?uts-proxy')[0]);
        fileName = `uts.sdk.modules.${(0, uni_cli_shared_1.camelize)(moduleId)}`;
    }
    const ident = (0, uni_cli_shared_1.genUTSComponentPublicInstanceIdent)(tagName);
    easyComponentAutoImports[fileName] = [ident, ident];
}
exports.addEasyComponentAutoImports = addEasyComponentAutoImports;
