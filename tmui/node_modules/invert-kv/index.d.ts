// TODO: Extend this to symbol when TS allows symbols in index signatures:
// https://github.com/Microsoft/TypeScript/issues/1863
/**
Invert the key/value of an object. Example: `{foo: 'bar'}` → `{bar: 'foo'}`.

@example
```
import invertKeyValue = require('invert-kv');

invertKeyValue({foo: 'bar', '🦄': '🌈'});
//=> {bar: 'foo', '🌈': '🦄'}
```
*/
declare function invertKeyValue<
	KeyType extends string | number,
	ValueType extends string | number | symbol
>(
	object: {[key in KeyType]: ValueType}
): {[key in ValueType]: KeyType extends number ? Exclude<KeyType, number> | string : KeyType};

export = invertKeyValue;
