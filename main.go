package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	// "time"
)

// TODO(#11): flags to specify display type. json, text,?
func displayDef(definition []string, traverses int, depth int) {
	defer recovery("invalid")

	if traverses == (len(definition)-1) || traverses == (depth-1) {
		fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
		return
	}

	fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
	displayDef(definition, traverses+1, depth)
}

func parseChan(tmpDef chan []string, size int) []string {

	if len(tmpDef) > 0 {
		tmpDict := make([]string, 0)

		for i := 0; i < size; i++ {
			tmp := <-tmpDef
			for j := 0; j < len(tmp); j++ {
				tmpDict = append(tmpDict, tmp[j])
			}
		}
		return tmpDict
	}

	return []string(nil)
}

func checkWeb(word string, apiConf *config, tmpDef chan<- []string) {
	var wg sync.WaitGroup

	for idx := range apiConf.Website {
		wg.Add(1)

		website, link, apiKey := parseConfig(apiConf, idx)

		requestLink, err := parseRequest(word, website, link, apiKey)
		if err != nil {
			log.Fatalln(err)
		} else {
			go func() {
				tmpDef <- callAPI(website, requestLink)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

func locateDef(word string, apiConf *config, dictFile string, dictionary map[string][]string) []string {

	definition, found := checkDict(word, dictionary)
	if !found {
		tmpDef := make(chan []string, len(apiConf.Website))
		checkWeb(word, apiConf, tmpDef)

		definition = parseChan(tmpDef, len(apiConf.Website))

		return definition
	}

	return definition
}

func parseFlags() (bool, int, []string) {
	verbose := flag.Bool("verbose", false, "Display word with definitions")
	depth := flag.Int("depth", 0, "Amount of definitions displayed")

	flag.Parse()

	return *verbose, *depth, flag.Args()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("invalid number of arguments\n")
		return
	}

	apiConf, defPath := getConfig()
	dictFile := getDictConf(apiConf)

	// TODO(#23): Should we update the dictionary to reflect definitions from multiple sources
	//  this may lead to over the top defintions, or repeats.
	dictionary := getDict(defPath + "/" + dictFile)

	verbose, depth, words := parseFlags()
	for index := 0; index < len(words); index++ {
		word := words[index]
		definition := locateDef(words[index], apiConf, dictFile, dictionary)

		failed := verifyDef(definition)
		if failed {
			fmt.Printf("\"%v\" - Not found \n", word)
			continue
		}

		updateDict(dictionary, word, definition)
		storeJSON(defPath+"/"+dictFile, dictionary)

		if verbose {
			fmt.Printf("\n%v:\n", word)
		}

		displayDef(definition, 0, depth)
	}
}
