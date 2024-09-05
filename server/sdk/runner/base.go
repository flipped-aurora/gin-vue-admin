package runner

var GetMap = map[string]func(runner *Context){
	"ping": func(c *Context) {
		c.ResponseOkWithText("ok")
	},
}
var POSTMap = map[string]func(runner *Context){}
