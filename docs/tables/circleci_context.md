# Table: circleci_context

CircleCI context provide a mechanism for securing and sharing environment variables across projects.

## Examples

### Contexts across my organizations

```sql
select
  id,
  name,
  organization_slug,
  created_at
from
  circleci_context;
```

### Context environment variables of an organization

```sql
select
  c.name as context,
  v.context_id,
  v.variable,
  v.created_at,
  v.updated_at
from
  circleci_context c
  join
    circleci_context_environment_variable v
    on v.context_id = c.id
where
  organization_slug = 'gh/fluent-cattle';
```
