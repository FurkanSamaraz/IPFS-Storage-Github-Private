package pulls

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

var (
	path = "./pullipfs/"
)

// Basic example of how to clone a repository using clone options.
func Pullrepo(url string) {

	// Clone the given repository to the given directory
	Info("git clone " + url)

	git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

}
func Pulssrepo(url string, username string, token string) {

	// Clone the given repository to the given directory
	Info("git clone %s ", url)

	r, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: token},
		URL:      url,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)

}
