// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
// DO NOT EDIT: this file is automatically generated by docgen
package types

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	FTWTestDoc     encoder.Doc
	FTWTestMetaDoc encoder.Doc
	TestDoc        encoder.Doc
	StageDoc       encoder.Doc
	StageDataDoc   encoder.Doc
	InputDoc       encoder.Doc
	FTWTestOutputDoc      encoder.Doc
	FTWTestLogDoc         encoder.Doc
)

func init() {
	FTWTestDoc.Type = "FTWTest"
	FTWTestDoc.Comments[encoder.LineComment] = " Welcome to the FTW YAMLFormat documentation."
	FTWTestDoc.Description = "Welcome to the FTW YAMLFormat documentation.\n In this document we will explain all the possible options that can be used within the YAML format.\n Generally this is the preferred format for writing tests in as they don't require any programming skills\n in order to understand and change. If you find a bug in this format please open an issue.\n\n\n FTWTest is the base type used when unmarshaling YAML tests files"
	FTWTestDoc.Fields = make([]encoder.Doc, 2)
	FTWTestDoc.Fields[0].Name = "meta"
	FTWTestDoc.Fields[0].Type = "FTWTestMeta"
	FTWTestDoc.Fields[0].Note = ""
	FTWTestDoc.Fields[0].Description = "Meta describes the metadata information of this yaml test file"
	FTWTestDoc.Fields[0].Comments[encoder.LineComment] = "Meta describes the metadata information of this yaml test file"
	FTWTestDoc.Fields[1].Name = "tests"
	FTWTestDoc.Fields[1].Type = "[]Test"
	FTWTestDoc.Fields[1].Note = ""
	FTWTestDoc.Fields[1].Description = "Tests is a list of FTW tests"
	FTWTestDoc.Fields[1].Comments[encoder.LineComment] = "Tests is a list of FTW tests"

	FTWTestDoc.Fields[1].AddExample("Tests", "the tests")

	FTWTestMetaDoc.Type = "FTWTestMeta"
	FTWTestMetaDoc.Comments[encoder.LineComment] = ""
	FTWTestMetaDoc.Description = ""
	FTWTestMetaDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "FTWTest",
			FieldName: "meta",
		},
	}
	FTWTestMetaDoc.Fields = make([]encoder.Doc, 5)
	FTWTestMetaDoc.Fields[0].Name = "author"
	FTWTestMetaDoc.Fields[0].Type = "string"
	FTWTestMetaDoc.Fields[0].Note = ""
	FTWTestMetaDoc.Fields[0].Description = "Author is the list of authors that added content to this file"
	FTWTestMetaDoc.Fields[0].Comments[encoder.LineComment] = "Author is the list of authors that added content to this file"

	FTWTestMetaDoc.Fields[0].AddExample("Author", "Felipe Zipitria")
	FTWTestMetaDoc.Fields[1].Name = "enabled"
	FTWTestMetaDoc.Fields[1].Type = "bool"
	FTWTestMetaDoc.Fields[1].Note = ""
	FTWTestMetaDoc.Fields[1].Description = "Enabled indicates if the tests are enabled to be run by the engine or not."
	FTWTestMetaDoc.Fields[1].Comments[encoder.LineComment] = "Enabled indicates if the tests are enabled to be run by the engine or not."

	FTWTestMetaDoc.Fields[1].AddExample("Enabled", false)
	FTWTestMetaDoc.Fields[2].Name = "name"
	FTWTestMetaDoc.Fields[2].Type = "string"
	FTWTestMetaDoc.Fields[2].Note = ""
	FTWTestMetaDoc.Fields[2].Description = "Name is the name of the tests contained in this file."
	FTWTestMetaDoc.Fields[2].Comments[encoder.LineComment] = "Name is the name of the tests contained in this file."

	FTWTestMetaDoc.Fields[2].AddExample("Name", "test01")
	FTWTestMetaDoc.Fields[3].Name = "description"
	FTWTestMetaDoc.Fields[3].Type = "string"
	FTWTestMetaDoc.Fields[3].Note = ""
	FTWTestMetaDoc.Fields[3].Description = "Description is a textual description of the tests contained in this file."
	FTWTestMetaDoc.Fields[3].Comments[encoder.LineComment] = "Description is a textual description of the tests contained in this file."

	FTWTestMetaDoc.Fields[3].AddExample("Description", "The tests here target SQL injection.")
	FTWTestMetaDoc.Fields[4].Name = "version"
	FTWTestMetaDoc.Fields[4].Type = "string"
	FTWTestMetaDoc.Fields[4].Note = ""
	FTWTestMetaDoc.Fields[4].Description = "Version is the version of the YAML Schema."
	FTWTestMetaDoc.Fields[4].Comments[encoder.LineComment] = "Version is the version of the YAML Schema."

	FTWTestMetaDoc.Fields[4].AddExample("Version", "v1")

	TestDoc.Type = "Test"
	TestDoc.Comments[encoder.LineComment] = ""
	TestDoc.Description = ""

	TestDoc.AddExample("Tests", "the tests")
	TestDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "FTWTest",
			FieldName: "tests",
		},
	}
	TestDoc.Fields = make([]encoder.Doc, 3)
	TestDoc.Fields[0].Name = "test_title"
	TestDoc.Fields[0].Type = "string"
	TestDoc.Fields[0].Note = ""
	TestDoc.Fields[0].Description = "TestTitle the title of this particular test. It is used for inclusion/exclusion of each run by the tool."
	TestDoc.Fields[0].Comments[encoder.LineComment] = "TestTitle the title of this particular test. It is used for inclusion/exclusion of each run by the tool."

	TestDoc.Fields[0].AddExample("TestTitle", "920100-1")
	TestDoc.Fields[1].Name = "desc"
	TestDoc.Fields[1].Type = "string"
	TestDoc.Fields[1].Note = ""
	TestDoc.Fields[1].Description = "TestDescription is the description for this particular test. Should be used to describe the internals of\nthe specific things this test is targeting."
	TestDoc.Fields[1].Comments[encoder.LineComment] = "TestDescription is the description for this particular test. Should be used to describe the internals of"

	TestDoc.Fields[1].AddExample("TestDescription", "This test targets something")
	TestDoc.Fields[2].Name = "stages"
	TestDoc.Fields[2].Type = "[]Stage"
	TestDoc.Fields[2].Note = ""
	TestDoc.Fields[2].Description = "Stages is the list of all the stages to perform this test."
	TestDoc.Fields[2].Comments[encoder.LineComment] = "Stages is the list of all the stages to perform this test."

	TestDoc.Fields[2].AddExample("Stages", exampleStages)

	StageDoc.Type = "Stage"
	StageDoc.Comments[encoder.LineComment] = ""
	StageDoc.Description = ""

	StageDoc.AddExample("Stages", exampleStages)
	StageDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Test",
			FieldName: "stages",
		},
	}
	StageDoc.Fields = make([]encoder.Doc, 1)
	StageDoc.Fields[0].Name = "stage"
	StageDoc.Fields[0].Type = "StageData"
	StageDoc.Fields[0].Note = ""
	StageDoc.Fields[0].Description = "StageData is an individual test stage"
	StageDoc.Fields[0].Comments[encoder.LineComment] = "StageData is an individual test stage"

	StageDoc.Fields[0].AddExample("StageData", exampleStageData)

	StageDataDoc.Type = "StageData"
	StageDataDoc.Comments[encoder.LineComment] = ""
	StageDataDoc.Description = ""

	StageDataDoc.AddExample("StageData", exampleStageData)
	StageDataDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Stage",
			FieldName: "stage",
		},
	}
	StageDataDoc.Fields = make([]encoder.Doc, 2)
	StageDataDoc.Fields[0].Name = "input"
	StageDataDoc.Fields[0].Type = "Input"
	StageDataDoc.Fields[0].Note = ""
	StageDataDoc.Fields[0].Description = "Input is the data that is passed to the test"
	StageDataDoc.Fields[0].Comments[encoder.LineComment] = "Input is the data that is passed to the test"

	StageDataDoc.Fields[0].AddExample("Input", exampleInput)
	StageDataDoc.Fields[1].Name = "output"
	StageDataDoc.Fields[1].Type = "Output"
	StageDataDoc.Fields[1].Note = ""
	StageDataDoc.Fields[1].Description = "Output is the data that is returned from the test"
	StageDataDoc.Fields[1].Comments[encoder.LineComment] = "Output is the data that is returned from the test"

	StageDataDoc.Fields[1].AddExample("Output", exampleOutput)

	InputDoc.Type = "Input"
	InputDoc.Comments[encoder.LineComment] = ""
	InputDoc.Description = ""

	InputDoc.AddExample("Input", exampleInput)
	InputDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "StageData",
			FieldName: "input",
		},
	}
	InputDoc.Fields = make([]encoder.Doc, 13)
	InputDoc.Fields[0].Name = "dest_addr"
	InputDoc.Fields[0].Type = "string"
	InputDoc.Fields[0].Note = ""
	InputDoc.Fields[0].Description = "DestAddr is the IP of the destination host that the test will send the message to."
	InputDoc.Fields[0].Comments[encoder.LineComment] = "DestAddr is the IP of the destination host that the test will send the message to."

	InputDoc.Fields[0].AddExample("DestAddr", "127.0.0.1")
	InputDoc.Fields[1].Name = "port"
	InputDoc.Fields[1].Type = "int"
	InputDoc.Fields[1].Note = ""
	InputDoc.Fields[1].Description = "Port allows you to declare which port on the destination host the test should connect to."
	InputDoc.Fields[1].Comments[encoder.LineComment] = "Port allows you to declare which port on the destination host the test should connect to."

	InputDoc.Fields[1].AddExample("Port", 80)
	InputDoc.Fields[2].Name = "protocol"
	InputDoc.Fields[2].Type = "string"
	InputDoc.Fields[2].Note = ""
	InputDoc.Fields[2].Description = "Protocol allows you to declare which protocol the test should use when sending the request."
	InputDoc.Fields[2].Comments[encoder.LineComment] = "Protocol allows you to declare which protocol the test should use when sending the request."

	InputDoc.Fields[2].AddExample("Protocol", "http")
	InputDoc.Fields[3].Name = "uri"
	InputDoc.Fields[3].Type = "string"
	InputDoc.Fields[3].Note = ""
	InputDoc.Fields[3].Description = "URI allows you to declare the URI the test should use as part of the request line."
	InputDoc.Fields[3].Comments[encoder.LineComment] = "URI allows you to declare the URI the test should use as part of the request line."

	InputDoc.Fields[3].AddExample("URI", "/get?hello=world")
	InputDoc.Fields[4].Name = "version"
	InputDoc.Fields[4].Type = "string"
	InputDoc.Fields[4].Note = ""
	InputDoc.Fields[4].Description = "Version allows you to declare the HTTP version the test should use as part of the request line."
	InputDoc.Fields[4].Comments[encoder.LineComment] = "Version allows you to declare the HTTP version the test should use as part of the request line."

	InputDoc.Fields[4].AddExample("Version", "1.1")
	InputDoc.Fields[5].Name = "method"
	InputDoc.Fields[5].Type = "string"
	InputDoc.Fields[5].Note = ""
	InputDoc.Fields[5].Description = "Method allows you to declare the HTTP method the test should use as part of the request line."
	InputDoc.Fields[5].Comments[encoder.LineComment] = "Method allows you to declare the HTTP method the test should use as part of the request line."

	InputDoc.Fields[5].AddExample("Method", "GET")
	InputDoc.Fields[6].Name = "headers"
	InputDoc.Fields[6].Type = "map[string]string"
	InputDoc.Fields[6].Note = ""
	InputDoc.Fields[6].Description = "Method allows you to declare headers that the test should send."
	InputDoc.Fields[6].Comments[encoder.LineComment] = "Method allows you to declare headers that the test should send."

	InputDoc.Fields[6].AddExample("Headers", exampleHeaders)
	InputDoc.Fields[7].Name = "data"
	InputDoc.Fields[7].Type = "string"
	InputDoc.Fields[7].Note = ""
	InputDoc.Fields[7].Description = "Data allows you to declare the payload that the test should in the request body."
	InputDoc.Fields[7].Comments[encoder.LineComment] = "Data allows you to declare the payload that the test should in the request body."

	InputDoc.Fields[7].AddExample("Data", "Bibitti bopi")
	InputDoc.Fields[8].Name = "save_cookie"
	InputDoc.Fields[8].Type = "bool"
	InputDoc.Fields[8].Note = ""
	InputDoc.Fields[8].Description = "SaveCookie allows you to automatically provide cookies if there are multiple stages and save cookie is set"
	InputDoc.Fields[8].Comments[encoder.LineComment] = "SaveCookie allows you to automatically provide cookies if there are multiple stages and save cookie is set"

	InputDoc.Fields[8].AddExample("SaveCookie", 80)
	InputDoc.Fields[9].Name = "stop_magic"
	InputDoc.Fields[9].Type = "bool"
	InputDoc.Fields[9].Note = ""
	InputDoc.Fields[9].Description = "StopMagic is deprecated."
	InputDoc.Fields[9].Comments[encoder.LineComment] = "StopMagic is deprecated."

	InputDoc.Fields[9].AddExample("StopMagic", false)
	InputDoc.Fields[10].Name = "autocomplete_headers"
	InputDoc.Fields[10].Type = "bool"
	InputDoc.Fields[10].Note = ""
	InputDoc.Fields[10].Description = "AutocompleteHeaders allows the test framework to automatically fill the request with Content-Type and Connection headers.\nDefaults to true."
	InputDoc.Fields[10].Comments[encoder.LineComment] = "AutocompleteHeaders allows the test framework to automatically fill the request with Content-Type and Connection headers."

	InputDoc.Fields[10].AddExample("StopMagic", false)
	InputDoc.Fields[11].Name = "encoded_request"
	InputDoc.Fields[11].Type = "string"
	InputDoc.Fields[11].Note = ""
	InputDoc.Fields[11].Description = "EncodedRequest will take a base64 encoded string that will be decoded and sent through as the request.\nIt will override all other settings"
	InputDoc.Fields[11].Comments[encoder.LineComment] = "EncodedRequest will take a base64 encoded string that will be decoded and sent through as the request."

	InputDoc.Fields[11].AddExample("EncodedRequest", "a")
	InputDoc.Fields[12].Name = "raw_request"
	InputDoc.Fields[12].Type = "string"
	InputDoc.Fields[12].Note = ""
	InputDoc.Fields[12].Description = "RAWRequest is deprecated."
	InputDoc.Fields[12].Comments[encoder.LineComment] = "RAWRequest is deprecated."

	InputDoc.Fields[12].AddExample("RAWRequest", "TXkgRGF0YQo=")

	FTWTestOutputDoc.Type = "Output"
	FTWTestOutputDoc.Comments[encoder.LineComment] = ""
	FTWTestOutputDoc.Description = ""

	FTWTestOutputDoc.AddExample("Output", exampleOutput)
	FTWTestOutputDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "StageData",
			FieldName: "output",
		},
	}
	FTWTestOutputDoc.Fields = make([]encoder.Doc, 6)
	FTWTestOutputDoc.Fields[0].Name = "status"
	FTWTestOutputDoc.Fields[0].Type = "int"
	FTWTestOutputDoc.Fields[0].Note = ""
	FTWTestOutputDoc.Fields[0].Description = "Status describes the HTTP status code expected in the response."
	FTWTestOutputDoc.Fields[0].Comments[encoder.LineComment] = "Status describes the HTTP status code expected in the response."

	FTWTestOutputDoc.Fields[0].AddExample("Status", 200)
	FTWTestOutputDoc.Fields[1].Name = "response_contains"
	FTWTestOutputDoc.Fields[1].Type = "string"
	FTWTestOutputDoc.Fields[1].Note = ""
	FTWTestOutputDoc.Fields[1].Description = "ResponseContains describes the text that should be contained in the HTTP response."
	FTWTestOutputDoc.Fields[1].Comments[encoder.LineComment] = "ResponseContains describes the text that should be contained in the HTTP response."

	FTWTestOutputDoc.Fields[1].AddExample("ResponseContains", "Hello, World")
	FTWTestOutputDoc.Fields[2].Name = "log_contains"
	FTWTestOutputDoc.Fields[2].Type = "string"
	FTWTestOutputDoc.Fields[2].Note = ""
	FTWTestOutputDoc.Fields[2].Description = "LogContains describes the text that should be contained in the WAF logs."
	FTWTestOutputDoc.Fields[2].Comments[encoder.LineComment] = "LogContains describes the text that should be contained in the WAF logs."

	FTWTestOutputDoc.Fields[2].AddExample("LogContains", "id 920100")
	FTWTestOutputDoc.Fields[3].Name = "no_log_contains"
	FTWTestOutputDoc.Fields[3].Type = "string"
	FTWTestOutputDoc.Fields[3].Note = ""
	FTWTestOutputDoc.Fields[3].Description = "NoLogContains describes the text that should not be contained in the WAF logs."
	FTWTestOutputDoc.Fields[3].Comments[encoder.LineComment] = "NoLogContains describes the text that should not be contained in the WAF logs."

	FTWTestOutputDoc.Fields[3].AddExample("NoLogContains", "id 920100")
	FTWTestOutputDoc.Fields[4].Name = "log"
	FTWTestOutputDoc.Fields[4].Type = "Log"
	FTWTestOutputDoc.Fields[4].Note = ""
	FTWTestOutputDoc.Fields[4].Description = "Log is used to configure expectations about the log contents."
	FTWTestOutputDoc.Fields[4].Comments[encoder.LineComment] = "Log is used to configure expectations about the log contents."

	FTWTestOutputDoc.Fields[4].AddExample("", exampleLog)
	FTWTestOutputDoc.Fields[5].Name = "expect_error"
	FTWTestOutputDoc.Fields[5].Type = "bool"
	FTWTestOutputDoc.Fields[5].Note = ""
	FTWTestOutputDoc.Fields[5].Description = "When `ExpectError` is true, we don't expect an answer from the WAF, just an error."
	FTWTestOutputDoc.Fields[5].Comments[encoder.LineComment] = "When `ExpectError` is true, we don't expect an answer from the WAF, just an error."

	FTWTestOutputDoc.Fields[5].AddExample("ExpectError", false)

	FTWTestLogDoc.Type = "Log"
	FTWTestLogDoc.Comments[encoder.LineComment] = ""
	FTWTestLogDoc.Description = ""

	FTWTestLogDoc.AddExample("", exampleLog)
	FTWTestLogDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Output",
			FieldName: "log",
		},
	}
	FTWTestLogDoc.Fields = make([]encoder.Doc, 4)
	FTWTestLogDoc.Fields[0].Name = "expect_id"
	FTWTestLogDoc.Fields[0].Type = "int"
	FTWTestLogDoc.Fields[0].Note = ""
	FTWTestLogDoc.Fields[0].Description = "description: |\n   Expect the given ID to be contained in the log output.\n examples:\n   - exampleLog.ExpectId"
	FTWTestLogDoc.Fields[0].Comments[encoder.LineComment] = " description: |"
	FTWTestLogDoc.Fields[1].Name = "expect_id"
	FTWTestLogDoc.Fields[1].Type = "int"
	FTWTestLogDoc.Fields[1].Note = ""
	FTWTestLogDoc.Fields[1].Description = "description: |\n   Expect the given ID _not_ to be contained in the log output.\n examples:\n   - exampleLog.NoExpectId"
	FTWTestLogDoc.Fields[1].Comments[encoder.LineComment] = " description: |"
	FTWTestLogDoc.Fields[2].Name = "match_regex"
	FTWTestLogDoc.Fields[2].Type = "string"
	FTWTestLogDoc.Fields[2].Note = ""
	FTWTestLogDoc.Fields[2].Description = "Expect the regular expression to match log content for the current test."
	FTWTestLogDoc.Fields[2].Comments[encoder.LineComment] = "Expect the regular expression to match log content for the current test."

	FTWTestLogDoc.Fields[2].AddExample("", exampleLog.MatchRegex)
	FTWTestLogDoc.Fields[3].Name = "no_match_regex"
	FTWTestLogDoc.Fields[3].Type = "string"
	FTWTestLogDoc.Fields[3].Note = ""
	FTWTestLogDoc.Fields[3].Description = "Expect the regular expression to _not_ match log content for the current test."
	FTWTestLogDoc.Fields[3].Comments[encoder.LineComment] = "Expect the regular expression to _not_ match log content for the current test."

	FTWTestLogDoc.Fields[3].AddExample("", exampleLog.NoMatchRegex)
}

// GetFTWTestDoc returns documentation for the file ./test_doc.go.
func GetFTWTestDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "FTWTest",
		Description: "",
		Structs: []*encoder.Doc{
			&FTWTestDoc,
			&FTWTestMetaDoc,
			&TestDoc,
			&StageDoc,
			&StageDataDoc,
			&InputDoc,
			&FTWTestOutputDoc,
			&FTWTestLogDoc,
		},
	}
}
