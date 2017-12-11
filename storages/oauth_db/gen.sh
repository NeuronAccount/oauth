#!/usr/bin/env bash

mysql-orm-gen -sql_file=./oauth_db.sql -orm_file=./oauth_db-gen.go -package_name="oauth_db"