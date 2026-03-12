package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type (
	PortConfig struct{}
	NodeConfig struct {
		NodeIndex int
		PortConfig
	}
)

type ComposeConfig struct {
	Nodes           []NodeConfig
	Image           string
	MySQLImage      string
	ProjectBasePath string
	BasePorts       PortConfig
}

func main() {
	// Get template file path
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	templatePath := filepath.Join(filepath.Dir(execPath), "docker-compose.tmpl")

	// If running with "go run", use current directory
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		templatePath = filepath.Join("cmd", "ci", "docker-compose.tmpl")
	}

	bp := PortConfig{}

	numNodes := 4

	var nodes []NodeConfig
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, NodeConfig{
			NodeIndex:  i,
			PortConfig: PortConfig{},
		})
	}

	config := ComposeConfig{
		Nodes:           nodes,
		Image:           "mocachain/moca-relayer",
		MySQLImage:      "mysql:8",
		ProjectBasePath: ".",
		BasePorts:       bp,
	}

	// Read and parse template file
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template file:", err)
		fmt.Println("Template path:", templatePath)
		return
	}

	// Create output file
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Execute template
	err = tpl.Execute(file, config)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Docker Compose file generated successfully!")
	fmt.Println("Template used:", templatePath)
}
