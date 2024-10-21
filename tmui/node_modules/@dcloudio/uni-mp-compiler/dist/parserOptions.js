"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parserOptions = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const shared_1 = require("@vue/shared");
exports.parserOptions = {
    isVoidTag: shared_1.isVoidTag,
    isNativeTag: (tag) => (0, shared_1.isHTMLTag)(tag) || (0, shared_1.isSVGTag)(tag),
    isPreTag: (tag) => tag === 'pre',
    // https://html.spec.whatwg.org/multipage/parsing.html#tree-construction-dispatcher
    getNamespace(tag, parent) {
        let ns = parent ? parent.ns : 0 /* DOMNamespaces.HTML */;
        if (parent && ns === 2 /* DOMNamespaces.MATH_ML */) {
            if (parent.tag === 'annotation-xml') {
                if (tag === 'svg') {
                    return 1 /* DOMNamespaces.SVG */;
                }
                if (parent.props.some((a) => a.type === compiler_core_1.NodeTypes.ATTRIBUTE &&
                    a.name === 'encoding' &&
                    a.value != null &&
                    (a.value.content === 'text/html' ||
                        a.value.content === 'application/xhtml+xml'))) {
                    ns = 0 /* DOMNamespaces.HTML */;
                }
            }
            else if (/^m(?:[ions]|text)$/.test(parent.tag) &&
                tag !== 'mglyph' &&
                tag !== 'malignmark') {
                ns = 0 /* DOMNamespaces.HTML */;
            }
        }
        else if (parent && ns === 1 /* DOMNamespaces.SVG */) {
            if (parent.tag === 'foreignObject' ||
                parent.tag === 'desc' ||
                parent.tag === 'title') {
                ns = 0 /* DOMNamespaces.HTML */;
            }
        }
        if (ns === 0 /* DOMNamespaces.HTML */) {
            if (tag === 'svg') {
                return 1 /* DOMNamespaces.SVG */;
            }
            if (tag === 'math') {
                return 2 /* DOMNamespaces.MATH_ML */;
            }
        }
        return ns;
    },
    parseMode: 'html',
};
