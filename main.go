package main

import ui "github.com/airking05/termui"
import (
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
		[]string{"IDs", "Container name"},
	}

	for _, container := range containers {
		cont := []string{
			container.ID,
			container.Names[0],
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
	docker_ps_table.Width = 200
	docker_ps_table.Height = 200

	ui.Render(docker_ps_table)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
