
Create the commit messages in English with the following instructions:

Each commit message should have at least two bullet points:

1. Summarize the changes made in the commit.
2. Explain the reason for each file change.
3. If applicable, mention any related issues or pull requests.
4. If the commit fixes a bug, include the issue number.
5. If the commit adds a feature, describe its purpose.
6. If the commit is a refactor, explain the motivation behind it.
7. If the commit is a chore, clarify what maintenance task was performed.
8. If the commit is a documentation update, specify what was documented.
9. If the commit is a style change, describe the formatting changes made.
10. If the commit is a test addition or modification, explain what was tested and why.
11. If the commit improves performance, detail the performance enhancements.
12. If the commit is a release, mention the version number and any significant changes.
13. If the commit is a hotfix, explain the urgency and the issue it addresses.

## Commit message rules
- Use the Conventional Commit message specification.
- Use the conventional commit format: `<type>(<scope>): <description>`
- Types: feat, fix, docs, style, refactor, test, chore, perf
- Keep the description concise (under 50 characters)
- Use imperative mood (e.g., "add" not "added" or "adds")
- Don't end with a period
- Use lowercase for the first word unless it's a proper noun
- Provide more details in the commit body if needed, separated by a blank line

## Branch naming conventions
- Use kebab-case (lowercase with hyphens)
- Follow the pattern: `<type>/<issue-number>-<short-description>`
- Types: feature, bugfix, hotfix, release, support
- Example: `feature/123-add-dark-mode`

## Pull request guidelines
- Link related issues using keywords (Fixes #123, Closes #456)
- Provide a clear description of changes
- Add screenshots for UI changes
- Ensure all CI checks pass before requesting review
- Keep PRs focused and small when possible