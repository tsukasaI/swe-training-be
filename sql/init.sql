create table `users`(
	`id` int unsigned auto_increment,
	`name` varchar(50) not null,
	`email` varchar(100) not null,
	`password` varchar(255) not null,
	`created_at` datetime,
	`updated_at` datetime,
	primary key(`id`)
);

insert into `users` (`name`, `email`, `password`, `created_at`, `updated_at`)
values
	('user1', 'user1@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user2', 'user2@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user3', 'user3@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user4', 'user4@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user5', 'user5@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user6', 'user6@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user7', 'user7@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user8', 'user8@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user9', 'user9@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP()),
	('user10', 'user10@a.com', 'password', '2022-05-18 09:00:00', CURRENT_TIMESTAMP())
;

create table `user_follows` (
    `user_id` int unsigned not null,
    `follow_id` int unsigned not null,
    primary key (`user_id`, `follow_id`),
    foreign key (`user_id`) references `users` (id) on delete cascade,
    foreign key (`follow_id`) references `users` (id) on delete cascade
);

insert into `user_follows` (`user_id`, `follow_id`)
values
	(1, 2),
	(1, 3),
	(1, 4),
	(1, 5),
	(1, 6),
	(2, 3),
	(2, 4),
	(2, 5),
	(2, 6),
	(2, 7)
;

create table `posts` (
	`id` int unsigned auto_increment,
    `comment` varchar(200) not null,
    `user_id` int unsigned not null,
    `created_at` datetime,
    `updated_at` datetime,
	primary key(`id`),
    foreign key (`user_id`) references `users` (id) on delete cascade
);

insert into `posts` (`comment`, `user_id`, `created_at`, `updated_at`)
values
	('hello from user1', 1, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user2', 2, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user3', 3, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user4', 4, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user5', 5, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user6', 6, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user7', 7, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user8', 8, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user9', 9, '2022-05-18 09:00:00', '2022-05-18 09:00:00'),
	('hello from user10', 10, '2022-05-18 09:00:00', '2022-05-18 09:00:00')
;
