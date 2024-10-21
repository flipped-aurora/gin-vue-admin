declare const rc4: {
    encrypt(key: string, str: string): string;
    decrypt(key: string, str: string): string;
};

export = rc4;
