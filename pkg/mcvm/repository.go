package mcvm

import "log"

// FIXME: Currently uses MCVM CLI to parse repository list
func ListRepositories() []string {
	result, err := ExecMCVM("pkg", "repo", "ls", "--raw")
	if err != nil {
		log.Panicf("Failed to list repositories: %v", err)
	}
	return result
}

// FIXME: Currently uses MCVM CLI to parse all packages
func ListAllPackages() []string {
	result, err := ExecMCVM("pkg", "list-all")
	if err != nil {
		log.Panicf("Failed to list all packages: %v", err)
	}
	return result
}

// FIXME: Currently uses MCVM CLI
func PackageMetadata(id string) (string, error) {
	return ExecMCVMRaw("pkg", "info", "-r", id)
}
