create table social.friend_relation
(
    relation_id int auto_increment
        primary key,
    user1_id    int not null,
    user2_id    int not null,
    constraint friend_relation_ibfk_1
        foreign key (user1_id) references social.user (user_id)
            on update cascade on delete cascade,
    constraint friend_relation_ibfk_2
        foreign key (user2_id) references social.user (user_id)
            on update cascade on delete cascade
);

create index user1_id
    on social.friend_relation (user1_id);

create index user2_id
    on social.friend_relation (user2_id);

