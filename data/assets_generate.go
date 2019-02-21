// +build ignore

package main

import (
	"log"

	"github.com/openshift-metalkube/kni-installer/data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "release",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
