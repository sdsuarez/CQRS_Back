#!/bin/bash
set -e

# Inicialización de la base de datos réplica
echo "Iniciando configuración de base de datos réplica..."
pg_basebackup -h write-db -D /var/lib/postgresql/data -U replicator -v --wal-method=stream
