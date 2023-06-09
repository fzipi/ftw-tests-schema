# FTW Tests YAML Schema

This repository contains the YAML Schema used in the FTW tool. It is being maintained as code, so it will be easy to include as a library.

## Usage

You will need the `dstdocgen` go generator for creating the documentation in markdown format. The tool is easily installed using:
```
go install github.com/projectdiscovery/yamldoc-go/cmd/docgen/dstdocgen@latest
```

Then you can run:
```
go run mage.go markdown
```

## Generated documentation versions


