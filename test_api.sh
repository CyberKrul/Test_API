# ==================================================

#              API TESTING SCRIPT

# ==================================================
# test_api.sh is a script to test the three operations of the API:
# 1. Device Registration
# 2. Updating the Mesh config
# 3. Device Information Retrieval
# 4. Unhappy paths for all three:
#   i. Malformed JSON
#   ii. Invalid Entries
#   iii. Specific cases

# Run with ./test_api.sh in GIT BASH only


# initializing the testing variables
SNO=98765432
FIRMWARE_VERSION=88
BASE_URL="http://127.0.0.1:8080"
NESNO=22345688 # device does not exist 
ISNO=89 # invalid number

echo "==================================================" 
echo " API Test Script"
echo " Target: $BASE_URL"
echo " Device S/N: $SNO"
echo "==================================================" 

echo 


# test 1: registering the device
echo "-----> Testing: POST/register"
echo "Happy Path: Register device w/ valid S/N: $SNO & valid Firmware Version: $FIRMWARE_VERSION..."
curl -s  -X POST $BASE_URL/register \
-H "Content-Type:application/json" \
-d '{"sno":'$SNO',"firmwareVersion":'$FIRMWARE_VERSION'}'
echo 

echo "Unhappy Path: Register device w/ invalid S/N: $ISNO & valid Firmware Version: $FIRMWARE_VERSION..."
curl -s  -X POST $BASE_URL/register \
-H "Content-Type:application/json" \
-d '{"sno":'$ISNO',"firmwareVersion":'$FIRMWARE_VERSION'}'
echo 

# test 2: updating the mesh config of a device
echo "-----> Testing: POST/update"
echo "Happy Path: Update device w/ valid S/N: $SNO..."
curl -s  -X POST $BASE_URL/update \
-H "Content-Type:application/json" \
-d '{"sno":'$SNO'}'
echo 

echo "Unhappy Path: Update device w/ invalid S/N: $ISNO..."
curl -s  -X POST $BASE_URL/update \
-H "Content-Type:application/json" \
-d '{"sno":'$ISNO'}'
echo 

echo "Unhappy Path: Update device w/ valid S/N not present in DB: $NESNO..."
curl -s  -X POST $BASE_URL/update \
-H "Content-Type:application/json" \
-d '{"sno":'$NESNO'}'
echo 

# test 3: retrieving the information of a device
echo "-----> Testing: POST/retrieve"
echo "Happy Path: Retrieve device w/ valid S/N: $SNO..."
curl -s  -X GET $BASE_URL/retrieve/$SNO 
echo 

echo "Unhappy Path: Retrieve device w/ invalid S/N: $ISNO..."
curl -s  -X GET $BASE_URL/retrieve/$SNO 
echo 

echo "Unhappy Path: Retrieve device w/ valid S/N not present in DB: $NESNO..."
curl -s  -X GET $BASE_URL/retrieve/$SNO 
echo 



echo "=================================================="
echo "          Test Script Finished"
echo "=================================================="
