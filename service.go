package go_service

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

type Program struct {
	RunFn func() //运行方法
}

func (p *Program) Start(s service.Service) error {
	fmt.Print("服务运行...")
	go p.RunFn()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	fmt.Print("服务停止。")
	return nil
}

// 以服务运行
func RunWithService(srvConfig *service.Config, run func()) (s service.Service) {

	prg := &Program{}
	prg.RunFn = run
	var err error
	s, err = service.New(prg, srvConfig)
	if err != nil {
		fmt.Printf("创建服务出错：%v", err)
	}
	var name = fmt.Sprintf("服务[%v]", srvConfig.Name)
	if len(os.Args) > 1 {
		serviceAction := os.Args[1]
		switch serviceAction {
		case "install":
			err := s.Install()
			if err != nil {
				fmt.Println(fmt.Sprintf("安装%v失败: ", name), err.Error())
			} else {
				fmt.Println(fmt.Sprintf("安装%v成功", name))
			}
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				fmt.Println(fmt.Sprintf("卸载%v失败: ", name), err.Error())
			} else {
				fmt.Println(fmt.Sprintf("卸载%v成功", name))
			}
		case "start":
			err := s.Start()
			if err != nil {
				fmt.Println(fmt.Sprintf("运行%v失败: ", name), err.Error())
			} else {
				fmt.Println(fmt.Sprintf("运行%v成功", name))
			}
		case "stop":
			err := s.Stop()
			if err != nil {
				fmt.Println(fmt.Sprintf("停止%v失败: ", name), err.Error())
			} else {
				fmt.Println(fmt.Sprintf("停止%v成功", name))
			}
		}
		return
	}

	//不带参数直接运行
	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
	return
}
