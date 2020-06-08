-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 08, 2020 at 04:03 AM
-- Server version: 10.1.38-MariaDB
-- PHP Version: 7.2.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `yalla_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `user_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `user_id`) VALUES
(1, 'product', 'description', 13),
(2, 'product', 'description', 13),
(3, 'product2', '', 13);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `first_name`, `last_name`, `image`, `cover`, `email`, `password`) VALUES
(13, 'basel', 'aly', 'uploads\\users\\UQDIWWMNPPQIMGEGMGKB643670114.jpg', 'uploads\\users\\GMWJGSAPGATLMODYLUMG333385551.jpg', 'basel.3ly@gmail.com', '$2a$14$jbKMgoWk7tbBswWEp0XFfeoRO/OkqKM8M1lQVbEWg5TbpprU5HYbe'),
(15, 'basel', 'aly', 'uploads\\users\\GMWJGSAPGATLMODYLUMG570726595.jpg', 'uploads\\users\\UQDIWWMNPPQIMGEGMGKB063113542.jpg', 'basel.rewly@gmail.com', '$2a$14$mRXzZkR6vhrH.4q7yOULgeV2XjqNSzesn92XHrDNYZeUfnxg1LaO.'),
(18, 'basel', 'aly', '', '', 'basel.wly@gmail.com', '$2a$14$rxg04kzj/5Bnj/Jcig24.ePEg/NAOEYoTLdXBnyUSZVEHf3gQO.1q'),
(20, 'basel', 'aly', '', '', 'basel.43244@gmail.com', '$2a$14$lGhQoi9f76bToz7uSZKGJeJS7cU9OtzbuCf73jYoZEu55.B0Q2glS'),
(21, 'basel', 'aly', 'uploads\\users\\GMWJGSAPGATLMODYLUMG647811687.jpg', '', 'basel.11@gmail.com', '$2a$14$MTtrawhqZkNbnfB4Z9gAQO7YYnB0VAKg12eqTEIrjRwkQ2b5a9Hxq'),
(22, 'basel', 'aly', '', '', 'basel.7878@gmail.com', '$2a$14$zvNQIMZU.3tjnjwlbMerc.bLuOpsRN77zM4inOw.YS6TSv2XYaFWW');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
