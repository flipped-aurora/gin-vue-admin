import type { Plugin } from 'vite';
import type { UniMiniProgramPluginOptions } from '.';
export declare function createConfigResolved({ cdn, style: { extname }, template: { component }, }: UniMiniProgramPluginOptions): Plugin['configResolved'];
