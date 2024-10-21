import type { ServerOptions, ViteDevServer } from 'vite';
import type { CliOptions } from '.';
export declare function createServer(options: CliOptions & ServerOptions): Promise<ViteDevServer>;
export declare function createSSRServer(options: CliOptions & ServerOptions): Promise<ViteDevServer>;
