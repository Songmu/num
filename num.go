package num

import (
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
)

const cmdName = "num"

// Run the num
func Run(argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}
	numStr := fs.Arg(0)
	if numStr == "" {
		return fmt.Errorf("no number specified")
	}

	dec, err := strconv.Atoi(numStr)
	if err != nil {
		return err
	}

	fmt.Fprintf(outStream, "bin: 0b%b\n", dec)
	fmt.Fprintf(outStream, "oct: 0o%o\n", dec)
	fmt.Fprintf(outStream, "dec: %d\n", dec)
	fmt.Fprintf(outStream, "hex: 0x%x\n", dec)

	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}
