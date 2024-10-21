import type { Plugin } from 'vite';
interface UniMovePluginOptions {
    apply: Plugin['apply'];
    enforce: Plugin['enforce'];
    /**
     * 原始根目录
     */
    cwd: string;
    /**
     * glob pattern 如：**\/*.js.mp)
     */
    pattern: string;
    /**
     * 目标目录
     */
    dest: string;
}
export declare function uniMovePlugin({ apply, enforce, cwd, pattern, dest, }: UniMovePluginOptions): Plugin;
export {};
