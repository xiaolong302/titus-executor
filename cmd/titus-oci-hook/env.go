package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Netflix/titus-executor/executor/runtime/docker"
	runtimeTypes "github.com/Netflix/titus-executor/executor/runtime/types"
)

// Creates the file $titusEnvironments/ContainerID.env filled with newline delimited set of environment variables
func createTitusEnvironmentFile(taskID string, env map[string]string) error {

	// TODO:
	// mountContainerProcPid1InTitusInits

	envFile := filepath.Join(runtimeTypes.TitusEnvironmentsDir, fmt.Sprintf("%s.env", taskID))
	f, err := os.OpenFile(envFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644) // nolint: gosec
	if err != nil {
		return err
	}
	defer f.Close()

	/* writeTitusEnvironmentFile closes the file for us */
	return docker.WriteTitusEnvironmentFile(env, f)
}
