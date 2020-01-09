package relay

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type HybridConnectionResourceID struct {
	ResourceGroup string
	Name          string
	NamespaceName string
}

type NamespaceResourceID struct {
	ResourceGroup string
	Name          string
}

func ParseHybridConnectionID(input string) (*HybridConnectionResourceID, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Hybrid Connection ID %q: %+v", input, err)
	}
	hybridConnection := HybridConnectionResourceID{
		ResourceGroup: id.ResourceGroup,
	}

	if hybridConnection.Name, err = id.PopSegment("hybridConnections"); err != nil {
		return nil, err
	}

	if hybridConnection.NamespaceName, err = id.PopSegment("namespaces"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &hybridConnection, nil
}

func ParseNamespaceID(input string) (*NamespaceResourceID, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Relay Namespace ID %q: %+v", input, err)
	}
	nameSpace := NamespaceResourceID{
		ResourceGroup: id.ResourceGroup,
	}

	if nameSpace.Name, err = id.PopSegment("namespaces"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &nameSpace, nil
}

// ValidateHybridConnectionID validates that the specified ID is a valid Relay Hybrid Connection ID
func ValidateHybridConnectionID(i interface{}, k string) (warnings []string, errors []error){
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
	}

	if _, err := ParseHybridConnectionID(v); err != nil {
		errors = append(errors, fmt.Errorf("Can not parse %q as a resource id: %v", v, err))
	}

	return warnings, errors
}

// ValidateNamespaceID validates that the specified ID is a valid Relay Namespace ID
func ValidateNamespaceID(i interface{}, k string) (warnings []string, errors []error){
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
	}

	if _, err := ParseNamespaceID(v); err != nil {
		errors = append(errors, fmt.Errorf("Can not parse %q as a resource id: %v", k, err))
		return
	}


	return warnings, errors
}