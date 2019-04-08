package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

const (
	defaultExpr = `ce.type == "com.github.pull_request.create"`
	defaultCE   = `
		{
			"specversion" : "0.2",
			"type" : "com.example.someevent",
			"source" : "/mycontext",
			"id" : "A234-1234-1234",
			"time" : "2018-04-05T17:31:00Z",
			"comexampleextension1" : "value",
			"comexampleextension2" : {
					"otherValue": 5,
					"stringValue": "value"
			},
			"contenttype" : "text/xml",
			"data" : "<much wow=\"xml\"/>"
		}`
)

var (
	expr   = flag.String("e", defaultExpr, "expression to evaluate")
	ceJSON = flag.String("ce", defaultCE, "CloudEvent as JSON")
)

func main() {
	flag.Parse()
	// Create the CEL environment with declarations for the input attributes and
	// the desired extension functions. In many cases the desired functionality will
	// be present in a built-in function.
	e, err := cel.NewEnv(
		//cel.Types(&dev_knative.CloudEvent{}),
		//cel.Types(&structpb.Struct{}),
		cel.Declarations(
			decls.NewIdent("ce", decls.NewObjectType("google.protobuf.Struct"), nil),
		),
	)

	if err != nil {
		log.Fatalf("environment creation error: %s\n", err)
	}

	// Parse and check the expression.
	p, iss := e.Parse(*expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalln(iss.Err())
	}
	pe, err := cel.AstToParsedExpr(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ParsedExpr:", pe.String())
	c, iss := e.Check(p)
	if iss != nil && iss.Err() != nil {
		log.Fatalln(iss.Err())
	}
	ce, err := cel.AstToCheckedExpr(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("CheckedExpr:", ce.String())

	// Create the program.
	prg, err := e.Program(c)
	if err != nil {
		log.Fatalf("program creation error: %s\n", err)
	}

	cloudEvent := structpb.Struct{}
	if err := jsonpb.Unmarshal(strings.NewReader(*ceJSON), &cloudEvent); err != nil {
		log.Fatalf("json parse error: %s\n", err)
	}
	fmt.Printf("cloudEvent Struct: %#v\n", cloudEvent)
	m := jsonpb.Marshaler{}
	jce, err := m.MarshalToString(&cloudEvent)
	if err != nil {
		log.Fatalf("json marshal error: %s\n", err)
	}
	fmt.Printf("cloudEvent marshaled JSON: %s\n", jce)

	// Evaluate the program against some inputs. Note: the details return is not used.
	out, _, err := prg.Eval(map[string]interface{}{
		// Native values are converted to CEL values under the covers.
		"ce": &cloudEvent,
	})
	if err != nil {
		log.Fatalf("runtime error: %s\n", err)
	}

	fmt.Printf("out: %#v\n", out)
}
