CREATE TABLE `users` (
  `id` char(26) COLLATE utf8mb4_bin NOT NULL,
  `first_name` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `last_name` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `first_name_kana` varchar(256) COLLATE utf8mb4_bin COMMENT 'アカウント氏名（名かな）',
  `last_name_kana` varchar(256) COLLATE utf8mb4_bin COMMENT 'アカウント氏名（性かな）',
  `registered_at` datetime NOT NULL,
  `registration_status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '仮登録ステータス',
  `account_type` varchar(45) COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(256) COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  KEY `users__email` (`email`),
  UNIQUE KEY `users__email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;