import type { Plugin, ResolvedConfig } from 'vite';
export declare function createConfigResolved({ createCssPostPlugin, }: {
    createCssPostPlugin: (config: ResolvedConfig) => Plugin;
}): Plugin['configResolved'];
