package git_test
	"github.com/sourcegraph/sourcegraph/internal/vcs/git"
		opt  git.RawLogDiffSearchOptions
		want []*git.LogCommitSearchResult
		opt: git.RawLogDiffSearchOptions{
			Query: git.TextSearchOptions{Pattern: "root"},
		want: []*git.LogCommitSearchResult{{
			Commit: git.Commit{
				Author:    git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:06Z")},
				Committer: &git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:06Z")},
			Diff:       &git.Diff{Raw: "diff --git a/f b/f\nindex d8649da..1193ff4 100644\n--- a/f\n+++ b/f\n@@ -1,1 +1,1 @@\n-root\n+branch1\n"},
			Commit: git.Commit{
				Author:    git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:05Z")},
				Committer: &git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:05Z")},
			Diff:       &git.Diff{Raw: "diff --git a/f b/f\nnew file mode 100644\nindex 0000000..d8649da\n--- /dev/null\n+++ b/f\n@@ -0,0 +1,1 @@\n+root\n"},
		opt: git.RawLogDiffSearchOptions{
			Query: git.TextSearchOptions{Pattern: ""},
		want: []*git.LogCommitSearchResult{{
			Commit: git.Commit{
				Author:    git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:06Z")},
				Committer: &git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:06Z")},
			Commit: git.Commit{
				Author:    git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:05Z")},
				Committer: &git.Signature{Name: "a", Email: "a@a.com", Date: MustParseTime(time.RFC3339, "2006-01-02T15:04:05Z")},
		opt: git.RawLogDiffSearchOptions{
			Paths: git.PathOptions{
			results, complete, err := git.RawLogDiffSearch(ctx, repo, test.opt)
		want map[*git.RawLogDiffSearchOptions][]*git.LogCommitSearchResult
			want: map[*git.RawLogDiffSearchOptions][]*git.LogCommitSearchResult{
					Paths: git.PathOptions{IncludePatterns: []string{"/xyz.txt"}, IsRegExp: true},
			results, complete, err := git.RawLogDiffSearch(ctx, test.repo, *opt)