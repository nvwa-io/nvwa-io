// Copyright 2019 - now The https://github.com/nvwa-io/nvwa-io Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package git

import (
	"errors"
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"strings"
)

func C() *Client {
	return new(Client)
}

type Client struct {
	Auth transport.AuthMethod
}

// Token Auth
func (t *Client) TokenAuth(token string) *Client {
	t.Auth = &http.TokenAuth{Token: token}
	return t
}

// Token Auth
func (t *Client) BasicAuth(username, password string) *Client {
	t.Auth = &http.BasicAuth{
		Username: username,
		Password: password,
	}

	return t
}

// clone repository
func (t *Client) Clone(url, path string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:  url,
		Auth: t.Auth,
	})
	if err != nil {
		logger.Errorf("Failed to clone %s, err=%s", url, err.Error())
		return err
	}

	return nil
}

// get all branches (short)
// e.g:
// fix/checkout-empty-repo
// master
// return(shortBranches, branchCommit, error)
func (t *Client) AllBranches(path string) ([]string, []string, error) {
	remoteBranches, branchCommit, err := t.LsRemoteBranches(path)
	if err != nil {
		return nil, nil, err
	}

	shortBranches := make([]string, 0)
	for _, v := range remoteBranches {
		tmpBranch := strings.Replace(v, "refs/remotes/origin/", "", -1)
		shortBranches = append(shortBranches, tmpBranch)
	}

	return shortBranches, branchCommit, nil
}

// List all remote branches
// e.g:
// refs/remotes/origin/fix/checkout-empty-repo
// refs/remotes/origin/master
// return(branches, branchCommit, error)
func (t *Client) LsRemoteBranches(path string) ([]string, []string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		logger.Errorf("Failed to open %s, err=%s", path, err.Error())
		return nil, nil, err
	}

	refs, err := r.Storer.IterReferences()
	if err != nil {
		logger.Errorf("Failed to IterReferences, err=%s", err.Error())
		return nil, nil, err
	}

	// filter all remote branches
	branches := make([]string, 0)
	branchCommit := make([]string, 0)
	bs := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs)
	err = bs.ForEach(func(b *plumbing.Reference) error {
		branches = append(branches, b.Name().String())
		branchCommit = append(branchCommit, b.Hash().String())
		return nil
	})
	if err != nil {
		logger.Errorf("Failed to ForEach branches, err=%s", err.Error())
		return nil, nil, err
	}

	return branches, branchCommit, nil
}

// get all tags
// e.g:
//v1.0.0
//v2.0.0
//v2.1.0
//v2.1.1
//v2.1.2
//v2.1.3
// return (tags, tagsCommits, err)
func (t *Client) AllTags(path string) ([]string, []string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		logger.Errorf("Failed to open %s, err=%s", path, err.Error())
		return nil, nil, err
	}

	refs, err := r.Storer.IterReferences()
	if err != nil {
		logger.Errorf("Failed to IterReferences, err=%s", err.Error())
		return nil, nil, err
	}

	// filter all tags
	tags := make([]string, 0)
	tagsCommits := make([]string, 0)
	itags := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsTag()
	}, refs)
	err = itags.ForEach(func(b *plumbing.Reference) error {
		tmpTag := strings.Replace(b.Name().String(), "refs/tags/", "", -1)
		tags = append(tags, tmpTag)
		tagsCommits = append(tagsCommits, b.Hash().String())
		return nil
	})
	if err != nil {
		logger.Errorf("Failed to ForEach tags, err=%s", err.Error())
		return nil, nil, err
	}

	return tags, tagsCommits, nil
}

// pull branch
func (t *Client) Pull(path, branch string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		logger.Errorf("Failed to open %s, err=%s", path, err.Error())
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		logger.Errorf("Failed to get work tree, err=%s", err.Error())
		return err
	}

	err = w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  false,
	})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate { // ignore this err
			return nil
		}

		logger.Errorf("Failed to pull, err=%s", err.Error())
		return err
	}

	// Print the latest commit that was just pulled
	//ref, err := r.Head()
	//if err != nil {
	//    logger.Errorf("Failed to get Head, err=%s", err.Error())
	//    return err
	//}
	//commit, err := r.CommitObject(ref.Hash())
	//if err != nil {
	//    logger.Errorf("Failed to get Commit, err=%s", err.Error())
	//    return err
	//}

	return nil
}

// fetch all branches
func (t *Client) FetchAll(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		logger.Errorf("Failed to open %s, err=%s", path, err.Error())
		return err
	}

	err = r.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{
			"refs/*:refs/*",
			"HEAD:refs/heads/HEAD",
			"+refs/heads/*:refs/remotes/origin/*",
		},
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		logger.Errorf("Failed to Fetch, err=%s", err.Error())
		return err
	}

	return nil
}

// checkout branch
func (t *Client) CheckoutBranch(path, branch string) error {
	return t.checkout(path, branch)
}

// checkout tag
func (t *Client) CheckoutTag(path, tag string) error {
	// 1. update repo
	err := t.FetchAll(path)
	if err != nil {
		return err
	}

	// 2. get tag commit id hash
	tags, commits, err := t.AllTags(path)
	if err != nil {
		return err
	}

	commit := ""
	for i, t := range tags {
		if t == tag {
			commit = commits[i]
		}
	}

	if commit == "" {
		return errors.New("Tag not exist: " + tag)
	}

	// 3. checkout by commit id
	return t.checkout(path, commit, false)
}

// checkout branch or checkout commit hash
func (t *Client) checkout(path, name string, isBranch ...bool) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		logger.Errorf("Failed to open %s, err=%s", path, err.Error())
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		logger.Errorf("Failed to get work tree, err=%s", err.Error())
		return err
	}

	// if isBranch is true or not set, checkout branch
	// if isBranch is false, checkout commit hash
	if len(isBranch) > 0 && !isBranch[0] {
		err = w.Checkout(&git.CheckoutOptions{
			Hash: plumbing.NewHash(name),
		})
		if err != nil {
			logger.Errorf("Failed to checkout hash %s, err=%s", name, err.Error())
			return err
		}
	} else {
		remoteRef, err := r.Reference(
			plumbing.ReferenceName(fmt.Sprintf("refs/remotes/origin/%s", name)),
			true,
		)
		if err != nil {
			logger.Errorf("Failed to Reference, err=%s", err.Error())
			return err
		}

		newRef := plumbing.NewHashReference(
			plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", name)),
			remoteRef.Hash(),
		)
		r.Storer.SetReference(newRef)
		err = w.Checkout(&git.CheckoutOptions{
			Branch: newRef.Name(),
			Create: false,
		})
	}

	return nil
}
