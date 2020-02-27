// Code generated by "generate.rb"; DO NOT EDIT.
package backlog

import (
	"context"
	"net/url"

	. "github.com/moutend/go-backlog/pkg/types"
)

func PostAttachment(filepath string) (*Attachment, error) {
	return backlogClient.PostAttachment(filepath)
}

func PostAttachmentContext(ctx context.Context, filepath string) (*Attachment, error) {
	return backlogClient.PostAttachmentContext(ctx, filepath)
}

func GetIssueComments(issueIdOrKey string, query url.Values) ([]*Comment, error) {
	return backlogClient.GetIssueComments(issueIdOrKey, query)
}

func GetIssueCommentsContext(ctx context.Context, issueIdOrKey string, query url.Values) ([]*Comment, error) {
	return backlogClient.GetIssueCommentsContext(ctx, issueIdOrKey, query)
}

func GetPullRequestComments(projectKeyOrId, repositoryNameOrId string, number int64, query url.Values) ([]*Comment, error) {
	return backlogClient.GetPullRequestComments(projectKeyOrId, repositoryNameOrId, number, query)
}

func GetPullRequestCommentsContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, number int64, query url.Values) ([]*Comment, error) {
	return backlogClient.GetPullRequestCommentsContext(ctx, projectKeyOrId, repositoryNameOrId, number, query)
}

func AddIssue(issue *Issue, notifiedUsers []*User) (*Issue, error) {
	return backlogClient.AddIssue(issue, notifiedUsers)
}

func AddIssueContext(ctx context.Context, issue *Issue, notifiedUsers []*User) (*Issue, error) {
	return backlogClient.AddIssueContext(ctx, issue, notifiedUsers)
}

func GetIssue(issueKeyOrId string) (*Issue, error) {
	return backlogClient.GetIssue(issueKeyOrId)
}

func GetIssueContext(ctx context.Context, issueKeyOrId string) (*Issue, error) {
	return backlogClient.GetIssueContext(ctx, issueKeyOrId)
}

func UpdateIssue(issue *Issue, notifiedUsers []*User, comment string) (*Issue, error) {
	return backlogClient.UpdateIssue(issue, notifiedUsers, comment)
}

func UpdateIssueContext(ctx context.Context, issue *Issue, notifiedUsers []*User, comment string) (*Issue, error) {
	return backlogClient.UpdateIssueContext(ctx, issue, notifiedUsers, comment)
}

func DeleteIssue(issueKeyOrId string) (*Issue, error) {
	return backlogClient.DeleteIssue(issueKeyOrId)
}

func DeleteIssueContext(ctx context.Context, issueKeyOrId string) (*Issue, error) {
	return backlogClient.DeleteIssueContext(ctx, issueKeyOrId)
}

func GetIssueTypes(projectIdOrKey string) ([]*IssueType, error) {
	return backlogClient.GetIssueTypes(projectIdOrKey)
}

func GetIssueTypesContext(ctx context.Context, projectIdOrKey string) ([]*IssueType, error) {
	return backlogClient.GetIssueTypesContext(ctx, projectIdOrKey)
}

func GetIssuesCount(query url.Values) (int64, error) {
	return backlogClient.GetIssuesCount(query)
}

func GetIssuesCountContext(ctx context.Context, query url.Values) (int64, error) {
	return backlogClient.GetIssuesCountContext(ctx, query)
}

func GetIssues(query url.Values) ([]*Issue, error) {
	return backlogClient.GetIssues(query)
}

func GetIssuesContext(ctx context.Context, query url.Values) ([]*Issue, error) {
	return backlogClient.GetIssuesContext(ctx, query)
}

func GetAllIssues(maxIssues int, query url.Values) ([]*Issue, error) {
	return backlogClient.GetAllIssues(maxIssues, query)
}

func GetAllIssuesContext(ctx context.Context, maxIssues int, query url.Values) ([]*Issue, error) {
	return backlogClient.GetAllIssuesContext(ctx, maxIssues, query)
}

func GetNotifications(query url.Values) ([]*Notification, error) {
	return backlogClient.GetNotifications(query)
}

func GetNotificationsContext(ctx context.Context, query url.Values) ([]*Notification, error) {
	return backlogClient.GetNotificationsContext(ctx, query)
}

func GetPriorities() ([]*Priority, error) {
	return backlogClient.GetPriorities()
}

