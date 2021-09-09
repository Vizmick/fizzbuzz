create table if not exists fb_request (
    `id` int not null auto_increment,
    `int1` int,
    `int2` int,
    `limit` int,
    `str1` varchar(250),
    `str2` varchar(250),
    primary key (`id`)
);