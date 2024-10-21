## 4.0.0

-  **New:** `variant` can have more than one value. Parsing the `variant` is too complex for this function, so any unknown tokens are assumed to be a `variant` and joined as one (e.g., `foo bold bar italic qux` parses a `variant` of `foo bar qux`).
-  **Fix:** Fixed `normal` overriding other declarations. The defaults are still `normal` and you can have any number of normals precede the `size` without overriding non-`normal` values.
-  **Breaking:** Removed parsing of CSS global keywords, as they don't seem to be supported. This is a fix, but it may break implementations that were relying on it.
-  **Breaking:** Error messaging has changed and new errors will throw when adding more than one `style`, `weight` or `stretch`.

## 3.0.4

-  Add BundlePhobia badges to npm (requires version bump).

## 3.0.3

-  Also exports a CommonJS module. Fixes `const parseCSSFont = require('parse-css-font');`.

## 3.0.2

-  Fix parsing of `1rem / 1.2` when the slash is surrounded by spaces.

## 3.0.1

-  Fix readme.
-  Add `package-lock.json`.

## 3.0.0

-  **Breaking:** Remove tcomb dependency. Breaking change if you were somehow relying on those errors.
-  Introduce TypeScript.
-  Provide TypeScript-generated type definitions.

## 2.0.2

-  Update dependencies.

## 2.0.1

-  Fix issue w/ slashes in functions.

## 2.0.0

-  Preserves functions with spaces and commas inside.

## 1.0.0

-  Initial release.
