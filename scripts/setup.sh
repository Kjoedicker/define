#!/bin/bash

#create the dictionary file
touch dictionary.json


#create the conf.yaml file
echo '
---
website: 
  - api:
      website: "dictionaryapi.com"
      link: "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"
      apikey: '$API1'
  - api:
      website: "api.dictionaryapi.dev"
      link: "https://api.dictionaryapi.dev/api/v2/entries/en/"
      apikey: NULL
dictionary: "dictionary.json"
' >> conf.yaml