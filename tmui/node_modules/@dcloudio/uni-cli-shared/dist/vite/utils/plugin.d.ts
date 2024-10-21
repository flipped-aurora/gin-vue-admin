import type { Plugin, ResolveFn, ResolvedConfig } from 'vite';
import { type CssUrlReplacer } from '../plugins/vitejs/plugins/css';
export type CreateUniViteFilterPlugin = (opts: UniViteFilterPluginOptions) => Plugin;
export interface UniViteFilterPluginOptions {
    resolvedConfig: ResolvedConfig;
    filter: (id: string) => boolean;
}
export declare function injectAssetPlugin(config: ResolvedConfig, options?: {
    isAndroidX: boolean;
}): void;
export declare function injectCssPlugin(config: ResolvedConfig, options?: {
    createUrlReplacer?: (resolve: ResolveFn) => CssUrlReplacer;
}): void;
export declare function injectCssPostPlugin(config: ResolvedConfig, newCssPostPlugin: Plugin): void;
export declare function replacePlugins(plugins: Plugin[], config: ResolvedConfig): void;
export declare function removePlugins(plugins: string | string[], config: ResolvedConfig): void;
export declare function insertBeforePlugin(plugin: Plugin, before: string, config: ResolvedConfig): void;
