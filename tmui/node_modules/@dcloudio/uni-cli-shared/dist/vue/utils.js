"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.advancePositionWithMutation = exports.advancePositionWithClone = exports.getInnerRange = exports.renameProp = exports.getBaseNodeTransforms = exports.createUniVueTransformAssetUrls = exports.createBindDirectiveNode = exports.createOnDirectiveNode = exports.createDirectiveNode = exports.addStaticClass = exports.createAttributeNode = exports.isUserComponent = exports.isVueSfcFile = exports.VUE_REF_IN_FOR = exports.VUE_REF = void 0;
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const compiler_core_1 = require("@vue/compiler-core");
const templateTransformAssetUrl_1 = require("./transforms/templateTransformAssetUrl");
const templateTransformSrcset_1 = require("./transforms/templateTransformSrcset");
const ast_1 = require("../vite/utils/ast");
const url_1 = require("../vite/utils/url");
const constants_1 = require("../constants");
exports.VUE_REF = 'r';
exports.VUE_REF_IN_FOR = 'r-i-f';
function isVueSfcFile(id) {
    const { filename, query } = (0, url_1.parseVueRequest)(id);
    return constants_1.EXTNAME_VUE_RE.test(filename) && !query.vue;
}
exports.isVueSfcFile = isVueSfcFile;
function isUserComponent(node, context) {
    return (node.type === compiler_core_1.NodeTypes.ELEMENT &&
        node.tagType === compiler_core_1.ElementTypes.COMPONENT &&
        !(0, uni_shared_1.isComponentTag)(node.tag) &&
        !(0, compiler_core_1.isCoreComponent)(node.tag) &&
        !context.isBuiltInComponent(node.tag));
}
exports.isUserComponent = isUserComponent;
function createAttributeNode(name, content) {
    return {
        type: compiler_core_1.NodeTypes.ATTRIBUTE,
        loc: compiler_core_1.locStub,
        nameLoc: compiler_core_1.locStub,
        name,
        value: {
            type: compiler_core_1.NodeTypes.TEXT,
            loc: compiler_core_1.locStub,
            content,
        },
    };
}
exports.createAttributeNode = createAttributeNode;
function createClassAttribute(clazz) {
    return createAttributeNode('class', clazz);
}
function addStaticClass(node, clazz) {
    const classProp = node.props.find((prop) => prop.type === compiler_core_1.NodeTypes.ATTRIBUTE && prop.name === 'class');
    if (!classProp) {
        return node.props.unshift(createClassAttribute(clazz));
    }
    if (classProp.value) {
        return (classProp.value.content = classProp.value.content + ' ' + clazz);
    }
    classProp.value = {
        type: compiler_core_1.NodeTypes.TEXT,
        loc: compiler_core_1.locStub,
        content: clazz,
    };
}
exports.addStaticClass = addStaticClass;
function createDirectiveNode(name, arg, exp) {
    return {
        type: compiler_core_1.NodeTypes.DIRECTIVE,
        name,
        modifiers: [],
        loc: compiler_core_1.locStub,
        arg: (0, compiler_core_1.createSimpleExpression)(arg, true),
        exp: (0, shared_1.isString)(exp) ? (0, compiler_core_1.createSimpleExpression)(exp, false) : exp,
    };
}
exports.createDirectiveNode = createDirectiveNode;
function createOnDirectiveNode(name, value) {
    return createDirectiveNode('on', name, value);
}
exports.createOnDirectiveNode = createOnDirectiveNode;
function createBindDirectiveNode(name, value) {
    return createDirectiveNode('bind', name, value);
}
exports.createBindDirectiveNode = createBindDirectiveNode;
function createUniVueTransformAssetUrls(base) {
    return {
        base,
        includeAbsolute: true,
        tags: {
            audio: ['src'],
            video: ['src', 'poster'],
            img: ['src'],
            image: ['src'],
            'cover-image': ['src'],
            // h5
            'v-uni-audio': ['src'],
            'v-uni-video': ['src', 'poster'],
            'v-uni-image': ['src'],
            'v-uni-cover-image': ['src'],
            // nvue
            'u-image': ['src'],
            'u-video': ['src', 'poster'],
        },
    };
}
exports.createUniVueTransformAssetUrls = createUniVueTransformAssetUrls;
function getBaseNodeTransforms(base) {
    const transformAssetUrls = createUniVueTransformAssetUrls(base);
    return [
        (0, templateTransformAssetUrl_1.createAssetUrlTransformWithOptions)(transformAssetUrls),
        (0, templateTransformSrcset_1.createSrcsetTransformWithOptions)(transformAssetUrls),
    ];
}
exports.getBaseNodeTransforms = getBaseNodeTransforms;
function renameProp(name, prop) {
    if (!prop) {
        return;
    }
    if ((0, ast_1.isDirectiveNode)(prop)) {
        if (prop.arg && (0, compiler_core_1.isStaticExp)(prop.arg)) {
            prop.arg.content = name;
        }
    }
    else {
        prop.name = name;
    }
}
exports.renameProp = renameProp;
// @vue/compiler-core 没有导出 getLoc，先使用旧版本的 getInnerRange
function getInnerRange(loc, offset, length) {
    const source = loc.source.slice(offset, offset + length);
    const newLoc = {
        source,
        start: advancePositionWithClone(loc.start, loc.source, offset),
        end: loc.end,
    };
    if (length != null) {
        newLoc.end = advancePositionWithClone(loc.start, loc.source, offset + length);
    }
    return newLoc;
}
exports.getInnerRange = getInnerRange;
function advancePositionWithClone(pos, source, numberOfCharacters = source.length) {
    return advancePositionWithMutation((0, shared_1.extend)({}, pos), source, numberOfCharacters);
}
exports.advancePositionWithClone = advancePositionWithClone;
// advance by mutation without cloning (for performance reasons), since this
// gets called a lot in the parser
function advancePositionWithMutation(pos, source, numberOfCharacters = source.length) {
    let linesCount = 0;
    let lastNewLinePos = -1;
    for (let i = 0; i < numberOfCharacters; i++) {
        if (source.charCodeAt(i) === 10 /* newline char code */) {
            linesCount++;
            lastNewLinePos = i;
        }
    }
    pos.offset += numberOfCharacters;
    pos.line += linesCount;
    pos.column =
        lastNewLinePos === -1
            ? pos.column + numberOfCharacters
            : numberOfCharacters - lastNewLinePos;
    return pos;
}
exports.advancePositionWithMutation = advancePositionWithMutation;
