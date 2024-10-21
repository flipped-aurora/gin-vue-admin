import $safeEls = require('./$safeEls');

declare namespace $property {
    interface IProperty {
        (element: $safeEls.El, value: string): void;
        (element: $safeEls.El): string;
    }
}
declare const $property: {
    html: $property.IProperty;
    val: $property.IProperty;
    text: $property.IProperty;
};

export = $property;
