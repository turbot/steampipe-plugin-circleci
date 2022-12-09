# Table: circleci_project

A CircleCI project shares the name of the code repository for which it automates workflows, tests, and deployment.

## Examples

### Projects that share secret environment variables with forks

```sql
select
  concat(username,'/',reponame) as repository
from
  circleci_project
where
  feature_flags ->> 'forks-receive-secret-env-vars' = 'true'
```

## Projects with builds running on main branch

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

## Project's last successful build (main branch)

```sql
select
  concat(username,'/',reponame) as repository,
  branches -> 'main' -> 'last_success' -> 'build_num'
from
  circleci_project
where
  (branches -> 'main') is not null
```
