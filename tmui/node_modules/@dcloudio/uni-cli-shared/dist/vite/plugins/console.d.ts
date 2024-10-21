import type { Plugin } from 'vite';
import { type FilterPattern } from '@rollup/pluginutils';
export interface ConsoleOptions {
    method: string;
    filename?: (filename: string) => string;
    include?: FilterPattern;
    exclude?: FilterPattern;
}
export declare function uniConsolePlugin(options: ConsoleOptions): Plugin;
