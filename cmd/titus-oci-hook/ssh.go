package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Netflix/titus-executor/executor/runtime/docker"
	podCommon "github.com/Netflix/titus-kube-common/pod"
	"github.com/coreos/go-systemd/dbus"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

// OCIPushSSHEnvironment mimics what titus-executor does for docker containers, except in an
// OCI-hook style way. See executor.runtime.docker.pushEnvironment for the docker version.
//
// Where possible this code uses the exact same functions as the docker version,
// to reduce drift as RK is rolled out, at the cost of some awkwardness
func OCIPushSSHEnvironment(spec *specs.Spec, pod *v1.Pod, env map[string]string) error {

	cwd, _ := os.Getwd()
	// Note: spec.Root.Path says it is supposed to be the absolute path, but it isn't?
	root := path.Join(cwd, spec.Root.Path)
	var imageEnv []string // TODO: how am I supposed to get this...

	var envTemplateBuf, tarBuf bytes.Buffer
	if err := docker.ExecuteEnvFileTemplate(env, imageEnv, &envTemplateBuf); err != nil {
		return err
	}

	taskID := pod.Name
	appName := pod.Labels[podCommon.LabelKeyApp]
	iamRole := pod.Annotations[podCommon.AnnotationKeyIAMRole]

	tw := tar.NewWriter(&tarBuf)

	if err := tw.WriteHeader(&tar.Header{
		Name:     "data",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}); err != nil {
		log.Fatal(err)
	}

	if err := tw.WriteHeader(&tar.Header{
		Name:     "logs",
		Mode:     0777,
		Typeflag: tar.TypeDir,
	}); err != nil {
		log.Fatal(err)
	}

	if err := tw.WriteHeader(&tar.Header{
		Name:     "titus",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}); err != nil {
		log.Fatal(err)
	}

	cfg := getTitusExecutorConfig()
	if err := docker.AddContainerSSHDConfig(taskID, appName, iamRole, tw, *cfg); err != nil {
		return err
	}

	path := "etc/nflx/base-environment.d/200titus"
	hdr := &tar.Header{
		Name: path,
		Mode: 0644,
		Size: int64(envTemplateBuf.Len()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		log.Fatalln(err)
	}
	if _, err := tw.Write(envTemplateBuf.Bytes()); err != nil {
		log.Fatalln(err)
	}
	// Make sure to check the error on Close.

	if err := tw.Close(); err != nil {
		return err
	}

	logrus.Infof("Untaring into %s. Tar size: %d", root, tarBuf.Len())
	reader := tar.NewReader(&tarBuf)
	err := untar(reader, root)
	if err != nil {
		return fmt.Errorf("Error untaring into %s: %w", root, err)
	}

	return nil
}

func launchSSHDSidecar(spec *specs.Spec, pod *v1.Pod, env map[string]string) error {
	conn, connErr := dbus.New()
	if connErr != nil {
		return connErr
	}
	defer conn.Close()

	cID := "unknown cid"
	runtime := "unkown runtime"

	svc := docker.ServiceOpts{
		HumanName: "ssh",
		UnitName:  "titus-sshd",
		Required:  true,
		//enabledCheck: func(cfg *config.Config, c runtimeTypes.Container) bool {
		//	return cfg.ContainerSSHD
		//},
	}

	// TODO : Create inits
	// TODO: Copy volumes

	
	if err := docker.StartSystemdUnit(context.TODO(), conn, pod.Name, cID, runtime, svc); err != nil {
		logrus.WithError(err).Errorf("Error starting %s service", svc.HumanName)
		return err
	}
	return nil
}
