declare function parseHtml(
    html: string,
    handlers: {
        start?: (tag: string, attrs: any, unary: boolean) => void;
        end?: (tag: string) => void;
        comment?: (text: string) => void;
        text?: (text: string) => void;
    }
): void;

export = parseHtml;
