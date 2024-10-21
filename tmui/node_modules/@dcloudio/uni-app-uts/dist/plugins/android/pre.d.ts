import type { Plugin } from 'vite';
import { type FilterPattern } from '@rollup/pluginutils';
export interface UniPrePluginOptions {
    include?: FilterPattern;
    exclude?: FilterPattern;
}
export declare function uniPrePlugin(options?: UniPrePluginOptions): Plugin;
