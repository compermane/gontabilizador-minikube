CREATE USER IF NOT EXISTS 'eugenio'@'%' IDENTIFIED BY 'aA@12345';
GRANT ALL PRIVILEGES ON gontabilizador.* TO 'eugenio'@'%';
FLUSH PRIVILEGES;

CREATE DATABASE IF NOT EXISTS gontabilizador;
USE gontabilizador;

-- Table: ensaio
CREATE TABLE IF NOT EXISTS ensaio (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  data_ensaio DATE NOT NULL,
  UNIQUE(nome, data_ensaio)
);

-- Table: ritmista
CREATE TABLE IF NOT EXISTS ritmista (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(100) NOT NULL UNIQUE,
  naipe VARCHAR(50) NOT NULL,
  modulo VARCHAR(50) NOT NULL
);

-- Table: presenca (relates ensaio and ritmista)
CREATE TABLE IF NOT EXISTS presenca (
  ensaio_id INT NOT NULL,
  ritmista_id INT NOT NULL,
  present BOOLEAN NOT NULL DEFAULT FALSE,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (ensaio_id, ritmista_id),
  FOREIGN KEY (ensaio_id) REFERENCES ensaio(id) ON DELETE CASCADE,
  FOREIGN KEY (ritmista_id) REFERENCES ritmista(id) ON DELETE CASCADE
);

