# Logic

This document outlines the processing logic.

## Main/Release Version Logic

1. If current branch is 'main', get the latest RC tag.
2. Trim the suffix of RC tag and set it as the official version

When parsing tags
- Get the latest qualifying RC tag that follows the format of vX.X.X-rc.X. return the first match
- With the qualifying version, remopve the RC labels. This will be the release version.
- Example: v0.0.1-rc.1234, release version should be v0.0.1


## Feature/RC Version Logic
1. if current branch is not 'main', get the latest release tag
2. Check for input increment type.
3. If there is no input increment type, attempt to check with Git service for increment type by checking if labels hae a 'major', 'minor' or 'patch' label. If labels are found, get the label with the highest gravity.
4. If input increment and labels cannot be found, then set increment type as 'patch'
5. Generate a RC tag using known unique variables for respective git service. Increment the version according to the increment type

When parsing tags
- If a PR is open, get the labels from PR first.
- If there are no labels, the default increment is patch.
- To know what version to patch, parse through tags to find the latest qualifying version that follows the format v1.0.0.
- Apply patch and append unique rc suffix on top of the retrieved version

## Tags
Tags might look something like
- v1.0.0
- v1.0.1-rc.1234
- v1.0.1
