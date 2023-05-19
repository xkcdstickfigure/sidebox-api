create table account (
    id           uuid          primary key,
    name         text,
    email        text          unique,
    google_id    text          unique,
    created_at   timestamptz
);

create table session (
    id             uuid          primary key,
    account_id     uuid          references account on delete cascade,
    token          text          unique,
    address        text,
    user_agent     text,
    created_at     timestamptz
);
