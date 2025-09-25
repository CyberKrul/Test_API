package service

import "strconv"

// validateSno checks if the serial number is exactly 8 digits.
func validateSno(sno int) error {
	if len(strconv.Itoa(sno)) != 8 {
		return ErrInvalidSno
	}
	return nil
}

// validateFirmwareVersion checks if the firmware version is within a valid range.
func validateFirmwareVersion(fv int) error {
	if fv <= 0 || fv > 100 {
		return ErrInvalidFV
	}
	return nil
}
