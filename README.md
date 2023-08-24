# Pip with multiple packages

Repository to test Cachi2 processing of multiple Pip packages. It is also used to validate Cachi2's reporting of missing hashes in requirements files.

## Included packages and dependencies

- first_pkg
  - requirements.in
    - pypi dependency (aiowsgi)
    - file dependency (appr)
- second_pkg
  - requirements.in
    - pypi dependency (aiowsgi, same as in first_pkg)
    - pypi dependency (charset-normalizer)
    - git dependency (appr)
  - requirements-extra.in
    - pypi dependency (charset-normalizer, same as in second_pkg/requirements.in)
- third_pkg
  - requirements.in
    - pypi dependency (click, also appears as a transitive dependency in second_pkg)

### Hash checking

- first_pkg/requirements.txt is resolved with hashes
- second_pkg/requirements.txt is resolved without hashes
- second_pkg/requirements-extra.txt is resolved with hashes
- third_pkg/requirements.txt is resolved without hashes

Cachi2 should report missing hashes in the SBOM accordingly. In case of repeated dependencies, it must report missing hashes **only** in the specific files in which they're missing. For example:

```
{
  "name": "aiowsgi",
  "version": "0.7",
  "purl": "pkg:pypi/aiowsgi@0.7",
  "properties": [
    {
      "name": "cachi2:missing_hash:in_file",
      "value": "second_package/requirements.txt"
    }
    // notice that the same dependency is also declared in first_package/requirements.txt,
    // but the hashes are present there
  ]
}
```
