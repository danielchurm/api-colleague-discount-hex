#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

#TODO Uncomment if using a DB
#source /smartshop/wait-for-postgres.sh

main() {
#TODO Uncomment if using a DB
#    wait_for_postgres

    echo "Running microservice"
   ./smartshop-service
}


main "$@"
