CREATE TABLE IF NOT EXISTS users (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(255) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
);

