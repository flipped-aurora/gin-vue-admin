import mime = require('./mime');

declare function fileType(
    input: Buffer | ArrayBuffer | Uint8Array
):
    | {
          ext: string;
          mime: string;
      }
    | undefined;

export = fileType;
