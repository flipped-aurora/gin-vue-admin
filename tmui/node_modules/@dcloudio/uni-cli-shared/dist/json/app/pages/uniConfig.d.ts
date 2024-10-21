export declare function normalizeAppUniConfig(pagesJson: UniApp.PagesJson, manifestJson: Record<string, any>): string;
export declare function parseEntryPagePath(pagesJson: UniApp.PagesJson): {
    entryPagePath: string;
    entryPageQuery: string;
    realEntryPagePath: string;
};
