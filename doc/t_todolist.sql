-- golang.t_todolist definition

CREATE TABLE `t_todolist` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `title` varchar(100) DEFAULT NULL,
                              `status` varchar(100) DEFAULT NULL,
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;