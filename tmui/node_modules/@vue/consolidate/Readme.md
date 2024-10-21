> This is a fork that removes the bluebird dependency to use native Node promises. Also requires Node 12+.

# Consolidate.js

  Template engine consolidation library.

## Installation

    $ npm install consolidate
    
## Supported Express versions

  - [Express 4.x](https://expressjs.com/fr/4x/api.html)
  - [Express 3.x](https://expressjs.com/fr/3x/api.html)

## Supported template engines

Some package has the same key name, consolidate will load them according to the order number. By example for dust, consolidate will try to use in this order: `dust`, `dustjs-helpers` and `dustjs-linkedin`. If `dust` is installed, `dustjs-linkedin` will not be used by consolidate.

| Name `cons.*` | Package Name / Order | Website / State |
| - | - | - |
| [atpl](https://github.com/soywiz/atpl.js) | [`npm install atpl`](https://www.npmjs.com/package/atpl) | - |
| [bracket](https://github.com/danlevan/bracket-template) | [`npm install bracket-template`](https://www.npmjs.com/package/bracket-template) | - |
| [dot](https://github.com/olado/doT) | [`npm install dot`](https://www.npmjs.com/package/dot) | [(website)](http://olado.github.io/doT/) |
| ~~[dust](https://github.com/akdubya/dustjs)~~ | [`npm install dust`](https://www.npmjs.com/package/dust) (1) | [(website)](http://akdubya.github.com/dustjs/) /  **(unmaintained)**<br>See: dustjs-linkedin |
| [dust](https://github.com/linkedin/dustjs) | [`npm install dustjs-helpers`](https://www.npmjs.com/package/dustjs-helpers) (2) or<br>[`npm install dustjs-linkedin`](https://www.npmjs.com/package/dustjs-linkedin) (3) | [(website)](http://linkedin.github.io/dustjs/) |
| ~~[eco](https://github.com/sstephenson/eco)~~ | [`npm install eco`](https://www.npmjs.com/package/eco) | **/!\ [Security issue](https://github.com/sstephenson/eco/pull/67)** |
| [ect](https://github.com/baryshev/ect) | [`npm install ect`](https://www.npmjs.com/package/ect) | [(website)](http://ectjs.com/) |
| [ejs](https://github.com/mde/ejs) | [`npm install ejs`](https://www.npmjs.com/package/ejs) | [(website)](http://ejs.co/) |
| [hamlet](https://github.com/gregwebs/hamlet.js) | [`npm install hamlet`](https://www.npmjs.com/package/hamlet) | - |
| [hamljs](https://github.com/visionmedia/haml.js) | [`npm install hamljs`](https://www.npmjs.com/package/hamljs) | - |
| [haml-coffee](https://github.com/netzpirat/haml-coffee) | [`npm install haml-coffee`](https://www.npmjs.com/package/haml-coffee) | - |
| [handlebars](https://github.com/wycats/handlebars.js/) | [`npm install handlebars`](https://www.npmjs.com/package/handlebars) | [(website)](http://handlebarsjs.com/) |
| [hogan](https://github.com/twitter/hogan.js) | [`npm install hogan.js`](https://www.npmjs.com/package/hogan.js) | [(website)](http://twitter.github.com/hogan.js/) |
| [htmling](https://github.com/codemix/htmling) | [`npm install htmling`](https://www.npmjs.com/package/htmling) | - |
| ~~[jade](https://github.com/visionmedia/jade)~~ | [`npm install jade`](https://www.npmjs.com/package/jade) | [(website)](http://jade-lang.com/) / **(renamed `pug`)** |
| [jazz](https://github.com/shinetech/jazz) | [`npm install jazz`](https://www.npmjs.com/package/jazz) | - |
| ~~[jqtpl](https://github.com/kof/jqtpl)~~ | [`npm install jqtpl`](https://www.npmjs.com/package/jqtpl) | **(deprecated)** |
| [just](https://github.com/baryshev/just) | [`npm install just`](https://www.npmjs.com/package/just) | - |
| [liquid](https://github.com/leizongmin/tinyliquid) | [`npm install tinyliquid`](https://www.npmjs.com/package/tinyliquid) | [(website)](http://liquidmarkup.org/)<br>**will never add any new features** |
| [liquor](https://github.com/chjj/liquor) | [`npm install liquor`](https://www.npmjs.com/package/liquor) | - |
| [lodash](https://github.com/bestiejs/lodash) | [`npm install lodash`](https://www.npmjs.com/package/lodash) | [(website)](http://lodash.com/) |
| [marko](https://github.com/marko-js/marko) | [`npm install marko`](https://www.npmjs.com/package/marko) | [(website)](http://markojs.com) |
| [mote](https://github.com/satchmorun/mote) | [`npm install mote`](https://www.npmjs.com/package/mote) | [(website)](http://satchmorun.github.io/mote/) |
| [mustache](https://github.com/janl/mustache.js) | [`npm install mustache`](https://www.npmjs.com/package/mustache) | - |
| [nunjucks](https://github.com/mozilla/nunjucks) | [`npm install nunjucks`](https://www.npmjs.com/package/nunjucks) | [(website)](https://mozilla.github.io/nunjucks) |
| [plates](https://github.com/flatiron/plates) | [`npm install plates`](https://www.npmjs.com/package/plates) | - |
| [pug](https://github.com/pugjs/pug) | [`npm install pug`](https://www.npmjs.com/package/pug) | [(website)](http://jade-lang.com/) / **(formerly jade)** |
| [qejs](https://github.com/jepso/QEJS) | [`npm install qejs`](https://www.npmjs.com/package/qejs) | - |
| [ractive](https://github.com/ractivejs/ractive) | [`npm install ractive`](https://www.npmjs.com/package/ractive) | - |
| [razor](https://github.com/kinogam/kino.razor) | [`npm install razor`](https://www.npmjs.com/package/razor) | - |
| [react](https://github.com/facebook/react) | [`npm install react`](https://www.npmjs.com/package/react) | - |
| [slm](https://github.com/slm-lang/slm) | [`npm install slm`](https://www.npmjs.com/package/slm) | - |
| [squirrelly](https://github.com/squirrellyjs/squirrelly) | [`npm install squirrelly`](https://www.npmjs.com/package/squirrelly) | [(website)](https://squirrelly.js.org) |
| ~~[swig](https://github.com/paularmstrong/swig)~~ | [`npm install swig`](https://www.npmjs.com/package/swig) (1) | **(unmaintained)**<br>See: swig-templates |
| [swig](https://github.com/node-swig/swig-templates) | [`npm install swig-templates`](https://www.npmjs.com/package/swig-templates) (2) | - |
| [teacup](https://github.com/goodeggs/teacup) | [`npm install teacup`](https://www.npmjs.com/package/teacup) | - |
| [templayed](https://github.com/archan937/templayed.js/) | [`npm install templayed`](https://www.npmjs.com/package/templayed) | [(website)](http://archan937.github.com/templayed.js/)  |
| [toffee](https://github.com/malgorithms/toffee) | [`npm install toffee`](https://www.npmjs.com/package/toffee) | - |
| [twig](https://github.com/justjohn/twig.js) | [`npm install twig`](https://www.npmjs.com/package/twig) | [(wiki)](https://github.com/twigjs/twig.js/wiki/Implementation-Notes) |
| [twing](https://github.com/NightlyCommit/twing) | [`npm install twing`](https://www.npmjs.com/package/twing) | [(website)](https://nightlycommit.github.io/twing/) |
| [underscore](https://github.com/documentcloud/underscore) | [`npm install underscore`](https://www.npmjs.com/package/underscore) | [(website)](http://underscorejs.org/#template) |
| [vash](https://github.com/kirbysayshi/vash) | [`npm install vash`](https://www.npmjs.com/package/vash) | - |
| [velocityjs](https://github.com/julianshapiro/velocity) | [BETA](https://www.npmjs.com/package/velocity-animate) | [(website)](http://velocityjs.org/) |
| [walrus](https://github.com/jeremyruppel/walrus) | [`npm install walrus`](https://www.npmjs.com/package/walrus) | [(website)](http://documentup.com/jeremyruppel/walrus/) |
| [whiskers](https://github.com/gsf/whiskers.js) | [`npm install whiskers`](https://www.npmjs.com/package/whiskers) | - | 

__NOTE__: you must still install the engines you wish to use, add them to your package.json dependencies.

## API

  All templates supported by this library may be rendered using the signature `(path[, locals], callback)` as shown below, which happens to be the signature that Express supports so any of these engines may be used within Express.

__NOTE__: All this example code uses cons.swig for the swig template engine. Replace swig with whatever templating you are using. For example, use cons.hogan for hogan.js, cons.jade for jade, etc. `console.log(cons)` for the full list of identifiers.

```js
var cons = require('consolidate');
cons.swig('views/page.html', { user: 'tobi' }, function(err, html){
  if (err) throw err;
  console.log(html);
});
```

  Or without options / local variables:

```js
var cons = require('consolidate');
cons.swig('views/page.html', function(err, html){
  if (err) throw err;
  console.log(html);
});
```

  To dynamically pass the engine, simply use the subscript operator and a variable:

```js
var cons = require('consolidate')
  , name = 'swig';

cons[name]('views/page.html', { user: 'tobi' }, function(err, html){
  if (err) throw err;
  console.log(html);
});
```

### Promises

  Additionally, all templates optionally return a promise if no callback function is provided. The promise represents the eventual result of the template function which will either resolve to a string, compiled from the template, or be rejected. Promises expose a `then` method which registers callbacks to receive the promise’s eventual value and a `catch` method which the reason why the promise could not be fulfilled. Promises allow more synchronous-like code structure and solve issues like race conditions.

```js
var cons = require('consolidate');

cons.swig('views/page.html', { user: 'tobi' })
  .then(function (html) {
    console.log(html);
  })
  .catch(function (err) {
    throw err;
  });
```

## Caching

 To enable caching simply pass `{ cache: true }`. Engines _may_ use this option to cache things reading the file contents, compiled `Function`s etc. Engines which do _not_ support this may simply ignore it. All engines that consolidate.js implements I/O for will cache the file contents, ideal for production environments.
 When using consolidate directly: `cons.swig('views/page.html', { user: 'tobi', cache:true }, callback);`
 Using supported Express versions: `app.locals.cache = true` or set NODE_ENV to 'production' and Express will do this for you.

## Express example

```js
var express = require('express')
  , cons = require('consolidate')
  , app = express();

// assign the swig engine to .html files
app.engine('html', cons.swig);

// set .html as the default extension
app.set('view engine', 'html');
app.set('views', __dirname + '/views');

var users = [];
users.push({ name: 'tobi' });
users.push({ name: 'loki' });
users.push({ name: 'jane' });

app.get('/', function(req, res){
  res.render('index', {
    title: 'Consolidate.js'
  });
});

app.get('/users', function(req, res){
  res.render('users', {
    title: 'Users',
    users: users
  });
});

app.listen(3000);
console.log('Express server listening on port 3000');
```

## Template Engine Instances

Template engines are exposed via the `cons.requires` object, but they are not instantiated until you've called the `cons[engine].render()` method. You can instantiate them manually beforehand if you want to add filters, globals, mixins, or other engine features.

```js
var cons = require('consolidate'),
  nunjucks = require('nunjucks');

// add nunjucks to requires so filters can be
// added and the same instance will be used inside the render method
cons.requires.nunjucks = nunjucks.configure();

cons.requires.nunjucks.addFilter('foo', function () {
  return 'bar';
});
```

## Notes

* If you're using Nunjucks, please take a look at the `exports.nunjucks.render` function in `lib.consolidate.js`.  You can pass your own engine/environment via `options.nunjucksEnv`, or if you want to support Express you can pass `options.settings.views`, or if you have another use case, pass `options.nunjucks` (see the code for more insight).
* You can pass **partials** with `options.partials`
* For using **template inheritance** with nunjucks, you can pass a loader
  with `options.loader`.
* To use **filters** with tinyliquid, use `options.filters` and specify an array of properties, each of which is a named filter function. A filter function takes a string as a parameter and returns a modified version of it.
* To use **custom tags** with tinyliquid, use `options.customTags` to specify an array of tag functions that follow the tinyliquid [custom tag](https://github.com/leizongmin/tinyliquid/wiki/Custom-Tag) definition.
* The default directory used with the **include** tag with tinyliquid is the current working directory. To override this, use `options.includeDir`.
* `React` To render content into a html base template (eg. `index.html` of your React app), pass the path of the template with `options.base`.

## Running tests

  Install dev deps:

    $ npm install -d

  Run the tests:

    $ make test

## License

(The MIT License)

Copyright (c) 2011-2016 TJ Holowaychuk &lt;tj@vision-media.ca&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
