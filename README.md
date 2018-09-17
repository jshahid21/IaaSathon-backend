# IaaSathon-backend
backend for OCI IaaSathon

# Prereqs
- the app is running on port 8000
- to run with Docker
```
$ docker run -d -p 8000:8000 -e ORACLE_USERNAME=c##<your db user> -e ORACLE_PASSWORD=<db password> -e ORACLE_SID=<SID of oracle db> schmidtp0740/iaasathon-backend
```
- the Oracle SID should be in this format
    ```
    <public ip address of the oracle database>:1521/<FQDN of the oracle database>
    ```
    one such example is
    ```
    129.213.40.22:1521/pollDB_iad159.sub06181613380.philiaasathonen.oraclevcn.com
    ```
# Configuring the database
- ssh into the instance
```
$ ssh opc@<public-ip-address>
```
- login as root
```
$ sudo su -
$ su - oracle
$ . oraenv
ORACLE_SID = [oracle] ? <db name>
$ srvctl config database -d <unique db name>
Database unique name: cdbm01 <<== DB unique name
Database name:
Oracle home: /u02/app/oracle/product/12.1.0/dbhome_2
Oracle user: oracle
Spfile: +DATAC1/cdbm01/spfilecdbm01.ora
Password file: +DATAC1/cdbm01/PASSWORD/passwd
Domain: data.customer1.oraclevcn.com
Start options: open
Stop options: immediate
Database role: PRIMARY
Management policy: AUTOMATIC
Server pools:
Disk Groups: DATAC1,RECOC1
Mount point paths:
Services:
Type: RAC
Start concurrency:
Stop concurrency:
OSDBA group: dba
OSOPER group: racoper
Database instances: cdbm011,cdbm012 <<== SID
Configured nodes: ed1db01,ed1db02
Database is administrator managed

$ sqlplus / as sysdba
 SQL> CREATE USER c##<your-name> IDENTIFIED BY <your-password>;
 SQL> GRANT CREATE SESSION TO c##<your-name>;
 SQL> GRANT DBA TO c##phil;
 SQL> exit
 $ sqlplus /nolog
 SQL> CONNECT c##<your-name>
 PASSWORD: <password>
 SQL> CREATE TABLE poll(cat number, dog number);
 SQL> INSERT INTO poll(cat, dog) VALUES (0,0);
 SQL> COMMIT;
 ```

