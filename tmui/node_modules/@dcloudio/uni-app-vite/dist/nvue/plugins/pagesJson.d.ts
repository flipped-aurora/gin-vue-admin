import type { Plugin } from 'vite';
interface NVuePages {
    [filename: string]: {
        disableScroll?: boolean;
        scrollIndicator?: 'none';
    };
}
export declare const nvuePagesCache: Map<Readonly<Omit<import("vite").UserConfig, "plugins" | "css" | "assetsInclude" | "optimizeDeps" | "worker" | "build"> & {
    configFile: string | undefined;
    configFileDependencies: string[];
    inlineConfig: import("vite").InlineConfig;
    root: string;
    base: string;
    publicDir: string;
    cacheDir: string;
    command: "build" | "serve";
    mode: string;
    isWorker: boolean;
    isProduction: boolean;
    envDir: string;
    env: Record<string, any>;
    resolve: Required<import("vite").ResolveOptions> & {
        alias: import("vite").Alias[];
    };
    plugins: readonly Plugin<any>[];
    css: import("vite").ResolvedCSSOptions;
    esbuild: false | import("vite").ESBuildOptions;
    server: import("vite").ResolvedServerOptions;
    build: import("vite").ResolvedBuildOptions;
    preview: import("vite").ResolvedPreviewOptions;
    ssr: import("vite").ResolvedSSROptions;
    assetsInclude: (file: string) => boolean;
    logger: import("vite").Logger;
    createResolver: (options?: Partial<import("vite").InternalResolveOptions> | undefined) => import("vite").ResolveFn;
    optimizeDeps: import("vite").DepOptimizationOptions;
    worker: import("vite").ResolvedWorkerOptions;
    appType: import("vite").AppType;
    experimental: import("vite").ExperimentalOptions;
} & import("vite").PluginHookUtils>, NVuePages>;
export declare function parseNVuePageOptions(filename: string): {
    disableScroll?: boolean | undefined;
    scrollIndicator?: "none" | undefined;
};
export declare function uniPagesJsonPlugin({ renderer, appService, }: {
    renderer?: 'native';
    appService: boolean;
}): Plugin;
export {};
