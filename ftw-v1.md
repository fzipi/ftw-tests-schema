



## FTWTest
Welcome to the FTW YAMLFormat documentation.
In this document we will explain all the possible options that can be used within the YAML format.
Generally this is the preferred format for writing tests in as they don't require any programming skills
in order to understand and change. If you find a bug in this format please open an issue.

FTWTest is the base type used when unmarshaling YAML tests files







<hr />

<div class="dd">

<code>meta</code>  <i><a href="#meta">Meta</a></i>

</div>
<div class="dt">

description: |
Meta describes the metadata information of this yaml test file


</div>

<hr />

<div class="dd">

<code>filename</code>  <i>string</i>

</div>
<div class="dt">

description: |
FileName is the name of the file where these tests are.
examples:
  - name: FileName
    value: test-1234.yaml


</div>

<hr />

<div class="dd">

<code>tests</code>  <i>[]<a href="#test">Test</a></i>

</div>
<div class="dt">

description: |
Tests is a list of FTW tests
examples:
  - name: Tests
    value: the tests


</div>

<hr />





## Meta
Meta describes the metadata information of this yaml test file

Appears in:


- <code><a href="#ftwtest">FTWTest</a>.meta</code>





<hr />

<div class="dd">

<code>author</code>  <i>string</i>

</div>
<div class="dt">

description: |
Author is the list of authors that added content to this file
examples:
  - name: Author
    value: Felipe Zipitria


</div>

<hr />

<div class="dd">

<code>enabled</code>  <i>bool</i>

</div>
<div class="dt">

description: |
Enabled indicates if the tests are enabled to be run by the engine or not.
examples:
  - name: Enabled
    value: false


</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

description: |
Name is the name of the tests contained in this file.
examples:
  - name: Name
    value: test01


</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

description: |
Description is a textual description of the tests contained in this file.
examples:
  - name: Description
    value: The tests here target SQL injection.


</div>

<hr />

<div class="dd">

<code>version</code>  <i>string</i>

</div>
<div class="dt">

description: |
Version is the version of the YAML Schema.
examples:
  - name: Version
    value: v1


</div>

<hr />





## Test
Test is an individual test

Appears in:


- <code><a href="#ftwtest">FTWTest</a>.tests</code>





<hr />

<div class="dd">

<code>test_title</code>  <i>string</i>

</div>
<div class="dt">

description: |
TestTitle the title of this particular test. It is used for inclusion/exclusion of each run by the tool.
examples:
  - name: TestTitle
    value: 920100-1


</div>

<hr />

<div class="dd">

<code>desc</code>  <i>string</i>

</div>
<div class="dt">

description: |
TestDescription is the description for this particular test. Should be used to describe the internals of
the specific things this test is targeting.
examples:
  - name: TestDescription
    value: This test targets something


</div>

<hr />

<div class="dd">

<code>stages</code>  <i>[]<a href="#stage">Stage</a></i>

</div>
<div class="dt">

description: |
Stages is the list of all the stages to perform this test.


</div>

<hr />





## Stage
Stage is an individual test stage

Appears in:


- <code><a href="#test">Test</a>.stages</code>





<hr />

<div class="dd">

<code>input</code>  <i><a href="#input">Input</a></i>

</div>
<div class="dt">

description: |
Input is the data that is passed to the test
examples:
  - name: Input
    value: test


</div>

<hr />

<div class="dd">

<code>output</code>  <i><a href="#output">Output</a></i>

</div>
<div class="dt">

description: |
Output is the data that is returned from the test
examples:
  - name: Output
    value: test


</div>

<hr />





## Input
Input represents the input request in a stage
The fields `Version`, `Method` and `URI` we want to explicitly now when they are set to ""


Appears in:


- <code><a href="#stage">Stage</a>.input</code>





<hr />

<div class="dd">

<code>dest_addr</code>  <i>string</i>

</div>
<div class="dt">

description: |
DestAddr is the IP of the destination host that the test will send the message to.
examples:
  - name: DestAddr
    value: "127.0.0.1"


</div>

<hr />

<div class="dd">

<code>port</code>  <i>int</i>

</div>
<div class="dt">

description: |
Port allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: Port
    value: 80


</div>

<hr />

<div class="dd">

<code>protocol</code>  <i>string</i>

</div>
<div class="dt">

description: |
Protocol allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: Protocol
    value: 80


</div>

<hr />

<div class="dd">

<code>uri</code>  <i>string</i>

</div>
<div class="dt">

description: |
URI allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: URI
    value: "/get?hello=world"


</div>

<hr />

<div class="dd">

<code>version</code>  <i>string</i>

</div>
<div class="dt">

description: |
Version it the HTTP version used.
examples:
  - name: Version
    value: "1.1"


</div>

<hr />

<div class="dd">

<code>method</code>  <i>string</i>

</div>
<div class="dt">

description: |
Headers is a map of headers to be sent.
examples:
  - name: Headers
    value: list
Headers ftwhttp.Header `yaml:"headers,omitempty" koanf:"headers,omitempty"`
description: |
Method allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: Method
    value: "GET"


</div>

<hr />

<div class="dd">

<code>data</code>  <i>string</i>

</div>
<div class="dt">

description: |
Data allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: Data
    value: "Bibitti bopi"


</div>

<hr />

<div class="dd">

<code>save_cookie</code>  <i>bool</i>

</div>
<div class="dt">

description: |
SaveCookie allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: SaveCookie
    value: 80


</div>

<hr />

<div class="dd">

<code>stop_magic</code>  <i>bool</i>

</div>
<div class="dt">

description: |
StopMagic allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: StopMagic
    value: false


</div>

<hr />

<div class="dd">

<code>encoded_request</code>  <i>string</i>

</div>
<div class="dt">

description: |
EncodedRequest allows you to declare which port on the destination host the tests should connect to.
examples:
  - name: EncodedRequest
    value: "a"


</div>

<hr />

<div class="dd">

<code>raw_request</code>  <i>string</i>

</div>
<div class="dt">

description: |
RAWRequest is deprecated.
examples:
  - name: RAWRequest
    value: asd


</div>

<hr />





## Output
Output is the response expected from the test

Appears in:


- <code><a href="#stage">Stage</a>.output</code>





<hr />

<div class="dd">

<code>status</code>  <i>[]int</i>

</div>
<div class="dt">

description: |
Status describes the HTTP status error code expected as response.
examples:
  - name: Status
    value: [200]


</div>

<hr />

<div class="dd">

<code>response_contains</code>  <i>string</i>

</div>
<div class="dt">

description: |
ResponseContains describes the text that should be contained in the HTTP response.
examples:
  - name: ResponseContains
    value: "Hello, World"


</div>

<hr />

<div class="dd">

<code>log_contains</code>  <i>string</i>

</div>
<div class="dt">

description: |
LogContains describes the text that should be contained in the WAF logs.
examples:
  - name: LogContains
    value: `id "920100"`


</div>

<hr />

<div class="dd">

<code>no_log_contains</code>  <i>string</i>

</div>
<div class="dt">

description: |
NoLogContains describes the text that should be contained in the WAF logs.
examples:
  - name: NoLogContains
    value: `id "920100"


</div>

<hr />

<div class="dd">

<code>expect_error</code>  <i>bool</i>

</div>
<div class="dt">

description: |
ExpectError describes the text that should be contained in the WAF logs.
examples:
  - name: ExpectError
    value: `id "920100"


</div>

<hr />




