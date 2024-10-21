interface ParseMiniProgramProjectJsonOptions {
    template: Record<string, any>;
    pagesJson: UniApp.PagesJson;
}
interface ProjectConfig {
    appid: string;
    projectname: string;
    condition?: {
        miniprogram?: UniApp.PagesJson['condition'];
    };
}
export declare function isMiniProgramProjectJsonKey(name: string): boolean;
export declare function parseMiniProgramProjectJson(jsonStr: string, platform: UniApp.PLATFORM, { template, pagesJson }: ParseMiniProgramProjectJsonOptions): ProjectConfig;
export {};
