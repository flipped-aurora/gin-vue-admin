declare function highlight(
    str: string,
    lang?: string,
    style?: {
        comment?: string;
        string?: string;
        number?: string;
        keyword?: string;
        operator?: string;
    }
): string;

export = highlight;
