package ksc

import (
	"github.com/gin-gonic/gin"
	"github.com/myksc/ksc-base/golib/env"
)

func main(){
	if(env.GetRunEnv() == env.RunEnvOnline){
		gin.SetMode(gin.ReleaseMode)
	}


}
