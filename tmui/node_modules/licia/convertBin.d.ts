import type = require('./type');

declare namespace convertBin {
    function blobToArrBuffer(blob: any): Promise<ArrayBuffer>;
}
declare function convertBin(bin: any, type: string): any;

export = convertBin;
