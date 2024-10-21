export declare const parseManifestJson: (inputDir: string) => any;
export declare const parseManifestJsonOnce: (inputDir: string) => any;
export declare const parseRpx2UnitOnce: (inputDir: string, platform?: UniApp.PLATFORM) => any;
interface CompilerCompatConfig {
    MODE?: 2 | 3;
}
declare function parseCompatConfig(_inputDir: string): CompilerCompatConfig;
export declare const parseCompatConfigOnce: typeof parseCompatConfig;
declare const defaultNetworkTimeout: {
    request: number;
    connectSocket: number;
    uploadFile: number;
    downloadFile: number;
};
export declare function normalizeNetworkTimeout(networkTimeout?: Partial<typeof defaultNetworkTimeout>): {
    request: number;
    connectSocket: number;
    uploadFile: number;
    downloadFile: number;
};
export declare function getUniStatistics(inputDir: string, platform: UniApp.PLATFORM): any;
export declare function isEnableUniPushV1(inputDir: string, platform: UniApp.PLATFORM): boolean;
export declare function isEnableUniPushV2(inputDir: string, platform: UniApp.PLATFORM): boolean;
export declare function isEnableSecureNetwork(inputDir: string, platform: UniApp.PLATFORM): boolean;
export declare function hasPushModule(inputDir: string): boolean;
export declare function isUniPushOffline(inputDir: string): boolean;
export declare function getRouterOptions(manifestJson: Record<string, any>): {
    mode?: 'history' | 'hash';
    base?: string;
};
export declare function isEnableTreeShaking(manifestJson: Record<string, any>): boolean;
export declare function getDevServerOptions(manifestJson: Record<string, any>): any;
export declare function getPlatformManifestJson(manifestJson: any, platform?: UniApp.PLATFORM): any;
export declare function getPlatformManifestJsonOnce(): any;
export declare function validateThemeValue(value: string): boolean;
export {};
