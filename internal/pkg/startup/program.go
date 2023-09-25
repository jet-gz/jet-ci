package startup

import (
	"jet-ci/internal/pkg/db"
)

type Program struct {
}

func (p *Program) Start() {
	p.run()
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	// hup := make(chan os.Signal, 2)
	// signal.Notify(hup, syscall.SIGHUP)
	// // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal, 2)
	// signal.Notify(quit, os.Interrupt, os.Kill)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-hup:
	// 		case <-quit:
	// 			//优雅地结束
	// 			_ = p.Stop()
	// 			//os.Exit(0)
	// 		}
	// 	}
	// }()

	// 加载db
	db.Start()

	// 启动web
	WebStart()

	select {}
}

func (p *Program) Stop() error {
	return p.shutdown()
}

func (p *Program) shutdown() error {
	//return db.Close()
	return nil
}
