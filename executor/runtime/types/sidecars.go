package types

var MandatoryVolumes = map[string]string{
	"titus-sidecar:latest.release": "/titus/sidecar",
}

var MandatoryServices = []string{
	"titus-metadata-proxy",
}
