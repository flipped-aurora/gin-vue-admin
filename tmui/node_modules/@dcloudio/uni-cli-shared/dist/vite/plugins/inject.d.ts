import type { Plugin } from 'vite';
import { type FilterPattern } from '@rollup/pluginutils';
type Injectment = string | [string, string];
export interface InjectOptions {
    enforce?: 'pre' | 'post';
    sourceMap?: boolean;
    callback?: (imports: Map<any, any>, mod: [string, string]) => void;
    include?: FilterPattern;
    exclude?: FilterPattern;
    [str: string]: Injectment | FilterPattern | Boolean | Function | undefined;
}
export declare function uniViteInjectPlugin(name: string, options: InjectOptions): Plugin;
export {};
