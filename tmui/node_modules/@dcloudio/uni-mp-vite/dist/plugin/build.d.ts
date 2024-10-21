import type { BuildOptions, UserConfig } from 'vite';
export declare function buildOptions(): UserConfig['build'];
export declare function createBuildOptions(inputDir: string, platform: UniApp.PLATFORM): BuildOptions;
export declare function notFound(filename: string): never;
