package ilogtail_plugins

import (
	"github.com/alibaba/ilogtail/pkg/models"
	"github.com/alibaba/ilogtail/pkg/pipeline"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

type MyProcessor struct {
}

func (p *MyProcessor) Init(context pipeline.Context) error {
	return nil
}

func (p *MyProcessor) Description() string {
	return "custom processor"
}

func (p *MyProcessor) Process(in *models.PipelineGroupEvents, context pipeline.PipelineContext) {
	in.Group.Tags.Add("proceed", "true")
	context.Collector().Collect(in.Group, in.Events...)
}

func (p *MyProcessor) ProcessLogs(logArray []*protocol.Log) []*protocol.Log {
	for _, log := range logArray {
		log.Contents = append(log.Contents, &protocol.Log_Content{Key: "proceed", Value: "true"})
	}
	return logArray
}

func init() {
	pipeline.AddProcessorCreator("my_processor", func() pipeline.Processor {
		return &MyProcessor{}
	})
}
