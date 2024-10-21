export declare function hasThemeJson(themeLocation: string): boolean;
export declare const parseThemeJson: (themeLocation?: string) => UniApp.ThemeJson;
export declare const normalizeThemeConfigOnce: (manifestJsonPlatform?: Record<string, any>) => UniApp.ThemeJson;
export declare function initTheme<T extends object>(manifestJson: Record<string, any>, pagesJson: T): T;
export declare class ThemeSassParser {
    _index: number;
    _input: string;
    _theme: Record<string, Record<string, any>>;
    constructor();
    parse(input: string): Record<string, Record<string, any>>;
    parseVariable(): void;
    parseVariableValue(): any;
    parseFunction(): void | string[];
    skipOtherValue(): void;
    parseString(): string;
    pushThemeValue(key: string, valuePair: string[]): void;
    consume(expected: string): void;
    get currentChar(): string;
    skipWhiteSpaceAndComments(): void;
    skipComment(): void;
    skipWhiteSpace(): void;
    isspace(str: string): boolean;
}
