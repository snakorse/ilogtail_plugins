package ilogtail_plugins

import (
	"time"

	"github.com/alibaba/ilogtail/pkg/models"
	"github.com/alibaba/ilogtail/pkg/pipeline"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

type MyServiceInput struct {
	Interval time.Duration `json:"interval" mapstructure:"interval"`

	collector   pipeline.Collector
	collectorV2 pipeline.PipelineCollector
	stopCh      chan struct{}
}

func (m *MyServiceInput) StartService(context pipeline.PipelineContext) error {
	m.collectorV2 = context.Collector()
	return nil
}

func (m *MyServiceInput) Init(context pipeline.Context) (int, error) {
	return 0, nil
}

func (m *MyServiceInput) Description() string {
	return "custom service input"
}

func (m *MyServiceInput) Stop() error {
	return nil
}

func (m *MyServiceInput) Start(collector pipeline.Collector) error {
	m.collector = collector
	go m.loop()
	return nil
}

func (m *MyServiceInput) loop() {
	for {
		select {
		case <-m.stopCh:
			return
		case <-time.After(m.Interval):
			if m.collector != nil {
				m.collector.AddRawLog(&protocol.Log{
					Time: uint32(time.Now().Unix()),
					Contents: []*protocol.Log_Content{
						{Key: "host", Value: "1.1.1.1"},
					},
				})
			}
			if m.collectorV2 == nil {
				continue
			}
			m.collectorV2.Collect(&models.GroupInfo{},
				models.NewSingleValueMetric("test_metric",
					models.MetricTypeCounter,
					models.NewTagsWithMap(map[string]string{"hello": "world"}),
					time.Now().UnixNano(),
					1),
			)
		}
	}
}

func init() {
	pipeline.AddServiceCreator("service_my_input", func() pipeline.ServiceInput {
		return &MyServiceInput{
			Interval: 5 * time.Second,
			stopCh:   make(chan struct{}, 1),
		}
	})
}
