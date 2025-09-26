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
echo " Invalid Device S/N: $ISNO"
echo "==================================================" 

echo 
echo

# test 1: registering the device
echo "-----> Testing: POST BASEURL/devices"
echo "Happy Path: Register device w/ valid S/N: $SNO & valid Firmware Version: $FIRMWARE_VERSION..."
curl -s  -X POST $BASE_URL/devices \
-H "Content-Type:application/json" \
-d '{"sno":'$SNO',"firmwareVersion":'$FIRMWARE_VERSION'}'
echo 

echo "Unhappy Path: Register device w/ invalid S/N: $ISNO & valid Firmware Version: $FIRMWARE_VERSION..."
curl -s  -X POST $BASE_URL/devices \
-H "Content-Type:application/json" \
-d '{"sno":'$ISNO',"firmwareVersion":'$FIRMWARE_VERSION'}'
echo 
echo

# test 2: updating the mesh config of a device
echo "-----> Testing: PATCH BASEURL/devices/:sno"
echo "Happy Path: Patch device w/ valid S/N: $SNO..."
curl -s  -X PATCH $BASE_URL/devices/$SNO 
echo 

echo "Unhappy Path: Update device w/ invalid S/N: $ISNO..."
curl -s  -X PATCH $BASE_URL/devices/$ISNO 
echo 

echo "Unhappy Path: Update device w/ valid S/N not present in DB: $NESNO..."
curl -s  -X PATCH $BASE_URL/devices/$NESNO 
echo 
echo

# test 3: retrieving the information of a device
echo "-----> Testing: GET BASEURL/devices/:sno"
echo "Happy Path: device w/ valid S/N: $SNO..."
curl -s  -X GET $BASE_URL/devices/$SNO 
echo 

echo "Unhappy Path: device w/ invalid S/N: $ISNO..."
curl -s  -X GET $BASE_URL/devices/$ISNO 
echo 

echo "Unhappy Path: device w/ valid S/N not present in DB: $NESNO..."
curl -s  -X GET $BASE_URL/devices/$NESNO
echo 



echo "=================================================="
echo "          Test Script Finished"
echo "=================================================="
