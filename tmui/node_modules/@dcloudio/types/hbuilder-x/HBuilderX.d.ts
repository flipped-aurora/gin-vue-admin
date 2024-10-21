declare namespace string {
  /**
   * @description 元素上的属性
   * @module dom
   */
  interface AttrString extends String {}

  /**
   * @description 元素上某个属性的值
   * @module dom
   */
  interface AttrValueString extends String {}

  /**
   * @description 元素全局属性`class`的值
   * @module dom
   */
  interface ClassString extends String {}

  /**
   * @description 元素全局属性`id`的值
   * @module dom
   */
  interface IDString extends String {}

  /**
   * @description 元素上的事件
   * @module dom
   */
  interface HTMLEventString extends String {}

  /**
   * @description CSS颜色的值
   * @module dom
   */
  interface ColorString extends String {}

  /**
   * @description 提示common模块 以及js文件路径
   * @module vue
   */
  interface RequireCommonString extends String {}

  /**
   * @description 国际化翻译的key值
   * @module vue
   */
  interface VueI18NKeyString extends String {}

  /**
   * @description vue默认参数data中的属性名称
   * @module vue
   */
  interface VueDataString extends String {}

  /**
   * @description vue组件中ref属性的值
   * @module vue
   */
  interface VueRefString extends String {}

  /**
   * @description vuex 中 actions 的名称
   * @module vue
   */
  interface VuexDispatchString extends String {}

  /**
   * @description vuex 中 mutations 的名称
   * @module vue
   */
  interface VuexCommitString extends String {}

  /**
   * @description vue, nvue, uvue页面文件的文件路径(根据项目自动匹配)
   * @module vue
   */
  interface PageURIString extends String {}

  /**
   * @description nvue页面文件的文件路径
   * @module vue
   */
  interface NPageURIString extends String {}

  /**
   * @description uvue页面文件的文件路径, 仅在uniappx中生效
   * @module uniappx
   */
  interface UPageURIString extends String {}

  /**
   * @description video 组件的 id, 仅在uniappx中生效
   * @module uniappx
   */
  interface VideoIdString extends String {}

  /**
   * @description web-view 组件的 id, 仅在uniappx中生效
   * @module uniappx
   */
  interface WebviewIdString extends String {}

  /**
   * @description uniCloud db schema中parentKey的值
   * @module uniCloud
   */
  interface ParentFieldString extends String {}

  /**
   * @description uniCloud db schema中required数组的值
   * @module uniCloud
   */
  interface SchemaFieldString extends String {}

  /**
   * @description uniCloud db schema中validateFunction的值
   * @module uniCloud
   */
  interface ValidateFunctionString extends String {}

  /**
   * @description uniCloud 云函数名
   * @module uniCloud
   */
  interface CloudFunctionString extends String {}

  /**
   * @description uniCloud 云对象名
   * @module uniCloud
   */
  interface CloudObjectString extends String {}

  /**
   * @description uniCloud 数据库集合的名称
   * @module uniCloud
   */
  interface DBCollectionString extends String {}

  /**
   * @description uniCloud 数据库字段名称
   * @module uniCloud
   */
  interface DBFieldString extends String {}

  /**
   * @description uniCloud 数据库要操作的集合, 要查询的字段
   * @module uniCloud
   */
  interface JQLString extends String {}

  /**
   * @description CSS属性的名称
   * @module jQuery
   */
  interface cssPropertyString extends String {}

  /**
   * @description CSS某个属性的值
   * @module jQuery
   */
  interface cssPropertyValueString extends String {}

  /**
   * @description CSS选择器的名称
   * @module jQuery
   */
  interface cssSelectorString extends String {}

  /**
   * @description 任意文件的文件路径
   * @module uri
   */
  interface URIString extends String {}

  /**
   * @description css文件的文件路径(后缀为`.css`的文件路径)
   * @module uri
   */
  interface CSSURIString extends String {}

  /**
   * @description js文件的文件路径(后缀为`.js`的文件路径)
   * @module uri
   */
  interface JSURIString extends String {}

  /**
   * @description html文件的文件路径(后缀为`.html`的文件路径)
   * @module uri
   */
  interface HTMLURIString extends String {}

  /**
   * @description markdown文件的文件路径(后缀为`.md`的文件路径)
   * @module uri
   */
  interface MarkdownURIString extends String {}

  /**
   * @description js, ts, uts引用文件或模块的文件路径(支持vue,nvue,uvue中script标签内容), 例: `import xxx from 'xxx'`
   * @module uri
   */
  interface ScriptImportURIString extends String {}

  /**
   * @description css文件可以引用的文件的文件路径, 后缀为`[".css"]`的文件路径 例: `@import url('xxx.css')`
   * @module uri
   */
  interface CssImportURIString extends String {}

  /**
   * @description scss文件可以引用的文件的文件路径, 后缀为`[".scss", ".css"]`的文件路径, 例: `@import 'xxx.scss'`
   * @module uri
   */
  interface ScssImportURIString extends String {}

  /**
   * @description less文件可以引用的文件的文件路径, 后缀为`[".less", ".css"]`的文件路径, 例: `@import 'xxx.less'`
   * @module uri
   */
  interface LessImportURIString extends String {}

  /**
   * @description 字体文件的文件路径
   * @module uri
   */
  interface FontURIString extends String {}

  /**
   * @description 图片文件的文件路径
   * @module uri
   */
  interface ImageURIString extends String {}

  /**
   * @description 音频文件的文件路径
   * @module uri
   */
  interface AudioURIString extends String {}

  /**
   * @description 视频文件的文件路径
   * @module uri
   */
  interface VideoURIString extends String {}
}
