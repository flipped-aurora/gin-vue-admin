declare const vlq: {
    encode(number: number | number[]): string;
    decode(string: string): number[];
};

export = vlq;
