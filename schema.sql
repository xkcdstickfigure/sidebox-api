create table account (
    id           uuid          primary key,
    name         text,
    email        text          unique,
    google_id    text          unique,
    created_at   timestamptz
);

create table session (
    id           uuid          primary key,
    account_id   uuid          references account on delete cascade,
    token        text          unique,
    address      text,
    user_agent   text,
    created_at   timestamptz
);

create table inbox (
    id           uuid          primary key,
    account_id   uuid          references account on delete cascade,
    code         text          unique,
    name         text,
    created_at   timestamptz
);

create table message (
    id             uuid          primary key,
    inbox_id       uuid          references inbox on delete cascade,
    message_id     text,
    from_name      text,
    from_address   text,
    subject        text,
    body           text,
    html           boolean,
    date           timestamptz
);