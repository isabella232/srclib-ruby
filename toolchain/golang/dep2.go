package golang

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"strings"

	"github.com/golang/gddo/gosrc"
	"sourcegraph.com/sourcegraph"
	"sourcegraph.com/sourcegraph/srcgraph/config"
	"sourcegraph.com/sourcegraph/srcgraph/container"
	"sourcegraph.com/sourcegraph/srcgraph/dep2"
	"sourcegraph.com/sourcegraph/srcgraph/task2"
	"sourcegraph.com/sourcegraph/srcgraph/unit"
)

func init() {
	dep2.RegisterLister(Package{}, dep2.DockerLister{defaultGoVersion})
	dep2.RegisterResolver(goImportPathTargetType, defaultGoVersion)
}

func (v *goVersion) BuildLister(dir string, unit unit.SourceUnit, c *config.Repository, x *task2.Context) (*container.Command, error) {
	goConfig := v.goConfig(c)
	pkg := unit.(*Package)

	dockerfile, err := v.baseDockerfile()
	if err != nil {
		return nil, err
	}

	containerDir := filepath.Join(containerGOPATH, "src", goConfig.BaseImportPath)
	cmd := container.Command{
		Container: container.Container{
			Dockerfile: dockerfile,
			RunOptions: []string{"-v", dir + ":" + containerDir},
			// TODO(sqs): include TestImports and XTestImports
			Cmd: []string{"go", "list", "-f", `[{{if .Imports}}"{{join .Imports "\",\""}}"{{end}}]`, pkg.ImportPath},
		},
		Transform: func(orig []byte) ([]byte, error) {
			var importPaths []string
			err := json.Unmarshal(orig, &importPaths)
			if err != nil {
				return nil, err
			}

			deps := make([]*dep2.RawDependency, len(importPaths))
			for i, importPath := range importPaths {
				deps[i] = &dep2.RawDependency{
					TargetType: goImportPathTargetType,
					Target:     goImportPath(importPath),
				}
			}

			return json.Marshal(deps)
		},
	}
	return &cmd, nil
}

// goImportPath represents a Go import path, such as "github.com/user/repo" or
// "net/http".
type goImportPath string

const goImportPathTargetType = "go-import-path"

func (v *goVersion) Resolve(dep *dep2.RawDependency, c *config.Repository, x *task2.Context) (*dep2.ResolvedTarget, error) {
	importPath := dep.Target.(string)
	return v.resolveGoImportDep(importPath, c)
}

func (v *goVersion) resolveGoImportDep(importPath string, c *config.Repository) (*dep2.ResolvedTarget, error) {
	// Look up in cache.
	resolvedTarget := func() *dep2.ResolvedTarget {
		v.resolveCacheMu.Lock()
		defer v.resolveCacheMu.Unlock()
		return v.resolveCache[importPath]
	}()
	if resolvedTarget != nil {
		return resolvedTarget, nil
	}

	// Check if this importPath is in this repository.
	goConfig := v.goConfig(c)
	if strings.HasPrefix(importPath, goConfig.BaseImportPath) {
		dir, err := filepath.Rel(goConfig.BaseImportPath, importPath)
		if err != nil {
			return nil, err
		}
		toUnit := Package{Dir: dir, ImportPath: importPath}
		return &dep2.ResolvedTarget{
			// TODO(sqs): this is a URI not a clone URL
			ToRepoCloneURL: string(c.URI),
			ToUnitID:       unit.MakeID(toUnit),
		}, nil
	}

	dir, err := gosrc.Get(sourcegraph.AuthenticatingAsNeededHTTPClient, string(importPath), "")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch information about Go package %q", importPath)
	}

	// gosrc returns code.google.com URLs ending in a slash. Remove it.
	dir.ProjectURL = strings.TrimSuffix(dir.ProjectURL, "/")

	toUnit := Package{ImportPath: dir.ImportPath}
	toUnit.Dir, err = filepath.Rel(dir.ProjectRoot, dir.ImportPath)
	if err != nil {
		return nil, err
	}

	resolvedTarget = &dep2.ResolvedTarget{
		ToRepoCloneURL: dir.ProjectURL,
		ToUnitID:       unit.MakeID(toUnit),
	}

	if gosrc.IsGoRepoPath(dir.ImportPath) {
		resolvedTarget.ToVersionString = v.VersionString
		resolvedTarget.ToRevSpec = v.VCSRevision
	}

	// Save in cache.
	v.resolveCacheMu.Lock()
	defer v.resolveCacheMu.Unlock()
	if v.resolveCache == nil {
		v.resolveCache = make(map[string]*dep2.ResolvedTarget)
	}
	v.resolveCache[importPath] = resolvedTarget

	return resolvedTarget, nil
}