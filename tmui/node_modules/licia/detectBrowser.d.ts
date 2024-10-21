declare function detectBrowser(
    ua?: string
): {
    name: string;
    version: number;
};

export = detectBrowser;
