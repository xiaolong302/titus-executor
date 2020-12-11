package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Netflix/titus-executor/utils/k8s"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"

	"github.com/wercker/journalhook"

	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	kubeletAPIPodsURL = "https://localhost:10250/pods"
)

func extractContainerAndPodNames(spec specs.Spec) (string, string) {
	envMap := getEnvMap(spec.Process.Env)
	logrus.Debugf("Env: %+v", envMap)
	cn, ok := envMap["CONTAINER_NAME"]
	if !ok {
		if spec.Process.Args[0] == "/pause" {
			cn = "pause"
		} else {
			cn = ""
		}
	}
	pn, ok := envMap["TITUS_TASK_ID"]
	if !ok {
		pn = spec.Hostname
	}
	return cn, pn
}

func getPod(podName string) (*v1.Pod, error) {
	var k8sArgs k8s.Args
	//k8sArgs.K8S_POD_NAME = podName.(types.UnmarshallableString)
	k8sArgs.K8S_POD_NAME = types.UnmarshallableString(podName)
	k8sArgs.K8S_POD_NAMESPACE = types.UnmarshallableString("default")
	return k8s.GetPod(context.TODO(), kubeletAPIPodsURL, k8sArgs)
}

func getEnvMap(e []string) (m map[string]string) {
	m = make(map[string]string)
	for _, s := range e {
		p := strings.SplitN(s, "=", 2)
		if len(p) != 2 {
			log.Panicln("environment error")
		}
		m[p[0]] = p[1]
	}
	return
}

func doPrestart() error {
	journalhook.Enable()
	bundleDirPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	logrus.Debugf("Using bundle file: %s\n", bundleDirPath+"/config.json")
	jsonFile, err := os.OpenFile(bundleDirPath+"/config.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Couldn't open OCI spec file: %w", err)
	}
	defer jsonFile.Close()

	jsonContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("Couldn't read OCI spec file: %w", err)
	}
	var spec specs.Spec
	err = json.Unmarshal(jsonContent, &spec)
	if err != nil {
		return fmt.Errorf("Couldn't unmarshal OCI spec file: %w", err)
	}

	containerName, podName := extractContainerAndPodNames(spec)
	logrus.Infof("Running on container %s for pod %s", containerName, podName)
	if podName == "" || containerName == "" {
		return fmt.Errorf("Not operating on container, is it a k8s pod container? Spec: %+v", spec)
	}

	if containerName != podName {
		logrus.Debugf("Skipping on container %s, only operating on the main container", containerName)
		return nil
	}

	pod, err := getPod(podName)
	if err != nil {
		return fmt.Errorf("Error getting pod for: %w", err)
	}
	logrus.Debugf("Got pod: %+v", pod)

	// TODO: Don't assume the first container is the one we want
	env := getEnvMapFromK8sEnv(pod.Spec.Containers[0].Env)

	err = createTitusEnvironmentFile(pod.Name, env)
	if err != nil {
		return fmt.Errorf("Failed when trying create titus-env files: %w", err)
	}
	err = OCIPushSSHEnvironment(&spec, pod, env)
	if err != nil {
		return fmt.Errorf("Failed when trying to update the environment: %w", err)
	}
	err = launchSSHDSidecar(&spec, pod, env)
	if err != nil {
		return fmt.Errorf("Failed to launch sshd sidecar: %w", err)
	}

	jsonOutput, err := json.Marshal(spec)
	if err != nil {
		return fmt.Errorf("Couldn't marshal OCI spec file: %w", err)
	}
	_, err = jsonFile.WriteAt(jsonOutput, 0)
	if err != nil {
		return fmt.Errorf("Couldn't write OCI spec file: %w", err)
	}

	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nCommands:\n")
	fmt.Fprintf(os.Stderr, "  prestart\n        run the prestart hook\n")
	fmt.Fprintf(os.Stderr, "  poststart\n       run the poststart hook\n")
	fmt.Fprintf(os.Stderr, "  poststop\n        run the poststop hook\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	switch args[0] {
	case "prestart":
		err := doPrestart()
		if err != nil {
			logrus.Fatal(err)
		}
		os.Exit(0)
	case "poststart":
		os.Exit(0)
	case "poststop":
		os.Exit(0)
	default:
		flag.Usage()
		os.Exit(2)
	}
}
