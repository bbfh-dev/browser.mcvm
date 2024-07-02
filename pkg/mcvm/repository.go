package mcvm

// FIXME: Currently uses MCVM CLI to parse repository list
func ListRepositories() []string {
	return ReadMCVMOutput("pkg", "repo", "ls", "--raw")
}

// FIXME: Currently uses MCVM CLI to parse all packages
func ListAllPackages() []string {
	return ReadMCVMOutput("pkg", "list-all")
}
