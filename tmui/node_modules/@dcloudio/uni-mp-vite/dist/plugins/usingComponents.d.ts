import type { Plugin } from 'vite';
import type { SFCScriptCompileOptions } from '@vue/compiler-sfc';
export declare function uniUsingComponentsPlugin(options?: {
    normalizeComponentName?: (name: string) => string;
    babelParserPlugins?: SFCScriptCompileOptions['babelParserPlugins'];
}): Plugin;
export declare function dynamicImport(name: string, value: string): string;
