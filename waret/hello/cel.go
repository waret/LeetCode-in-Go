package main

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"log"
)

func main() {
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewIdent("name", decls.String, nil),
			decls.NewIdent("group", decls.String, nil)))

	if env != nil {
		parsed, issues := env.Parse(`startsWith("/groups/")`)
		if issues != nil && issues.Err() != nil {
			log.Fatalf("parse error: %s", issues.Err())
		}
		checked, issues := env.Check(parsed)
		if issues != nil && issues.Err() != nil {
			log.Fatalf("type-check error: %s", issues.Err())
		}
		prg, err := env.Program(checked)
		if err != nil {
			log.Fatalf("program construction error: %s", err)
		}
		val, detail, err := prg.Eval(nil)
		if val != nil {
			log.Println(val.Type(), val.Value())
		}
		if detail != nil {
			log.Println(detail.State())
		}
	} else if err != nil {
		log.Println(err)
	}
}
