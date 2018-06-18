# IaaSathon-backend
backend for OCI IaaSathon

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
 ```

