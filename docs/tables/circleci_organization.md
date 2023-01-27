# Table: circleci_organization

CircleCI organization is a representation of a VCS account ownership.

## Examples

### Organizations I have access to

```sql
select
  name,
  vcs_type
from
  circleci_organization
order by
  name;
```

### Context environment variables not updated for more then 90 days across my organizations

```sql
select
  c.organization_slug,
  v.context_id,
  c.name as context,
  v.variable,
  v.created_at,
  v.updated_at
from
  circleci_context_environment_variable v
  join
    circleci_context c
    on c.id = v.context_id
where
  updated_at < current_date - interval '90' day;
```
