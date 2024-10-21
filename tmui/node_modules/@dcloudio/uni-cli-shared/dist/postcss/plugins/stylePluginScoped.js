"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const postcss_selector_parser_1 = __importDefault(require("postcss-selector-parser"));
const scopedPlugin = () => {
    return {
        postcssPlugin: 'uni-sfc-scoped',
        prepare({ processor: { plugins } }) {
            const hasVueSfcScoped = !!plugins.find((plugin) => plugin.postcssPlugin === 'vue-sfc-scoped');
            return {
                Rule(rule) {
                    processRule(rule, hasVueSfcScoped);
                },
            };
        },
    };
};
const processedRules = new WeakSet();
function processRule(rule, hasVueSfcScoped) {
    if (processedRules.has(rule) ||
        (rule.parent &&
            rule.parent.type === 'atrule' &&
            /-?keyframes$/.test(rule.parent.name))) {
        return;
    }
    processedRules.add(rule);
    rule.selector = (0, postcss_selector_parser_1.default)((selectorRoot) => {
        selectorRoot.each((selector) => {
            hasVueSfcScoped
                ? rewriteDeprecatedSelector(selector)
                : rewriteSelector(selector, selectorRoot);
        });
    }).processSync(rule.selector);
}
/**
 * @param selector
 * @returns
 */
function rewriteDeprecatedSelector(selector) {
    const nodes = [];
    let deepNode;
    selector.each((n) => {
        if (deepNode) {
            nodes.push(n);
            selector.removeChild(n);
        }
        else {
            const { type, value } = n;
            if (type === 'pseudo' && value === '::v-deep') {
                deepNode = n;
            }
            else if (type === 'combinator' &&
                (value === '>>>' || value === '/deep/')) {
                deepNode = n;
            }
        }
    });
    if (!deepNode) {
        return;
    }
    if (deepNode.type === 'combinator') {
        const index = selector.index(deepNode);
        if (index > 0) {
            selector.insertBefore(deepNode, postcss_selector_parser_1.default.combinator({ value: ' ' }));
        }
    }
    // remove first combinator
    // ::v-deep a{color:red;} => :deep(a){color:red;}
    const firstNode = nodes[0];
    if (firstNode && firstNode.type === 'combinator' && firstNode.value === ' ') {
        nodes.shift();
    }
    selector.insertBefore(deepNode, postcss_selector_parser_1.default.pseudo({
        value: ':deep',
        nodes: [postcss_selector_parser_1.default.selector({ value: '', nodes })],
    }));
    selector.removeChild(deepNode);
}
function rewriteSelector(selector, selectorRoot) {
    let node = null;
    // find the last child node to insert attribute selector
    selector.each((n) => {
        // DEPRECATED ">>>" and "/deep/" combinator
        if (n.type === 'combinator' &&
            (n.value === '>>>' || n.value === '/deep/')) {
            n.value = ' ';
            n.spaces.before = n.spaces.after = '';
            // warn(
            //   `the >>> and /deep/ combinators have been deprecated. ` +
            //     `Use :deep() instead.`
            // )
            return false;
        }
        if (n.type === 'pseudo') {
            const { value } = n;
            // deep: inject [id] attribute at the node before the ::v-deep
            // combinator.
            if (value === ':deep' || value === '::v-deep') {
                if (n.nodes.length) {
                    // .foo ::v-deep(.bar) -> .foo[xxxxxxx] .bar
                    // replace the current node with ::v-deep's inner selector
                    let last = n;
                    n.nodes[0].each((ss) => {
                        selector.insertAfter(last, ss);
                        last = ss;
                    });
                    // insert a space combinator before if it doesn't already have one
                    const prev = selector.at(selector.index(n) - 1);
                    if (!prev || !isSpaceCombinator(prev)) {
                        selector.insertAfter(n, postcss_selector_parser_1.default.combinator({
                            value: ' ',
                        }));
                    }
                    selector.removeChild(n);
                }
                else {
                    // DEPRECATED usage
                    // .foo ::v-deep .bar -> .foo[xxxxxxx] .bar
                    // warn(
                    //   `::v-deep usage as a combinator has ` +
                    //     `been deprecated. Use :deep(<inner-selector>) instead.`
                    // )
                    const prev = selector.at(selector.index(n) - 1);
                    if (prev && isSpaceCombinator(prev)) {
                        selector.removeChild(prev);
                    }
                    selector.removeChild(n);
                }
                return false;
            }
            // slot: use selector inside `::v-slotted` and inject [id + '-s']
            // instead.
            // ::v-slotted(.foo) -> .foo[xxxxxxx-s]
            if (value === ':slotted' || value === '::v-slotted') {
                rewriteSelector(n.nodes[0], selectorRoot);
                let last = n;
                n.nodes[0].each((ss) => {
                    selector.insertAfter(last, ss);
                    last = ss;
                });
                // selector.insertAfter(n, n.nodes[0])
                selector.removeChild(n);
                // since slotted attribute already scopes the selector there's no
                // need for the non-slot attribute.
                return false;
            }
            // global: replace with inner selector and do not inject [id].
            // ::v-global(.foo) -> .foo
            if (value === ':global' || value === '::v-global') {
                selectorRoot.insertAfter(selector, n.nodes[0]);
                selectorRoot.removeChild(selector);
                return false;
            }
        }
        if (n.type !== 'pseudo' && n.type !== 'combinator') {
            node = n;
        }
    });
    if (node) {
        ;
        node.spaces.after = '';
    }
    else {
        // For deep selectors & standalone pseudo selectors,
        // the attribute selectors are prepended rather than appended.
        // So all leading spaces must be eliminated to avoid problems.
        selector.first.spaces.before = '';
    }
}
function isSpaceCombinator(node) {
    return node.type === 'combinator' && /^\s+$/.test(node.value);
}
scopedPlugin.postcss = true;
exports.default = scopedPlugin;
