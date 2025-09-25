-- deviceDBTwo is the table to be used for this project
CREATE TABLE deviceDBTwo (
    sno INT PRIMARY KEY,
    firmware_version INT,
    current_firmware_version BOOL,
    mesh_configuration BOOL,
    app_configuration BOOL,
    kc_configuration BOOL
)

SELECT * FROM deviceDBTwo