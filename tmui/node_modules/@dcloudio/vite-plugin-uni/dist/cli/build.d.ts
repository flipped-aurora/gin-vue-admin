import type { BuildOptions, InlineConfig } from 'vite';
import type { CliOptions } from '.';
import type { RollupWatcher } from 'rollup';
export declare function buildByVite(inlineConfig: InlineConfig): Promise<import("rollup").RollupOutput | import("rollup").RollupOutput[] | RollupWatcher>;
export declare function build(options: CliOptions): Promise<RollupWatcher | void>;
export declare function buildSSR(options: CliOptions): Promise<void>;
export declare function initBuildOptions(options: CliOptions, build: BuildOptions): InlineConfig;
export declare function buildApp(options: CliOptions): Promise<RollupWatcher | void>;
