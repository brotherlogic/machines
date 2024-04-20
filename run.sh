#!/usr/bin/zsh
cd /home/simon/machines
git fetch -pq
git merge -q origin/main
if diff chead .git/refs/heads/main; then
else
    go build
    ./machines kclust1:31290 > hosts
  curl http://myexternalip.com/raw
  if ansible-playbook --vault-password-file=.vault_pass install.yml; then
    cp .git/refs/heads/main chead
  fi
fi
