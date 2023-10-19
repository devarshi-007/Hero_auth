# Concept
Many people have watch One Punchman. Concept of this repo is This contains website for hero and willain organization. Hero and willain can authorize with session or token.

# Frontend

# Backend

## Database
1. Backend is using postgres database
2. You can connect to pgadmin on your localhost port 8080, when you do docker-compose up
3. Connect to DB, password can be found in docker-compose file.
4. run make migrateup, to make a tables in database. (your current directory should be backend)
5. CSV files are given, import it in your tables

Note:- db_design.txt file contains postgres for making tables.

# Commit conventions

- feat: Add a new feature to the codebase (MINOR in semantic versioning).
- fix: Fix a bug (equivalent to a PATCH in Semantic Versioning).
- docs: Documentation changes.
- style: Code style change (semicolon, indentation...).
- refactor: Refactor code without changing public API.
- perf: Update code performances.
- test: Add test to an existing feature.
- chore: CHange Outside Runtime Environment - Update something without impacting the code