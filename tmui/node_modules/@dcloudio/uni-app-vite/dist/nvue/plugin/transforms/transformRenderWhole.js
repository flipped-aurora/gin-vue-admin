"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addRenderWhole = exports.transformRenderWhole = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const compiler_core_1 = require("@vue/compiler-core");
const transformRenderWhole = (node, _) => {
    if (!(0, uni_cli_shared_1.isElementNode)(node)) {
        return;
    }
    const prop = (0, compiler_core_1.findProp)(node, 'render-whole');
    if (!prop) {
        return;
    }
    // render-whole => append
    (0, uni_cli_shared_1.renameProp)('append', prop);
};
exports.transformRenderWhole = transformRenderWhole;
//
const RENDER_WHOLE_TAGS = [
    'view',
    'scroll-view',
    'swiper',
    'match-media',
    'movable-area',
    'movable-view',
    'cover-view',
    'cover-image',
    'form',
    'picker',
    'picker-view',
    'navigator',
    'map',
];
/**
 * 仅当根节点只有一个，标签在白名单，且开发者未主动配置的情况下，才补充
 * @param node
 */
function addRenderWhole(node) {
    if (node.children.length === 1) {
        const element = node.children[0];
        if ((0, uni_cli_shared_1.isElementNode)(element) && RENDER_WHOLE_TAGS.includes(element.tag)) {
            if (!(0, compiler_core_1.findProp)(element, 'render-whole') &&
                !(0, compiler_core_1.findProp)(element, 'append') &&
                !(0, compiler_core_1.findProp)(element, 'appendAsTree')) {
                element.props.push((0, uni_cli_shared_1.createBindDirectiveNode)('render-whole', (0, compiler_core_1.createSimpleExpression)('true')));
            }
        }
    }
}
exports.addRenderWhole = addRenderWhole;
