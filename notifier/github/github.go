package github

import (
	"context"

	"github.com/google/go-github/github"
)

// API is GitHub API interface
type API interface {
	IssuesCreateComment(ctx context.Context, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
	IssuesDeleteComment(ctx context.Context, commentID int64) (*github.Response, error)
	IssuesListComments(ctx context.Context, number int, opt *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)
	RepositoriesCreateComment(ctx context.Context, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
	RepositoriesListCommits(ctx context.Context, opt *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	RepositoriesGetCommit(ctx context.Context, sha string) (*github.RepositoryCommit, *github.Response, error)
}

// GitHub represents the attribute information necessary for requesting GitHub API
type GitHub struct {
	*github.Client
	owner, repo string
}

// IssuesCreateComment is a wrapper of https://godoc.org/github.com/google/go-github/github#IssuesService.CreateComment
func (g *GitHub) IssuesCreateComment(ctx context.Context, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	return g.Client.Issues.CreateComment(ctx, g.owner, g.repo, number, comment)
}

// IssuesDeleteComment is a wrapper of https://godoc.org/github.com/google/go-github/github#IssuesService.DeleteComment
func (g *GitHub) IssuesDeleteComment(ctx context.Context, commentID int64) (*github.Response, error) {
	return g.Client.Issues.DeleteComment(ctx, g.owner, g.repo, int64(commentID))
}

// IssuesListComments is a wrapper of https://godoc.org/github.com/google/go-github/github#IssuesService.ListComments
func (g *GitHub) IssuesListComments(ctx context.Context, number int, opt *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	return g.Client.Issues.ListComments(ctx, g.owner, g.repo, number, opt)
}

// RepositoriesCreateComment is a wrapper of https://godoc.org/github.com/google/go-github/github#RepositoriesService.CreateComment
func (g *GitHub) RepositoriesCreateComment(ctx context.Context, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error) {
	return g.Client.Repositories.CreateComment(ctx, g.owner, g.repo, sha, comment)
}

// RepositoriesListCommits is a wrapper of https://godoc.org/github.com/google/go-github/github#RepositoriesService.ListCommits
func (g *GitHub) RepositoriesListCommits(ctx context.Context, opt *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error) {
	return g.Client.Repositories.ListCommits(ctx, g.owner, g.repo, opt)
}

// RepositoriesGetCommit is a wrapper of https://godoc.org/github.com/google/go-github/github#RepositoriesService.GetCommit
func (g *GitHub) RepositoriesGetCommit(ctx context.Context, sha string) (*github.RepositoryCommit, *github.Response, error) {
	return g.Client.Repositories.GetCommit(ctx, g.owner, g.repo, sha)
}
