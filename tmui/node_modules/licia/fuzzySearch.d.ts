declare function fuzzySearch(
    needle: string,
    haystack: any[],
    options?: {
        caseSensitive?: boolean;
        key?: string | string[];
    }
): any[];

export = fuzzySearch;
