export declare const rewriteCompileScriptOnce: typeof rewriteCompileScript;
export declare const rewriteCompilerSfcParseOnce: typeof rewriteCompilerSfcParse;
declare function rewriteCompileScript(): void;
/**
 * 重写 parse，解决相同内容被缓存，未触发 template 编译的问题
 */
declare function rewriteCompilerSfcParse(): void;
export {};
