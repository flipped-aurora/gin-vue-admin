import type { Plugin } from 'vite';
interface UniUTSPluginOptions {
    x?: boolean;
    extApis?: Record<string, [string, string]>;
    isSingleThread?: boolean;
}
export declare function getCurrentCompiledUTSPlugins(): Set<string>;
export declare function uniUTSAppUniModulesPlugin(options?: UniUTSPluginOptions): Plugin;
export declare function buildUniExtApis(): Promise<void>;
export declare function resolveExtApiProvider(pkg: Record<string, any>): {
    name?: string | undefined;
    plugin?: string | undefined;
    service: string;
    servicePlugin: string;
} | undefined;
export declare function uniDecryptUniModulesPlugin(): Plugin;
export {};
