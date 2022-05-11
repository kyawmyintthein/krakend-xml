package gin

import (
	"github.com/clbanning/mxj"
	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/proxy"
)

const (
	_header = "<?xml version='1.0' standalone='yes'?>"
)

// Render marshals the proxy response and passes the resulting xml to the response writer
func Render(c *gin.Context, response *proxy.Response) {
	status := c.Writer.Status()
	if response == nil {
		c.XML(status, nil)
		return
	}
	mxj.XMLEscapeChars(true)
	mv := mxj.Map(response.Data)
	data, _ := mv.Xml()
	data = []byte(_header + string(data))

	c.Header("Content-Type", gin.MIMEXML)
	c.Writer.Write(data)
}
