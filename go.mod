module github.com/coreruleset/ftw-tests-schema

go 1.21

require github.com/magefile/mage v1.15.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)

require (
	// Goccy verion is expected to stay at 1.9.2 (Same of go-ftw) because of https://github.com/goccy/go-yaml/issues/325
	github.com/goccy/go-yaml v1.9.2
	github.com/projectdiscovery/yamldoc-go v1.0.4
)

require (
	github.com/fatih/color v1.10.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/stretchr/testify v1.8.4
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
