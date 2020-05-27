package main

import (
	"log"

	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	myjaeger "github.com/xiaobudongzhang/micro-plugins/tracer/myjaeger"
	"github.com/xiaobudongzhang/micro-plugins/tracer/myopentracing/stdhttp"
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
	stdhttp.SetSamplingFrequency(50)
	t, io, err := myjaeger.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}

	defer io.Close()
	opentracing.SetGlobalTracer(t)

	cmd.Init()
}
