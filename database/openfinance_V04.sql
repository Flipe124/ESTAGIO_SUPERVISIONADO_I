-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 07-Dez-2022 às 13:29
-- Versão do servidor: 10.4.24-MariaDB
-- versão do PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `openfinance`
--

-- --------------------------------------------------------

--
-- Estrutura da tabela `account`
--

CREATE TABLE `account` (
  `id` int(11) NOT NULL,
  `name` varchar(45) NOT NULL,
  `balance` double NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Extraindo dados da tabela `account`
--

INSERT INTO `account` (`id`, `name`, `balance`, `created_at`, `updated_at`) VALUES
(1, 'nubank', 1000, '2022-11-21 10:48:17', '2022-11-21 10:48:17');

-- --------------------------------------------------------

--
-- Estrutura da tabela `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL,
  `name` varchar(45) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `type` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Extraindo dados da tabela `category`
--

INSERT INTO `category` (`id`, `name`, `icon`, `type`, `created_at`, `updated_at`) VALUES
(1, 'casa', '/', 'EXPENSE', '2022-11-21 10:37:40', '2022-11-21 10:37:40'),
(2, 'serviço', '/', 'EXPENSE', '2022-11-21 10:39:24', '2022-11-21 10:39:24'),
(3, 'mercado', '/', 'EXPENSE', '2022-11-21 10:40:07', '2022-11-21 10:40:07'),
(4, 'Salário', '/', 'REVENUE', '2022-11-27 20:06:06', '2022-11-27 20:06:06');

-- --------------------------------------------------------

--
-- Estrutura da tabela `finance`
--

CREATE TABLE `finance` (
  `id` int(11) NOT NULL,
  `account_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `value` double NOT NULL,
  `type` varchar(45) NOT NULL,
  `status` varchar(45) NOT NULL,
  `description` varchar(50) NOT NULL,
  `date` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Extraindo dados da tabela `finance`
--

INSERT INTO `finance` (`id`, `account_id`, `category_id`, `value`, `type`, `status`, `description`, `date`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 100, 'EXPENSE', 'PAID', 'Aluguel', '2022-11-17 10:49:05', '2022-11-21 10:49:55', '2022-11-27 19:46:47'),
(3, 1, 2, 200, 'EXPENSE', 'NOT_PAID', 'Gasolina', '2022-11-09 11:41:38', '2022-11-21 11:42:25', '2022-11-27 19:46:47'),
(5, 1, 3, 1000, 'EXPENSE', 'PAID', 'Supermercado', '2022-11-10 16:32:27', '2022-11-22 16:33:13', '2022-11-27 19:53:19'),
(6, 1, 4, 1200.88, 'REVENUE', 'PAID', 'salário', '2022-11-27 20:06:43', '2022-11-27 20:08:02', '2022-11-27 20:08:02');

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `account`
--
ALTER TABLE `account`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- Índices para tabela `finance`
--
ALTER TABLE `finance`
  ADD PRIMARY KEY (`id`,`account_id`,`category_id`),
  ADD KEY `fk_finance_category_idx` (`category_id`),
  ADD KEY `fk_finance_account1_idx` (`account_id`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `account`
--
ALTER TABLE `account`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT de tabela `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT de tabela `finance`
--
ALTER TABLE `finance`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Restrições para despejos de tabelas
--

--
-- Limitadores para a tabela `finance`
--
ALTER TABLE `finance`
  ADD CONSTRAINT `fk_finance_account1` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `fk_finance_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
