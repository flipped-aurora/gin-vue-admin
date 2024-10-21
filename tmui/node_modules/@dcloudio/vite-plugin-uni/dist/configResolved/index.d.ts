import type { Plugin, ResolvedConfig } from 'vite';
import type { VitePluginUniResolvedOptions } from '..';
export declare function createConfigResolved(options: VitePluginUniResolvedOptions): Plugin['configResolved'];
export declare function initLogger({ logger, nvue, }: ResolvedConfig & {
    nvue?: boolean;
}): void;
