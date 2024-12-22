package monitoring

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    ActiveUsers = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "active_users",
        Help: "The current number of active users",
    })


    DataProcessingDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name: "data_processing_duration_seconds",
        Help: "Time spent processing data",
        Buckets: prometheus.LinearBuckets(0, 0.1, 10),
    })
)

func RecordDataProcessingDuration(duration float64) {
    DataProcessingDuration.Observe(duration)
}

func UpdateActiveUsers(count int) {
    ActiveUsers.Set(float64(count))
}
