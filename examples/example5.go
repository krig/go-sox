// Simple example of using SoX libraries.
package main

// Use this URL to import the go-sox library
import "github.com/krig/go-sox"
import "log"
import "flag"

const (
	MAX_SAMPLES = 2048
)

// Example of reading and writing audio files stored
// in memory buffers rather than actual files.
func main() {
	var samples [MAX_SAMPLES]sox.Sample

	flag.Parse()
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	defer sox.Quit()

	// Open the input file.
	in := sox.OpenRead(flag.Arg(0))
	if in == nil {
		log.Fatal("Failed to open input file")
	}

	// Set up the memory buffer for writing
	buf := sox.NewMemstream()
	out := sox.OpenMemstreamWrite(buf, in.Signal(), nil, "sox")
	if out == nil {
		log.Fatal("Failed to open memory buffer")
	}

	for number_read := in.Read(samples[:], MAX_SAMPLES); number_read > 0; number_read = in.Read(samples[:], MAX_SAMPLES) {
		out.Write(samples[:], uint(number_read))
	}
	out.Close()
	in.Close()

	in = sox.OpenMemRead(buf)
	if in == nil {
		log.Fatal("Failed to open memory buffer for reading")
	}
	out = sox.OpenWrite(flag.Arg(1), in.Signal(), nil, nil)
	if out == nil {
		log.Fatal("Failed to open file for writing")
	}
	for number_read := in.Read(samples[:], MAX_SAMPLES); number_read > 0; number_read = in.Read(samples[:], MAX_SAMPLES) {
		out.Write(samples[:], uint(number_read))
	}
	out.Close()
	in.Close()
}
