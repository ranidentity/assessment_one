package middleware

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/siongui/gojianfan"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func GinBodyLogMiddleware(c *gin.Context) {
	lang, _ := c.GetQuery("lang")
	if lang == "zh_HANT" {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			//ok this is an request with error, let's make a record for it
			// now print body (or log in your preferred way)
			fmt.Println("Response body: " + blw.body.String())
		} else {
			data := blw.body.String()
			traditional_chinese := gojianfan.S2T(data)
			blw.ResponseWriter.WriteString(traditional_chinese)
			blw.body.Reset()
		}
	}
}
