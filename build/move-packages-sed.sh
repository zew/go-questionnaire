# https://stackoverflow.com/questions/6758963/find-and-replace-with-sed-in-directory-and-sub-directories

# find files
find ./ -type f

find ./ -type f -exec ls -lAh {} \;

sed -i -e 's/apple/orange/g' main.go

#  -i edit in place, -e the script
find ./ -type f -exec sed -i -e 's/apple/orange/g' {} \;
find ./ -type f -exec sed -i -e 's|apple|orange|g' {} \;


# test on just one file: main.go
sed -i -e 's|"github.com/zew/go-questionnaire/pkg/handlers|"github.com/zew/go-questionnaire/internal/handlers|g' main.go



find ./ -type f -exec sed -i -e 's|"github.com/zew/go-questionnaire/pkg/handlers|"github.com/zew/go-questionnaire/internal/handlers|g' {} \;