package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
	ExcelApi
	FileUploadAndDownloadApi
	SimpleUploaderApi
}

var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
var customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
var excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
var simpleUploaderService = service.ServiceGroupApp.ExampleServiceGroup.SimpleUploaderService
