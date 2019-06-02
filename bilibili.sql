/*
 Navicat Premium Data Transfer

 Source Server         : test.xiaotuiai.xyz
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : test.xiaotuiai.xyz:3306
 Source Schema         : scripts

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 02/06/2019 13:35:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bilibili
-- ----------------------------
DROP TABLE IF EXISTS `bilibili`;
CREATE TABLE `bilibili` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '直播平台视频ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '视频标题',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '直播平台视频URL',
  `duration` int(11) NOT NULL DEFAULT '0' COMMENT '视频时长',
  `view` int(11) NOT NULL DEFAULT '0' COMMENT '访问量',
  `danmaku` int(11) NOT NULL DEFAULT '0' COMMENT '弹幕量',
  `reply` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `favorite` int(11) NOT NULL DEFAULT '0' COMMENT '关注量',
  `coin` int(11) NOT NULL DEFAULT '0',
  `share` int(11) NOT NULL DEFAULT '0' COMMENT '分享数',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '喜欢数',
  `now_rank` int(11) NOT NULL DEFAULT '0',
  `his_rank` int(11) NOT NULL DEFAULT '0',
  `keywords` varchar(500) NOT NULL DEFAULT '' COMMENT '直播平台标签',
  `action_tag` varchar(500) NOT NULL DEFAULT '' COMMENT '行为标签',
  `emotion_tag` varchar(500) NOT NULL DEFAULT '' COMMENT '情绪标签',
  `scene_tag` varchar(500) NOT NULL DEFAULT '' COMMENT '场景标签',
  `star_tag` varchar(500) NOT NULL DEFAULT '' COMMENT '明星标签',
  `dialog_tag` varchar(500) NOT NULL DEFAULT '' COMMENT '对话标签',
  `update_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新次数',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `aid` (`aid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5996 DEFAULT CHARSET=utf8mb4 COMMENT='哔哩哔哩';

SET FOREIGN_KEY_CHECKS = 1;
