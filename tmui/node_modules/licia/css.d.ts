declare const css: {
    parse(css: string): object;
    stringify(stylesheet: object, options?: { indent?: string }): string;
};

export = css;
