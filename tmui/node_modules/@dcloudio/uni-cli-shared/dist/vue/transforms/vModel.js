"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createTransformModel = exports.defaultMatch = void 0;
const utils_1 = require("../utils");
const vOn_1 = require("./vOn");
function defaultMatch(node, context) {
    return (0, utils_1.isUserComponent)(node, context);
}
exports.defaultMatch = defaultMatch;
/**
 * 百度、快手小程序的自定义组件，不支持动态事件绑定，故 v-model 也需要调整
 * @param baseTransformModel
 * @returns
 */
function createTransformModel(baseTransformModel, { match } = {
    match: defaultMatch,
}) {
    return (dir, node, context, augmentor) => {
        const res = baseTransformModel(dir, node, context, augmentor);
        if (!match(node, context)) {
            return res;
        }
        const props = res.props;
        if (props[1]) {
            // input,textarea 的 v-model 事件可能会被合并到已有的 input 中
            const { arg, exp } = props[1];
            (0, vOn_1.addEventOpts)(arg.content, exp, node, context);
            props[1].exp = (0, vOn_1.createCustomEventExpr)();
        }
        return res;
    };
}
exports.createTransformModel = createTransformModel;
