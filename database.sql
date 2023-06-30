-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: reserva_turnos
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- -----------------------------------------------------
-- Schema reserva_turnos
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `reserva_turnos` ;

-- -----------------------------------------------------
-- Schema reserva_turnos
-- -----------------------------------------------------
CREATE SCHEMA  IF NOT EXISTS `reserva_turnos` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `reserva_turnos`;

--
-- Table structure for table `dentista`
--

DROP TABLE IF EXISTS `dentista`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dentista` (
  `id` int NOT NULL AUTO_INCREMENT,
  `apellido` varchar(45) NOT NULL,
  `nombre` varchar(45) NOT NULL,
  `matricula` varchar(45) NOT NULL,
  `activo` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dentista`
--

LOCK TABLES `dentista` WRITE;
/*!40000 ALTER TABLE `dentista` DISABLE KEYS */;
INSERT INTO `dentista` VALUES (1,'Acedo','Natali','05299',1),(2,'Acetti','Noemi','02396',1),(3,'Bassanesi','Luis','03190',1),(4,'Bertuzzi','Emilce','04592',1),(5,'Calvet','Marina','204576',1),(6,'Cardona','Sergio','03505',1),(7,'Dietrich','Guillermo','03176',1),(8,'Gassibe','Valentina','06095',1),(9,'Ledesma','Veronica','02854',1),(10,'Rucci','Gustavo','01685',1);
/*!40000 ALTER TABLE `dentista` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paciente`
--

DROP TABLE IF EXISTS `paciente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paciente` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) NOT NULL,
  `apellido` varchar(45) NOT NULL,
  `domicilio` varchar(45) NOT NULL,
  `dni` varchar(45) NOT NULL,
  `fechaAlta` varchar(45) NOT NULL,
  `activo` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paciente`
--

LOCK TABLES `paciente` WRITE;
/*!40000 ALTER TABLE `paciente` DISABLE KEYS */;
INSERT INTO `paciente` VALUES (1,'Fabiana','Aguero','Santa Fe 167','37715650','22/04/2023',1),(2,'Mercedes','Alvarez','Belgrano 79','28657827','2/03/2022',1),(3,'Facundo','Battaglino','Cordoba 571','42891027','14/06/2023',1),(4,'Luciana','Faricelli','Libertad 699','34332105','12/12/2022',1),(5,'Alvaro','Fernandez','General Paz 379','30587951','30/05/2023',1),(6,'Walter','Ghiglione','Independencia 274','42187452','10/01/2023',1),(7,'Noelia','Heffner','Mitre 847','29514021','16/10/2022',1),(8,'Yanina','Martinez','Laprida 150','36450508','1/04/2023',1),(9,'Maximiliano','Molina','Cordoba 601','20501489','2/03/2023',1),(10,'Gloria','Ortiz','9 de Julio 489','14991748','26/02/2023',1),(11,'Miguel','Paulucci','San Martin 5813','27592186','11/05/2022',1),(12,'Analia','Pretti','Av. Forestal 580','35861381','18/05/2023',1),(13,'Eduardo','Roccia','San Luis 254','20158748','17/01/2023',1),(14,'Mauricio','Ruffino','Avellaneda 560','23437943','24/08/2022',1),(15,'Lucas','Vogliotti','Buenos Aire 569','40202446','27/06/2023',1);
/*!40000 ALTER TABLE `paciente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `turno`
--

DROP TABLE IF EXISTS `turno`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `turno` (
  `id` int NOT NULL AUTO_INCREMENT,
  `idPaciente` int NOT NULL,
  `idDentista` int NOT NULL,
  `fecha` varchar(45) NOT NULL,
  `hora` varchar(45) NOT NULL,
  `descripcion` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `pacienteId_idx` (`idPaciente`),
  KEY `odontologoId_idx` (`idDentista`),
  CONSTRAINT `idDentista` FOREIGN KEY (`idDentista`) REFERENCES `dentista` (`id`),
  CONSTRAINT `idPaciente` FOREIGN KEY (`idPaciente`) REFERENCES `paciente` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `turno`
--

LOCK TABLES `turno` WRITE;
/*!40000 ALTER TABLE `turno` DISABLE KEYS */;
INSERT INTO `turno` VALUES (1,3,8,'05/06/2023','09:00','Dolor de muela'),(2,8,2,'07/06/2023','09:00','Tratamiento de conducto'),(3,10,8,'05/06/2023','09:30','Blanqueamiento dental'),(4,1,5,'10/06/2023','11:30','Ortodoncia'),(5,2,4,'17/06/2023','10:15','Consulta'),(6,4,3,'20/06/2023','09:45','Tratamiento de caries'),(7,5,1,'26/06/2023','12:15','Sensibilidad dental'),(8,15,10,'23/06/2023','15:00','Blanqueamiento dental'),(9,1,5,'10/07/2023','14:00','Control ortodoncia'),(10,10,8,'05/07/2023','11:00','Consulta');
/*!40000 ALTER TABLE `turno` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

SET SQL_MODE = '';
DROP USER IF EXISTS user1@localhost;
SET SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'password1';

GRANT ALL ON `reserva_turnos`.* TO 'user1'@'localhost';

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-28 14:46:52
