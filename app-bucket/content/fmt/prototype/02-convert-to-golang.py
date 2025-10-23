from pathlib import Path

jsonPath = Path("treatment.json")
goSnippetPath = Path("treatment_go_map_snippet.txt")

# --------------------------------------------
# Re-open JSON as text and adapt for Go map literal
# - Drop the outermost braces
# - Add trailing commas to the final key-value line inside each inner object
# - Leave everything else unchanged (minimal, textual transformation)
# --------------------------------------------
with open(jsonPath, "r", encoding="utf-8") as f:
    jsonText = f.read()

lines = jsonText.splitlines()

# Remove the first line that contains only the opening '{' and the last line that contains only the closing '}'
startIndex = 0
while startIndex < len(lines) and lines[startIndex].strip() == "":
    startIndex += 1

endIndex = len(lines) - 1
while endIndex >= 0 and lines[endIndex].strip() == "":
    endIndex -= 1

if startIndex <= endIndex and lines[startIndex].strip() == "{":
    startIndex += 1
else:
    print("Warning: JSON does not start with '{' on its own line as expected.")

if endIndex >= 0 and lines[endIndex].strip() == "}":
    endIndex -= 1
else:
    print("Warning: JSON does not end with '}' on its own line as expected.")

innerLines = []
for i in range(startIndex, endIndex + 1):
    innerLines.append(lines[i])

# Add a trailing comma to the last key-value line inside each inner object:
# Detect a block ending with '}' or '},' and ensure the previous non-empty line ends with a comma.
processedLines = []
for idx in range(len(innerLines)):
    line = innerLines[idx]
    stripped = line.strip()

    if stripped == "}" or stripped == "},":
        j = len(processedLines) - 1
        while j >= 0 and processedLines[j].strip() == "":
            j -= 1

        if j >= 0:
            prevLine = processedLines[j].rstrip()
            if not prevLine.endswith(","):
                processedLines[j] = prevLine + ","
        processedLines.append(line)
    else:
        processedLines.append(line)

# Optionally wrap into a ready-to-paste Go snippet
# Note: This does NOT convert inner objects to map[string]interface{} literals.
# It performs only the minimal textual adjustments requested.
goSnippetLines = []
goSnippetLines.append('var forecastData = map[string]map[string]interface{}{\n')
for line in processedLines:
    line = line.replace(": null,", ": nil,")    
    goSnippetLines.append(line + "\n")
goSnippetLines.append('}\n')

with open(goSnippetPath, "w", encoding="utf-8") as g:
    for line in goSnippetLines:
        g.write(line)

print(f"Go map snippet written to: {goSnippetPath.resolve()}")

# If you want the bare inner body (without the var/outer braces), you can also print it:
print("----- BEGIN Go map inner body -----")
for line in processedLines:
    print(line)
print("----- END Go map inner body -----")