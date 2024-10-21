"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformRootNode = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformRenderWhole_1 = require("./transformRenderWhole");
const pagesJson_1 = require("../../plugins/pagesJson");
const SCROLLER_COMPONENTS = [
    'list',
    'scroller',
    'scroll-view',
    'waterfall',
    'recycle-list',
];
const transformRootNode = (node, context) => {
    if (node.type !== compiler_core_1.NodeTypes.ROOT) {
        return;
    }
    const pageOptions = (0, pagesJson_1.parseNVuePageOptions)((0, uni_cli_shared_1.normalizePath)(context.filename));
    if (!pageOptions) {
        // 非页面组件，自动为根节点补充 render-whole
        return (0, transformRenderWhole_1.addRenderWhole)(node);
    }
    const { disableScroll, scrollIndicator } = pageOptions;
    // 禁用滚动，或已包含滚动元素
    if (disableScroll || hasScrollerElement(node)) {
        return wrapperByView(node);
    }
    return wrapperByScrollView(node, { scrollIndicator });
};
exports.transformRootNode = transformRootNode;
function hasScrollerElement(node) {
    return node.children.some((child) => {
        if (child.type === compiler_core_1.NodeTypes.ELEMENT) {
            return SCROLLER_COMPONENTS.includes(child.tag);
        }
    });
}
function wrapperByScrollView(node, { scrollIndicator }) {
    node.children = [
        createElement('scroll-view', createScrollViewProps({ scrollIndicator }), node.children),
    ];
}
const trueExpr = (0, compiler_core_1.createSimpleExpression)('true');
const falseExpr = (0, compiler_core_1.createSimpleExpression)('false');
function createScrollViewProps({ scrollIndicator, }) {
    return [
        (0, uni_cli_shared_1.createBindDirectiveNode)('scrollY', trueExpr),
        (0, uni_cli_shared_1.createBindDirectiveNode)('showScrollbar', scrollIndicator === 'none' ? falseExpr : trueExpr),
        (0, uni_cli_shared_1.createBindDirectiveNode)('enableBackToTop', trueExpr),
        (0, uni_cli_shared_1.createAttributeNode)('bubble', 'true'),
        (0, uni_cli_shared_1.createBindDirectiveNode)('style', `{flexDirection:'column'}`),
    ];
}
/**
 * 目前暂不支持多节点，故发现多节点时，自动补充一个 view 根节点
 * @param node
 */
function wrapperByView(node) {
    if (node.children.length > 1) {
        node.children = [createElement('view', [], node.children)];
    }
}
function createElement(tag, props, children) {
    return {
        type: compiler_core_1.NodeTypes.ELEMENT,
        ns: 0,
        tag,
        isSelfClosing: false,
        props,
        children,
        tagType: compiler_core_1.ElementTypes.ELEMENT,
        codegenNode: undefined,
        loc: compiler_core_1.locStub,
    };
}
