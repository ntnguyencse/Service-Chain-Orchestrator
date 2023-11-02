package controllers

import (
	"context"
	"fmt"

	github "github.com/google/go-github/v56/github"
)

type GitCOnfig struct {
	username string
	token    string
}
type GitRepository struct {
	Owner    string
	RepoName string
}

var token string = "ghp_"

func main() {
	client := github.NewClient(nil).WithAuthToken(token)
	ctx := context.Background()
	repo := GitRepository{
		Owner:    "SFC-Demo",
		RepoName: "edge",
	}
	filePath := "README.md"
	fileContent := "This is the content of my file\nand the 2nd line of it version 2"
	CreateAFile(client, ctx, repo, filePath, &fileContent)

}
func GetMetaDataOfFile(client *github.Client, ctx context.Context, path string) {
	// repo, res, err := client.Repositories.Get(ctx, "SFC-Demo", "edge")

}
func CreateAFile(client *github.Client, ctx context.Context, repo GitRepository, path string, content *string) error {
	// In this example we're creating a new file in a repository using the
	// Contents API. Only 1 file per commit can be managed through that API.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://godoc.org/github.com/google/go-github/github#hdr-Authentication

	sha, _ := GetFileSHA(client, ctx, repo.Owner, repo.RepoName, path)
	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("SFC Controller updated file"),
		Content:   []byte(*content),
		Branch:    github.String("dev"),
		SHA:       &sha,
		Committer: &github.CommitAuthor{Name: github.String("SFC Controller"), Email: github.String("sfc@controller.com")},
	}
	_, _, err := client.Repositories.CreateFile(ctx, repo.Owner, repo.RepoName, path, opts)
	if err != nil {
		fmt.Println("Error when update github file", err)
		return err
	}
	fmt.Println("Updated github file")
	return nil

}

func GetFileSHA(client *github.Client, ctx context.Context, owner, repo, path string) (string, error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: "dev",
	}
	_, readme, _, err := client.Repositories.DownloadContentsWithMeta(ctx, owner, repo, path, opts)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return readme.GetSHA(), nil

}
