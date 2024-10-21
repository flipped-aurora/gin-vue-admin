declare function download(
    data: Blob | File | string | any[],
    name: string,
    type?: string
): void;

export = download;
