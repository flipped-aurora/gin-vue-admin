import $safeEls = require('./$safeEls');

declare namespace $insert {
    type IInsert = (element: $safeEls.El, content: string | Element) => void;
}
declare const $insert: {
    before: $insert.IInsert;
    after: $insert.IInsert;
    append: $insert.IInsert;
    prepend: $insert.IInsert;
};

export = $insert;
