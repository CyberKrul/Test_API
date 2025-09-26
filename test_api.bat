@echo off
REM ==================================================
REM
REM              API TESTING SCRIPT
REM
REM ==================================================
REM test_api.bat is a script to test the three operations of the API:
REM 1. Device Registration
REM 2. Updating the Mesh config
REM 3. Device Information Retrieval
REM 4. Unhappy paths for all three:
REM   i. Malformed JSON
REM   ii. Invalid Entries
REM   iii. Specific cases
REM
REM Run with "test_api.bat" in a Windows Command Prompt (cmd.exe)


REM initializing the testing variables
set SNO=98765432
set FIRMWARE_VERSION=88
set BASE_URL=http://127.0.0.1:8080
set NESNO=12345678
set ISNO=89

echo ==================================================
echo  API Test Script
echo  Target: %BASE_URL%
echo  Device S/N: %SNO%
echo ==================================================
echo.


REM test 1: registering the device
echo -----> Testing: POST/register
echo Happy Path: Register device w/ valid S/N: %SNO% ^& valid Firmware Version: %FIRMWARE_VERSION%...
curl -X POST %BASE_URL%/register ^
-H "Content-Type:application/json" ^
-d "{\"sno\":%SNO%,\"firmwareVersion\":%FIRMWARE_VERSION%}"
echo.

echo Unhappy Path: Register device w/ invalid S/N: %ISNO% ^& valid Firmware Version: %FIRMWARE_VERSION%...
curl -X POST %BASE_URL%/register ^
-H "Content-Type:application/json" ^
-d "{\"sno\":%ISNO%,\"firmwareVersion\":%FIRMWARE_VERSION%}"
echo.

REM test 2: updating the mesh config of a device
echo -----> Testing: POST/update
echo Happy Path: Update device w/ valid S/N: %SNO%...
curl -X POST %BASE_URL%/update ^
-H "Content-Type:application/json" ^
-d "{\"sno\":%SNO%}"
echo.

echo Unhappy Path: Update device w/ invalid S/N: %ISNO%...
curl -X POST %BASE_URL%/update ^
-H "Content-Type:application/json" ^
-d "{\"sno\":%ISNO%}"
echo.

echo Unhappy Path: Update device w/ valid S/N not present in DB: %NESNO%...
curl -X POST %BASE_URL%/update ^
-H "Content-Type:application/json" ^
-d "{\"sno\":%NESNO%}"
echo.

REM test 3: retrieving the information of a device
echo -----> Testing: POST/retrieve
echo Happy Path: Retrieve device w/ valid S/N: %SNO%...
curl -X GET %BASE_URL%/retrieve/%SNO%
echo.

echo Unhappy Path: Retrieve device w/ invalid S/N: %ISNO%...
echo.

echo Unhappy Path: Retrieve device w/ valid S/N not present in DB: %NESNO%...
curl -X GET %BASE_URL%/retrieve/%NESNO%
echo.

echo ==================================================
echo           Test Script Finished
echo ==================================================
