package k8sHealthCheck

type HealthCheck interface {
	Run() error
}
