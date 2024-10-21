import types = require('./types');

declare function compressImg(file: File | Blob | string, cb: types.AnyFn): void;
declare function compressImg(
    file: File | Blob | string,
    options?: {
        maxWidth?: number;
        maxHeight?: number;
        width?: number;
        height?: number;
        mimeType?: string;
        quality?: number;
    },
    cb?: types.AnyFn
): void;

export = compressImg;
