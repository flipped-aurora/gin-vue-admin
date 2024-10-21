import type { BuildOptions } from 'esbuild';
export declare function transformWithEsbuild(code: string, filename: string, options: BuildOptions): Promise<import("esbuild").BuildResult<BuildOptions>>;
export declare function esbuild(options: BuildOptions): Promise<import("esbuild").BuildResult<BuildOptions>>;
