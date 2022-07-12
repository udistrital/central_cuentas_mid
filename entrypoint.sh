#!/usr/bin/env bash

set -e
set -u
set -o pipefail

# if [ -n "${PARAMETER_STORE:-}" ]; then
#   export CENTRAL_CUENTAS_MID__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/central_cuentas_mid/db/username --output text --query Parameter.Value)"
#   export CENTRAL_CUENTAS_MID__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/central_cuentas_mid/db/password --output text --query Parameter.Value)"
# fi

exec ./main "$@"
