set -a
. .env
set +a
sshpass -f pass_file rsync \
  --exclude='.git' \
  --exclude='.gitignore' \
  --exclude='pass_file' \
  -avz -e ssh . ${HOSTNAME}@${HOSTADDR}:${HOSTPATH}
