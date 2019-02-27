package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	descpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func parseProtos() {

}

const ceEnvelopeProto = `
syntax = "proto3";

package dev.knative;

message CloudEvent {
  string specversion = 1;
  string type = 2;
  string source = 3;
  string id = 4;
  string time = 5;
}
`

func main() {
	// Generate the proto object for CE envelope
	fd := &descpb.FileDescriptorProto{}
	desc := proto.FileDescriptor("ce.proto")
	// desc := []byte(ceEnvelopeProto)
	if err := proto.Unmarshal(desc, fd); err != nil {
		log.Fatalf("bad descriptor: %v\n", err)
	}

	// Create the CEL environment with declarations for the input attributes and
	// the desired extension functions. In many cases the desired functionality will
	// be present in a built-in function.
	e, err := cel.NewEnv(
		cel.Container("dev.knative"),
		cel.IsolateTypes(),
		cel.Types(fd),
		cel.Declarations(
			decls.NewIdent("ce", decls.NewObjectType("dev.knative.CloudEvent"), nil),
		),
	)

	if err != nil {
		log.Fatalf("environment creation error: %s\n", err)
	}

	// Parse and check the expression.
	exp := `ce.type == "com.github.pull_request.create"`
	fmt.Println("expr:", exp)
	p, iss := e.Parse(exp)
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

	// Evaluate the program against some inputs. Note: the details return is not used.
	out, _, err := prg.Eval(cel.Vars(map[string]interface{}{
		// Native values are converted to CEL values under the covers.
		"ce": map[string]interface{}{
			"type":   "com.github.pull_request.create",
			"source": "github.com/knative/eventing/pulls/21",
		},
	}))
	if err != nil {
		log.Fatalf("runtime error: %s\n", err)
	}

	fmt.Println("out:", out)
}
