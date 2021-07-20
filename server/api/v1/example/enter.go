package example

import "gin-vue-admin/service"

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
