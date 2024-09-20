package runner

type InstallInfo struct {
	TempPath     string //软件安装时候临时目录，下载到该目录，然后copy到所属目录
	RootPath     string //存储根路径
	Pc           string //软件平台
	Name         string //软件名称
	FullName     string //软件名称,带后缀
	User         string //所属用户
	DownloadPath string //软件的云端地址
	InstallPath  string //安装后的所属目录
	Version      string //安装的软件版本
}

type UnInstallInfo struct {
}
