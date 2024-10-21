import lcidCodes = require('./lcid.json');

declare const lcid: {
	/**
	Get a [standard locale identifier](https://en.wikipedia.org/wiki/Locale_(computer_software)) from a [Windows locale identifier (LCID)](http://en.wikipedia.org/wiki/Locale#Specifics_for_Microsoft_platforms).

	@example
	```
	import lcid = require('lcid');

	lcid.from(1044);
	//=> 'nb_NO'
	```
	*/
	from(lcidCode: number): string;

	/**
	Get a [Windows locale identifier (LCID)](https://en.wikipedia.org/wiki/Locale#Specifics_for_Microsoft_platforms) from a [standard locale identifier](https://en.wikipedia.org/wiki/Locale_(computer_software)).

	@example
	```
	import lcid = require('lcid');

	lcid.to('nb_NO');
	//=> 1044
	```
	*/
	to(localeId: string): number;

	/**
	Mapping between [standard locale identifiers](https://en.wikipedia.org/wiki/Locale_(computer_software)) and [Windows locale identifiers (LCID)](http://en.wikipedia.org/wiki/Locale#Specifics_for_Microsoft_platforms).

	@example
	```
	import lcid = require('lcid');

	lcid.all;
	//=> {'af_ZA': 1078, …}
	```
	*/
	readonly all: typeof lcidCodes;
};

export = lcid;
