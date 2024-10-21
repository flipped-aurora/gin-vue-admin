declare function openFile(options?: {
    accept?: string;
    multiple?: boolean;
}): Promise<File[]>;

export = openFile;
