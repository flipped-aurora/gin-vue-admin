declare namespace selector {
    interface IToken {
        type: string;
        value: string;
    }
}
declare const selector: {
    parse(selector: string): Array<selector.IToken[]>;
    stringify(selector: Array<selector.IToken[]>): string;
};

export = selector;
