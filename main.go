package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

const usage = `mydocker is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	app := cli.NewApp()
//	app.Name = "ls"
//	app.Usage = "List directory contents"
//	app.Action = func(c *cli.Context) error {
//		dir, err := os.Open(".") // 默认打开当前目录
//		if err != nil {
//			return err
//		}
//		defer dir.Close()
//
//		fileInfos, err := dir.Readdir(-1) // 列出所有文件
//		if err != nil {
//			return err
//		}
//
//		for _, fi := range fileInfos {
//			fmt.Println(fi.Name())
//		}
//		return nil
//	}
//
//	err := app.Run(os.Args)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
