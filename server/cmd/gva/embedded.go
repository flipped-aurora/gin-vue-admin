package main

// embeddedManifest 默认为空，CLI 运行时从 --manifest 或配置文件加载。
// 后端编译“开箱即用”版本时，会用 //go:embed 版本覆盖本文件，
// 把 manifest 内嵌进二进制，从而无需额外文件即可运行。
var embeddedManifest []byte
