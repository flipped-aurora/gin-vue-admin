declare const html: {
    parse(html: string): any[];
    stringify(tree: any[]): string;
};

export = html;
