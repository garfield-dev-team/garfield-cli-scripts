package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

func init() {
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release the current package to NPM",
	Long:  `Release the current package to NPM`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		absPath := path.Join(cwd, "package.json")
		log.Println(absPath)

		content, err := os.ReadFile(absPath)
		if err != nil {
			log.Fatal(err)
		}

		var pkg map[string]any
		if err := json.Unmarshal(content, &pkg); err != nil {
			log.Fatal(err)
		}

		//if value, ok := pkg["peerDependencies"]; ok {
		//	pkg.PeerDependencies = make(map[string]string)
		//}
		//
		//dependencies := pkg.Dependencies
		//peerDependencies := pkg.PeerDependencies
		//
		//for _, dep := range externals {
		//	temp := dependencies[dep]
		//	peerDependencies[dep] = temp
		//	delete(dependencies, dep)
		//}
		//bytes, err := json.Marshal(pkg)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//output := path.Join(cwd, "_package")
		//err = os.WriteFile(output, bytes, fs.ModePerm)
		//if err != nil {
		//	log.Fatal(err)
		//}
	},
}
