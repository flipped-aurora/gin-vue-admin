declare const fs: {
    readFile(path: string, encoding: string): Promise<string>;
    readFile(path: string): Promise<Buffer>;
    exists(path: string): Promise<boolean>;
    unlink(path: string): Promise<void>;
    writeFile(path: string, data: string, options?: string): Promise<void>;
    writeFile(path: string, data: Buffer): Promise<void>;
    readdir(path: string): Promise<string[]>;
    rmdir(path: string): Promise<void>;
    [key: string]: any;
};

export = fs;
