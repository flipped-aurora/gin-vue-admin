import type { ResolverFunction, UserConfig } from 'vite';
import type { VitePluginUniResolvedOptions } from '..';
export declare const customResolver: ResolverFunction;
export declare function createResolve(options: VitePluginUniResolvedOptions, _config: UserConfig): UserConfig['resolve'];
