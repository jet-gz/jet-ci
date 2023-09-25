package startup

import (
	"errors"
	formvalidator "jet-ci/internal/pkg/form-validator"
	"jet-ci/internal/user/api"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

type WebStartup struct {
}

func WebStart() {
	web := &WebStartup{}
	web.Start()
}

func (web *WebStartup) Start() {

	Load()

	formvalidator.Register()

	web.initGin()

	web.reginsterPage()

	web.RegisterRouting()
	web.webServe()

}

func (web *WebStartup) initGin() {
	if !options.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.New()
	//TODO: w
	engine.Use(gin.Recovery())
	if options.IsDebug {
		engine.Use(gin.Logger())
	}

	// 是否跨域
	if true {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowCredentials = true
		engine.Use(cors.New(c))
	}

	//启用session
	engine.Use(sessions.Sessions("jet-ci", cookie.NewStore([]byte("jet-ci"))))

	//开启压缩
	if options.IsGzip {
		engine.Use(gzip.Gzip(gzip.DefaultCompression))
	}
}

func (web *WebStartup) webServe() error {
	//log.Info("Web服务启动 ", options.Addr)

	addr := ":" + strconv.Itoa(options.Port)
	err := engine.Run(addr)
	if err != nil {
		//log.Fatal("HTTP 服务启动错误", err)
		return errors.New("启动错误...")
	}
	return nil
}

func (web *WebStartup) RegisterRouting() {

	rootR := engine.Group("/api")

	api.UserRouting(rootR)
	api.LoginRouting(rootR)

}

func (web *WebStartup) reginsterPage() {

	engine.Static("/web", "./web/dist")

}
