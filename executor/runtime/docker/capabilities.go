package docker

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Netflix/titus-executor/executor/runtime/docker/seccomp"
	runtimeTypes "github.com/Netflix/titus-executor/executor/runtime/types"
	"github.com/docker/docker/api/types/container"
)

const (
	SYS_ADMIN = "SYS_ADMIN" // nolint: golint
	NET_ADMIN = "NET_ADMIN" // nolint: golint
)

func setupAdditionalCapabilities(c runtimeTypes.Container, hostCfg *container.HostConfig) error {
	seccompProfile := "default.json"
	apparmorProfile := "unconfined"

	hostCfg.CapAdd = append(hostCfg.CapAdd, SYS_ADMIN)

	// Tell Tini to exec systemd so it's pid 1
	c.SetEnv("TINI_HANDOFF", trueString)

	asset := seccomp.MustAsset(seccompProfile)
	var buf bytes.Buffer
	err := json.Compact(&buf, asset)
	if err != nil {
		return fmt.Errorf("Could not JSON compact seccomp profile string: %w", err)
	}

	hostCfg.SecurityOpt = append(hostCfg.SecurityOpt, fmt.Sprintf("seccomp=%s", buf.String()))
	hostCfg.SecurityOpt = append(hostCfg.SecurityOpt, "apparmor:"+apparmorProfile)

	return nil
}
