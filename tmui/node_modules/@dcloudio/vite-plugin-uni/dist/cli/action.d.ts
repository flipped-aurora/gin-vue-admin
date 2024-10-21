import { type BuildOptions, type ServerOptions } from 'vite';
import type { CliOptions } from '.';
export declare function runDev(options: CliOptions & ServerOptions): Promise<undefined>;
export declare function runBuild(options: CliOptions & BuildOptions): Promise<void>;
export declare const stopProfiler: (log: (message: string) => void) => void | Promise<void>;