func GetPrioritiesContext(ctx context.Context) ([]*Priority, error) {
	return backlogClient.GetPrioritiesContext(ctx)
}

func GetProjectStatuses(projectIdOrKey string) ([]*ProjectStatus, error) {
	return backlogClient.GetProjectStatuses(projectIdOrKey)
}

func GetProjectStatusesContext(ctx context.Context, projectIdOrKey string) ([]*ProjectStatus, error) {
	return backlogClient.GetProjectStatusesContext(ctx, projectIdOrKey)
}

func GetProject(projectKeyOrId string) (*Project, error) {
	return backlogClient.GetProject(projectKeyOrId)
}

func GetProjectContext(ctx context.Context, projectKeyOrId string) (*Project, error) {
	return backlogClient.GetProjectContext(ctx, projectKeyOrId)
}

func GetProjects(query url.Values) ([]*Project, error) {
	return backlogClient.GetProjects(query)
}

func GetProjectsContext(ctx context.Context, query url.Values) ([]*Project, error) {
	return backlogClient.GetProjectsContext(ctx, query)
}

func AddPullRequest(pullRequest *PullRequest, notifiedUsers []*User) (*PullRequest, error) {
	return backlogClient.AddPullRequest(pullRequest, notifiedUsers)
}

func AddPullRequestContext(ctx context.Context, pullRequest *PullRequest, notifiedUsers []*User) (*PullRequest, error) {
	return backlogClient.AddPullRequestContext(ctx, pullRequest, notifiedUsers)
}

func GetPullRequest(projectIdOrKey, repositoryIdOrName string, number int64) (*PullRequest, error) {
	return backlogClient.GetPullRequest(projectIdOrKey, repositoryIdOrName, number)
}

func GetPullRequestContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, number int64) (*PullRequest, error) {
	return backlogClient.GetPullRequestContext(ctx, projectIdOrKey, repositoryIdOrName, number)
}

func UpdatePullRequest(pullRequest *PullRequest, notifiedUsers []*User, comment string) (*PullRequest, error) {
	return backlogClient.UpdatePullRequest(pullRequest, notifiedUsers, comment)
}

func UpdatePullRequestContext(ctx context.Context, pullRequest *PullRequest, notifiedUsers []*User, comment string) (*PullRequest, error) {
	return backlogClient.UpdatePullRequestContext(ctx, pullRequest, notifiedUsers, comment)
}

func GetPullRequestsCount(projectIdOrKey, repositoryIdOrName string) (int64, error) {
	return backlogClient.GetPullRequestsCount(projectIdOrKey, repositoryIdOrName)
}

func GetPullRequestsCountContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string) (int64, error) {
	return backlogClient.GetPullRequestsCountContext(ctx, projectIdOrKey, repositoryIdOrName)
}

func GetPullRequests(projectIdOrKey, repositoryIdOrName string, query url.Values) ([]*PullRequest, error) {
	return backlogClient.GetPullRequests(projectIdOrKey, repositoryIdOrName, query)
}

func GetPullRequestsContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, query url.Values) ([]*PullRequest, error) {
	return backlogClient.GetPullRequestsContext(ctx, projectIdOrKey, repositoryIdOrName, query)
}

func GetAllPullRequests(projectIdOrKey, repositoryIdOrName string) ([]*PullRequest, error) {
	return backlogClient.GetAllPullRequests(projectIdOrKey, repositoryIdOrName)
}

func GetAllPullRequestsContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string) ([]*PullRequest, error) {
	return backlogClient.GetAllPullRequestsContext(ctx, projectIdOrKey, repositoryIdOrName)
}

func GetRepositories(projectKeyOrId string, query url.Values) ([]*Repository, error) {
	return backlogClient.GetRepositories(projectKeyOrId, query)
}

func GetRepositoriesContext(ctx context.Context, projectKeyOrId string, query url.Values) ([]*Repository, error) {
	return backlogClient.GetRepositoriesContext(ctx, projectKeyOrId, query)
}

func GetRepository(projectKeyOrId, repositoryNameOrId string) (*Repository, error) {
	return backlogClient.GetRepository(projectKeyOrId, repositoryNameOrId)
}

func GetRepositoryContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string) (*Repository, error) {
	return backlogClient.GetRepositoryContext(ctx, projectKeyOrId, repositoryNameOrId)
}

func GetSpace() (*Space, error) {
	return backlogClient.GetSpace()
}

func GetSpaceContext(ctx context.Context) (*Space, error) {
	return backlogClient.GetSpaceContext(ctx)
}

