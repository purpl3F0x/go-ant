/*
 * constants.go
 *
 * Copyright (c) 2021-2022 Stavros Avramidis (@purpl3F0x). All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 */

package ant

// TODO: add to ant+ package

func AntPlusNetworkKey() [8]uint8 {
	return [8]uint8{0xB9, 0xA5, 0x21, 0xFB, 0xBD, 0x72, 0xC3, 0x45}
}

const (
	// ////////////////////////////////////////////
	// ANT Message Packet Size
	// ////////////////////////////////////////////

	ANT_STANDARD_DATA_PAYLOAD_SIZE uint8 = 8

	//////////////////////////////////////////////
	// ANT LIBRARY Extended Data Message Fields
	// NOTE: You must check the extended message
	// bitfield first to find out which fields
	// are present before accessing them!
	//////////////////////////////////////////////

	ANT_EXT_MESG_DEVICE_ID_FIELD_SIZE uint8 = 4
	ANT_EXT_STRING_SIZE               uint8 = 27 // increase buffer size to support longer messages (32 extra bytes after ANT standard payload)

	//////////////////////////////////////////////
	// ANT Extended Data Message Bifield Definitions
	//////////////////////////////////////////////

	ANT_EXT_MESG_BITFIELD_DEVICE_ID uint8 = 0x80 // first field after bitfield

	// 4 bits free reserved set to 0
	ANT_EXT_MESG_BIFIELD_EXTENSION uint8 = 0x01

	// extended message input bitfield defines
	ANT_EXT_MESG_BITFIELD_OVERWRITE_SHARED_ADR uint8 = 0x10
	ANT_EXT_MESG_BITFIELD_TRANSMISSION_TYPE    uint8 = 0x08

	//////////////////////////////////////////////
	// ANT Library Config
	//////////////////////////////////////////////

	ANT_LIB_CONFIG_MASK_ALL                uint8 = 0xFF
	ANT_LIB_CONFIG_RADIO_CONFIG_ALWAYS     uint8 = 0x01
	ANT_LIB_CONFIG_MESG_OUT_INC_TIME_STAMP uint8 = 0x20
	ANT_LIB_CONFIG_MESG_OUT_INC_RSSI       uint8 = 0x40
	ANT_LIB_CONFIG_MESG_OUT_INC_DEVICE_ID  uint8 = 0x80

	//////////////////////////////////////////////
	// ID Definitions
	//////////////////////////////////////////////

	ANT_ID_SIZE                      uint8 = 4
	ANT_ID_TRANS_TYPE_OFFSET         uint8 = 3
	ANT_ID_DEVICE_TYPE_OFFSET        uint8 = 2
	ANT_ID_DEVICE_NUMBER_HIGH_OFFSET uint8 = 1
	ANT_ID_DEVICE_NUMBER_LOW_OFFSET  uint8 = 0
	ANT_ID_DEVICE_TYPE_PAIRING_FLAG  uint8 = 0x80

	ANT_TRANS_TYPE_SHARED_ADDR_MASK      uint8 = 0x03
	ANT_TRANS_TYPE_1_BYTE_SHARED_ADDRESS uint8 = 0x02
	ANT_TRANS_TYPE_2_BYTE_SHARED_ADDRESS uint8 = 0x03

	//////////////////////////////////////////////
	// Assign Channel Parameters
	//////////////////////////////////////////////

	PARAMETER_RX_NOT_TX                     uint8 = 0x00
	PARAMETER_TX_NOT_RX                     uint8 = 0x10
	PARAMETER_SHARED_CHANNEL                uint8 = 0x20
	PARAMETER_NO_TX_GUARD_BAND              uint8 = 0x40
	PARAMETER_ALWAYS_RX_WILD_CARD_SEARCH_ID uint8 = 0x40 //Pre-AP2
	PARAMETER_RX_ONLY                       uint8 = 0x40

	//////////////////////////////////////////////
	// Ext. Assign Channel Parameters
	//////////////////////////////////////////////

	EXT_PARAM_ALWAYS_SEARCH     uint8 = 0x01
	EXT_PARAM_FREQUENCY_AGILITY uint8 = 0x04

	//////////////////////////////////////////////
	// Radio TX Power Definitions
	//////////////////////////////////////////////

	RADIO_TX_POWER_LVL_MASK uint8 = 0x03

	RADIO_TX_POWER_LVL_0 uint8 = 0x00 //(formerly: RADIO_TX_POWER_MINUS20DB) lowest
	RADIO_TX_POWER_LVL_1 uint8 = 0x01 //(formerly: RADIO_TX_POWER_MINUS10DB)
	RADIO_TX_POWER_LVL_2 uint8 = 0x02 //(formerly: RADIO_TX_POWER_MINUS5DB)
	RADIO_TX_POWER_LVL_3 uint8 = 0x03 //(formerly: RADIO_TX_POWER_0DB) highest

	//////////////////////////////////////////////
	// Channel Status
	//////////////////////////////////////////////

	STATUS_CHANNEL_STATE_MASK uint8 = 0x03
	STATUS_UNASSIGNED_CHANNEL uint8 = 0x00
	STATUS_ASSIGNED_CHANNEL   uint8 = 0x01
	STATUS_SEARCHING_CHANNEL  uint8 = 0x02
	STATUS_TRACKING_CHANNEL   uint8 = 0x03

	//////////////////////////////////////////////
	// Standard capabilities defines
	//////////////////////////////////////////////

	CAPABILITIES_NO_RX_CHANNELS    uint8 = 0x01
	CAPABILITIES_NO_TX_CHANNELS    uint8 = 0x02
	CAPABILITIES_NO_RX_MESSAGES    uint8 = 0x04
	CAPABILITIES_NO_TX_MESSAGES    uint8 = 0x08
	CAPABILITIES_NO_ACKD_MESSAGES  uint8 = 0x10
	CAPABILITIES_NO_BURST_TRANSFER uint8 = 0x20

	//////////////////////////////////////////////
	// Advanced capabilities defines
	//////////////////////////////////////////////

	CAPABILITIES_OVERUN_UNDERRUN              uint8 = 0x01 // Support for this functionality has been dropped
	CAPABILITIES_NETWORK_ENABLED              uint8 = 0x02
	CAPABILITIES_AP1_VERSION_2                uint8 = 0x04 // This Version of the AP1 does not support transmit and only had a limited release
	CAPABILITIES_SERIAL_NUMBER_ENABLED        uint8 = 0x08
	CAPABILITIES_PER_CHANNEL_TX_POWER_ENABLED uint8 = 0x10
	CAPABILITIES_LOW_PRIORITY_SEARCH_ENABLED  uint8 = 0x20
	CAPABILITIES_SCRIPT_ENABLED               uint8 = 0x40
	CAPABILITIES_SEARCH_LIST_ENABLED          uint8 = 0x80

	//////////////////////////////////////////////
	// Advanced capabilities 2 defines
	//////////////////////////////////////////////

	CAPABILITIES_LED_ENABLED         uint8 = 0x01
	CAPABILITIES_EXT_MESSAGE_ENABLED uint8 = 0x02
	CAPABILITIES_SCAN_MODE_ENABLED   uint8 = 0x04
	CAPABILITIES_RESERVED            uint8 = 0x08
	CAPABILITIES_PROX_SEARCH_ENABLED uint8 = 0x10
	CAPABILITIES_EXT_ASSIGN_ENABLED  uint8 = 0x20
	CAPABILITIES_FS_ANTFS_ENABLED    uint8 = 0x40
	CAPABILITIES_FIT1_ENABLED        uint8 = 0x80

	//////////////////////////////////////////////
	// Advanced capabilities 3 defines
	//////////////////////////////////////////////

	CAPABILITIES_ADVANCED_BURST_ENABLED             uint8 = 0x01
	CAPABILITIES_EVENT_BUFFERING_ENABLED            uint8 = 0x02
	CAPABILITIES_EVENT_FILTERING_ENABLED            uint8 = 0x04
	CAPABILITIES_HIGH_DUTY_SEARCH_MODE_ENABLED      uint8 = 0x08
	CAPABILITIES_ACTIVE_SEARCH_SHARING_MODE_ENABLED uint8 = 0x10
	CAPABILITIES_SELECTIVE_DATA_UPDATE_ENABLED      uint8 = 0x40
	CAPABILITIES_ENCRYPTED_CHANNEL_ENABLED          uint8 = 0x80

	//////////////////////////////////////////////
	// Burst Message Sequence
	//////////////////////////////////////////////

	CHANNEL_NUMBER_MASK      uint8 = 0x1F
	SEQUENCE_NUMBER_MASK     uint8 = 0xE0
	SEQUENCE_NUMBER_ROLLOVER uint8 = 0x60
	SEQUENCE_FIRST_MESSAGE   uint8 = 0x00
	SEQUENCE_LAST_MESSAGE    uint8 = 0x80
	SEQUENCE_NUMBER_INC      uint8 = 0x20

	//////////////////////////////////////////////
	// Advanced Burst Config defines
	//////////////////////////////////////////////

	ADV_BURST_CONFIG_FREQ_HOP uint32 = 0x00000001

	//////////////////////////////////////////////
	// Extended Message ID Mask
	//////////////////////////////////////////////

	MSG_EXT_ID_MASK uint8 = 0xE0

	//////////////////////////////////////////////
	// Control Message Flags
	//////////////////////////////////////////////

	BROADCAST_CONTROL_BYTE    uint8 = 0x00
	ACKNOWLEDGED_CONTROL_BYTE uint8 = 0xA0

	//////////////////////////////////////////////
	// Response / Event Codes
	//////////////////////////////////////////////

	RESPONSE_NO_ERROR uint8 = 0x00
	NO_EVENT          uint8 = 0x00

	EVENT_RX_SEARCH_TIMEOUT     uint8 = 0x01
	EVENT_RX_FAIL               uint8 = 0x02
	EVENT_TX                    uint8 = 0x03
	EVENT_TRANSFER_RX_FAILED    uint8 = 0x04
	EVENT_TRANSFER_TX_COMPLETED uint8 = 0x05
	EVENT_TRANSFER_TX_FAILED    uint8 = 0x06
	EVENT_CHANNEL_CLOSED        uint8 = 0x07
	EVENT_RX_FAIL_GO_TO_SEARCH  uint8 = 0x08
	EVENT_CHANNEL_COLLISION     uint8 = 0x09
	EVENT_TRANSFER_TX_START     uint8 = 0x0A // a pending transmit transfer has begun

	EVENT_CHANNEL_ACTIVE uint8 = 0x0F

	EVENT_TRANSFER_TX_NEXT_MESSAGE uint8 = 0x11 // only enabled in FIT1

	CHANNEL_IN_WRONG_STATE uint8 = 0x15 // returned on attempt to perform an action from the wrong channel state
	CHANNEL_NOT_OPENED     uint8 = 0x16 // returned on attempt to communicate on a channel that is not open
	CHANNEL_ID_NOT_SET     uint8 = 0x18 // returned on attempt to open a channel without setting the channel ID
	CLOSE_ALL_CHANNELS     uint8 = 0x19 // returned when attempting to start scanning mode, when channels are still open

	TRANSFER_IN_PROGRESS           uint8 = 0x1F // returned on attempt to communicate on a channel with a TX transfer in progress
	TRANSFER_SEQUENCE_NUMBER_ERROR uint8 = 0x20 // returned when sequence number is out of order on a Burst transfer
	TRANSFER_IN_ERROR              uint8 = 0x21
	TRANSFER_BUSY                  uint8 = 0x22

	INVALID_MESSAGE_CRC        uint8 = 0x26 // returned if there is a framing error on an incomming message
	MESSAGE_SIZE_EXCEEDS_LIMIT uint8 = 0x27 // returned if a data message is provided that is too large
	INVALID_MESSAGE            uint8 = 0x28 // returned when the message has an invalid parameter
	INVALID_NETWORK_NUMBER     uint8 = 0x29 // returned when an invalid network number is provided
	INVALID_LIST_ID            uint8 = 0x30 // returned when the provided list ID or size exceeds the limit
	INVALID_SCAN_TX_CHANNEL    uint8 = 0x31 // returned when attempting to transmit on channel 0 when in scan mode
	INVALID_PARAMETER_PROVIDED uint8 = 0x33 // returned when an invalid parameter is specified in a configuration message

	EVENT_SERIAL_QUE_OVERFLOW uint8 = 0x34
	EVENT_QUE_OVERFLOW        uint8 = 0x35 // ANT event que has overflowed and lost 1 or more events

	EVENT_CLK_ERROR     uint8 = 0x36 // debug event for XOSC16M on LE1
	EVENT_STATE_OVERRUN uint8 = 0x37

	EVENT_ENCRYPT_NEGOTIATION_SUCCESS uint8 = 0x38 // ANT stack generated event when connecting to an encrypted channel has succeeded
	EVENT_ENCRYPT_NEGOTIATION_FAIL    uint8 = 0x39 // ANT stack generated event when connecting to an encrypted channel has failed

	SCRIPT_FULL_ERROR         uint8 = 0x40 // error writing to script, memory is full
	SCRIPT_WRITE_ERROR        uint8 = 0x41 // error writing to script, bytes not written correctly
	SCRIPT_INVALID_PAGE_ERROR uint8 = 0x42 // error accessing script page
	SCRIPT_LOCKED_ERROR       uint8 = 0x43 // the scripts are locked and can't be dumped

	NO_RESPONSE_MESSAGE uint8 = 0x50 // returned to the Command_SerialMessageProcess function, so no reply message is generated
	RETURN_TO_MFG       uint8 = 0x51 // default return to any mesg when the module determines that the mfg procedure has not been fully completed

	FIT_ACTIVE_SEARCH_TIMEOUT uint8 = 0x60 // Fit1 only event added for timeout of the pairing state after the Fit module becomes active
	FIT_WATCH_PAIR            uint8 = 0x61 // Fit1 only
	FIT_WATCH_UNPAIR          uint8 = 0x62 // Fit1 only

	USB_STRING_WRITE_FAIL uint8 = 0x70

	// Internal only events below this point
	INTERNAL_ONLY_EVENTS uint8 = 0x80
	EVENT_RX             uint8 = 0x80 // INTERNAL: Event for a receive message
	EVENT_NEW_CHANNEL    uint8 = 0x81 // INTERNAL: EVENT for a new active channel
	EVENT_PASS_THRU      uint8 = 0x82 // INTERNAL: Event to allow an upper stack events to pass through lower stacks

	EVENT_BLOCKED uint8 = 0xFF // INTERNAL: Event to replace any event we do not wish to go out, will also zero the size of the Tx message

	///////////////////////////////////////////////////////
	// Script Command Codes
	///////////////////////////////////////////////////////

	SCRIPT_CMD_FORMAT             uint8 = 0x00
	SCRIPT_CMD_DUMP               uint8 = 0x01
	SCRIPT_CMD_SET_DEFAULT_SECTOR uint8 = 0x02
	SCRIPT_CMD_END_SECTOR         uint8 = 0x03
	SCRIPT_CMD_END_DUMP           uint8 = 0x04
	SCRIPT_CMD_LOCK               uint8 = 0x05

	///////////////////////////////////////////////////////
	// USB Descriptor Strings
	///////////////////////////////////////////////////////

	USB_DESCRIPTOR_VID_PID             uint8 = 0x00
	USB_DESCRIPTOR_MANUFACTURER_STRING uint8 = 0x01
	USB_DESCRIPTOR_DEVICE_STRING       uint8 = 0x02
	USB_DESCRIPTOR_SERIAL_STRING       uint8 = 0x03

	///////////////////////////////////////////////////////
	// Reset Mesg Codes
	///////////////////////////////////////////////////////

	RESET_FLAGS_MASK uint8 = 0xE0
	RESET_SUSPEND    uint8 = 0x80 // this must follow bitfield def
	RESET_SYNC       uint8 = 0x40 // this must follow bitfield def
	RESET_CMD        uint8 = 0x20 // this must follow bitfield def
	RESET_WDT        uint8 = 0x02
	RESET_RST        uint8 = 0x01
	RESET_POR        uint8 = 0x00

	//////////////////////////////////////////////
	// PC Application Event Codes
	///////////////////////////////////////

	///////

	//NOTE: These events are not generated by the embedded ANT module

	EVENT_RX_BROADCAST    uint8 = 0x9A // returned when module receives broadcast data
	EVENT_RX_ACKNOWLEDGED uint8 = 0x9B // returned when module receives acknowledged data
	EVENT_RX_BURST_PACKET uint8 = 0x9C // returned when module receives burst data

	EVENT_RX_EXT_BROADCAST    uint8 = 0x9D // returned when module receives broadcast data
	EVENT_RX_EXT_ACKNOWLEDGED uint8 = 0x9E // returned when module receives acknowledged data
	EVENT_RX_EXT_BURST_PACKET uint8 = 0x9F // returned when module receives burst data

	EVENT_RX_RSSI_BROADCAST    uint8 = 0xA0 // returned when module receives broadcast data
	EVENT_RX_RSSI_ACKNOWLEDGED uint8 = 0xA1 // returned when module receives acknowledged data
	EVENT_RX_RSSI_BURST_PACKET uint8 = 0xA2 // returned when module receives burst data

	EVENT_RX_FLAG_BROADCAST    uint8 = 0xA3 // returned when module receives broadcast data with flag attached
	EVENT_RX_FLAG_ACKNOWLEDGED uint8 = 0xA4 // returned when module receives acknowledged data with flag attached
	EVENT_RX_FLAG_BURST_PACKET uint8 = 0xA5 // returned when module receives burst data with flag attached

	//////////////////////////////////////////////
	// Integrated ANT-FS Client Response/Event Codes
	//////////////////////////////////////////////

	MESG_FS_ANTFS_EVENT_PAIR_REQUEST      uint8 = 0x01
	MESG_FS_ANTFS_EVENT_DOWNLOAD_START    uint8 = 0x02
	MESG_FS_ANTFS_EVENT_UPLOAD_START      uint8 = 0x03
	MESG_FS_ANTFS_EVENT_DOWNLOAD_COMPLETE uint8 = 0x04
	MESG_FS_ANTFS_EVENT_UPLOAD_COMPLETE   uint8 = 0x05
	MESG_FS_ANTFS_EVENT_ERASE_COMPLETE    uint8 = 0x06
	MESG_FS_ANTFS_EVENT_LINK_STATE        uint8 = 0x07
	MESG_FS_ANTFS_EVENT_AUTH_STATE        uint8 = 0x08
	MESG_FS_ANTFS_EVENT_TRANSPORT_STATE   uint8 = 0x09
	MESG_FS_ANTFS_EVENT_CMD_RECEIVED      uint8 = 0x0A
	MESG_FS_ANTFS_EVENT_CMD_PROCESSED     uint8 = 0x0B

	FS_NO_ERROR_RESPONSE                          uint8 = 0x00
	FS_MEMORY_UNFORMATTED_ERROR_RESPONSE          uint8 = 0x01
	FS_MEMORY_NO_FREE_SECTORS_ERROR_RESPONSE      uint8 = 0x02
	FS_MEMORY_READ_ERROR_RESPONSE                 uint8 = 0x03
	FS_MEMORY_WRITE_ERROR_RESPONSE                uint8 = 0x04
	FS_MEMORY_ERASE_ERROR_RESPONSE                uint8 = 0x05
	FS_TOO_MANY_FILES_OPEN_RESPONSE               uint8 = 0x06
	FS_FILE_INDEX_INVALID_ERROR_RESPONSE          uint8 = 0x07
	FS_FILE_INDEX_EXISTS_ERROR_RESPONSE           uint8 = 0x08
	FS_AUTO_INDEX_FAILED_TRY_AGAIN_ERROR_RESPONSE uint8 = 0x09
	FS_FILE_ALREADY_OPEN_ERROR_RESPONSE           uint8 = 0x0A
	FS_FILE_NOT_OPEN_ERROR_RESPONSE               uint8 = 0x0B
	FS_DIR_CORRUPTED_ERROR_RESPONSE               uint8 = 0x0C
	FS_INVALID_OFFSET_ERROR_RESPONSE              uint8 = 0x0D
	FS_BAD_PERMISSIONS_ERROR_RESPONSE             uint8 = 0x0E
	FS_EOF_REACHED_ERROR_RESPONSE                 uint8 = 0x0F
	FS_INVALID_FILE_HANDLE_ERROR_RESPONSE         uint8 = 0x10

	FS_CRYPTO_OPEN_PERMISSION_ERROR_RESPONSE  uint8 = 0x32
	FS_CRYPTO_HANDLE_ALREADY_IN_USE_RESPONSE  uint8 = 0x33
	FS_CRYPTO_USER_KEY_NOT_SPECIFIED_RESPONSE uint8 = 0x34
	FS_CRYPTO_USER_KEY_ADD_ERROR_RESPONSE     uint8 = 0x35
	FS_CRYPTO_USER_KEY_FETCH_ERROR_RESPONSE   uint8 = 0x36
	FS_CRYPTO_IVNONE_READ_ERROR_RESPONSE      uint8 = 0x37
	FS_CRYPTO_BLOCK_OFFSET_ERROR_RESPONSE     uint8 = 0x38
	FS_CRYPTO_BLOCK_SIZE_ERROR_RESPONSE       uint8 = 0x39
	FS_CRYPTO_CFG_TYPE_NOT_SUPPORTED_RESPONSE uint8 = 0x40

	FS_FIT_FILE_HEADER_ERROR_RESPONSE           uint8 = 0x64
	FS_FIT_FILE_SIZE_INTEGRITY_ERROR_RESPONSE   uint8 = 0x65
	FS_FIT_FILE_CRC_ERROR_RESPONSE              uint8 = 0x66
	FS_FIT_FILE_CHECK_PERMISSION_ERROR_RESPONSE uint8 = 0x67
	FS_FIT_FILE_CHECK_FILE_TYPE_ERROR_RESPONSE  uint8 = 0x68
	FS_FIT_FILE_OP_ABORT_ERROR_RESPONSE         uint8 = 0x69
)

