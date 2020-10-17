package main

import (
	"fmt"
	"log"
	"os"
	// "time"
)

// TODO(#11): flags to specify display type. json, text,?
func displayDef(definition []string, traverses int) {
	defer recovery("invalid")

	if traverses == (len(definition) - 1) {
		fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
		return
	}

	fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
	displayDef(definition, traverses+1)
}

func parseDef(tmpDef chan []string, size int) []string {
	tmpDict := make([]string, 0)

	for i := 0; i < size; i++ {
		tmp := <-tmpDef
		for j := 0; j < len(tmp); j++ {
			tmpDict = append(tmpDict, tmp[j])
		}
	}

	return tmpDict
}

func procWord(word string, verbosity int) {

	if verbosity == 1 {
		fmt.Printf("\n%v:\n", word)
	}

	apiConf, defPath := getConfig()
	dictFile := getDictConf(apiConf)

	// TODO(#23): Should we update the dictionary to reflect definitions from multiple sources
	//  this may lead to over the top defintions, or repeats.
	dictionary := getDict(defPath + "/" + dictFile)
	dictDefinition, err := checkDict(word, dictionary)
	if err != nil {
		tmpDef := make(chan []string, len(apiConf.Website))

		for idx := range apiConf.Website {

			website, link, apiKey := parseConfig(apiConf, idx)

			requestLink, err := parseRequest(word, website, link, apiKey)
			if err != nil {
				log.Fatalln(err)
			} else {
				go func() {
					tmpDef <- callAPI(website, requestLink)
				}()
			}
		}

		definition := parseDef(tmpDef, len(apiConf.Website))

		err := updateDict(dictionary, word, definition)
		if err != false {
			fmt.Printf("\"%v\" - not in dictionary - %v\n", word)
		} else {
			storeJSON(defPath+"/"+dictFile, dictionary)
			displayDef(definition, 0)
			return
		}
	}

	displayDef(dictDefinition, 0)
	return
}

func checkFlag(flag string) int {

	if flag == "-v" {
		return 1
	}
	return 0
}

// TODO(#2): Implement more flags, is there a better way to parse flags?
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("invalid number of arguments\n")
		return
	}

	verbosity := checkFlag(os.Args[1])
	for index := verbosity + 1; index < len(os.Args); index++ {
		procWord(os.Args[index], verbosity)
	}
}
