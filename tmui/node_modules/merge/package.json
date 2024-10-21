{
	"name": "merge",
	"version": "2.1.1",
	"author": "yeikos",
	"description": "(recursive)? merging of (cloned)? objects.",
	"main": "lib/src/index.js",
	"files": [
		"lib/src/index.d.ts"
	],
	"license": "MIT",
	"homepage": "https://github.com/yeikos/js.merge",
	"repository": {
		"type": "git",
		"url": "https://github.com/yeikos/js.merge.git"
	},
	"keywords": [
		"merge",
		"recursive",
		"extend",
		"clone",
		"object",
		"browser"
	],
	"scripts": {
		"build": "npm run build:ts && npm run build:wp",
		"dev": "./node_modules/.bin/concurrently --kill-others \"npm run dev:ts\" \"npm run dev:wp\"",
		"test": "./node_modules/.bin/mocha lib/test/index.js",
		"build:ts": "./node_modules/.bin/tsc -p tsconfig.json",
		"build:wp": "./node_modules/.bin/webpack --config webpack.config.js",
		"dev:ts": "./node_modules/.bin/tsc -p tsconfig.json -w",
		"dev:wp": "./node_modules/.bin/webpack --config webpack.config.js -w"
	},
	"devDependencies": {
		"@types/chai": "^4.2.14",
		"@types/mocha": "^8.0.4",
		"@types/node": "^14.14.7",
		"chai": "^4.2.0",
		"concurrently": "^5.3.0",
		"mocha": "^8.2.1",
		"typescript": "^4.0.5",
		"webpack": "^5.4.0",
		"webpack-cli": "^4.2.0"
	}
}
