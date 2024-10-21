"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformAssetUrl = exports.createAssetUrlTransformWithOptions = exports.normalizeOptions = exports.defaultAssetUrlOptions = void 0;
const path_1 = __importDefault(require("path"));
const compiler_core_1 = require("@vue/compiler-core");
const templateUtils_1 = require("./templateUtils");
const shared_1 = require("@vue/shared");
exports.defaultAssetUrlOptions = {
    base: null,
    includeAbsolute: false,
    tags: {
        video: ['src', 'poster'],
        source: ['src'],
        img: ['src'],
        image: ['xlink:href', 'href'],
        use: ['xlink:href', 'href'],
    },
};
const normalizeOptions = (options) => {
    if (Object.keys(options).some((key) => (0, shared_1.isArray)(options[key]))) {
        // legacy option format which directly passes in tags config
        return {
            ...exports.defaultAssetUrlOptions,
            tags: options,
        };
    }
    return {
        ...exports.defaultAssetUrlOptions,
        ...options,
    };
};
exports.normalizeOptions = normalizeOptions;
const createAssetUrlTransformWithOptions = (options) => {
    return (node, context) => exports.transformAssetUrl(node, context, options);
};
exports.createAssetUrlTransformWithOptions = createAssetUrlTransformWithOptions;
/**
 * A `@vue/compiler-core` plugin that transforms relative asset urls into
 * either imports or absolute urls.
 *
 * ``` js
 * // Before
 * createVNode('img', { src: './logo.png' })
 *
 * // After
 * import _imports_0 from './logo.png'
 * createVNode('img', { src: _imports_0 })
 * ```
 */
const transformAssetUrl = (node, context, options = exports.defaultAssetUrlOptions) => {
    if (node.type === compiler_core_1.NodeTypes.ELEMENT) {
        if (!node.props.length) {
            return;
        }
        const tags = options.tags || exports.defaultAssetUrlOptions.tags;
        const attrs = tags[node.tag];
        const wildCardAttrs = tags['*'];
        if (!attrs && !wildCardAttrs) {
            return;
        }
        const assetAttrs = (attrs || []).concat(wildCardAttrs || []);
        node.props.forEach((attr, index) => {
            if (attr.type !== compiler_core_1.NodeTypes.ATTRIBUTE ||
                !assetAttrs.includes(attr.name) ||
                !attr.value ||
                (0, templateUtils_1.isExternalUrl)(attr.value.content) ||
                (0, templateUtils_1.isDataUrl)(attr.value.content) ||
                attr.value.content[0] === '#' ||
                (!options.includeAbsolute && !(0, templateUtils_1.isRelativeUrl)(attr.value.content))) {
                return;
            }
            const url = (0, templateUtils_1.parseUrl)(attr.value.content);
            if (options.base && attr.value.content[0] === '.') {
                // explicit base - directly rewrite relative urls into absolute url
                // to avoid generating extra imports
                // Allow for full hostnames provided in options.base
                const base = (0, templateUtils_1.parseUrl)(options.base);
                const protocol = base.protocol || '';
                const host = base.host ? protocol + '//' + base.host : '';
                const basePath = base.path || '/';
                // when packaged in the browser, path will be using the posix-
                // only version provided by rollup-plugin-node-builtins.
                attr.value.content =
                    host +
                        (path_1.default.posix || path_1.default).join(basePath, url.path + (url.hash || ''));
                return;
            }
            // otherwise, transform the url into an import.
            // this assumes a bundler will resolve the import into the correct
            // absolute url (e.g. webpack file-loader)
            const exp = getImportsExpressionExp(url.path, url.hash, attr.loc, context);
            node.props[index] = {
                type: compiler_core_1.NodeTypes.DIRECTIVE,
                name: 'bind',
                arg: (0, compiler_core_1.createSimpleExpression)(attr.name, true, attr.loc),
                exp,
                modifiers: [],
                loc: attr.loc,
            };
        });
    }
};
exports.transformAssetUrl = transformAssetUrl;
function getImportsExpressionExp(path, hash, loc, context) {
    if (path) {
        let name;
        let exp;
        const existingIndex = context.imports.findIndex((i) => i.path === path);
        if (existingIndex > -1) {
            name = `_imports_${existingIndex}`;
            exp = context.imports[existingIndex].exp;
        }
        else {
            name = `_imports_${context.imports.length}`;
            exp = (0, compiler_core_1.createSimpleExpression)(name, false, loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
            // We need to ensure the path is not encoded (to %2F),
            // so we decode it back in case it is encoded
            context.imports.push({
                exp,
                path: decodeURIComponent(path),
            });
        }
        if (!hash) {
            return exp;
        }
        const hashExp = `${name} + '${hash}'`;
        const finalExp = (0, compiler_core_1.createSimpleExpression)(hashExp, false, loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
        if (!context.hoistStatic) {
            return finalExp;
        }
        const existingHoistIndex = context.hoists.findIndex((h) => {
            return (h &&
                h.type === compiler_core_1.NodeTypes.SIMPLE_EXPRESSION &&
                !h.isStatic &&
                h.content === hashExp);
        });
        if (existingHoistIndex > -1) {
            return (0, compiler_core_1.createSimpleExpression)(`_hoisted_${existingHoistIndex + 1}`, false, loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
        }
        return context.hoist(finalExp);
    }
    else {
        return (0, compiler_core_1.createSimpleExpression)(`''`, false, loc, compiler_core_1.ConstantTypes.CAN_STRINGIFY);
    }
}
