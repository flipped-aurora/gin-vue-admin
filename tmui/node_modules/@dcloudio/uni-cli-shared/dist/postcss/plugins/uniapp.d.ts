import type { Plugin, Root } from 'postcss';
export interface UniAppCssProcessorOptions {
    unit?: string;
    unitRatio?: number;
    unitPrecision?: number;
}
export declare function filterPrefersColorScheme(root: Root, force?: boolean): void;
declare const uniapp: {
    (opts?: UniAppCssProcessorOptions): Plugin;
    postcss: boolean;
};
export default uniapp;
