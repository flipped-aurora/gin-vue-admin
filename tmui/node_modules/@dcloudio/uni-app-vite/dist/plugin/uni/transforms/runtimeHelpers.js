"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.WXS_PROP = void 0;
const compiler_core_1 = require("@vue/compiler-core");
exports.WXS_PROP = Symbol(`wxsProp`);
(0, compiler_core_1.registerRuntimeHelpers)({
    [exports.WXS_PROP]: 'wp',
});
