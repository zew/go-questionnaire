# grep 
#    0  means found
#    1  means not found
# following should return 1
# see stackoverflow.com/questions/73833 for $'\r'
grep -U --quiet $'\r' ./deploy-new-release.sh ; echo $?

if grep -U --quiet $'\r' ./deploy-new-release.sh; then
    echo "contains windows newlines"
    exit
fi

echo 'no windows newlines found - continue'

chmod +x ./deploy-new-release.sh


FILE=/opt/go-questionnaire/go-questionnaire-new

if test -f "$FILE"; then

    echo "$FILE exists. doing reploy"

    sudo systemctl stop     go-questionnaire.service

    mv   --force /opt/go-questionnaire/go-questionnaire-new  /opt/go-questionnaire/go-questionnaire
    sudo setcap cap_net_bind_service=+eip /opt/go-questionnaire/go-questionnaire
    sudo chmod +x /opt/go-questionnaire/go-questionnaire


    sudo systemctl start    go-questionnaire.service
    sudo systemctl status   go-questionnaire.service   --no-pager

fi

