package init

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	// 我们一般使用系统时间的不确定性来进行初始化
	rand.Seed(time.Now().UnixNano())
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
