declare function onChange(callback: Function): void;
declare function offChange(callback: Function): void;
declare var safeAreaInsets: {
    readonly support: boolean;
    readonly top: number;
    readonly left: number;
    readonly right: number;
    readonly bottom: number;
    onChange: typeof onChange;
    offChange: typeof offChange;
};
export = safeAreaInsets;
