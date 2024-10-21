declare function isHidden(
    el: Element,
    options?: {
        display?: boolean;
        visibility?: boolean;
        opacity?: boolean;
        size?: boolean;
        viewport?: boolean;
        overflow?: boolean;
    }
): boolean;

export = isHidden;
