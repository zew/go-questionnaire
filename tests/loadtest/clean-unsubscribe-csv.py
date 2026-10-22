import sys
from pathlib import Path
from urllib.parse import parse_qs

def stackTrace(e):
    print(e)


inactiveProjects = [
    "unsubscribe", 
    "lix", 
    "pds",
    "muni",
    "muni",
]


def restoreObfuscation(rawStr):
    # restoring historical custom ascii obfuscation from go-massmail
    res = rawStr.replace("mmpp", "&")
    res = res.replace("qquu", "=")
    res = res.replace("ddtt", ".")
    res = res.replace("aatt", "@")
    res = res.replace("hhyy", "-")
    res = res.replace("pct40", "@")
    return res


def extractPathEmail(pthUrl):
    # splitting url path segments
    segs = pthUrl.split("/")
    if len(segs) > 1:
        # cutting trailing dummy segment
        segs = segs[:-1]

    if len(segs) > 3:
        lastSeg = restoreObfuscation(segs[-1])
        # checking for email marker in path segment
        if "@" in lastSeg:
            return lastSeg

    return ""


def checkOldRowGarbled(qry):
    if qry == "":
        return False, ""

    qryNorm = restoreObfuscation(qry)
    validKeys = {"project", "task", "email", "all"}
    qryDict = parse_qs(qryNorm, keep_blank_values=True)
    qryKeys = list(qryDict.keys())

    for idx1, keyName in enumerate(qryKeys):
        if keyName not in validKeys:
            vals = qryDict[keyName]
            valFirst = ""
            if len(vals) > 0:
                valFirst = vals[0]

            # matching cipher length signature for project=fmt (7=3 chars)
            if len(keyName) == 7 and len(valFirst) == 3:
                return True, f"garbled_project_param({keyName}={valFirst})"

            # matching cipher length signature for email param (5 chars)
            if len(keyName) == 5 and "@" in valFirst:
                return True, f"garbled_email_param({keyName}={valFirst})"

            # matching cipher length signature for task param (4 chars)
            if len(keyName) == 4:
                return True, f"garbled_task_param({keyName}={valFirst})"

            # catching any other scrambled query parameter name
            return True, f"unknown_query_param({keyName})"

    return False, ""


def checkNewRowGarbled(emailPath, qry):
    if qry == "":
        return False, ""

    qryNorm = restoreObfuscation(qry)
    validKeys = {"project", "task", "email", "all"}
    expectedKeys = {"project", "task", "email"}
    qryDict = parse_qs(qryNorm, keep_blank_values=True)
    qryKeys = list(qryDict.keys())

    # checking for scrambled parameter keys
    for idx1, keyName in enumerate(qryKeys):
        if keyName not in validKeys:
            return True, f"scrambled_query_key({keyName})"

    # verifying presence of required keys in non-empty query
    for idx1, reqKey in enumerate(expectedKeys):
        if reqKey not in qryDict:
            return True, f"missing_expected_key({reqKey})"

    emailVals = qryDict.get("email", [""])
    emailQuery = ""
    if len(emailVals) > 0:
        emailQuery = restoreObfuscation(emailVals[0])

    # comparing query email against path email
    if emailQuery != emailPath:
        return True, f"email_mismatch(path={emailPath} vs query={emailQuery})"

    return False, ""


def cleanUnsubscribeCsv(pthIn, pthOut):
    try:
        if not pthIn.exists():
            print(f"input file not found: {pthIn}")
            sys.exit(-1)

        rawText = pthIn.read_text(encoding="utf-8")
        lines = rawText.splitlines()

        keptLines = []
        removedTplList = []
        cntTotal = 0
        cntInactive = 0

        for idx1, lineStr in enumerate(lines):
            # preserving header line
            if idx1 == 0 or lineStr.startswith("project;task;"):
                keptLines.append(lineStr)
                continue

            if lineStr.strip() == "":
                continue

            cntTotal += 1
            tpl = lineStr.split(";")
            if len(tpl) < 6:
                keptLines.append(lineStr)
                continue

            proj = tpl[0]
            tsk = tpl[1]
            emailCsv = tpl[2]
            dtStr = tpl[3]
            pthUrl = tpl[4]
            qry = tpl[5]

            # discarding records from inactive projects
            if proj in inactiveProjects:
                cntInactive += 1
                continue

            emailPath = extractPathEmail(pthUrl)
            emailDisplay = restoreObfuscation(emailCsv)

            isGarbled = False
            reasonStr = ""

            if emailPath != "":
                # evaluating row with email in url path
                isGarbled, reasonStr = checkNewRowGarbled(emailPath, qry)
            else:
                # evaluating legacy row without email in url path
                isGarbled, reasonStr = checkOldRowGarbled(qry)

            if isGarbled:
                removedTplList.append((idx1 + 1, dtStr, proj, tsk, emailDisplay, pthUrl, qry, reasonStr))
            else:
                keptLines.append(lineStr)

        # writing cleaned output file
        outContent = "\n".join(keptLines) + "\n"
        pthOut.write_text(outContent, encoding="utf-8")

        # printing summary and removal report for colleagues
        print(f"total data rows processed : {cntTotal}")
        print(f"rows kept                 : {len(keptLines) - 1}")
        print(f"rows removed (inactive)   : {cntInactive} ({', '.join(inactiveProjects)})")
        print(f"rows removed (scanners)   : {len(removedTplList)}")
        print("-" * 110)
        print(f"{'Line':<6} | {'Date':<19} | {'Proj':<6} | {'Task':<12} | {'Wrongly Unsubscribed Email':<32} | {'Reason'}")
        print("-" * 110)

        for idx1, remTpl in enumerate(removedTplList):
            lineNo, dtStr, proj, tsk, emailDisplay, pthUrl, qry, reasonStr = remTpl
            print(f"{lineNo:<6} | {dtStr:<19} | {proj:<6} | {tsk:<12} | {emailDisplay:<32} | {reasonStr}")

    except Exception as exc:
        stackTrace(exc)
        print("error cleaning unsubscribe csv")
        sys.exit(-1)


if __name__ == "__main__":
    pthSource = Path("unsubscribe.csv")
    pthCleaned = Path("unsubscribe_cleaned.csv")
    if len(sys.argv) > 1:
        pthSource = Path(sys.argv[1])
    if len(sys.argv) > 2:
        pthCleaned = Path(sys.argv[2])

    cleanUnsubscribeCsv(pthSource, pthCleaned)