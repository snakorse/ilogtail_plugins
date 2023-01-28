package ilogtail_plugins

import (
	"github.com/alibaba/ilogtail/pkg/models"
	"github.com/alibaba/ilogtail/pkg/pipeline"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

type MyAggregator struct {
}

func (m *MyAggregator) Init(context pipeline.Context, queue pipeline.LogGroupQueue) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) Description() string {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) Reset() {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) Add(log *protocol.Log, ctx map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) Flush() []*protocol.LogGroup {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) Record(events *models.PipelineGroupEvents, context pipeline.PipelineContext) error {
	//TODO implement me
	panic("implement me")
}

func (m *MyAggregator) GetResult(context pipeline.PipelineContext) error {
	//TODO implement me
	panic("implement me")
}

func init() {
	pipeline.AddAggregatorCreator("my_aggregator", func() pipeline.Aggregator {
		return &MyAggregator{}
	})
}
