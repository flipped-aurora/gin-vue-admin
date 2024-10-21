import { defineConfig } from "vite";
import uni from "@dcloudio/vite-plugin-uni";
import vueJsx from "@vitejs/plugin-vue-jsx";
import { resolve } from "path"
import Components from "unplugin-vue-components/vite"
export default defineConfig({
	define: {
		__VUE_I18N_FULL_INSTALL__: true,
		__VUE_I18N_LEGACY_API__: true,
		__VUE_I18N_PROD_DEVTOOLS__: false,
	},
	build: {
		target: "es6"
	},
	css: {
		preprocessorOptions: {
			scss: {
				api: 'modern-compiler',
				silenceDeprecations: ['legacy-js-api']
			}
		}
	},
	resolve: {
		alias: [
			{
				find: "@",
				replacement: resolve(__dirname, 'src')
			}
		]
	},
	server: {
		port: 1314,
		// 选项写法
		proxy: {
			'/pag': {
				target: 'https://cdn.tmui.design',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/api/, '/api')
			},
		}
	},
	plugins: [
		uni(),
		vueJsx(),
		Components({
			dts: true,
			resolvers: [
				// example of importing Vant
				(componentName) => {
					if (componentName.startsWith('tm')) {
						return { name: componentName.slice(2), from: 'tm' }
					}
				},
			],
			dirs: ['./src/uni_modules/tm-ui/components'],
			include: [/\.vue$/, /\.uvue$/]
		})
	]
});