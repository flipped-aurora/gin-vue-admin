import type { AppJson, NetworkTimeout, PageWindowOptions } from './types';
interface ParsePagesJsonOptions {
    debug?: boolean;
    darkmode?: boolean;
    subpackages: boolean;
    windowOptionsMap?: Record<string, string>;
    tabBarOptionsMap?: Record<string, string>;
    tabBarItemOptionsMap?: Record<string, string>;
    networkTimeout?: NetworkTimeout;
}
export declare function parseMiniProgramPagesJson(jsonStr: string, platform: UniApp.PLATFORM, options?: ParsePagesJsonOptions): {
    appJson: AppJson;
    pageJsons: Record<string, PageWindowOptions>;
    nvuePages: string[];
};
export declare function mergeMiniProgramAppJson(appJson: Record<string, any>, platformJson?: Record<string, any>): void;
export {};
