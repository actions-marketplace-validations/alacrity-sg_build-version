# Logic

1. If current branch is 'main', get tags. If tag does not contain major, minor or patch, then increment version to patch
2. If current branch is not 'main', create a rc tag based on the latest release tag with suffixes.
  - if latest release is 1.0.0, then the rc version should be 1.0.0-rc.${SOME-UNIQUE-ID}


## Tags
Tags might look something like
- v1.0.0
- v1.0.1-rc.1234
- v1.0.1

When parsing tags for (1)
- Get the latest qualifying RC tag that follows the format of vX.X.X-rc.X. return the first match
- With the qualifying version, remopve the RC labels. This will be the release version.
- Example: v0.0.1-rc.1234, release version should be v0.0.1
When parsing tags for (2),
- If a PR is open, get the labels from PR first.
- If there are no labels, the default increment is patch.
- To know what version to patch, parse through tags to find the latest qualifying version that follows the format v1.0.0.
- Apply patch and append unique rc suffix on top of the retrieved version
