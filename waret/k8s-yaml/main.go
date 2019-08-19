package k8s_yaml

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"os"
	"path/filepath"
	//"time"

	//"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	//"k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/deprecated-dynamic"
	"k8s.io/client-go/restmapper"
	//"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	file, err := os.Open("beyondmesh.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yamlFile := yaml.NewYAMLOrJSONDecoder(file, 4096)
	dd := clientset.Discovery()
	apigroups, err := restmapper.GetAPIGroupResources(dd)
	if err != nil {
		log.Fatal(err)
	}

	//restmapper := restmapper.NewDiscoveryRESTMapper(apigroups)
	discoveryRestMapper := restmapper.NewDiscoveryRESTMapper(apigroups)

	for {
		// https://github.com/kubernetes/apimachinery/blob/master/pkg/runtime/types.go
		ext := runtime.RawExtension{}
		if err := yamlFile.Decode(&ext); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("raw: ", string(ext.Raw))
		versions := &runtime.VersionedObjects{}
		obj, gvk, err := unstructured.UnstructuredJSONScheme.Decode(ext.Raw, nil, versions)
		fmt.Println("obj: ", obj)

		// https://github.com/kubernetes/apimachinery/blob/master/pkg/api/meta/interfaces.go
		mapping, err := discoveryRestMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatal(err)
		}

		restconfig := config
		restconfig.GroupVersion = &schema.GroupVersion{
			Group:   mapping.GroupVersionKind.Group,
			Version: mapping.GroupVersionKind.Version,
		}

		dclient, err := dynamic.NewClient(restconfig, *restconfig.GroupVersion)
		if err != nil {
			log.Fatal(err)
		}

		// https://github.com/kubernetes/client-go/blob/master/discovery/discovery_client.go
		apiresourcelist, err := dd.ServerResources()
		if err != nil {
			log.Fatal(err)
		}
		var myapiresource metav1.APIResource
		for _, apiresourcegroup := range apiresourcelist {
			if apiresourcegroup.GroupVersion == mapping.GroupVersionKind.Version {
				for _, apiresource := range apiresourcegroup.APIResources {
					//fmt.Println(apiresource)
					if apiresource.Name == mapping.Resource.String() && apiresource.Kind == mapping.GroupVersionKind.Kind {
						myapiresource = apiresource
					}
				}
			}
		}
		fmt.Println(myapiresource)
		// https://github.com/kubernetes/client-go/blob/master/dynamic/client.go

		var unstruct unstructured.Unstructured
		unstruct.Object = make(map[string]interface{})
		var blob interface{}
		if err := json.Unmarshal(ext.Raw, &blob); err != nil {
			log.Fatal(err)
		}
		unstruct.Object = blob.(map[string]interface{})
		fmt.Println("unstruct:", unstruct)
		ns := "default"
		if md, ok := unstruct.Object["metadata"]; ok {
			metadata := md.(map[string]interface{})
			if internalns, ok := metadata["namespace"]; ok {
				ns = internalns.(string)
			}
		}
		res := dclient.Resource(&myapiresource, ns)
		fmt.Println(res)
		us, err := res.Create(&unstruct)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("unstruct response:", us)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
