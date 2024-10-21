/*
  @license
	Rollup.js v4.24.0
	Wed, 02 Oct 2024 09:36:48 GMT - commit d3c000f4fd453e39a354299f0cfaa6831f56d7d8

	https://github.com/rollup/rollup

	Released under the MIT License.
*/
export { version as VERSION, defineConfig, rollup, watch } from './shared/node-entry.js';
import './shared/parseAst.js';
import '../native.js';
import 'node:path';
import 'path';
import 'node:process';
import 'node:perf_hooks';
import 'node:fs/promises';
import 'tty';
