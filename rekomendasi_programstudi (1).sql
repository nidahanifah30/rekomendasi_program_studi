-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 17, 2025 at 05:49 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rekomendasi_programstudi`
--

-- --------------------------------------------------------

--
-- Table structure for table `data_seluruh_siswa`
--

CREATE TABLE `data_seluruh_siswa` (
  `No` int(11) NOT NULL,
  `Kelas` varchar(10) NOT NULL,
  `Jumlah Siswa` int(11) NOT NULL,
  `IPA` int(11) NOT NULL,
  `IPS` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `data_seluruh_siswa`
--

INSERT INTO `data_seluruh_siswa` (`No`, `Kelas`, `Jumlah Siswa`, `IPA`, `IPS`) VALUES
(1, '10', 428, 322, 106),
(2, '11', 427, 320, 107),
(3, '12', 426, 318, 108);

-- --------------------------------------------------------

--
-- Table structure for table `data_siswa`
--

CREATE TABLE `data_siswa` (
  `nisn` int(20) NOT NULL,
  `nama_siswa` varchar(100) NOT NULL,
  `jurusan` varchar(10) NOT NULL,
  `N1` float NOT NULL,
  `N2` float NOT NULL,
  `N3` float NOT NULL,
  `N4` float NOT NULL,
  `N5` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `data_siswa`
--

INSERT INTO `data_siswa` (`nisn`, `nama_siswa`, `jurusan`, `N1`, `N2`, `N3`, `N4`, `N5`) VALUES
(1234567890, 'Aurellia Callista Dewi', 'IPA', 65, 70, 73, 75, 86),
(1237654098, 'Nida Hanifah', 'IPA', 70, 90, 80, 85, 70),
(1239876345, 'Naila Amelia Shahada', 'IPA', 60, 95, 95, 70, 90),
(1239988456, 'Ruth Starlita Sandra', 'IPA', 85, 80, 82, 70, 75);

-- --------------------------------------------------------

--
-- Table structure for table `program_studi`
--

CREATE TABLE `program_studi` (
  `kode` varchar(225) NOT NULL,
  `nama_prodi` varchar(225) NOT NULL,
  `N1` float NOT NULL,
  `N2` float NOT NULL,
  `N3` float NOT NULL,
  `N4` float NOT NULL,
  `N5` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `program_studi`
--

INSERT INTO `program_studi` (`kode`, `nama_prodi`, `N1`, `N2`, `N3`, `N4`, `N5`) VALUES
('ARS1', 'Arsitektur', 78, 80, 82, 75, 70),
('BDS1', 'Bisnis Digital', 75, 80, 75, 78, 80),
('HMS1', 'Hukum', 80, 65, 60, 70, 85),
('IFS1', 'Informatika', 80, 86, 75, 60, 72),
('MJS1', 'Manajemen', 60, 65, 85, 65, 80),
('PMS1', 'Pendidikan Matematika', 75, 78, 80, 70, 70),
('TES1', 'Teknik Elektro', 75, 85, 88, 70, 70),
('TMS1', 'Teknik Mesin', 66, 61, 72, 68, 74),
('TPS1', 'Teknologi Pangan', 70, 78, 80, 75, 75),
('TSS1', 'Teknik Sipil', 86, 75, 88, 60, 72);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `data_seluruh_siswa`
--
ALTER TABLE `data_seluruh_siswa`
  ADD PRIMARY KEY (`No`);

--
-- Indexes for table `data_siswa`
--
ALTER TABLE `data_siswa`
  ADD PRIMARY KEY (`nisn`);

--
-- Indexes for table `program_studi`
--
ALTER TABLE `program_studi`
  ADD PRIMARY KEY (`kode`) USING BTREE;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `data_seluruh_siswa`
--
ALTER TABLE `data_seluruh_siswa`
  MODIFY `No` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
