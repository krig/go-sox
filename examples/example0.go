package main

import "github.com/krig/go-sox"
import "log"

func main() {
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	defer sox.Quit()

	in := sox.OpenRead("./test-input.wav")
	if in == nil {
		log.Fatal("Failed to open test-input.wav")
	}
	out := sox.OpenWrite("./test-output.wav", in.Signal())
	if in == nil {
		log.Fatal("Failed to open test-output.wav")
	}

	chain := sox.CreateEffectsChain(in.Encoding(), out.Encoding())

	e := sox.CreateEffect(sox.FindEffect("input"))
	e.Options(in)
	chain.AddEffect(e, in.Signal(), in.Signal())
	e.Free()

	e = sox.CreateEffect(sox.FindEffect("vol"))
	e.Options("3dB")
	chain.AddEffect(e, in.Signal(), in.Signal())
	e.Free()

	e = sox.CreateEffect(sox.FindEffect("flanger"))
	e.Options()
	chain.AddEffect(e, in.Signal(), in.Signal())
	e.Free()

	e = sox.CreateEffect(sox.FindEffect("output"))
	e.Options(out)
	chain.AddEffect(e, in.Signal(), in.Signal())
	e.Free()

	chain.FlowEffects()

	chain.Delete()
	out.Close()
	in.Close()
}