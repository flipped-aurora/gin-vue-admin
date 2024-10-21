"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformShow = void 0;
const errors_1 = require("./errors");
const transformShow = (dir, node, context) => {
    context.onError((0, errors_1.createNVueCompilerError)(0 /* NVueErrorCodes.X_V_SHOW */, dir.loc));
    return {
        props: [],
    };
};
exports.transformShow = transformShow;
