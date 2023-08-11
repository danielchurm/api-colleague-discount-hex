#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

source /smartshop/wait-for-postgres.sh

main() {
    wait_for_postgres

    echo "Running migrations"
    make migrate
}


main "$@"
