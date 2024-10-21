import type { Plugin } from 'vite';
import { type FilterPattern } from '@rollup/pluginutils';
interface UniEasycomPluginOptions {
    include?: FilterPattern;
    exclude?: FilterPattern;
}
export declare function uniEasycomPlugin(options: UniEasycomPluginOptions): Plugin;
export {};
