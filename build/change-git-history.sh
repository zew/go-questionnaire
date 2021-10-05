# https://www.git-tower.com/learn/git/faq/change-author-name-email
git filter-branch --env-filter '
WRONG_EMAIL="Peter Buchmann <PBU@zew.local>"
WRONG_EMAIL="PBU@zew.local"
NEW_NAME="Peter Buchmann"
NEW_EMAIL="peter.buchmann@zew.de"

if [ "$GIT_COMMITTER_EMAIL" = "$WRONG_EMAIL" ]
then
    export GIT_COMMITTER_NAME="$NEW_NAME"
    export GIT_COMMITTER_EMAIL="$NEW_EMAIL"
fi
if [ "$GIT_AUTHOR_EMAIL" = "$WRONG_EMAIL" ]
then
    export GIT_AUTHOR_NAME="$NEW_NAME"
    export GIT_AUTHOR_EMAIL="$NEW_EMAIL"
fi
' --tag-name-filter cat -- --branches --tags


git filter-branch --force --tree-filter 'rm -rf   ./responses/*.json'     HEAD
