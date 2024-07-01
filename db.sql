CREATE DATABASE IF NOT EXISTS amartha;

USE amartha;

CREATE TABLE `customer` (
  `cif_id` int(11) NOT NULL,
  `fullname` varchar(50) NULL,
  `ktp_no` varchar(15) NULL,
  `address` text  NULL,
  `create_date` datetime  NULL,
  `update_date` datetime  NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `loans` (
  `loan_Id` int(11) NOT NULL,
  `cif_id` int(11) NOT NULL,
  `total_amount` int(11)  NULL,
  `outstanding_amount` int(11)  NULL,
  `delinquent` tinyint(1)  NULL,
  `create_time` datetime  NULL,
  `missed_payments` int(11)  NULL,
  `total_weeks` int(11)  NULL,
  `interest_rate` int(3)  NULL,
  `status` varchar(10)  NULL,
  `weekly_payment` int(10)  NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `payment` (
  `payment_id` int(11)  NULL,
  `loan_id` varchar(10)  NULL,
  `payment_date` date  NULL,
  `amount` int(11)  NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



ALTER TABLE `customer`
  ADD PRIMARY KEY (`cif_id`);

ALTER TABLE `loans`
  ADD PRIMARY KEY (`loan_Id`);


ALTER TABLE `payment`
  ADD PRIMARY KEY (`payment_id`);


ALTER TABLE `customer`
  MODIFY `cif_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;


ALTER TABLE `loans`
  MODIFY `loan_Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

ALTER TABLE `payment`
  MODIFY `payment_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;