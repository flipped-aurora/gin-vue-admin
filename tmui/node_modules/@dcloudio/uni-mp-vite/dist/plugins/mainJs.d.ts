import type { SFCScriptCompileOptions } from '@vue/compiler-sfc';
export declare function uniMainJsPlugin(options?: {
    normalizeComponentName?: (name: string) => string;
    babelParserPlugins?: SFCScriptCompileOptions['babelParserPlugins'];
}): import("vite").Plugin<any>;
