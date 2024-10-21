"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.rewriteCompilerSfcParseOnce = exports.rewriteCompileScriptOnce = void 0;
const shared_1 = require("@vue/shared");
const uni_shared_1 = require("@dcloudio/uni-shared");
const uni_cli_shared_1 = require("@dcloudio/uni-cli-shared");
exports.rewriteCompileScriptOnce = (0, uni_shared_1.once)(rewriteCompileScript);
exports.rewriteCompilerSfcParseOnce = (0, uni_shared_1.once)(rewriteCompilerSfcParse);
function rewriteCompileScript() {
    const compiler = require((0, uni_cli_shared_1.resolveBuiltIn)('@vue/compiler-sfc'));
    const { compileScript, compileTemplate, compileStyle, compileStyleAsync } = compiler;
    compiler.compileStyle = (options) => {
        // https://github.com/dcloudio/uni-app/issues/4076
        options.isProd = true;
        return compileStyle(options);
    };
    compiler.compileStyleAsync = (options) => {
        // https://github.com/dcloudio/uni-app/issues/4076
        options.isProd = true;
        return compileStyleAsync(options);
    };
    // script-setup + v-bind
    compiler.compileScript = (sfc, options) => {
        if (options?.templateOptions?.compilerOptions) {
            ;
            options.templateOptions.compilerOptions.bindingCssVars =
                sfc.cssVars || [];
        }
        // 强制生产模式，确保 cssVar 的生成使用 hash
        // https://github.com/dcloudio/uni-app/issues/4076
        // dev模式下，会生成：{ "83a5a03c-style.color": style.color}
        options.isProd = true;
        return compileScript(sfc, options);
    };
    // script + v-bind
    compiler.compileTemplate = (options) => {
        if (options?.compilerOptions) {
            ;
            options.compilerOptions.bindingCssVars =
                options.ssrCssVars || [];
        }
        // 同上
        options.isProd = true;
        return compileTemplate(options);
    };
}
/**
 * 重写 parse，解决相同内容被缓存，未触发 template 编译的问题
 */
function rewriteCompilerSfcParse() {
    const compilerSfc = require((0, uni_cli_shared_1.resolveBuiltIn)('@vue/compiler-sfc'));
    const { parse } = compilerSfc;
    compilerSfc.parse = (source, options) => {
        const res = parse(source, options);
        // template 中，先<view>hello</view>，然后修改为<view></view>，再恢复为<view>hello</view>，
        // 此时因为 descriptor 被缓存，不会触发 compileTemplate，故 parse 时，每次生成一个全新的 descriptor
        // https://github.com/vitejs/vite/blob/v2.9.13/packages/plugin-vue/src/script.ts#L44
        // https://github.com/dcloudio/uni-app/issues/3685
        res.descriptor = (0, shared_1.extend)({}, res.descriptor);
        return res;
    };
}
