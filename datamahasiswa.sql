-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 02, 2020 at 03:48 AM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `datamahasiswa`
--

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `Nama` varchar(25) NOT NULL,
  `Nim` varchar(14) NOT NULL,
  `Jurusan` varchar(15) NOT NULL,
  `Angkatan` int(10) NOT NULL,
  `username` varchar(25) NOT NULL,
  `password` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`Nama`, `Nim`, `Jurusan`, `Angkatan`, `username`, `password`) VALUES
('Made Adi Widyananda', '1301174158', 'informatika', 2017, 'adiwid', 'adiwid'),
('hahahaasasasas', '', 'Informatika', 2018, 'adiwid', 'adiww'),
('hahahaasasasas', '', 'Informatika', 2018, 'adiwid', 'adiww'),
('hahahaasasasas', '', 'Informatika', 2018, 'adiwid', 'adiww'),
('bubububu', '1301177', 'Informatika', 2018, 'adiwid', 'adiww'),
('hahaha', '1301155555', 'Informatika', 2018, 'adiwid', 'adiwidaaaa'),
('hahahaasasasas', '1301155555', 'Informatika', 2018, 'adiwid', 'adiwidaaaa'),
('aaaa', '122222', 'Industri', 2020, 'hahaha', 'hihihii'),
('aaaabbbbb', '122222', 'Industri', 2020, 'hahaha', 'hihihii'),
('aaaabbbbbcccc', '122222', 'Industri', 2020, 'hahaha', 'hihihii');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
