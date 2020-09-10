package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// 参考 https://book.kubebuilder.io/reference/metrics.html 来实现自定义prometheus metrics的注册使用
var (
	// 全局可调用
	Goobers = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "goobers_total",
			Help: "Number of goobers proccessed",
		},
	)
	GooberFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "goober_failures_total",
			Help: "Number of failed goobers",
		},
	)
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(Goobers, GooberFailures)
}
