"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformSrcset = exports.createSrcsetTransformWithOptions = void 0;
const path_1 = __importDefault(require("path"));
const compiler_core_1 = require("@vue/compiler-core");
const templateUtils_1 = require("./templateUtils");
const transformAssetUrl_1 = require("./transformAssetUrl");
const srcsetTags = ['img', 'source'];
// http://w3c.github.io/html/semantics-embedded-content.html#ref-for-image-candidate-string-5
const escapedSpaceCharacters = /( |\\t|\\n|\\f|\\r)+/g;
const createSrcsetTransformWithOptions = (options) => {
    return (node, context) => exports.transformSrcset(node, context, options);
};
exports.createSrcsetTransformWithOptions = createSrcsetTransformWithOptions;
const transformSrcset = (node, context, options = transformAssetUrl_1.defaultAssetUrlOptions) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
        if (srcsetTags.includes(node.tag) && node.props.length) {
            node.props.forEach((attr, index) => {
                if (attr.name === 'srcset' && attr.type === compiler_core_1.NodeTypes.ATTRIBUTE) {
                    if (!attr.value)
                        return;
                    const value = attr.value.content;
                    if (!value)
                        return;
                    const imageCandidates = value
                        .split(',')
                        .map((s) => {
                        // The attribute value arrives here with all whitespace, except
                        // normal spaces, represented by escape sequences
                        const [url, descriptor] = s
                            .replace(escapedSpaceCharacters, ' ')
                            .trim()
                            .split(' ', 2);
                        return { url, descriptor };
                    });
                    // data urls contains comma after the encoding so we need to re-merge
                    // them
                    for (let i = 0; i < imageCandidates.length; i++) {
                        const { url } = imageCandidates[i];
                        if ((0, templateUtils_1.isDataUrl)(url)) {
                            imageCandidates[i + 1].url =
                                url + ',' + imageCandidates[i + 1].url;
                            imageCandidates.splice(i, 1);
                        }
                    }
                    const shouldProcessUrl = (url) => {
                        return (!(0, templateUtils_1.isExternalUrl)(url) &&
                            !(0, templateUtils_1.isDataUrl)(url) &&
                            (options.includeAbsolute || (0, templateUtils_1.isRelativeUrl)(url)));
                    };
                    // When srcset does not contain any qualified URLs, skip transforming
                    if (!imageCandidates.some(({ url }) => shouldProcessUrl(url))) {
                        return;
                    }
                    if (options.base) {
                        const base = options.base;
                        const set = [];
                        let needImportTransform = false;
                        imageCandidates.forEach((candidate) => {
                            let { url, descriptor } = candidate;
                            descriptor = descriptor ? ` ${descriptor}` : ``;
                            if (url[0] === '.') {
                                candidate.url = (path_1.default.posix || path_1.default).join(base, url);
                                set.push(candidate.url + descriptor);
                            }
                            else if (shouldProcessUrl(url)) {
                                needImportTransform = true;
                            }
                            else {
                                set.push(url + descriptor);
                            }
                        });
                        if (!needImportTransform) {
                            attr.value.content = set.join(', ');
                            return;
                        }
                    }
                    const compoundExpression = (0, compiler_core_1.createCompoundExpression)([], attr.loc);
                    imageCandidates.forEach(({ url, descriptor }, index) => {
                        if (shouldProcessUrl(url)) {
                            const { path } = (0, templateUtils_1.parseUrl)(url);
                            let exp;
                            if (path) {
                                const existingImportsIndex = context.imports.findIndex((i) => i.path === path);
                                if (existingImportsIndex > -1) {
                                    exp = (0, compiler_core_1.createSimpleExpression)(`_imports_${existingImportsIndex}`, false, attr.loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
                                }
                                else {
                                    exp = (0, compiler_core_1.createSimpleExpression)(`_imports_${context.imports.length}`, false, attr.loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
                                    context.imports.push({ exp, path });
                                }
                                compoundExpression.children.push(exp);
                            }
                        }
                        else {
                            const exp = (0, compiler_core_1.createSimpleExpression)(`"${url}"`, false, attr.loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
                            compoundExpression.children.push(exp);
                        }
                        const isNotLast = imageCandidates.length - 1 > index;
                        if (descriptor && isNotLast) {
                            compoundExpression.children.push(` + ' ${descriptor}, ' + `);
                        }
                        else if (descriptor) {
                            compoundExpression.children.push(` + ' ${descriptor}'`);
                        }
                        else if (isNotLast) {
                            compoundExpression.children.push(` + ', ' + `);
                        }
                    });
                    let exp = compoundExpression;
                    if (context.hoistStatic) {
                        exp = context.hoist(compoundExpression);
                        exp.constType = compiler_core_1.ConstantTypes.CAN_STRINGIFY;
                    }
                    node.props[index] = {
                        type: compiler_core_1.NodeTypes.DIRECTIVE,
                        name: 'bind',
                        arg: (0, compiler_core_1.createSimpleExpression)('srcset', true, attr.loc),
                        exp,
                        modifiers: [],
                        loc: attr.loc,
                    };
                }
            });
        }
    }
};
exports.transformSrcset = transformSrcset;
