"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformShow = void 0;
const runtimeHelpers_1 = require("../runtimeHelpers");
const transformShow = (dir, node, context) => {
    return {
        props: [],
        needRuntime: context.helper(runtimeHelpers_1.V_SHOW),
    };
};
exports.transformShow = transformShow;
