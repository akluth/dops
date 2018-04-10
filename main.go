package main

import (
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	ui "github.com/airking05/termui"
)

func getDockerVersion() string {
	cmd := exec.Command("docker", "version")
	stdout, _ := cmd.Output()

	return string(stdout)
}

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	container_table_rows := [][]string{
		[]string{"ID", "Name", "Image", "Command", "Created"},
	}

	for _, container := range containers {
		cont := []string{
			container.ID[:12],
			container.Names[0],
			container.Image,
			container.State + ", " + container.Status,
			string(time.Unix(container.Created, 0).String()),
		}

		container_table_rows = append(container_table_rows, cont)
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	docker_ps_table := ui.NewTable()
	docker_ps_table.Rows = container_table_rows
	docker_ps_table.Width = ui.TermWidth()
	docker_ps_table.Height = ui.TermHeight()
	docker_ps_table.BorderLabel = "Containers"
	docker_ps_table.BorderLabelFg = ui.ColorWhite
	docker_ps_table.BorderFg = ui.ColorGreen
	docker_ps_table.Separator = false

	docker_ps_table.Analysis()
	docker_ps_table.SetSize()
	docker_version_box := ui.NewPar(getDockerVersion())
	docker_version_box.Width = ui.TermWidth()
	docker_version_box.Height = ui.TermHeight()
	docker_version_box.BorderLabel = "Docker Information"
	docker_version_box.BorderFg = ui.ColorCyan
	docker_version_box.BorderLabelFg = ui.ColorWhite

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(4, 0, docker_version_box),
			ui.NewCol(8, 0, docker_ps_table),
		),
	)

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
