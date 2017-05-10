// A simple example of using SoX libraries
package main

// Use this URL to import the go-sox library
import "github.com/krig/go-sox"
import "log"
import "flag"

const (
	//The (maximum) number of samples that we shall read/write at a time;
	// chosen as a rough match to typical operating system I/O buffer size:
	MAX_SAMPLES = 2048
)

func check(cond bool, test string) {
	if !cond {
		log.Fatal(test + " failed\n")
	}
}

// Concatenate audio files.  Note that the files must have the same number
// of channels and the same sample rate.
//
// Usage: example4 input-1 input-2 [... input-n] output
func main() {
	flag.Parse()

	// Need at least 2 input files + 1 output file.
	if flag.NArg() < 2+1 {
		log.Fatal("Usage: <input1> <input2> ... <output>")
	}

	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	// Make sure to call Quit before terminating
	defer sox.Quit()

	var input *sox.Format
	var output *sox.Format
	var signal *sox.SignalInfo

	// For each input file...
	for i := 0; i < flag.NArg()-1; i++ {
		//initialize the Slice, such that when we pass it to C, it is an Array with sufficient Space
		samples := make([]sox.Sample, MAX_SAMPLES)

		input = sox.OpenRead(flag.Arg(i))
		if input == nil {
			log.Fatal("OpenRead failed")
		}
		// If this is the first input file...
		if i == 0 {
			output = sox.OpenWrite(flag.Arg(flag.NArg()-1),
				input.Signal(), input.Encoding(),
				nil)
			if output == nil {
				log.Fatal("OpenWrite failed")
			}
			defer output.Release()
			signal = input.Signal()
		} else {
			// If this is not the first input file, the number of Channels
			// and the Rate must match that of the first file (WHY?)
			check(input.Signal().Channels() == signal.Channels(), "channels")
			check(input.Signal().Rate() == signal.Rate(), "rate")
		}

		// Continue here!
		for number_read := input.Read(samples, MAX_SAMPLES); //
		number_read > 0;                                     //
		number_read = input.Read(samples, MAX_SAMPLES) {
			check(output.Write(samples, uint(number_read)) == number_read, "write")
		}
		input.Release()
	}
}
