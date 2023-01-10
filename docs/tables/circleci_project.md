# Table: circleci_project

A CircleCI project shares the name of the code repository for which it automates workflows, tests, and deployment.

## Examples

### Projects with builds running on the main branch

```sql
select
  concat(username,'/',reponame) as repository,
  'main' as branch,
  jsonb_array_length(branches -> 'main' -> 'running_builds') as running_builds
from
  circleci_project
where
  (branches -> 'main') is not null
```

### Project's last successful build (main branch)

```sql
select
  concat(username,'/',reponame) as repository,
  branches -> 'main' -> 'last_success' -> 'build_num' as build_num
from
  circleci_project
where
  (branches -> 'main') is not null
```
