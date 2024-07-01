CREATE DATABASE IF NOT EXISTS amartha;

USE amartha;

CREATE TABLE `customer` (
  `cif_id` int(11) NOT NULL,
  `fullname` varchar(50) NOT NULL,
  `ktp_no` varchar(15) NOT NULL,
  `address` text NOT NULL,
  `create_date` datetime NOT NULL,
  `update_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `loans` (
  `loan_Id` int(11) NOT NULL,
  `cif_id` int(11) NOT NULL,
  `total_amount` int(11) NOT NULL,
  `outstanding_amount` int(11) NOT NULL,
  `delinquent` tinyint(1) NOT NULL,
  `create_time` datetime NOT NULL,
  `missed_payments` int(11) NOT NULL,
  `total_weeks` int(11) NOT NULL,
  `interest_rate` int(3) NOT NULL,
  `status` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `payment` (
  `payment_id` int(11) NOT NULL,
  `load_id` varchar(10) NOT NULL,
  `payment_date` date NOT NULL,
  `amount` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `customer`
  ADD PRIMARY KEY (`cif_id`);

ALTER TABLE `loans`
  ADD PRIMARY KEY (`loan_Id`);

ALTER TABLE `payment`
  ADD PRIMARY KEY (`payment_id`);


ALTER TABLE `customer`
  MODIFY `cif_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

ALTER TABLE `loans`
  MODIFY `loan_Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

ALTER TABLE `payment`
  MODIFY `payment_id` int(11) NOT NULL AUTO_INCREMENT;