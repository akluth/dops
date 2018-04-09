package main

import (
	"time"

	ui "github.com/airking05/termui"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

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
	docker_ps_table.Y = 0
	docker_ps_table.X = 0
	docker_ps_table.Width = ui.TermWidth()
	docker_ps_table.Height = ui.TermHeight()
	docker_ps_table.BorderLabel = "Containers"
	docker_ps_table.BorderLabelFg = ui.ColorWhite
	docker_ps_table.BorderFg = ui.ColorGreen

	ui.Body.AddRows(
		ui.NewRow(
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
