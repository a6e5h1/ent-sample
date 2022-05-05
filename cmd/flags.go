package cmd

import "github.com/urfave/cli/v2"

func Flags(flags ...[]cli.Flag) []cli.Flag {
	l := make([]cli.Flag, 0)
	for _, v := range flags {
		l = append(l, v...)
	}
	return l
}
