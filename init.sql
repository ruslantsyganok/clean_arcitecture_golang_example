create table person
(
    id           bigserial primary key,
    first_name   varchar        not null,
    last_name    varchar,
    email        varchar unique not null,
    password     varchar        not null,
    phone_number varchar,
    role         varchar,
    verified     bool,
    email_code   int,
    balance      int default 0
);

create table course
(
    id          bigserial primary key,
    user_id     bigint not null,
    title       varchar,
    price       int,
    description varchar,
    foreign key (user_id) references person (id)
        on delete cascade
);

create table course_section
(
    id          bigserial primary key,
    course_id   bigint,
    title       varchar,
    description varchar,
    file_name   varchar,
    foreign key (course_id) references course (id)
        on delete cascade
);

create table indicator
(
    id          bigserial primary key,
    title       varchar not null,
    description varchar not null
);

create table question
(
    id           bigserial primary key,
    indicator_id bigint,
    title        varchar not null,
    foreign key (indicator_id)
        references indicator (id)
        on delete cascade
);

create table answer
(
    id          bigserial primary key,
    question_id bigint  not null,
    answer      varchar not null,
    score       int     not null,
    foreign key (question_id)
        references question (id)
        on delete cascade
);

create table score
(
    id           bigserial primary key,
    user_id      bigint not null,
    indicator_id int,
    score        int,
    foreign key (user_id) references person (id) on delete cascade,
    foreign key (indicator_id) references indicator (id) on delete cascade
);

create table review
(
    id         bigserial primary key,
    user_id    bigint  not null,
    course_id  bigint  not null,
    feedback   varchar not null,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    foreign key (user_id) references person (id) on delete cascade,
    foreign key (course_id) references course (id) on delete cascade
);

create table transaction
(
    id        bigserial,
    user_id   bigint not null,
    course_id bigint not null,
    status    varchar,
    foreign key (user_id) references person (id) on delete cascade,
    foreign key (course_id) references course (id) on delete cascade
);