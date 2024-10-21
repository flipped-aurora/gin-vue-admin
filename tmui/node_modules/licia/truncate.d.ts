declare function truncate(
    txt: string,
    width: number,
    options?: {
        ellipsis?: string;
        separator: string;
    }
): string;

export = truncate;
