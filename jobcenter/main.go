package main

import (
	"flag"
	"io"
	"jobcenter/internal/config"
	"jobcenter/internal/svc"
	"jobcenter/internal/task"
	"log"
	"os"
	"os/signal"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	// 检查配置文件是否存在，若不存在则优雅引导用户配置
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		if *configFile == "etc/conf.yaml" {
			templateFile := "etc/conf-template.yaml"
			if _, terr := os.Stat(templateFile); terr == nil {
				// 自动从模版文件复制一份
				if copyFile(templateFile, *configFile) {
					log.Printf("配置文件 %s 不存在，已为您从模版 %s 自动创建。请修改配置后再重新运行。\n", *configFile, templateFile)
					os.Exit(1)
				}
			}
		}
		log.Printf("配置文件 %s 不存在，请参考 etc/conf-template.yaml 创建该文件并配置。\n", *configFile)
		os.Exit(1)
	}

	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	t := task.NewTask(ctx)
	t.Run()

	// 优雅退出
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		select {
		case <-exit:
			log.Println("任务中心中断执行，开始 clear 资源")
			t.Stop()
		}
	}()
	t.StartBlocking()
}

// 辅助复制文件函数
func copyFile(src, dst string) bool {
	sourceFile, err := os.Open(src)
	if err != nil {
		return false
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return false
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err == nil
}
