goaml-api
-
still error on : 
-
ERROR: permission denied for sequence sessions_id_seq (SQLSTATE 42501)        
[10.363ms] [rows:0] INSERT INTO "sessions" ("username","customer_id","session_id") VALUES ('testuser1',1,'123456') RETURNING "id"
exit status 0xc000013a
-
[SOLVED] the cause of error is there's no grant to user and data in both tables (users & session)
still need another check if the result is appropriate
-
[FIX] output is match
-
[Update] Add handler if data not exist etc
