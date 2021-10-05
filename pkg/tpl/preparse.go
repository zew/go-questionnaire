package tpl

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/zew/go-questionnaire/pkg/cloudio"
)

// TemplatesPreparse - parsing templates is expensive;
// concurrent access is expensive;
// we parse all templates in this http handler
// which is also used for initialization at app start in bootstrap.
func TemplatesPreparse(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// reading the bundling info
	err := cloudio.MarshalWriteFile(bundles, path.Join(".", "templates", "bundles-example.json"))
	if err != nil {
		log.Panicf("failed to create 'templates/bundles-example.json': %v", err)
	}
	err = cloudio.ReadFileUnmarshal(path.Join(".", "templates", "bundles.json"), &bundles)
	if err != nil {
		log.Panicf("failed to read/unmarshal 'templates/bundles.json': %v", err)
	}

	// reading all template files from the directory
	// and pushing them to tpl()
	// tpl() adds the bundled templates - reading the files redunantly via bundle()
	lo, err := cloudio.ReadDir("templates")
	if err != nil {
		fmt.Fprintf(w, "cannot read directory 'templates': %v \n", err)
		return
	}
	for _, o := range *lo {

		if o.IsDir {
			continue
		}

		pth := path.Join(".", o.Key)
		pth = strings.ReplaceAll(pth, "\\", "/") // cloudio always has forward slash

		tName := path.Base(pth)
		if strings.HasPrefix(tName, "tmp-") {
			fmt.Fprintf(w, "skipping tmp-* entry %v \n", tName)
			continue
		}

		t, err := Get(tName, true)
		if err != nil {
			fmt.Fprintf(w, "preparse failure template %-30v: %v\n", tName, err)
			continue
		}
		cache[tName] = t // cache write only here

		// logging the relationships
		if false {
			fmt.Fprintf(w, "preparse success template  %-30v\n", tName)
			for _, tSub := range t.Templates() {
				if tSub.Name() != tName {
					fmt.Fprintf(w, "  preparse success subtpl  %-30v\n", tSub.Name())
				}
			}
		}
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Templates pre-parsed and cached. \n")

}
