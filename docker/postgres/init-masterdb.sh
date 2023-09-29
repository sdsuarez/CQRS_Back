#!/bin/bash
set -e

# Configuración de la base de datos maestra
echo "Configurando base de datos maestra..."
echo "wal_level = replica" >> /var/lib/postgresql/data/postgresql.conf
echo "max_replication_slots = 2" >> /var/lib/postgresql/data/postgresql.conf
echo "max_wal_senders = 2" >> /var/lib/postgresql/data/postgresql.conf
