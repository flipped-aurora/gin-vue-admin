export declare function parseScripts(name: string, pkgPath: string): {
    name: string;
    platform: keyof UniApp.PagesJsonPagePlatformStyle;
    define: {
        [name: string]: string;
    };
    context: {
        [name: string]: boolean;
    };
} | undefined;
