declare function dateFormat(
    date: Date,
    mask: string,
    utc?: boolean,
    gmt?: boolean
): string;
declare function dateFormat(mask: string, utc?: boolean, gmt?: boolean): string;

export = dateFormat;
