declare const ini: {
    parse(ini: string): any;
    stringify(
        obj: any,
        options?: {
            section?: string;
            whitespace: boolean;
        }
    ): string;
};

export = ini;
