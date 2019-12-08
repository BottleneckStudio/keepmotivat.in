#!/bin/bash
echo ''

# Create the file inside .git/pre-commit
cat ./scripts/pre-commit.sh >> .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit

# Init message
echo $(date) '[ Git Hooks ] Symlinking commit hook scripts...'

# Check if file already symlinked
if [ -h .git/hooks/pre-commit ]; then
  echo $(date) '[ Git Hooks ] Commit hook scripts already symlinked'
else
  ln -svf $(pwd)/script/pre-commit.sh .git/hooks/pre-commit
  echo $(date) '[ Git Hooks ] Already symlinked. DONE!'
fi

