package signaltest

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitSignal() {
	// 1、修改了配置文件后，希望在不重启进程的情况下重新加载配置文件；
	// 2、当用 Ctrl + C 强制关闭应用后，做一些必要的处理；
	// golang中对信号的处理主要使用os/signal包中的两个方法：
	//   一个是notify方法用来监听收到的信号；
	//   一个是 stop方法用来取消监听。

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 第一个参数表示接收信号的管道
	// 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	go func() {
		sig := <-sigs // 阻塞直至有信号传入
		log.Println("receive signaltest: ", sig)

		//if utils.LogRotater != nil {
		//	utils.LogRotater.Rotate()
		//}

		os.Exit(0)
	}()
}

// golang中处理信号非常简单
// 关于信号本身需要了解的还有很多，建议可以参考《Unix环境高级编程》中的信号篇章。
