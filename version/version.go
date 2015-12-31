package version

var (
	Name      = "meld"
	Version   = "0.1.0"
	GitCommit = "HEAD"
)

func FullName() string {
	return Name
}

func FullVersion() string {
	return Version + " (" + GitCommit + ")"
}
