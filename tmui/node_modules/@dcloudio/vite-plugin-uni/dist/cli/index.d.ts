import type { LogLevel } from 'vite';
export interface CliOptions {
    '--'?: string[];
    c?: boolean | string;
    config?: string;
    platform?: string;
    p?: string;
    ssr?: boolean;
    base?: string;
    debug?: boolean | string;
    d?: boolean | string;
    filter?: string;
    f?: string;
    logLevel?: LogLevel;
    l?: LogLevel;
    m?: string;
    mode?: string;
    clearScreen?: boolean;
    autoHost?: string;
    autoPort?: number;
    devtools?: boolean;
    devtoolsHost?: string;
    devtoolsPort?: number;
    subpackage?: string;
    plugin?: boolean;
}
