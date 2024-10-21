"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.STRINGIFY_JSON = exports.ATTR_DATASET_EVENT_OPTS = exports.addEventOpts = exports.createCustomEventExpr = exports.createTransformOn = exports.defaultMatch = void 0;
const uni_shared_1 = require("@dcloudio/uni-shared");
const compiler_core_1 = require("@vue/compiler-core");
const utils_1 = require("../utils");
function defaultMatch(name, node, context) {
    return isCustomEvent(name) && (0, utils_1.isUserComponent)(node, context);
}
exports.defaultMatch = defaultMatch;
/**
 * 百度、快手小程序的自定义组件，不支持动态事件绑定，故转换为静态事件 + dataset
 * @param baseTransformOn
 * @returns
 */
function createTransformOn(baseTransformOn, { match } = {
    match: defaultMatch,
}) {
    return (dir, node, context, augmentor) => {
        const res = baseTransformOn(dir, node, context, augmentor);
        const { name, arg, exp } = dir;
        if (name !== 'on' || !arg || !exp || !(0, compiler_core_1.isStaticExp)(arg)) {
            return res;
        }
        if (!match(arg.content, node, context)) {
            return res;
        }
        const value = res.props[0].value;
        res.props[0].value = createCustomEventExpr();
        addEventOpts(node.tagType === compiler_core_1.ElementTypes.COMPONENT
            ? (0, uni_shared_1.customizeEvent)(arg.content)
            : arg.content, value, node, context);
        return res;
    };
}
exports.createTransformOn = createTransformOn;
function createCustomEventExpr() {
    return (0, compiler_core_1.createSimpleExpression)('__e', true);
}
exports.createCustomEventExpr = createCustomEventExpr;
function addEventOpts(event, value, node, context) {
    const attrName = node.tagType === compiler_core_1.ElementTypes.COMPONENT
        ? ATTR_DATA_EVENT_OPTS
        : exports.ATTR_DATASET_EVENT_OPTS;
    const opts = (0, compiler_core_1.findProp)(node, attrName, true);
    if (!opts) {
        node.props.push(createDataEventOptsProp(attrName, event, value, context));
    }
    else {
        const children = opts.exp.children;
        children.splice(children.length - 2, 0, createDataEventOptsProperty(event, value));
    }
}
exports.addEventOpts = addEventOpts;
const ATTR_DATA_EVENT_OPTS = 'eO';
exports.ATTR_DATASET_EVENT_OPTS = 'data-e-o';
function createDataEventOptsProperty(event, exp) {
    return (0, compiler_core_1.createCompoundExpression)([`'${event}'`, ': ', exp, ',']);
}
exports.STRINGIFY_JSON = Symbol(`stringifyJson`);
function createDataEventOptsProp(name, event, exp, context) {
    const children = [];
    const stringify = name === ATTR_DATA_EVENT_OPTS;
    if (stringify) {
        children.push(context.helperString(exports.STRINGIFY_JSON) + '(');
    }
    children.push('{', createDataEventOptsProperty(event, exp), '}');
    if (stringify) {
        children.push(')');
    }
    return {
        type: compiler_core_1.NodeTypes.DIRECTIVE,
        name: 'bind',
        loc: compiler_core_1.locStub,
        modifiers: [],
        arg: (0, compiler_core_1.createSimpleExpression)(name, true),
        exp: (0, compiler_core_1.createCompoundExpression)(children),
    };
}
const builtInEvents = [
    '__l', // 快手使用了该事件
    'tap',
    'longtap',
    'longpress',
    'touchstart',
    'touchmove',
    'touchcancel',
    'touchend',
    'touchforcechange',
    'transitionend',
    'animationstart',
    'animationiteration',
    'animationend',
];
function isCustomEvent(name) {
    return !builtInEvents.includes(name);
}
