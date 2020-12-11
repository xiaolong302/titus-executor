package main

import (
	"github.com/Netflix/titus-executor/config"
	v1 "k8s.io/api/core/v1"
)

func getEnvMapFromK8sEnv(e []v1.EnvVar) (m map[string]string) {
	m = make(map[string]string)
	for _, s := range e {
		m[s.Name] = s.Value
	}
	return
}

// getTitusExecutorConfig returns a config object that mimics
// what titus-executor would have if run via the command line
// in the classic way
func getTitusExecutorConfig() *config.Config {
	cfg, _ := config.NewConfig()
	cfg.ContainerSSHD = true
	cfg.ContainerSSHDCAFile = "/etc/titus-executor/titus_sshd_ca.pub"
	return cfg
}
