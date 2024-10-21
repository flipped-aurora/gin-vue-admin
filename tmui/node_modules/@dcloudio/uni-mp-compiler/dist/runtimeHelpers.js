"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.WITH_MODEL_MODIFIERS = exports.TO_DISPLAY_STRING = exports.NORMALIZE_CLASS = exports.STRINGIFY_STYLE = exports.WITH_SCOPED_SLOT = exports.DYNAMIC_SLOT = exports.RENDER_SLOT = exports.RENDER_PROPS = exports.HYPHENATE = exports.CAMELIZE = exports.SET_REF = exports.EXTEND = exports.V_FOR = exports.V_ON = void 0;
const compiler_core_1 = require("@vue/compiler-core");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
exports.V_ON = Symbol(`vOn`);
exports.V_FOR = Symbol(`vFor`);
exports.EXTEND = Symbol(`extend`);
exports.SET_REF = Symbol(`setRef`);
exports.CAMELIZE = Symbol(`camelize`);
exports.HYPHENATE = Symbol(`hyphenate`);
exports.RENDER_PROPS = Symbol(`renderProps`);
exports.RENDER_SLOT = Symbol(`renderSlot`);
exports.DYNAMIC_SLOT = Symbol(`dynamicSlot`);
exports.WITH_SCOPED_SLOT = Symbol(`withScopedSlot`);
exports.STRINGIFY_STYLE = Symbol(`stringifyStyle`);
exports.NORMALIZE_CLASS = Symbol(`normalizeClass`);
exports.TO_DISPLAY_STRING = Symbol(`toDisplayString`);
exports.WITH_MODEL_MODIFIERS = Symbol(`withModelModifiers`);
(0, compiler_core_1.registerRuntimeHelpers)({
    [exports.V_ON]: 'o',
    [exports.V_FOR]: 'f',
    [exports.EXTEND]: 'e',
    [exports.SET_REF]: 'sr',
    [exports.CAMELIZE]: 'c',
    [exports.HYPHENATE]: 'h',
    [exports.RENDER_PROPS]: 'p',
    [exports.RENDER_SLOT]: 'r',
    [exports.DYNAMIC_SLOT]: 'd',
    [exports.WITH_SCOPED_SLOT]: 'w',
    [exports.STRINGIFY_STYLE]: 's',
    [exports.NORMALIZE_CLASS]: 'n',
    [exports.TO_DISPLAY_STRING]: 't',
    [exports.WITH_MODEL_MODIFIERS]: 'm',
    [uni_cli_shared_1.STRINGIFY_JSON]: 'j',
});
