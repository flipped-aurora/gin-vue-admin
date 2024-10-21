"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseWxsCode = exports.parseWxsNodes = exports.parseBlockCode = exports.parseVueCode = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const magic_string_1 = __importDefault(require("magic-string"));
const ast_1 = require("../vite/utils/ast");
const BLOCK_RE = /<\/block>/;
const WXS_LANG_RE = /lang=["|'](renderjs|wxs|sjs)["|']/;
const WXS_ATTRS = ['wxs', 'renderjs', 'sjs'];
function parseVueCode(code, isNVue = false) {
    const hasBlock = BLOCK_RE.test(code);
    const hasWxs = WXS_LANG_RE.test(code);
    if (!hasBlock && !hasWxs) {
        return { code };
    }
    const errors = [];
    const files = [];
    let ast = (0, ast_1.parseVue)(code, errors);
    if (hasBlock) {
        code = parseBlockCode(ast, code);
        // 重新解析新的 code
        ast = (0, ast_1.parseVue)(code, errors);
    }
    if (!isNVue && hasWxs) {
        const wxsNodes = parseWxsNodes(ast);
        code = parseWxsCode(wxsNodes, code);
        // add watch
        for (const wxsNode of wxsNodes) {
            const srcProp = wxsNode.props.find((prop) => prop.type === compiler_core_1.NodeTypes.ATTRIBUTE && prop.name === 'src');
            if (srcProp && srcProp.value) {
                files.push(srcProp.value.content);
            }
        }
    }
    return { code, files, errors };
}
exports.parseVueCode = parseVueCode;
function traverseChildren({ children }, blockNodes) {
    children.forEach((node) => traverseNode(node, blockNodes));
}
function traverseNode(node, blockNodes) {
    if ((0, ast_1.isElementNode)(node) && node.tag === 'block') {
        blockNodes.push(node);
    }
    if (node.type === compiler_core_1.NodeTypes.IF_BRANCH ||
        node.type === compiler_core_1.NodeTypes.FOR ||
        node.type === compiler_core_1.NodeTypes.ELEMENT ||
        node.type === compiler_core_1.NodeTypes.ROOT) {
        traverseChildren(node, blockNodes);
    }
}
function parseBlockCode(ast, code) {
    const blockNodes = [];
    traverseNode(ast, blockNodes);
    if (blockNodes.length) {
        return parseBlockNode(code, blockNodes);
    }
    return code;
}
exports.parseBlockCode = parseBlockCode;
const BLOCK_END_LEN = '</block>'.length;
const BLOCK_START_LEN = '<block'.length;
function parseBlockNode(code, blocks) {
    const magicString = new magic_string_1.default(code);
    blocks.forEach(({ loc }) => {
        const startOffset = loc.start.offset;
        const endOffset = loc.end.offset;
        magicString.overwrite(startOffset, startOffset + BLOCK_START_LEN, '<template');
        magicString.overwrite(endOffset - BLOCK_END_LEN, endOffset, '</template>');
    });
    return magicString.toString();
}
function parseWxsNodes(ast) {
    return ast.children.filter((node) => node.type === compiler_core_1.NodeTypes.ELEMENT &&
        node.tag === 'script' &&
        node.props.find((prop) => prop.name === 'lang' &&
            prop.type === compiler_core_1.NodeTypes.ATTRIBUTE &&
            prop.value &&
            WXS_ATTRS.includes(prop.value.content)));
}
exports.parseWxsNodes = parseWxsNodes;
function parseWxsCode(wxsNodes, code) {
    if (wxsNodes.length) {
        code = parseWxsNode(code, wxsNodes);
    }
    return code;
}
exports.parseWxsCode = parseWxsCode;
const SCRIPT_END_LEN = '</script>'.length;
const SCRIPT_START_LEN = '<script'.length;
function parseWxsNode(code, nodes) {
    const magicString = new magic_string_1.default(code);
    nodes.forEach(({ loc, props }) => {
        const langAttr = props.find((prop) => prop.name === 'lang');
        const moduleAttr = props.find((prop) => prop.name === 'module');
        const startOffset = loc.start.offset;
        const endOffset = loc.end.offset;
        const lang = langAttr.value.content;
        const langStartOffset = langAttr.loc.start.offset;
        magicString.overwrite(startOffset, startOffset + SCRIPT_START_LEN, '<' + lang); // <renderjs or <wxs
        magicString.overwrite(langStartOffset, langStartOffset + ('lang="' + lang + '"').length, ''); // remove lang="renderjs" or lang="wxs"
        magicString.overwrite(endOffset - SCRIPT_END_LEN, endOffset, '</' + lang + '>'); //</renderjs> or </wxs>
        if (moduleAttr) {
            const moduleStartOffset = moduleAttr.loc.start.offset;
            magicString.overwrite(moduleStartOffset, moduleStartOffset + 'module'.length, 'name'); // module="echarts" => name="echarts"
        }
    });
    return magicString.toString();
}
