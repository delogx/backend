## tables needed
- users
- apps
- dashboard_users
- sessions
- page_visits
- components
- string_components
- child_components
- component_children
- transactions

## apps
- id auto increment 
- hostname string unique
- created_at timestamp default now
- updated_at timestamp default now

## dashboard_users
- id auto increment
- email string
- name string
- is_admin boolean default false
- app_id fk apps.id
- created_at timestamp default now
- updated_at timestamp default now

## users
- id auto increment
- ip string nullable
- app_id fk apps.id
- user_id_from_app string nullable
- created_at timestamp default now
- updated_at timestamp default now

## sessions
- id
- ip
- dashboard_user_id
- created_at

## page_visits
- id
- session_id
- url
- created_at

## components
- id
- tag_name string
- attributes jsonb
- text_content string
- children jsonb

## string_components
- id
- value string
