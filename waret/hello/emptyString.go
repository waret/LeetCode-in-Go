package main

import "fmt"

type K8sObject struct {

	Group     string
	Kind      string
	Name      string
	Namespace string

	json []byte
	yaml []byte
}

func main() {
	o := &K8sObject{
	}

	if o.Namespace == "" {
		fmt.Println("empty")
	}

	if o == nil {
		fmt.Println("nil")
	}
}
