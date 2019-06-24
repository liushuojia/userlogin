-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2019-06-24 09:20:46
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `TUserDB`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE `admin` (
  `admin_id` int(11) NOT NULL,
  `name_en` char(32) NOT NULL DEFAULT '' COMMENT '英文名',
  `realname` char(16) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `admin_email` char(32) NOT NULL DEFAULT '' COMMENT '邮件',
  `admin_email_flag` tinyint(4) NOT NULL COMMENT '邮件验证状态',
  `admin_mobile` char(12) NOT NULL DEFAULT '' COMMENT '手机',
  `admin_mobile_flag` tinyint(4) NOT NULL COMMENT '手机验证状态',
  `admin_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态1 正常 0 停用',
  `admin_verify` char(32) NOT NULL COMMENT '用户密钥',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `admin_role` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否管理员 1是 其他 否',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1 已删 其他 未删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`admin_id`, `name_en`, `realname`, `admin_email`, `admin_email_flag`, `admin_mobile`, `admin_mobile_flag`, `admin_status`, `admin_verify`, `create_time`, `admin_role`, `is_delete`) VALUES
(1, 'hiloy', 'xxx', 'xxx@qq.com', 0, '13001300000', 0, 1, '72D026A9C920F3BD30B41058219E8736', 0, 1, 0),
(2, 'apple', 'xxx', 'xxx@qq.com', 0, '13001311111', 0, 1, '001', 1553093222, 1, 0);

-- --------------------------------------------------------

--
-- 表的结构 `admin_desc_ext`
--

CREATE TABLE `admin_desc_ext` (
  `id` int(11) UNSIGNED NOT NULL,
  `admin_id` int(11) NOT NULL,
  `phone` char(32) NOT NULL DEFAULT '' COMMENT '固话',
  `remark` text NOT NULL COMMENT '描述',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  `create_admin_id` int(11) NOT NULL COMMENT '创建人id',
  `create_realname` char(32) NOT NULL COMMENT '创建人名',
  `stop_time` int(11) NOT NULL COMMENT '离职时间',
  `headimg` char(64) NOT NULL COMMENT '头像图片'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `admin_desc_ext`
--

INSERT INTO `admin_desc_ext` (`id`, `admin_id`, `phone`, `remark`, `update_time`, `create_admin_id`, `create_realname`, `stop_time`, `headimg`) VALUES
(0, 1, '', '', 0, 0, '', 0, '');

