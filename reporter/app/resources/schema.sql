create table not2_bcc
(
    not2_mail_request_id integer,
    bcc                  varchar
);
create table not2_cc
(
    not2_mail_request_id integer,
    cc                   varchar
);
create table not2_mail_item
(
    id           serial
        primary key
        unique,
    receiver     varchar(256),
    send_date    date,
    meta_data    varchar(4000),
    message_id   varchar,
    state        integer default 0 not null,
    inserttime   date,
    updatetime   date,
    active       boolean default true,
    mail_request integer
);
create table not2_mail_limitation
(
    id                       integer not null
        constraint not2_mail_limitation_pk
            primary key,
    receiver                 varchar not null,
    sender                   varchar not null,
    last_update_time_counter date,
    opt_lock_version         integer default 0,
    counter                  integer not null,
    constraint not2_mail_limitation_pk2
        unique (receiver, sender)
);
create table not2_mail_request
(
    id            integer not null
        constraint not2_mail_request_pk
            primary key,
    inserttime    date,
    updatetime    date,
    active        boolean default true,
    request_id    integer,
    body          varchar,
    text          varchar,
    subject       varchar,
    mail_config   integer,
    reply_address varchar
);
create table not2_send_mail_config
(
    id         integer not null
        constraint not2_send_mail_config_pk
            primary key,
    inserttime date,
    updatetime date,
    active     boolean default true,
    user_id    integer,
    host       varchar(256),
    port       varchar(10),
    password   varchar(256),
    name       varchar(256),
    address    varchar(256)
);