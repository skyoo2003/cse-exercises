#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

"""
Example of controlling temporary file and directory
"""

from tempfile import TemporaryFile, NamedTemporaryFile, TemporaryDirectory

# Example-1) Create a nonamed temp file. write a text and read it
with TemporaryFile('w+t') as f:
    f.write('Hello, World!')
    f.seek(0)
    print(f.read())

# Example-2) Create a named temp file, write a text and read it
with NamedTemporaryFile('w+t') as f:
    f.write("Hello, World!!!")
    f.seek(0)
    print(f.read())

# Example-3) Same as 'Example-2'. but disable the 'delete' option
with NamedTemporaryFile('w+t', delete=False) as f:
    f.write("Hello, World!!!")
    f.seek(0)
    print(f.read())

# Example-4) Create a named temp file in a specific location.
with NamedTemporaryFile('w+t', prefix="example", suffix=".txt", dir="/tmp") as f:
    print(f.name)

# Example-5) Create a nonamed temp directory
with TemporaryDirectory() as tmpdir:
    print(tmpdir)
