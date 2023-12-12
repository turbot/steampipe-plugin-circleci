---
title: "Steampipe Table: circleci_context_environment_variable - Query CircleCI Context Environment Variables using SQL"
description: "Allows users to query CircleCI Context Environment Variables, specifically providing details about each variable's name, context, and value."
---

# Table: circleci_context_environment_variable - Query CircleCI Context Environment Variables using SQL

CircleCI Context Environment Variables are resources within CircleCI that allow you to store environment variable data. These variables can be shared across multiple projects within a CircleCI organization. They are used to store sensitive information and can be used across multiple jobs, providing a secure and efficient way to manage environment-specific data.

## Table Usage Guide

The `circleci_context_environment_variable` table provides insights into the environment variables within CircleCI Contexts. As a DevOps engineer, explore environment variable-specific details through this table, including their names, contexts, and values. Utilize it to uncover information about the variables, such as those that are shared across multiple projects, and efficiently manage environment-specific data.

## Examples

### Environment variables in a context
Explore the environment variables associated with a specific context in CircleCI to understand when they were created and last updated. This can be useful for auditing and managing configurations across projects.

```sql+postgres
select
  context_id,
  variable,
  created_at,
  updated_at
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b';
```

```sql+sqlite
select
  context_id,
  variable,
  created_at,
  updated_at
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b';
```

### Environment variables not updated for more than 90 days
Determine the areas in which CircleCI environment variables have not been updated for more than 90 days. This could be useful for auditing and maintaining best practices for regular updates and security.

```sql+postgres
select
  context_id,
  variable,
  created_at,
  updated_at 
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b'
  and updated_at < current_date - interval '90' day;
```

```sql+sqlite
select
  context_id,
  variable,
  created_at,
  updated_at 
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b'
  and updated_at < date('now','-90 day');
```