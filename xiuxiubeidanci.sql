CREATE TABLE `word` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '单词ID',
  `spelling` varchar(128) NOT NULL COMMENT '单词拼写',
  `in_usa` varchar(128) NOT NULL DEFAULT '' COMMENT '美音音标',
  `in_uk` varchar(128) NOT NULL DEFAULT '' COMMENT '英音音标',
  `w_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0(普通词) 88(词根词缀) ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=138982 DEFAULT CHARSET=utf8mb4 COMMENT='单词库';

CREATE TABLE `word_class` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='单词本分类';

CREATE TABLE `word_class_book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `word_class_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1483 DEFAULT CHARSET=utf8mb4 COMMENT='单词本';

CREATE TABLE `word_class_book_word` (
  `id` int(10) unsigned NOT NULL,
  `word_class_book_id` int(10) unsigned NOT NULL,
  `word_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`word_class_book_id`,`word_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单词本词库';

CREATE TABLE `word_definition` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `word_id` int(10) unsigned NOT NULL COMMENT '单词ID',
  `part_of_speech` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '词性: 1动词;2及物动词;3不及物动词;4形容词;5名词;6可数名词;7不可数名词;8专有名词;9代词;10副词;11介词;12连词;13数词;14基数词;15序数词;16系动词;17助动词;18情态动词;19感叹词;20冠词;21词组短语;22简写缩略词;88词根词缀',
  `def_data` varchar(512) NOT NULL DEFAULT '' COMMENT '词性的定义数据',
  `ref_from` varchar(16) NOT NULL DEFAULT '' COMMENT '来源：user用户；wb韦氏；ox牛津',
  `sort` tinyint(4) unsigned NOT NULL COMMENT '单词定义的排序',
  `cati` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '分类：1学生；2成人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56195 DEFAULT CHARSET=utf8mb4 COMMENT='单词的意思(定义)库';

CREATE TABLE `word_definition_sentence` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `word_definition_id` int(10) unsigned NOT NULL COMMENT '指定单词意思的ID',
  `ex_data` varchar(512) NOT NULL,
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '定义内的排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94169 DEFAULT CHARSET=utf8mb4 COMMENT='单词特定意思使用的例句';

CREATE TABLE `word_definition_sentence_translation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `word_definition_sentence_id` int(10) unsigned NOT NULL,
  `trs_data` varchar(512) NOT NULL,
  `lang` char(8) NOT NULL DEFAULT '' COMMENT '被翻译成的目标语言',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94162 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `word_definition_translation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `word_definition_id` int(10) unsigned NOT NULL,
  `trs_data` varchar(512) NOT NULL,
  `lang` char(8) NOT NULL COMMENT '被翻译成的目标语言',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25346 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `word_relationship` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `major` varchar(128) NOT NULL COMMENT '主单词',
  `minor` varchar(128) NOT NULL COMMENT '从单词',
  `relationship` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '从单词是主单词的什么关系: 1形容词比较级;2形容词最高级;3副词比较级;4副词最高级;5过去式;6过去分词;7现在分词;8第三人称单数形式;9反义词;10同(近)义词;11派生词;12同源词;13形近词;14关联短语或词组;15名词复数;16组合词',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=147233 DEFAULT CHARSET=utf8mb4 COMMENT='单词与单词间关系， 见words_lib_config.php';

CREATE TABLE `word_show_name` (
  `word_id` int(10) unsigned NOT NULL,
  `spelling` varchar(128) NOT NULL,
  PRIMARY KEY (`word_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='某些单词的显示名，主要一些专有名词和缩写词，显示和词库不一样，显示需要保持大小区别，在词库里都是小写的';

