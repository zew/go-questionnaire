# https://stackoverflow.com/questions/6758963/find-and-replace-with-sed-in-directory-and-sub-directories

# find files in entire subtree
find ./ -type f

# find files in entire subtree and execute command on it
find ./ -type f -exec ls -lAh {} \;

# replacing
sed -i -e 's/apple/orange/g' main.go

# replacing in entire subree
#  -i edit in place, -e the script
find ./ -type f -exec sed -i -e 's/apple/orange/g' {} \;
find ./ -type f -exec sed -i -e 's|apple|orange|g' {} \;

# test replacement on just one file: main.go
sed -i -e 's|"github.com/zew/go-questionnaire/pkg/handlers|"github.com/zew/go-questionnaire/internal/handlers|g' main.go

# move single package from pkg to internal
find ./ -type f -exec sed -i -e 's|"github.com/zew/go-questionnaire/pkg/handlers|"github.com/zew/go-questionnaire/internal/handlers|g' {} \;

# all project packages from app root into dir pkg
find ./ -type f -exec sed -i -e 's|"github.com/zew/go-questionnaire/|"github.com/zew/go-questionnaire/pkg/|g' {} \;

# all project packages from dir pkg  into dir internal
find ./ -type f -exec sed -i -e 's|"github.com/zew/go-questionnaire/pkg/|"github.com/zew/go-questionnaire/internal/|g' {} \;