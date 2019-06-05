// Package provisioners is collection of all the terraform provisioners that are used/required by installer.
package provisioners

// KnownProvisioners is a map of all the known provisioner names to their exec functions.
var KnownProvisioners = map[string]func(){}
