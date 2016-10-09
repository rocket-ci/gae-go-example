#!/usr/bin/env bash
cat <<EOF >> app.yaml
env_variables:
  SLACK_TOKEN: ${SLACK_TOKEN}
EOF
aedeploy gcloud app deployv
