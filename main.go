package main

import (
	"log"

	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	tracer "github.com/xiaobudongzhang/micro-plugins/tracer/jaeger"
	"github.com/xiaobudongzhang/micro-plugins/tracer/opentracing/stdhttp"
)

func init() {
	plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
}

const name = "API gateway"

func main() {
	stdhttp.SetSampingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}

	defer io.Close()
	opentracing.SetGlobalTracer(t)

	cmd.Init()
}
