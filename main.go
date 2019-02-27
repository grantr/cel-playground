package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func main() {
	// Create the CEL environment with declarations for the input attributes and
	// the desired extension functions. In many cases the desired functionality will
	// be present in a built-in function.
	decls := cel.Declarations(
		// Identifiers used within this expression.
		decls.NewIdent("eventType", decls.String, nil),
		decls.NewIdent("source", decls.String, nil),
	)
	e, err := cel.NewEnv(decls)
	if err != nil {
		log.Fatalf("environment creation error: %s\n", err)
	}

	// Parse and check the expression.
	exp := `eventType == "com.github.pull_request.create"`
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
		"eventType": "com.github.pull_request.create",
		"source":    "github.com/knative/eventing/pulls/21",
	}))
	if err != nil {
		log.Fatalf("runtime error: %s\n", err)
	}

	fmt.Println("out:", out)
}
