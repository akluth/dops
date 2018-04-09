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

	var container_ids []string
	var container_names []string

	for _, container := range containers {
		container_ids = append(container_ids, container.ID)
		container_names = append(container_names, container.Names[0])
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	table_rows := [][]string{
		[]string{"IDs", "Container name"},
		container_ids,
		container_names,
	}

	docker_ps_table := ui.NewTable()
	docker_ps_table.Rows = table_rows
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
