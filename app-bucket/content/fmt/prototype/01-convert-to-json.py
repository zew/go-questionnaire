from pathlib import Path
import openpyxl
import json

# Define paths
excelPath = Path("treatment.xlsx")
jsonPath = Path("treatment.json")

# Load the workbook and sheet
workbook = openpyxl.load_workbook(excelPath)
sheet = workbook["Sheet1"]

# Get all column headers
headers = []
for cell in sheet[1]:
    headers.append(cell.value)

# Find the index of 'person_id'
if "person_id" not in headers:
    raise ValueError("Column 'person_id' not found in header row.")

personIdIndex = headers.index("person_id")

# Prepare output dictionary
dataDict = {}

# Keys to round
keysToRound = {"consensus", "distance_to_consensus", "share_lower_Q42025"}
keysToDrop  = {
        
        "abgabedatum",
        "finished",
        "lang_code",
        
        "pprwbipq1",
        "pprwbipq2",
        "pprwbipq3",
        "pprwbipq4",
        
        "n_lower_Q42025",
        "Q42025_quartile",
        "var", 
        "var2",
 }

# Iterate through data rows (starting from row 2)
for row in sheet.iter_rows(min_row=2, values_only=True):
    personId = row[personIdIndex]
    if personId is None:
        print("Skipping row with missing person_id:", row)
        continue

    rowDict = {}
    for i in range(len(headers)):
        if i == personIdIndex:
            continue

        header = headers[i]
        value = row[i]

        if header in keysToRound and isinstance(value, (int, float)):
            value = round(value, 5)
        if header in keysToDrop:
            continue

        rowDict[header] = value

    dataDict[personId] = rowDict

# Write to JSON
with open(jsonPath, "w", encoding="utf-8") as jsonFile:
    json.dump(dataDict, jsonFile, indent=4, ensure_ascii=False)

print(f"JSON file created: {jsonPath.resolve()}")




