package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/grantr/cel-playground/dev_knative"
)

func main() {
	// Create the CEL environment with declarations for the input attributes and
	// the desired extension functions. In many cases the desired functionality will
	// be present in a built-in function.
	e, err := cel.NewEnv(
		cel.Container("dev.knative"),
		cel.Types(&dev_knative.CloudEvent{}),
		cel.Declarations(
			decls.NewIdent("ce", decls.NewObjectType("dev.knative.CloudEvent"), nil),
		),
	)

	if err != nil {
		log.Fatalf("environment creation error: %s\n", err)
	}

	var exp string
	if len(os.Args) > 1 {
		exp = strings.Join(os.Args[1:], " ")
	} else {
		exp = `ce.type == "com.github.pull_request.create"`
	}

	// Parse and check the expression.
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

	cloudEvent := &dev_knative.CloudEvent{
		Type:   "com.github.pull_request.create",
		Source: "github.com/grantr/cel-playground/pulls/21",
	}
	fmt.Println("cloudEvent: ", cloudEvent.String())

	// Evaluate the program against some inputs. Note: the details return is not used.
	out, _, err := prg.Eval(cel.Vars(map[string]interface{}{
		// Native values are converted to CEL values under the covers.
		"ce": cloudEvent,
	}))
	if err != nil {
		log.Fatalf("runtime error: %s\n", err)
	}

	fmt.Printf("out: %#v\n", out)
	fmt.Printf("pass: %#v\n", out == types.True)
}
