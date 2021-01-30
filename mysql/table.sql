drop table clouddist;
create database clouddist;
use clouddist;
set foreign_key_checks=0;

create table `user` (
    `id` int(10) unsigned not null auto_increment,
    `username` varchar(50) default '' comment '账号',
    `password` varchar(50) default '' comment '密码',
    `email` varchar(255) default '' comment '邮箱',
    `privilege` int(5) default 0 comment "用户角色",
    `created_at` DATETIME comment '创建时间',
    `updated_at` DATETIME comment '修改时间',
    `deleted_at` DATETIME comment '删除时间',
    primary key(`id`)
)engine=InnoDB auto_increment=2 default charset=utf8;