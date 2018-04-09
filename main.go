package main

import "os/exec"
import ui "github.com/airking05/termui"

func main() {

	app := "docker"
	arg := "ps"

	cmd := exec.Command(app, arg)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	docker_ps := ui.NewPar(string(stdout))
	docker_ps.Height = 20
	docker_ps.Width = 100
	docker_ps.Y = 0
	docker_ps.Border = false

	ui.Render(docker_ps)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
