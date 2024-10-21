import type { Plugin, ResolvedConfig } from 'vite';
import type { UniPluginFilterOptions } from '.';
/**
 * 提供static等目录静态资源加载
 * @param _options
 * @param config
 * @returns
 */
export declare function uniStaticPlugin(_options: UniPluginFilterOptions, config: ResolvedConfig): Plugin;
