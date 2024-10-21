import type { SFCDescriptor } from '@vue/compiler-sfc';
import type { ExistingRawSourceMap, RollupError } from 'rollup';
import type * as PostCSS from 'postcss';
import type { Plugin } from '../plugin';
import type { ResolvedConfig } from '../config';
import type { ResolveFn } from '../';
import * as Postcss from 'postcss';
export interface CSSOptions {
    /**
     * https://github.com/css-modules/postcss-modules
     */
    modules?: CSSModulesOptions | false;
    preprocessorOptions?: Record<string, any>;
    postcss?: string | (Postcss.ProcessOptions & {
        plugins?: Postcss.Plugin[];
    });
}
export interface CSSModulesOptions {
    getJSON?: (cssFileName: string, json: Record<string, string>, outputFileName: string) => void;
    scopeBehaviour?: 'global' | 'local';
    globalModulePaths?: RegExp[];
    generateScopedName?: string | ((name: string, filename: string, css: string) => string);
    hashPrefix?: string;
    /**
     * default: null
     */
    localsConvention?: 'camelCase' | 'camelCaseOnly' | 'dashes' | 'dashesOnly' | null;
}
export declare const cssLangRE: RegExp;
export declare const commonjsProxyRE: RegExp;
export declare const isCSSRequest: (request: string) => boolean;
export declare const isDirectCSSRequest: (request: string) => boolean;
/**
 * Plugin applied before user plugins
 */
export declare function cssPlugin(config: ResolvedConfig, options?: {
    isAndroidX: boolean;
    getDescriptor?(filename: string): SFCDescriptor | undefined;
    createUrlReplacer?: (resolve: ResolveFn) => CssUrlReplacer;
}): Plugin;
/**
 * Plugin applied after user plugins
 */
export declare function cssPostPlugin(config: ResolvedConfig, { platform, isJsCode, preserveModules, chunkCssFilename, chunkCssCode, includeComponentCss, }: {
    platform: UniApp.PLATFORM;
    isJsCode?: boolean;
    preserveModules?: boolean;
    chunkCssFilename: (id: string) => string | void;
    chunkCssCode: (filename: string, cssCode: string) => Promise<string> | string;
    includeComponentCss?: boolean;
}): Plugin;
export declare function formatPostcssSourceMap(rawMap: ExistingRawSourceMap, file: string): ExistingRawSourceMap;
export type CssUrlReplacer = (url: string, importer?: string, source?: PostCSS.Source) => string | Promise<string>;
export declare const cssUrlRE: RegExp;
export declare const cssDataUriRE: RegExp;
export declare const importCssRE: RegExp;
export declare function minifyCSS(css: string, config: ResolvedConfig): Promise<string>;
export declare function hoistAtRules(css: string): Promise<string>;
export interface StylePreprocessorResults {
    code: string;
    map?: ExistingRawSourceMap | undefined;
    additionalMap?: ExistingRawSourceMap | undefined;
    errors: RollupError[];
    deps: string[];
}
/**
 * 重写 readFileSync
 * 目前主要解决 scss 文件被 @import 的条件编译
 */
export declare function rewriteScssReadFileSync(): void;
