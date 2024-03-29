package pork

import (
	"fmt"
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// GHRepo contains the repository
type GHRepo struct {
	RepoDir string
	owner   string
	project string
	repo    *git.Repository
}

// NewGHRepo creates a new github repository
func NewGHRepo(repository string) (*GHRepo, error) {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return nil, fmt.Errorf("repository must be in format owner/project")
	}

	return &GHRepo{
		owner:   values[0],
		project: values[1],
	}, nil
}

// RepositoryURL return the repository URL
func (g *GHRepo) RepositoryURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

// Clone clones the repository
func (g *GHRepo) Clone(dest string) error {
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: g.RepositoryURL(),
	})
	if err != nil {
		return err
	}

	g.repo = repo
	g.RepoDir = fullPath
	return nil
}

// Checkout checkouts a repository
func (g *GHRepo) Checkout(ref string, create bool) error {
	opts := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref)),
		Create: create,
	}
	if create {
		head, err := g.repo.Head()
		if err != nil {
			return err
		}
		opts.Hash = head.Hash()
	}

	tree, err := g.repo.Worktree()
	if err != nil {
		return err
	}

	return tree.Checkout(opts)
}

func (g *GHRepo) AddUpstream(repository *GHRepo) error {
	_, err := g.repo.CreateRemote(&config.RemoteConfig{
		Name: "upstream",
		URLs: []string{repository.RepositoryURL()},
	})
	return err
}
