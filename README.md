# Cachi2 RPMs .repo file test

This scenario covers the generation of a proper .repo file when the `inject-files` command is invoked after the prefetch. It contains three valid RPMs:
- two from different repo ids
- one missing a repo id, so that the repo id generation can be tested

The test steps are as follow:
- prefetch the dependencies with Cachi2
- inject the project files
- validate that the generated .repo file has the expected content
