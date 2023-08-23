create table social.user
(
    user_id           int auto_increment
        primary key,
    username          varchar(255) not null,
    self_introduction varchar(255) null,
    longitude         double       null,
    latitude          double       null,
    is_online         tinyint(1)   not null
);

