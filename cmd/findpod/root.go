package findpod

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var allNamespaces bool
var namespace string

func filter(command string, input func(in io.WriteCloser)) []string {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := cmd.Output()
	return strings.Split(string(result), "\n")
}

func filterOne(command string, input func(in io.WriteCloser)) string {
	var r string
	command = fmt.Sprintf("%s | fzf", command)
	filtered := filter(command, input)
	if len(filtered) > 0 {
		r = filtered[0]
	}
	return r
}

func filterMany(command string, input func(in io.WriteCloser)) []string {
	command = fmt.Sprintf("%s | fzf -m ", command)
	return filter(command, input)
}

var rootCmd = &cobra.Command{
	Use:   "findpod",
	Short: "Search for a pod",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var podPattern string
		if len(args) == 1 {
			podPattern = args[0]
		}
		fmt.Println(podPattern)

		filtered := filterOne("kubectl get pod --all-namespaces -o=jsonpath='{range .items[*]}{.metadata.namespace}{\"\\t\"}{.metadata.name}{\"\\n\"}{end}'", func(in io.WriteCloser) {})

		if filtered == "" {
			fmt.Println("Selected nothing")
			os.Exit(0)
		}

		p := strings.Split(filtered, "\t")
		namespace := p[0]
		name := p[1]

		fmt.Println(name, namespace)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "a", false, "look in all namespaces")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace pattern")
}
