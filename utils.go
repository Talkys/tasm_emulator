package main

/*Convert uint32 to 4 uint8*/
func ui32to4ui8(num uint32) [4]uint8 {
	var output [4]uint8

	output[0] = uint8(num >> THREE_BYTES)
	output[1] = uint8((num >> TWO_BYTES) & ONE_BYTE_MASK)
	output[2] = uint8((num >> ONE_BYTE) & ONE_BYTE_MASK)
	output[3] = uint8(num & ONE_BYTE_MASK)

	return output
}

func ui32fr4ui8(num [4]uint8) uint32 {
	output := uint32(num[0]) << THREE_BYTES
	output += uint32(num[1]) << TWO_BYTES
	output += uint32(num[2]) << ONE_BYTE
	output += uint32(num[3])

	return output
}
