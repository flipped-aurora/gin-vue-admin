import type { Plugin, ResolvedConfig } from 'vite';
import type { UniMiniProgramPluginOptions } from '../plugin';
export declare function getNVueCssPaths(config: ResolvedConfig): string[] | undefined;
export declare function uniPagesJsonPlugin(options: UniMiniProgramPluginOptions): Plugin;
