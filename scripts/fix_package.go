// go:build ignore
//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var types = [...]string{
	"DriveArray",
	"InstanceArray",
	"Infrastructure",
	"InfrastructureOperation",
	"DriveArrayOperation",
	"InstanceArrayOperation",
	"ServerType",
	"Network",
	"NetworkOperation",
	"VolumeTemplate",
	"HardwareConfiguration",
	"ShutdownOptions",
	"Instance",
	"Secret",
	"Variable",
	"User",
	"OSAsset",
	"OSTemplate",
	"ServerSearchResult",
	"Server",
	"ServerComponent",
	"StageDefinition",
	"Workflow",
	"WorkflowStageAssociation",
	"WorkflowStageDefinitionReference",
	"WorkflowReference",
	"Drive",
	"Snapshot",
	"ServerTypeMatches",
	"ServerTypeMatch",
	"Datacenter",
	"DatacenterConfig",
	"SharedDrive",
	"OSTemplateOSAssetData",
	"SharedDriveOperation",
	"SwitchDevice",
	"SwitchDeviceLink",
	"SubnetPool",
	"SubnetPoolUtilization",
	"SwitchInterfaceSearchResult",
	"StoragePoolSearchResult",
	"InfrastructuresSearchResult",
	"ServerEditInventory",
	"ServerEditRack",
}

const packageName = "metalcloud"
const sdkPackage = "github.com/metalsoft-io/metal-cloud-sdk-go/v2"

func main() {

	var input = flag.String("input", "", "input")

	flag.Parse()

	s, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(-1)
	}

	// r := regexp.MustCompile("package ([a-zA-Z0-9_]*)")

	// m := r.FindAllStringSubmatch(string(s), -1)

	// currentPackageName := m[0][1]
	// packageAndImport := fmt.Sprintf("package %s\n\n\nimport %s \"%s\"\n", currentPackageName, packageName, sdkPackage)
	// s = r.ReplaceAll(s, []byte(packageAndImport))

	//     for _, v := range types {

	// 	r := regexp.MustCompile(fmt.Sprintf("\\b%s\\b", v))

	// 	p := fmt.Sprintf("%s.%s", packageName, v)
	// 	s = r.ReplaceAll(s, []byte(p))
	// }

	r := regexp.MustCompile("import\\s*\\([\\s\n]*\\. \"github\\.com\\/onsi\\/gomega\"[\\s\n]*\\)")

	s = r.ReplaceAll(s, []byte(""))

	err = ioutil.WriteFile(*input, s, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(-1)
	}

}
