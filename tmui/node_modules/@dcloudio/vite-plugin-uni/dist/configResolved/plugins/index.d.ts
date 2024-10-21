import type { ResolvedConfig } from 'vite';
import type { FilterPattern } from '@rollup/pluginutils';
import type { VitePluginUniResolvedOptions } from '../..';
export interface UniPluginFilterOptions extends VitePluginUniResolvedOptions {
    include?: FilterPattern;
    exclude?: FilterPattern;
}
export declare function initPlugins(config: ResolvedConfig, options: VitePluginUniResolvedOptions): void;
