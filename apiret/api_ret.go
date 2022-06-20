package apiret

import "github.com/gin-gonic/gin"

const StatusCodeOk int = 200
const StatusCode204 int = 204
const ResultOK int = 0
const ResultParamFailed = 1
const ResultServiceFailed = 2
const ResultServerWarn = 3
const ResultTokenFailed = 4
const ResultTokenExpireTime = 10000

// CJsonOk the http api handle ok
func CJsonOk(c *gin.Context, message string) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultOK,
		"message": message,
	})
}

// CJsonOkData the http api handle ok, and return data
func CJsonOkData(c *gin.Context, message string, data interface{}) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultOK,
		"message": message,
		"data":    data,
	})
}

// CJsonParamFailed parse param failed
func CJsonParamFailed(c *gin.Context, message string) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultParamFailed,
		"message": message,
	})
}

// CJsonServerFailed parse param failed
func CJsonServerFailed(c *gin.Context, message string) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultParamFailed,
		"message": message,
	})
}

func CJsonServerWarn(c *gin.Context, message string) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultServerWarn,
		"message": message,
	})
}

func CJsonServerWarnData(c *gin.Context, message string, data interface{}) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultServerWarn,
		"message": message,
		"data":    data,
	})
}

func CJsonTokenFailed(c *gin.Context, message string) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultTokenFailed,
		"message": message,
	})
}

func CJsonTokenExpire(c *gin.Context) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultTokenExpireTime,
		"message": "token is expire time",
	})
}

func CJsonTokenUserIsLogined(c *gin.Context) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultTokenFailed,
		"message": "user is logining in other place",
	})
}

// CJson206 parse tb service failed
func CJson206(c *gin.Context, message string, data interface{}) {
	c.JSON(StatusCodeOk, gin.H{
		"result":  ResultServiceFailed,
		"message": message,
		"data":    data,
	})
}
