declare const dataUrl: {
    parse(
        dataUrl: string
    ): { data: string; mime: string; charset: string; base64: boolean } | null;
    stringify(
        data: any,
        mime: string,
        options?: { base64?: boolean; charset?: string }
    ): string;
};

export = dataUrl;
