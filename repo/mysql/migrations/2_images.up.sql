CREATE TABLE IF NOT EXISTS images (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `image_url` varchar(255) DEFAULT NULL,
    `upload_time` timestamp,
    PRIMARY KEY (`id`) USING BTREE
);