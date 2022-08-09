package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ShangRui-hash/notifier/wx"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var configFile string

func main() {
	author := cli.Author{
		Name:  "无在无不在",
		Email: "2227627947@qq.com",
	}
	app := &cli.App{
		Name:      "notifier",
		Usage:     "The tool is used to read standard input line by line and send it to enterprise wechat",
		UsageText: "echo 'happy hacking' | notifier ",
		Version:   "v0.1",
		Authors:   []*cli.Author{&author},
		Flags:     []cli.Flag{},
		Action:    run,
	}
	//启动app
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
	}
}

func run(c *cli.Context) (err error) {
	//从标准输入中读
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//获取输入的内容
		line := scanner.Text()
		//发送到wx
		if err := wx.SendMsg(line); err != nil {
			logrus.Error("wx.SendMsg failed,err:", err)
		}
		//输出到标准输出
		fmt.Println(line)
	}
	return nil
}
