declare class Trie {
    add(word: string): void;
    remove(word: string): void;
    has(word: string): boolean;
    words(prefix: string): string[];
    clear(): void;
}

export = Trie;
