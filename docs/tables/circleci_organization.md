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

### Contexts of an organizations

```sql
select
  c ->> 'id' as id,
  c ->> 'name' as name,
  c ->> 'created_at' as created_at
from
  circleci_organization o,
  jsonb_array_elements(o.contexts) c
where
  o.slug = 'gh/fluent-cattle';
```
