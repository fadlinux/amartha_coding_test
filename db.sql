-- CREATE DATABASE IF NOT EXISTS amartha;

-- USE amartha;

-- CREATE TABLE `customer` (
--   `cif_id` int(11) NOT NULL,
--   `fullname` varchar(50) NOT NULL,
--   `ktp_no` varchar(15) NOT NULL,
--   `address` text NOT NULL,
--   `create_date` datetime DEFAULT NULL,
--   `update_date` datetime DEFAULT NULL
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- CREATE TABLE `loans` (
--   `loan_Id` int(11) NOT NULL,
--   `cif_id` int(11) NOT NULL,
--   `total_amount` int(11) DEFAULT NULL,
--   `outstanding_amount` int(11) DEFAULT NULL,
--   `delinquent` tinyint(1) DEFAULT NULL,
--   `create_time` datetime DEFAULT NULL,
--   `missed_payments` int(11) DEFAULT NULL,
--   `total_weeks` int(11) DEFAULT NULL,
--   `interest_rate` int(3) NOT NULL,
--   `status` varchar(10) NOT NULL
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ALTER TABLE `customer`
--   ADD PRIMARY KEY (`cif_id`);

-- ALTER TABLE `loans`
--   ADD PRIMARY KEY (`loan_Id`);

-- ALTER TABLE `customer`
--   MODIFY `cif_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

-- ALTER TABLE `loans`
--   MODIFY `loan_Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;