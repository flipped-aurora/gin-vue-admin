"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createNVueCompiler = void 0;
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
const shared_1 = require("@vue/shared");
/**
 * nvue 需要移除 scoped
 * @param nvuePages
 */
function createNVueCompiler() {
    const compileSfc = require((0, uni_cli_shared_1.resolveBuiltIn)('@vue/compiler-sfc'));
    const { parse } = compileSfc;
    return (0, shared_1.extend)({}, compileSfc, {
        parse(source, options = {}) {
            const result = parse(source, options);
            result.descriptor.styles.forEach((style) => {
                if (style.scoped) {
                    delete style.scoped;
                }
            });
            return result;
        },
    });
}
exports.createNVueCompiler = createNVueCompiler;
