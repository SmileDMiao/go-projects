package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var prefix string

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename files",
	Long:  `批量重命名`,

	Run: func(cmd *cobra.Command, args []string) {
		prefix, _ = cmd.Flags().GetString("prefix")

		currentDir, _ := os.Getwd()
		paths := make([]string, 0)

		filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})

		number := 1
		for _, path := range paths {
			newPath := generatePath(prefix, number, filepath.Ext(path), path)

			for {
				_, err := os.Stat(newPath)
				if err == nil {
					number++
					newPath = generatePath(prefix, number, filepath.Ext(path), path)
				}
				if os.IsNotExist(err) {
					break
				}
			}

			os.Rename(path, newPath)
			number++

			color.Yellow(path)
			color.Yellow(newPath)
		}
	},
}

func generatePath(prefix string, number int, ext string, path string) string {
	var builder strings.Builder
	builder.WriteString(prefix)
	builder.WriteString(strconv.Itoa(number))

	return filepath.Join(filepath.Dir(path), builder.String()+filepath.Ext(path))
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func init() {
	renameCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "prefix required")
	renameCmd.MarkPersistentFlagRequired("prefix")
	rootCmd.AddCommand(renameCmd)
}