// Message Constants
/////////////////////////////////////////////////////////////////////////////
// Message Format
// Messages are in the format:
//
// AX XX YY -------- CK
//
// where: AX    is the 1 byte sync byte either transmit or recieve
//        XX    is the 1 byte size of the message (0-249) NOTE: THIS WILL BE LIMITED BY THE EMBEDDED RECEIVE BUFFER SIZE
//        YY    is the 1 byte ID of the message (1-255, 0 is invalid)
//        ----- is the data of the message (0-249 bytes of data)
//        CK    is the 1 byte Checksum of the message
/////////////////////////////////////////////////////////////////////////////

const (
	MESG_TX_SYNC          = 0xA4
	MESG_RX_SYNC          = 0xA5
	MESG_SYNC_SIZE        = 1
	MESG_SIZE_SIZE        = 1
	MESG_ID_SIZE          = 1
	MESG_CHANNEL_NUM_SIZE = 1
	MESG_EXT_MESG_BF_SIZE = 1 // NOTE: this could increase in the future
	MESG_CHECKSUM_SIZE    = 1
	MESG_DATA_SIZE        = 9

	// MESG_ANT_MAX_PAYLOAD_SIZE The largest serial message is an ANT data message with all of the extended fields
	MESG_ANT_MAX_PAYLOAD_SIZE = ANT_STANDARD_DATA_PAYLOAD_SIZE

	MESG_MAX_EXT_DATA_SIZE = ANT_EXT_MESG_DEVICE_ID_FIELD_SIZE + ANT_EXT_STRING_SIZE // ANT device ID (4 bytes) +  Padding for ANT EXT string size(27 bytes)

	MESG_MAX_DATA_SIZE  = (MESG_ANT_MAX_PAYLOAD_SIZE + MESG_EXT_MESG_BF_SIZE + MESG_MAX_EXT_DATA_SIZE) // ANT data payload (8 bytes) + extended bitfield (1 byte) + extended data (31 bytes) = 40 bytes
	MESG_MAX_SIZE_VALUE = (MESG_MAX_DATA_SIZE + MESG_CHANNEL_NUM_SIZE)                                 // this is the maximum value that the serial message size value is allowed to be (40 + 1 = 41 bytes)
	MESG_BUFFER_SIZE    = (MESG_SIZE_SIZE + MESG_ID_SIZE + MESG_CHANNEL_NUM_SIZE + MESG_MAX_DATA_SIZE + MESG_CHECKSUM_SIZE)
	MESG_FRAMED_SIZE    = (MESG_ID_SIZE + MESG_CHANNEL_NUM_SIZE + MESG_MAX_DATA_SIZE)
	MESG_HEADER_SIZE    = (MESG_SYNC_SIZE + MESG_SIZE_SIZE + MESG_ID_SIZE)
	MESG_FRAME_SIZE     = (MESG_HEADER_SIZE + MESG_CHECKSUM_SIZE)
	MESG_MAX_SIZE       = (MESG_MAX_DATA_SIZE + MESG_FRAME_SIZE)

	MESG_SIZE_OFFSET             = (MESG_SYNC_SIZE)
	MESG_ID_OFFSET               = (MESG_SYNC_SIZE + MESG_SIZE_SIZE)
	MESG_DATA_OFFSET             = (MESG_HEADER_SIZE)
	MESG_RECOMMENDED_BUFFER_SIZE = 64 // This is the recommended size for serial message buffers if there are no RAM restrictions on the system

	//////////////////////////////////////////////
	// Message ID's
	//////////////////////////////////////////////

	MESG_INVALID_ID = 0x00
	MESG_EVENT_ID   = 0x01

	MESG_VERSION_ID        = 0x3E
	MESG_RESPONSE_EVENT_ID = 0x40

	MESG_UNASSIGN_CHANNEL_ID       = 0x41
	MESG_ASSIGN_CHANNEL_ID         = 0x42
	MESG_CHANNEL_MESG_PERIOD_ID    = 0x43
	MESG_CHANNEL_SEARCH_TIMEOUT_ID = 0x44
	MESG_CHANNEL_RADIO_FREQ_ID     = 0x45
	MESG_NETWORK_KEY_ID            = 0x46
	MESG_RADIO_TX_POWER_ID         = 0x47
	MESG_RADIO_CW_MODE_ID          = 0x48
	MESG_SEARCH_WAVEFORM_ID        = 0x49
	MESG_SYSTEM_RESET_ID           = 0x4A
	MESG_OPEN_CHANNEL_ID           = 0x4B
	MESG_CLOSE_CHANNEL_ID          = 0x4C
	MESG_REQUEST_ID                = 0x4D

	MESG_BROADCAST_DATA_ID    = 0x4E
	MESG_ACKNOWLEDGED_DATA_ID = 0x4F
	MESG_BURST_DATA_ID        = 0x50

	MESG_CHANNEL_ID_ID     = 0x51
	MESG_CHANNEL_STATUS_ID = 0x52
	MESG_RADIO_CW_INIT_ID  = 0x53
	MESG_CAPABILITIES_ID   = 0x54

	MESG_STACKLIMIT_ID = 0x55

	MESG_SCRIPT_DATA_ID = 0x56
	MESG_SCRIPT_CMD_ID  = 0x57

	MESG_ID_LIST_ADD_ID           = 0x59
	MESG_CRYPTO_ID_LIST_ADD_ID    = 0x59
	MESG_ID_LIST_CONFIG_ID        = 0x5A
	MESG_CRYPTO_ID_LIST_CONFIG_ID = 0x5A
	MESG_OPEN_RX_SCAN_ID          = 0x5B

	MESG_EXT_CHANNEL_RADIO_FREQ_ID = 0x5C // OBSOLETE: (for 905 radio)
	MESG_EXT_BROADCAST_DATA_ID     = 0x5D
	MESG_EXT_ACKNOWLEDGED_DATA_ID  = 0x5E
	MESG_EXT_BURST_DATA_ID         = 0x5F

	MESG_CHANNEL_RADIO_TX_POWER_ID    = 0x60
	MESG_GET_SERIAL_NUM_ID            = 0x61
	MESG_GET_TEMP_CAL_ID              = 0x62
	MESG_SET_LP_SEARCH_TIMEOUT_ID     = 0x63
	MESG_SET_TX_SEARCH_ON_NEXT_ID     = 0x64
	MESG_SERIAL_NUM_SET_CHANNEL_ID_ID = 0x65
	MESG_RX_EXT_MESGS_ENABLE_ID       = 0x66
	MESG_RADIO_CONFIG_ALWAYS_ID       = 0x67
	MESG_ENABLE_LED_FLASH_ID          = 0x68
	MESG_XTAL_ENABLE_ID               = 0x6D
	MESG_ANTLIB_CONFIG_ID             = 0x6E
	MESG_STARTUP_MESG_ID              = 0x6F
	MESG_AUTO_FREQ_CONFIG_ID          = 0x70
	MESG_PROX_SEARCH_CONFIG_ID        = 0x71

	MESG_ADV_BURST_DATA_ID         = 0x72
	MESG_EVENT_BUFFERING_CONFIG_ID = 0x74

	MESG_SET_SEARCH_CH_PRIORITY_ID = 0x75

	MESG_HIGH_DUTY_SEARCH_MODE_ID = 0x77
	MESG_CONFIG_ADV_BURST_ID      = 0x78
	MESG_EVENT_FILTER_CONFIG_ID   = 0x79
	MESG_SDU_CONFIG_ID            = 0x7A
	MESG_SDU_SET_MASK_ID          = 0x7B
	MESG_USER_CONFIG_PAGE_ID      = 0x7C
	MESG_ENCRYPT_ENABLE_ID        = 0x7D
	MESG_SET_CRYPTO_KEY_ID        = 0x7E
	MESG_SET_CRYPTO_INFO_ID       = 0x7F
	MESG_CUBE_CMD_ID              = 0x80

	MESG_ACTIVE_SEARCH_SHARING_ID = 0x81
	MESG_NVM_CRYPTO_KEY_OPS_ID    = 0x83

	MESG_GET_PIN_DIODE_CONTROL_ID = 0x8D
	MESG_PIN_DIODE_CONTROL_ID     = 0x8E
	MESG_FIT1_SET_AGC_ID          = 0x8F

	MESG_FIT1_SET_EQUIP_STATE_ID = 0x91 // *** CONFLICT: w/ Sensrcore, Fit1 will never have sensrcore enabled

	// Sensrcore Messages
	MESG_SET_CHANNEL_INPUT_MASK_ID     = 0x90
	MESG_SET_CHANNEL_DATA_TYPE_ID      = 0x91
	MESG_READ_PINS_FOR_SECT_ID         = 0x92
	MESG_TIMER_SELECT_ID               = 0x93
	MESG_ATOD_SETTINGS_ID              = 0x94
	MESG_SET_SHARED_ADDRESS_ID         = 0x95
	MESG_ATOD_EXTERNAL_ENABLE_ID       = 0x96
	MESG_ATOD_PIN_SETUP_ID             = 0x97
	MESG_SETUP_ALARM_ID                = 0x98
	MESG_ALARM_VARIABLE_MODIFY_TEST_ID = 0x99
	MESG_PARTIAL_RESET_ID              = 0x9A
	MESG_OVERWRITE_TEMP_CAL_ID         = 0x9B
	MESG_SERIAL_PASSTHRU_SETTINGS_ID   = 0x9C

	MESG_BIST_ID             = 0xAA
	MESG_UNLOCK_INTERFACE_ID = 0xAD
	MESG_SERIAL_ERROR_ID     = 0xAE
	MESG_SET_ID_STRING_ID    = 0xAF

	MESG_PORT_GET_IO_STATE_ID = 0xB4
	MESG_PORT_SET_IO_STATE_ID = 0xB5

	MESG_RSSI_ID                   = 0xC0
	MESG_RSSI_BROADCAST_DATA_ID    = 0xC1
	MESG_RSSI_ACKNOWLEDGED_DATA_ID = 0xC2
	MESG_RSSI_BURST_DATA_ID        = 0xC3
	MESG_RSSI_SEARCH_THRESHOLD_ID  = 0xC4
	MESG_SLEEP_ID                  = 0xC5
	MESG_GET_GRMN_ESN_ID           = 0xC6
	MESG_SET_USB_INFO_ID           = 0xC7

	MESG_HCI_COMMAND_COMPLETE = 0xC8

	// 0xE0 - 0xEF reserved for extended ID
	MESG_EXT_ID_0 = 0xE0
	MESG_EXT_ID_1 = 0xE1
	MESG_EXT_ID_2 = 0xE2

	// 0xE0 extended IDs
	MESG_EXT_RESPONSE_ID = 0xE000

	// 0xE1 extended IDs
	MESG_EXT_REQUEST_ID = 0xE100

	// 0xE2 extended IDs
	MESG_FS_INIT_MEMORY_ID                  = 0xE200
	MESG_FS_FORMAT_MEMORY_ID                = 0xE201
	MESG_FS_GET_USED_SPACE_ID               = 0xE202
	MESG_FS_GET_FREE_SPACE_ID               = 0xE203
	MESG_FS_FIND_FILE_INDEX_ID              = 0xE204
	MESG_FS_DIRECTORY_READ_ABSOLUTE_ID      = 0xE205
	MESG_FS_DIRECTORY_READ_ENTRY_ID         = 0xE206
	MESG_FS_DIRECTORY_SAVE_ID               = 0xE207
	MESG_FS_DIRECTORY_GET_SIZE_ID           = 0xE208
	MESG_FS_DIRECTORY_REBUILD_ID            = 0xE209
	MESG_FS_FILE_CREATE_ID                  = 0xE20A
	MESG_FS_FILE_OPEN_ID                    = 0xE20B
	MESG_FS_FILE_DELETE_ID                  = 0xE20C
	MESG_FS_FILE_CLOSE_ID                   = 0xE20D
	MESG_FS_FILE_READ_ABSOLUTE_ID           = 0xE20E
	MESG_FS_FILE_READ_RELATIVE_ID           = 0xE20F
	MESG_FS_FILE_WRITE_ABSOLUTE_ID          = 0xE210
	MESG_FS_FILE_WRITE_RELATIVE_ID          = 0xE211
	MESG_FS_FILE_SET_SPECIFIC_FLAGS_ID      = 0xE212
	MESG_FS_FILE_GET_SIZE_ID                = 0xE213
	MESG_FS_FILE_GET_SPECIFIC_FILE_FLAGS_ID = 0xE214
	MESG_FS_FILE_GET_SIZE_IN_MEM_ID         = 0xE215
	MESG_FS_DIRECTORY_READ_LOCK_ID          = 0xE216

	MESG_FS_FILE_SET_GENERAL_FLAGS_ID   = 0xE21E
	MESG_FS_DIRECTORY_WRITE_ABSOLUTE_ID = 0xE21F
	// reserved
	MESG_MEMDEV_EEPROM_INIT_ID = 0xE220
	MESG_MEMDEV_FLASH_INIT_ID  = 0xE221
	//reserved
	MESG_FS_ANTFS_EVENT_ID                 = 0xE230
	MESG_FS_ANTFS_OPEN_ID                  = 0xE231
	MESG_FS_ANTFS_CLOSE_ID                 = 0xE232
	MESG_FS_ANTFS_CONFIG_BEACON_ID         = 0xE233
	MESG_FS_ANTFS_SET_AUTH_STRING_ID       = 0xE234
	MESG_FS_ANTFS_SET_BEACON_STATE_ID      = 0xE235
	MESG_FS_ANTFS_PAIR_RESPONSE_ID         = 0xE236
	MESG_FS_ANTFS_SET_LINK_FREQ_ID         = 0xE237
	MESG_FS_ANTFS_SET_BEACON_TIMEOUT_ID    = 0xE238
	MESG_FS_ANTFS_SET_PAIRING_TIMEOUT_ID   = 0xE239
	MESG_FS_ANTFS_REMOTE_FILE_CREATE_EN_ID = 0xE23A
	MESG_FS_ANTFS_GET_CMD_PIPE_ID          = 0xE23B
	MESG_FS_ANTFS_SET_CMD_PIPE_ID          = 0xE23C
	MESG_FS_SYSTEM_TIME_ID                 = 0xE23D
	MESG_FS_ANTFS_SET_ANTFS_STATE_ID       = 0xE23E
	// reserved
	MESG_FS_CRYPTO_ADD_USER_KEY_INDEX_ID = 0xE245
	MESG_FS_CRYPTO_SET_USER_KEY_INDEX_ID = 0xE246
	MESG_FS_CRYPTO_SET_USER_KEY_VAL_ID   = 0xE247
	// reserved
	MESG_FS_FIT_FILE_INTEGRITY_CHECK_ID = 0xE250

	//////////////////////////////////////////////
	// Message Sizes
	//////////////////////////////////////////////

	MESG_INVALID_SIZE = 0

	MESG_VERSION_SIZE        = 13
	MESG_RESPONSE_EVENT_SIZE = 3
	MESG_CHANNEL_STATUS_SIZE = 2

	MESG_UNASSIGN_CHANNEL_SIZE       = 1
	MESG_ASSIGN_CHANNEL_SIZE         = 3
	MESG_CHANNEL_ID_SIZE             = 5
	MESG_CHANNEL_MESG_PERIOD_SIZE    = 3
	MESG_CHANNEL_SEARCH_TIMEOUT_SIZE = 2
	MESG_CHANNEL_RADIO_FREQ_SIZE     = 2
	MESG_CHANNEL_RADIO_TX_POWER_SIZE = 2
	MESG_NETWORK_KEY_SIZE            = 9
	MESG_RADIO_TX_POWER_SIZE         = 2
	MESG_RADIO_CW_MODE_SIZE          = 3
	MESG_RADIO_CW_INIT_SIZE          = 1
	MESG_SYSTEM_RESET_SIZE           = 1
	MESG_OPEN_CHANNEL_SIZE           = 1
	MESG_CLOSE_CHANNEL_SIZE          = 1
	MESG_REQUEST_SIZE                = 2
	MESG_REQUEST_USER_NVM_SIZE       = 5

	MESG_CAPABILITIES_SIZE = 8
	MESG_STACKLIMIT_SIZE   = 2

	MESG_SCRIPT_DATA_SIZE = 10
	MESG_SCRIPT_CMD_SIZE  = 3

	MESG_ID_LIST_ADD_SIZE            = 6
	MESG_ID_LIST_CONFIG_SIZE         = 3
	MESG_OPEN_RX_SCAN_SIZE           = 1
	MESG_EXT_CHANNEL_RADIO_FREQ_SIZE = 3

	MESG_RADIO_CONFIG_ALWAYS_SIZE   = 2
	MESG_RX_EXT_MESGS_ENABLE_SIZE   = 2
	MESG_SET_TX_SEARCH_ON_NEXT_SIZE = 2
	MESG_SET_LP_SEARCH_TIMEOUT_SIZE = 2

	MESG_SERIAL_NUM_SET_CHANNEL_ID_SIZE = 3
	MESG_ENABLE_LED_FLASH_SIZE          = 2
	MESG_GET_SERIAL_NUM_SIZE            = 4
	MESG_GET_TEMP_CAL_SIZE              = 4
	MESG_CONFIG_ADV_BURST_SIZE          = 9
	MESG_CONFIG_ADV_BURST_SIZE_EXT      = 12
	MESG_ANTLIB_CONFIG_SIZE             = 2
	MESG_XTAL_ENABLE_SIZE               = 1
	MESG_STARTUP_MESG_SIZE              = 1
	MESG_AUTO_FREQ_CONFIG_SIZE          = 4
	MESG_PROX_SEARCH_CONFIG_SIZE        = 2
	MESG_EVENT_BUFFERING_CONFIG_SIZE    = 6
	MESG_EVENT_FILTER_CONFIG_SIZE       = 3
	MESG_HIGH_DUTY_SEARCH_MODE_SIZE     = 3
	MESG_SDU_CONFIG_SIZE                = 3
	MESG_SDU_SET_MASK_SIZE              = 9

	MESG_USER_CONFIG_PAGE_SIZE = 3

	MESG_ACTIVE_SEARCH_SHARING_SIZE = 2

	MESG_SET_SEARCH_CH_PRIORITY_SIZE = 2

	MESG_ENCRYPT_ENABLE_SIZE        = 4
	MESG_SET_CRYPTO_KEY_SIZE        = 17
	MESG_SET_CRYPTO_ID_SIZE         = 5
	MESG_SET_CRYPTO_USER_INFO_SIZE  = 20
	MESG_SET_CRYPTO_RNG_SEED_SIZE   = 17
	MESG_NVM_CRYPTO_KEY_LOAD_SIZE   = 3
	MESG_NVM_CRYPTO_KEY_STORE_SIZE  = 18
	MESG_CRYPTO_ID_LIST_ADD_SIZE    = 6
	MESG_CRYPTO_ID_LIST_CONFIG_SIZE = 3

	MESG_GET_PIN_DIODE_CONTROL_SIZE = 1
	MESG_PIN_DIODE_CONTROL_ID_SIZE  = 2
	MESG_FIT1_SET_EQUIP_STATE_SIZE  = 2
	MESG_FIT1_SET_AGC_SIZE          = 4

	MESG_BIST_SIZE               = 6
	MESG_UNLOCK_INTERFACE_SIZE   = 1
	MESG_SET_SHARED_ADDRESS_SIZE = 3

	MESG_GET_GRMN_ESN_SIZE = 5

	MESG_PORT_SET_IO_STATE_SIZE = 5

	MESG_SLEEP_SIZE = 1

	MESG_EXT_DATA_SIZE = 13

	MESG_RSSI_SIZE                  = 5
	MESG_RSSI_DATA_SIZE             = 17
	MESG_RSSI_RESPONSE_SIZE         = 7
	MESG_RSSI_SEARCH_THRESHOLD_SIZE = 2

	MESG_MEMDEV_EEPROM_INIT_SIZE              = 0x04
	MESG_FS_INIT_MEMORY_SIZE                  = 0x01
	MESG_FS_FORMAT_MEMORY_SIZE                = 0x05
	MESG_FS_DIRECTORY_SAVE_SIZE               = 0x01
	MESG_FS_DIRECTORY_REBUILD_SIZE            = 0x01
	MESG_FS_FILE_DELETE_SIZE                  = 0x02
	MESG_FS_FILE_CLOSE_SIZE                   = 0x02
	MESG_FS_FILE_SET_SPECIFIC_FLAGS_SIZE      = 0x03
	MESG_FS_DIRECTORY_READ_LOCK_SIZE          = 0x02
	MESG_FS_SYSTEM_TIME_SIZE                  = 0x05
	MESG_FS_CRYPTO_ADD_USER_KEY_INDEX_SIZE    = 0x34
	MESG_FS_CRYPTO_SET_USER_KEY_INDEX_SIZE    = 0x05
	MESG_FS_CRYPTO_SET_USER_KEY_VAL_SIZE      = 0x33
	MESG_FS_FIT_FILE_INTEGRITY_CHECK_SIZE     = 0x02
	MESG_FS_ANTFS_OPEN_SIZE                   = 0x01
	MESG_FS_ANTFS_CLOSE_SIZE                  = 0x01
	MESG_FS_ANTFS_CONFIG_BEACON_SIZE          = 0x09
	MESG_FS_ANTFS_SET_AUTH_STRING_SIZE        = 0x02
	MESG_FS_ANTFS_SET_BEACON_STATE_SIZE       = 0x02
	MESG_FS_ANTFS_PAIR_RESPONSE_SIZE          = 0x02
	MESG_FS_ANTFS_SET_LINK_FREQ_SIZE          = 0x03
	MESG_FS_ANTFS_SET_BEACON_TIMEOUT_SIZE     = 0x02
	MESG_FS_ANTFS_SET_PAIRING_TIMEOUT_SIZE    = 0x02
	MESG_FS_ANTFS_REMOTE_FILE_CREATE_EN_SIZE  = 0x02
	MESG_FS_GET_USED_SPACE_SIZE               = 0x03
	MESG_FS_GET_FREE_SPACE_SIZE               = 0x03
	MESG_FS_FIND_FILE_INDEX_SIZE              = 0x07
	MESG_FS_DIRECTORY_READ_ABSOLUTE_SIZE      = 0x08
	MESG_FS_DIRECTORY_READ_ENTRY_SIZE         = 0x05
	MESG_FS_DIRECTORY_GET_SIZE_SIZE           = 0x03
	MESG_FS_FILE_CREATE_SIZE                  = 0x0B
	MESG_FS_FILE_OPEN_SIZE                    = 0x06
	MESG_FS_FILE_READ_ABSOLUTE_SIZE           = 0x09
	MESG_FS_FILE_READ_RELATIVE_SIZE           = 0x05
	MESG_FS_FILE_WRITE_ABSOLUTE_SIZE          = 0x09
	MESG_FS_FILE_WRITE_RELATIVE_SIZE          = 0x05
	MESG_FS_FILE_GET_SIZE_SIZE                = 0x04
	MESG_FS_FILE_GET_SIZE_IN_MEM_SIZE         = 0x04
	MESG_FS_FILE_GET_SPECIFIC_FILE_FLAGS_SIZE = 0x04
	MESG_FS_SYSTEM_TIME_REQUEST_SIZE          = 0x03
	MESG_FS_ANTFS_GET_CMD_PIPE_SIZE           = 0x05
	MESG_FS_ANTFS_SET_CMD_PIPE_SIZE           = 0x05
	MESG_FS_REQUEST_RESPONSE_SIZE             = 0x03
)
