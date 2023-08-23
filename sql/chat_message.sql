create table social.chat_message
(
    message_id      int auto_increment
        primary key,
    message_content varchar(255) null,
    from_user_id    int          not null,
    to_user_id      int          not null,
    constraint chat_message_ibfk_1
        foreign key (from_user_id) references social.user (user_id)
            on update cascade on delete cascade,
    constraint chat_message_ibfk_2
        foreign key (to_user_id) references social.user (user_id)
            on update cascade on delete cascade
);

create index from_user_id
    on social.chat_message (from_user_id);

create index to_user_id
    on social.chat_message (to_user_id);

