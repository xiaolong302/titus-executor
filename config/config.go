package config

import (
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

const (
	defaultLogsTmpDir = "/var/lib/titus-container-logs"
)

// Config contains the executor configuration
type Config struct {
	// nolint: maligned

	// PrivilegedContainersEnabled returns whether to give tasks CAP_SYS_ADMIN
	PrivilegedContainersEnabled bool
	// UseNewNetworkDriver returns which network driver to use
	UseNewNetworkDriver bool
	// DisableMetrics makes it so we don't send metrics to Atlas
	DisableMetrics bool
	// LogUpload returns settings about the log uploader
	//LogUpload logUpload
	LogsTmpDir string
	// Stack returns the stack configuration variable
	Stack string
	// Docker returns the Docker-specific configuration settings
	DockerHost     string
	DockerRegistry string

	// This is the image used for execution.
	ExecutorImage string

	// CopiedFromHost indicates which environment variables to lift from the current config
	copiedFromHostEnv cli.StringSlice
	hardCodedEnv      cli.StringSlice
}

// NewConfig generates a configuration and a set of flags to passed to urfave/cli
func NewConfig() (*Config, []cli.Flag) {
	cfg := &Config{
		copiedFromHostEnv: []string{
			"NETFLIX_ENVIRONMENT",
			"NETFLIX_ACCOUNT",
			"NETFLIX_STACK",
			"EC2_INSTANCE_ID",
			"EC2_REGION",
			"EC2_AVAILABILITY_ZONE",
			"EC2_OWNER_ID",
			"EC2_RESERVATION_ID",
		},
		hardCodedEnv: []string{
			"NETFLIX_APPUSER=appuser",
			"EC2_DOMAIN=amazonaws.com",
			/* See:
			 * - https://docs.aws.amazon.com/cli/latest/topic/config-vars.html
			 * - https://github.com/jtblin/kube2iam/issues/31
			 * AWS_METADATA_SERVICE_TIMEOUT, and AWS_METADATA_SERVICE_NUM_ATTEMPTS are respected by all AWS standard SDKs
			 * as timeouts for connecting to the metadata service.
			 */
			"AWS_METADATA_SERVICE_TIMEOUT=5",
			"AWS_METADATA_SERVICE_NUM_ATTEMPTS=3",
		},
	}

	flags := []cli.Flag{
		cli.BoolFlag{
			Name:        "privileged-containers-enabled",
			EnvVar:      "PRIVILEGED_CONTAINERS_ENABLED",
			Destination: &cfg.PrivilegedContainersEnabled,
		},
		cli.BoolFlag{
			Name:        "use-new-network-driver",
			EnvVar:      "USE_NEW_NETWORK_DRIVER",
			Destination: &cfg.UseNewNetworkDriver,
		},
		cli.BoolFlag{
			Name:        "disable-metrics",
			EnvVar:      "DISABLE_METRICS,SHORT_CIRCUIT_QUITELITE",
			Destination: &cfg.DisableMetrics,
		},
		cli.StringFlag{
			Name:        "logs-tmp-dir",
			Value:       defaultLogsTmpDir,
			EnvVar:      "LOGS_TMP_DIR",
			Destination: &cfg.LogsTmpDir,
		},
		cli.StringFlag{
			Name:        "stack",
			Value:       "mainvpc",
			EnvVar:      "STACK,NETFLIX_STACK",
			Destination: &cfg.Stack,
		},
		cli.StringFlag{
			Name: "docker-host",
			// In prod this is tcp://127.0.0.1:4243
			Value:       "unix:///var/run/docker.sock",
			Destination: &cfg.DockerHost,
			EnvVar:      "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:        "docker-registry",
			Value:       "docker.io",
			Destination: &cfg.DockerRegistry,
			EnvVar:      "DOCKER_REGISTRY",
		},
		cli.StringFlag{
			Name:        "executor-image",
			Value:       "titusops/pause",
			Destination: &cfg.ExecutorImage,
			EnvVar:      "EXECUTOR_IMAGE",
		},
		cli.StringSliceFlag{
			Name:  "copied-from-host-env",
			Value: &cfg.copiedFromHostEnv,
		},
		cli.StringSliceFlag{
			Name:  "hard-coded-env",
			Value: &cfg.hardCodedEnv,
		},
	}

	return cfg, flags
}

func (c *Config) GetEnvFromHost() map[string]string {
	fromHost := make(map[string]string)

	for _, hostKey := range c.copiedFromHostEnv {
		if hostKey == "NETFLIX_STACK" {
			// Add agent's stack as TITUS_STACK so platform libraries can
			// determine agent stack, if needed
			addElementFromHost(fromHost, hostKey, "TITUS_STACK")
		} else {
			addElementFromHost(fromHost, hostKey, hostKey)
		}
	}
	return fromHost
}

func addElementFromHost(addTo map[string]string, hostEnvVarName string, containerEnvVarName string) {
	hostVal := os.Getenv(hostEnvVarName)
	if hostVal != "" {
		addTo[containerEnvVarName] = hostVal
	}
}

func (c *Config) GetHardcodedEnv() map[string]string {
	env := make(map[string]string)

	for _, line := range c.hardCodedEnv {
		kv := strings.SplitN(line, "=", 2)
		env[kv[0]] = kv[1]
	}

	return env
}

// GenerateConfiguration is only meant to validate the behaviour of parsing command line arguments
func GenerateConfiguration(args []string) (*Config, error) {
	cfg, flags := NewConfig()

	app := cli.NewApp()
	app.Flags = flags
	app.Action = func(c *cli.Context) error {
		return nil
	}
	if args == nil {
		args = []string{}
	}

	args = append([]string{"fakename"}, args...)

	return cfg, app.Run(args)
}
