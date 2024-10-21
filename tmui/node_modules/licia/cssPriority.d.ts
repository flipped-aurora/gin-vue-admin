import selector = require('./selector');

declare namespace cssPriority {
    function compare(p1: number[], p2: number[]): number;
}
declare function cssPriority(
    selector: string,
    options?: {
        important?: boolean;
        inlineStyle?: boolean;
        position?: number;
    }
): number[];

export = cssPriority;
