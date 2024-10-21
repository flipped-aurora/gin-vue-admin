import type { ResolvedConfig } from 'vite';
export * from './ast';
export * from './url';
export * from './plugin';
export * from './utils';
export declare const buildInCssSet: Set<string>;
export declare function isCombineBuiltInCss(config: ResolvedConfig): boolean;
