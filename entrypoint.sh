#! /usr/bin/sh

PASSWORD=$(aws ecr get-login-password --region eu-west-1)
AUTH=$(base64 -w0 <<<AWS:${PASSWORD})

cat <<EOF > "$1"
{
  "auths": {
    "194182751502.dkr.ecr.eu-west-1.amazonaws.com": {
      "auth": "${AUTH}"
    },
    "426105708615.dkr.ecr.eu-west-1.amazonaws.com": {
      "auth": "${AUTH}"
    }
  }
}
EOF
