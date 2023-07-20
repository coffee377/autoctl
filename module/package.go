package module

import "path"

// Package https://github.com/lerna/lerna/blob/main/libs/core/src/lib/package.ts

type PackageManager string

const (
	NPM  PackageManager = "npm"
	YARN PackageManager = "yarn"
	PNPM PackageManager = "pnpm"
)

type PublishConfig struct {
	Access    string `json:"access"`
	Tag       string `json:"tag"`
	Registry  string `json:"registry"`
	Directory string `json:"directory"`
}

type Repository struct {
	Type      string
	Url       string
	Directory string
}

type RawManifest struct {
	Name                 string            `json:"name"`
	Version              string            `json:"version"`
	Description          string            `json:"description"`
	Private              bool              `json:"private"`
	Scripts              map[string]string `json:"scripts"`
	Repository           Repository        `json:"repository"`
	Dependencies         map[string]string `json:"dependencies"`
	DevDependencies      map[string]string `json:"devDependencies"`
	OptionalDependencies map[string]string `json:"optionalDependencies"`
	PeerDependencies     map[string]string `json:"peerDependencies"`
	PublishConfig        PublishConfig     `json:"publishConfig"`
	PackageManager       PackageManager    `json:"packageManager"`
	Workspaces           []string          `json:"workspaces"`
}

type Package struct {
	Name     string
	Pkg      *RawManifest
	location string
	//resolved interface{}
	rootPath string
	scripts  map[string]string
	contents string
}

func NewPackage(pkg RawManifest, location string, rootPath string) Package {
	return Package{
		Name:     pkg.Name,
		Pkg:      &pkg,
		location: location,
		rootPath: rootPath,
		scripts:  pkg.Scripts,
	}
}

func (p *Package) IsPrivate() bool {
	return false
}

func (p *Package) SetPrivate(private bool) *Package {
	p.Pkg.Private = private
	return p
}

func (p *Package) GetVersion() string {
	return p.Pkg.Version
}

func (p *Package) SetVersion(version string) *Package {
	p.Pkg.Version = version
	return p
}

func (p *Package) GetLocation() string {
	return p.location
}

func (p *Package) GetRootPath() string {
	return p.rootPath
}

func (p *Package) GetScripts() map[string]string {
	return p.scripts
}

func (p *Package) GetBinLocation() string {
	return path.Join(p.location, "node_modules", ".bin")
}

func (p *Package) GetManifestLocation() string {
	return path.Join(p.location, "package.json")
}

func (p *Package) GetNodeModulesLocation() string {
	return path.Join(p.location, "node_modules")
}

func (p *Package) GetDependencies() map[string]string {
	return p.Pkg.Dependencies
}

func (p *Package) GetDevDependencies() map[string]string {
	return p.Pkg.DevDependencies
}

func (p *Package) GetOptionalDependencies() map[string]string {
	return p.Pkg.OptionalDependencies
}

func (p *Package) GetPeerDependencies() map[string]string {
	return p.Pkg.PeerDependencies
}

func (p *Package) ToJSON() string {
	return ""
}

// Refresh internal state from disk (e.g., changed by external lifecycles)
func (p *Package) Refresh() *Package {
	// todo
	return p
}

// Serialize write manifest changes to disk
func (p *Package) Serialize() *Package {
	// todo
	return p
}
