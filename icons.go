package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"go.i3wm.org/i3"
)

type Icons map[string]string

type Config struct {
	icons     Icons
	separator string
}

func (c *Config) init(confFile string) {

	jsonFile, err := os.Open(confFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &c.icons)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) iterate(node *i3.Node) {

	// ignore __i3_scratch workspace
	if node.Type == "workspace" && node.Name != "__i3_scratch" {
		windows := node.Nodes
		windows = append(windows, node.FloatingNodes...)
		c.editIcons(node, windows)
	}

	for _, n := range node.Nodes {
		c.iterate(n)
	}

	for _, n := range node.FloatingNodes {
		c.iterate(n)
	}
}

func (c *Config) editIcons(workspace *i3.Node, windows []*i3.Node) {

	workspaceIcons := make([]string, 0)

	for _, win := range windows {
		found_icon := false

		for identifier, icon := range c.icons {
			if strings.Contains(strings.ToLower(win.Name), identifier) {
				found_icon = true
				workspaceIcons = append(workspaceIcons, icon)
			}
		}

		if !found_icon {
			workspaceIcons = append(workspaceIcons, c.icons["_wildcard"])
		}
	}

	_, err := i3.RunCommand(c.generateRenameCommand(workspace.Name, workspaceIcons))
	if err != nil {
		log.Println(err)
	}
}

func (c *Config) generateRenameCommand(workspace string, workspaceIcons []string) string {
	r := regexp.MustCompile(`^[0-9]+`)
	workspaceNumber := r.FindString(workspace)
	icons_display := strings.Join(workspaceIcons, c.separator)

	if icons_display == "" {
		return fmt.Sprintf(`rename workspace "%s" to "%s"`, workspace, workspaceNumber)
	}

	return fmt.Sprintf(`rename workspace "%s" to "%s:%s"`, workspace, workspaceNumber, icons_display)
}