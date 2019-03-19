// +build !windows

package cmd

import (
	"os"

	"github.com/prvst/philosopher/lib/ext/msconvert"
	"github.com/prvst/philosopher/lib/met"
	"github.com/prvst/philosopher/lib/sys"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// msconvertCmd represents the msconvert command
var msconvertCmd = &cobra.Command{
	Use:   "msconvert",
	Short: "Convert mass spec data file formats",
	Run: func(cmd *cobra.Command, args []string) {

		logrus.Info("Executing Msconvert ", Version)

		m, e := msconvert.Run(m, args)
		if e != nil {
			logrus.Fatal(e.Error())
		}

		m.Serialize()

		// clean tmp
		met.CleanTemp(m.Temp)

		logrus.Info("Done")
		return

	},
}

func init() {

	if len(os.Args) > 1 && os.Args[1] == "msconvert" {

		m.Restore(sys.Meta())

		msconvertCmd.Flags().StringVarP(&m.Msconvert.Format, "format", "", "", "mzML, mzXML, mz5, mgf, text, ms1, cms1, ms2, cms2")
		msconvertCmd.Flags().BoolVarP(&m.Msconvert.Zlib, "zlib", "", false, "use zlib compression for binary data")
		msconvertCmd.Flags().BoolVarP(&m.Msconvert.NoIndex, "noindex", "", false, "do not write index")
		msconvertCmd.Flags().StringVarP(&m.Msconvert.MZBinaryEncoding, "mzenconding", "", "64", "MZ default binary encoding")
		msconvertCmd.Flags().StringVarP(&m.Msconvert.IntensityBinaryEncoding, "intenconding", "", "64", "Intensity default binary encoding")
		//msconvertCmd.Flags().StringVarP(&m.Msconvert.Output, "output", "", "", "override the name of output file")
		//msconvertCmd.Flags().StringVarP(&m.Msconvert.Input, "input", "", "", "override the name of output file")

	}
	RootCmd.AddCommand(msconvertCmd)
}