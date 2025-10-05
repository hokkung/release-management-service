package githuby

import (
	"context"
	"sort"
	"time"

	"github.com/google/go-github/v75/github"
)

func (g *Github) ListTags(ctx context.Context, user, repoName string) ([]*github.RepositoryTag, error) {
	opt := &github.ListOptions{PerPage: 50}
	var all []*github.RepositoryTag
	for {
		tags, resps, err := g.client.Repositories.ListTags(ctx, user, repoName, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, tags...)
		if resps.NextPage == 0 {
			break
		}
		opt.Page = resps.NextPage
	}
	return all, nil
}

func (g *Github) getLatestTag(ctx context.Context, owner, repo string) (*github.RepositoryTag, error) {
	opt := &github.ListOptions{Page: 1, PerPage: 50}
	tags, _, err := g.client.Repositories.ListTags(ctx, owner, repo, opt)
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, nil
	}
	return tags[0], nil
}

func (g *Github) getTags(ctx context.Context, client *github.Client, owner, repo string) ([]*github.RepositoryTag, error) {
	opt := &github.ListOptions{Page: 1, PerPage: 50}
	var all []*github.RepositoryTag
	for {
		tags, resp, err := client.Repositories.ListTags(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, tags...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// We have tag names + commit SHA. We'll enrich each tag with commit date if available to sort.
	tagPairs := make([]struct {
		Tag  *github.RepositoryTag
		Date time.Time
	}, 0, len(all))

	for _, t := range all {
		sha := t.GetCommit().GetSHA()
		commit, _, err := client.Git.GetCommit(ctx, owner, repo, sha)
		if err != nil || commit == nil || commit.Author == nil || commit.Author.Date == nil {
			// fallback to zero time
			tagPairs = append(tagPairs, struct {
				Tag  *github.RepositoryTag
				Date time.Time
			}{Tag: t, Date: time.Time{}})
			continue
		}
		tagPairs = append(tagPairs, struct {
			Tag  *github.RepositoryTag
			Date time.Time
		}{Tag: t, Date: commit.Author.GetDate().Time})
	}

	// sort by Date desc — zero dates will appear at the end
	sort.Slice(tagPairs, func(i, j int) bool {
		iDate := tagPairs[i].Date
		jDate := tagPairs[j].Date
		if iDate.Equal(jDate) {
			// fallback to tag name comparison
			return tagPairs[i].Tag.GetName() > tagPairs[j].Tag.GetName()
		}
		return iDate.After(jDate)
	})

	out := make([]*github.RepositoryTag, 0, len(tagPairs))
	for _, p := range tagPairs {
		out = append(out, p.Tag)
	}
	return out, nil
}
