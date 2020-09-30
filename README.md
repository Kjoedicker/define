# define
A cli tool for defining words, and building a local dictionary.

## About
Allows you to define, store, and reference previously defined words.

## Installation

```console
$ go get github.com/Kjoedicker/define
```

You will need to create a developer account with [dictionary](dictionary.com)

This is required inorder to gain acess to an api key, which you will need to include in ```conf.yaml```

## conf.yaml

```conf.yaml``` file is defined in the executable path

The following is to be included:

```yaml
---
website: dictionary.com [0]
link: https://www.dictionaryapi.com/api/v3/references/collegiate/json/
apiKey: <insert apikey>
dictionary: dictionary.json [1]
```

**[0]** At this stage in development, there is only one api that is interfaced with.

**[1]** Another filename can be specified.