func GetSpaceDiskUsage() (*TotalDiskUsage, error) {
	return backlogClient.GetSpaceDiskUsage()
}

func GetSpaceDiskUsageContext(ctx context.Context) (*TotalDiskUsage, error) {
	return backlogClient.GetSpaceDiskUsageContext(ctx)
}

func GetUsers() ([]*User, error) {
	return backlogClient.GetUsers()
}

func GetUsersContext(ctx context.Context) ([]*User, error) {
	return backlogClient.GetUsersContext(ctx)
}

func GetMyself() (*User, error) {
	return backlogClient.GetMyself()
}

func GetMyselfContext(ctx context.Context) (*User, error) {
	return backlogClient.GetMyselfContext(ctx)
}

func AddWiki(wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return backlogClient.AddWiki(wiki, mailNotify)
}

func AddWikiContext(ctx context.Context, wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return backlogClient.AddWikiContext(ctx, wiki, mailNotify)
}

func GetWiki(wikiId uint64) (*Wiki, error) {
	return backlogClient.GetWiki(wikiId)
}

func GetWikiContext(ctx context.Context, wikiId uint64) (*Wiki, error) {
	return backlogClient.GetWikiContext(ctx, wikiId)
}

func UpdateWiki(wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return backlogClient.UpdateWiki(wiki, mailNotify)
}

func UpdateWikiContext(ctx context.Context, wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return backlogClient.UpdateWikiContext(ctx, wiki, mailNotify)
}

func DeleteWiki(wikiId uint64) (*Wiki, error) {
	return backlogClient.DeleteWiki(wikiId)
}

func DeleteWikiContext(ctx context.Context, wikiId uint64) (*Wiki, error) {
	return backlogClient.DeleteWikiContext(ctx, wikiId)
}

func GetWikis(projectIdOrKey string, query url.Values) ([]*Wiki, error) {
	return backlogClient.GetWikis(projectIdOrKey, query)
}

func GetWikisContext(ctx context.Context, projectIdOrKey string, query url.Values) ([]*Wiki, error) {
	return backlogClient.GetWikisContext(ctx, projectIdOrKey, query)
}

func GetWikiCount(projectIdOrKey string) (int, error) {
	return backlogClient.GetWikiCount(projectIdOrKey)
}

func GetWikiCountContext(ctx context.Context, projectIdOrKey string) (int, error) {
	return backlogClient.GetWikiCountContext(ctx, projectIdOrKey)
}

func GetWikiTags(projectIdOrKey string) ([]*Tag, error) {
	return backlogClient.GetWikiTags(projectIdOrKey)
}

func GetWikiTagsContext(ctx context.Context, projectIdOrKey string) ([]*Tag, error) {
	return backlogClient.GetWikiTagsContext(ctx, projectIdOrKey)
}

func GetWikiAttachments(wikiId uint64) ([]*Attachment, error) {
	return backlogClient.GetWikiAttachments(wikiId)
}

func GetWikiAttachmentsContext(ctx context.Context, wikiId uint64) ([]*Attachment, error) {
	return backlogClient.GetWikiAttachmentsContext(ctx, wikiId)
}

func DeleteWikiAttachment(wikiId, attachmentId uint64) (*Attachment, error) {
	return backlogClient.DeleteWikiAttachment(wikiId, attachmentId)
}

func DeleteWikiAttachmentContext(ctx context.Context, wikiId, attachmentId uint64) (*Attachment, error) {
	return backlogClient.DeleteWikiAttachmentContext(ctx, wikiId, attachmentId)
}

func AddWikiAttachments(wikiId uint64, filepaths ...string) ([]*Attachment, error) {
	return backlogClient.AddWikiAttachments(wikiId, filepaths...)
}

func AddWikiAttachmentsContext(ctx context.Context, wikiId uint64, filepaths ...string) ([]*Attachment, error) {
	return backlogClient.AddWikiAttachmentsContext(ctx, wikiId, filepaths...)
}

func DownloadWikiAttachment(wikiId, attachmentId uint64) (data []byte, filename string, err error) {
	return backlogClient.DownloadWikiAttachment(wikiId, attachmentId)
}

func DownloadWikiAttachmentContext(ctx context.Context, wikiId, attachmentId uint64) (data []byte, filename string, err error) {
	return backlogClient.DownloadWikiAttachmentContext(ctx, wikiId, attachmentId)
}
