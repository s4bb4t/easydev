-- +goose Up
create table public.roles
(
    user_id integer,
    role    text[]
);

-- +goose Down
DROP TABLE IF EXISTS public.roles;
