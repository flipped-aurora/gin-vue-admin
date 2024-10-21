import type { Plugin } from 'vite';
import type { UniMiniProgramPluginOptions } from '../plugin';
export declare function virtualPagePath(filepath: string): string;
export declare function virtualComponentPath(filepath: string): string;
export declare function parseVirtualPagePath(uniPageUrl: string): string;
export declare function parseVirtualComponentPath(uniComponentUrl: string): string;
export declare function isUniPageUrl(id: string): boolean;
export declare function isUniComponentUrl(id: string): boolean;
export declare function uniEntryPlugin({ global, }: UniMiniProgramPluginOptions): Plugin;
