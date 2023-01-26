#!/usr/bin/zsh
cd /home/simon/machines
git fetch -pq
git merge origin/main
if diff chead .git/refs/heads/main; then
else
  if ansible-playbook --vault-password-file=.vault_pass install.yml; then
    cp .git/refs/heads/main chead
  fi
fi