# Table: circleci_project

A CircleCI project shares the name of the code repository for which it automates workflows, tests, and deployment.

## Examples

### List checkout keys of a project

```sql
select
  k ->> 'fingerprint' as fingerprint,
  k ->> 'preferred' as preferred,
  k ->> 'public_key' as public_key,
  k ->> 'time' as time,
  k ->> 'type' as type,
  k ->> 'login' as login
from
  circleci_project p,
  jsonb_array_elements(p.checkout_keys) k
where
  p.slug = 'gh/fluent-cattle/sp-plugin-test'
order by
  k ->> 'time';
```

### Projects with builds running on the main branch

```sql
select
  concat(username, '/', reponame) as repository,
  'main' as branch,
  jsonb_array_length(branches -> 'main' -> 'running_builds') as running_builds
from
  circleci_project
where
  (
    branches -> 'main'
  )
  is not null;
```

### Project's last successful build (main branch)

```sql
select
  concat(username, '/', reponame) as repository,
  branches -> 'main' -> 'last_success' -> 'build_num' as build_num
from
  circleci_project
where
  (
    branches -> 'main'
  )
  is not null;
```
