create table social.friend_request
(
    request_id      int auto_increment
        primary key,
    from_user_id    int          not null,
    to_user_id      int          not null,
    request_message varchar(255) null,
    constraint friend_request_ibfk_1
        foreign key (from_user_id) references social.user (user_id)
            on update cascade on delete cascade,
    constraint friend_request_ibfk_2
        foreign key (to_user_id) references social.user (user_id)
            on update cascade on delete cascade
);

create index from_user_id
    on social.friend_request (from_user_id);

create index to_user_id
    on social.friend_request (to_user_id);

