package controllers

import (
	"context"
	"fmt"
	"os"

	github "github.com/google/go-github/v56/github"
)

type GitCOnfig struct {
	username string
	token    string
}
type GitRepository struct {
	Owner    string
	RepoName string
	Branch   string
}

var GithubTokenFilePath string = "/home/ubuntu/github/Service-Chain-Orchestrator/config/git/config"

func GetGithubToken(path string) (string, error) {
	token, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error when get github token", path, err)
		return string(token), err
	}
	return string(token), err
}
func main() {
	token, err := GetGithubToken(GithubTokenFilePath)
	if err != nil {
		fmt.Println("Error when get github token", GithubTokenFilePath, err)
		return
	}
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
func GetMetaDataOfFile(client *github.Client, ctx context.Context, repo GitRepository, path string) (*github.RepositoryContent, error) {
	// repo, res, err := client.Repositories.Get(ctx, "SFC-Demo", "edge")

	opts := &github.RepositoryContentGetOptions{
		Ref: repo.Branch,
	}
	_, content, _, err := client.Repositories.DownloadContentsWithMeta(ctx, repo.Owner, repo.RepoName, path, opts)
	if err != nil {
		fmt.Println(err)
		return content, err
	}
	return content, nil

}
func CreateAFile(client *github.Client, ctx context.Context, repo GitRepository, path string, content *string) error {
	// In this example we're creating a new file in a repository using the
	// Contents API. Only 1 file per commit can be managed through that API.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://godoc.org/github.com/google/go-github/github#hdr-Authentication

	sha, _ := GetFileSHA(client, ctx, repo, path)
	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("SFC Controller updated file"),
		Content:   []byte(*content),
		Branch:    &repo.Branch,
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

func GetFileSHA(client *github.Client, ctx context.Context, repo GitRepository, path string) (string, error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: repo.Branch,
	}
	_, readme, _, err := client.Repositories.DownloadContentsWithMeta(ctx, repo.Owner, repo.RepoName, path, opts)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return readme.GetSHA(), nil

}
// Example path=github.com/test/folder1/newfolder
// newfolder is a name of new folder which will be created
func CreateFolderAtPath(client *github.Client, ctx context.Context, repo GitRepository, path string) error {
	// Example:
	// 1. Create a README file at that folder
	// 2. Delete README file (if needed)
	fpath := path + "/README.md"
	err := CreateAFile(client, ctx, repo, fpath, github.String(""))
	if err != nil {
		fmt.Println("Error when create new folder")
		return err
	}

	return nil
}