-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Jun 24, 2024 at 12:33 AM
-- Server version: 8.0.33
-- PHP Version: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `casbin-golang`
--

-- --------------------------------------------------------

--
-- Table structure for table `casbin_rule`
--

CREATE TABLE `casbin_rule` (
  `id` bigint UNSIGNED NOT NULL,
  `ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `v1` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `v2` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `casbin_rule`
--
-- OPEN ISSUE: converting NULL to string is unsupported"
-- SOLUTION: to never use NULL, please use an empty string -> ""
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
(5, 'g', '1', 'administrator', '', '', '', ''),
(4, 'p', 'administrator', 'user', 'delete', '', '', ''),
(3, 'p', 'administrator', 'user', 'update', '', '', ''),
(2, 'p', 'administrator', 'user', 'write', '', '', ''),
(1, 'p', 'administrator', 'user', 'read', '', '', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `casbin_rule`
--
ALTER TABLE `casbin_rule`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `casbin_rule`
--
ALTER TABLE `casbin_rule`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
