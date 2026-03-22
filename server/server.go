package server

import (
	"dayz-server-tools/logger"
	"dayz-server-tools/steam"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func StartServer() {
	// 获取服务器路径和启动参数(只需获取一次)
	slp := GetServerLaunchParameters()
	dayzServerExe := steam.GetDayZServerExecutable()

	args := []string{
		"-port=" + strconv.Itoa(slp.Port),
		"-mission=" + slp.Mission,
		"-profiles=" + slp.Profiles,
		"-mod=" + slp.ClientMods,
		"-serverMod=" + slp.ServerMods,
		"-config=" + slp.Config,
		"-dologs",
		"-adminlog",
		"-netlog",
		"-freezecheck",
	}

	for {
		logger.Info("正在启动服务器...")
		logger.Info("如果遇到BattleEye相关报错，请确保防火墙，杀毒等工具已经关闭，代码完全开源无毒，请放心使用")
		cmd := exec.Command(dayzServerExe, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// 启动进程
		if err := cmd.Start(); err != nil {
			logger.Error("DayZ 服务器启动失败", "err", err.Error())
			fmt.Println("3 秒后自动退出...")
			os.Exit(1)
			return
		}

		logger.Info("DayZ 服务器已启动", "PID", cmd.Process.Pid)

		// 等待进程结束
		if err := cmd.Wait(); err != nil {
			logger.Error("服务器进程关闭", "err", err.Error())
		}

		logger.Info("⏳ 等待 3 秒后重启服务器...")
		time.Sleep(3 * time.Second)
	}
}
