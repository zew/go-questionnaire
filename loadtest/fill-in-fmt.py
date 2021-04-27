import threading
import time
import requests
import sys



# 'dev.zew.de' should point to your development server - i.e. localhost
baseURL = "https://survey2.zew.de/"

loginQueries = [
    "?u=10000&sid=fmt&wid=2021-05&p=1&h=5y60OnjnS7oTrM6G9qhuz8Nv6R7oFecvJ_srMNqL8rg",
    "?u=10001&sid=fmt&wid=2021-05&p=1&h=6ddRSTs0PSYCy_7SqMXybjadhhVawsV1YQnJrbK-iCU",
    "?u=10002&sid=fmt&wid=2021-05&p=1&h=EO50-rRdu8aA8Xw-FxiUEPAgDrfiJi-esb3v9MwcMTg",
    "?u=10003&sid=fmt&wid=2021-05&p=1&h=9nHtDfUXv-HJwb2nGAJAq9jSh4qx-Dd7tgiUjUCrmSs",
]


pagesToFields = {
    0 : [
        "selbst",
        "contact",
    ],

    1 : [

        "y0_ez",
        "y0_deu",
        "y0_usa",
        "y0_chn",

        "y_ez",
        "y_deu",
        "y_usa",
        "y_chn",

        {"y_probgood":      33},
        {"y_probnormal":    34},
        {"y_probbad":       33},
        {"y_recession_q0":  10},
        {"y_recession_q1":  15},
    ],

    2: [

        "pi_ez",
        "pi_deu",
        "pi_usa",
        "pi_chn",

        "i_deu",
        "i_usa",
        "i_chn",

        "r_deu",
        "r_usa",
        "r_chn",
    ],

    3: [

        "sto_ez",
        "sto_dax",
        "sto_usa",
        "sto_sse_comp_chn",

        {"dax_erw": 14500},
        {"dax_min": 11000},
        {"dax_max": 17000},

        "dax_fund",
    ],

    4: [
        "fx_usa",
        "fx_chn",
    ],

    5: [
        "sec_banks",
        "sec_insur",
        "sec_cars",
        "sec_chemi",
        "sec_steel",
        "sec_elect",
        "sec_mecha",
        "sec_consu",
        "sec_const",
        "sec_utili",
        "sec_servi",
        "sec_telec",
        "sec_infor",
    ],

}




def requestToken(responseText):
    # Extract form request token from response
    loc1 = 'type="hidden" name="token" value="'
    if loc1 not in responseText:
        print("Request token at location '%s' not found" % loc1)
        return
    pos1 = responseText.find(loc1)
    pos2 = responseText.find("\"", len(loc1)+pos1+1)
    tokn = responseText[pos1+len(loc1):pos2]
    # print("Pos1 %d - pos2 %d. Request token is %s" % (pos1, pos2, tokn))
    # print(responseText[pos1+len(loc1)-2:pos2+2])
    if len(tokn) > 65 or len(tokn) < 63:
        # ffdcac226db52edeb447149299f01ea96c7a1fbeead51f168653bc0994335dd4
        raise Exception("token seems fish %s" % tokn)
    return tokn



def printResponseData(resp, testString):
    print("\nHeaders %s" % resp.headers)
    print("Url %s - Status Code %d" % (resp.url, resp.status_code))
    if testString not in resp.text:
        # print(r.text)
        print("MISSING %s" % testString)
        sys.exit()
    else:
        print("FOUND %s" % testString)

    print("'%s' fetched in %ss" % (resp.url, (time.time() - start)))






def fillQuestionnaire(userID,queryStr,value):

    print("start  user %s - value %s" % (str(userID), str(value)))

    s = requests.session()
    loginURL = baseURL + queryStr
    resp = s.get(loginURL, params={}, allow_redirects=True)
    if resp.status_code != 200:
        print("status code %d for %s" % (resp.status_code, loginURL))
        return

    tokn = requestToken(resp.text)
    dictParams = {}
    dictParams["token"] = tokn
    dictParams["submitBtn"] = 1
    resp = s.post(baseURL, params=dictParams, allow_redirects=True)
    if resp.status_code != 200:
        print("status code %d for %s with %s" % (resp.status_code, baseURL, dictParams))
        return
    else:
        print("successfully proceeded to page 1")


    for idx1, pageIdx in enumerate(pagesToFields):

        fields = pagesToFields[pageIdx]
        if pageIdx == 0:
            continue

        dictParams = {}
        tokn = requestToken(resp.text)
        dictParams["token"] = tokn
        dictParams["submitBtn"] = "next"
        for idx2, field in enumerate(fields):
            if type(field) is str:
                dictParams[field] = value
            if type(field) is dict:
                for key in field:
                    dictParams[key] = field[key]
        resp = s.post(baseURL, params=dictParams, allow_redirects=True)
        if resp.status_code != 200:
            print("status code %d for loop %d - %s with %s" % (resp.status_code, pageIdx, baseURL, dictParams))
            return
        else:
            print("successfully proceeded to page %d" % (pageIdx+1))

    print("finish user %s" % str(userID))
    print(" ")







start = time.time()

for idx, queryStr in enumerate(loginQueries):
    # 0 => 1
    # 1 => 1
    # 2 => 2
    # 3 => 4
    value = idx
    if idx == 0:
        value = 1
    if idx == 3:
        value = 4
    fillQuestionnaire(idx+1, queryStr, value)

print("\nElapsed Time: %s" % (time.time() - start))




