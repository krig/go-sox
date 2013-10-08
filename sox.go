package sox

/*
#cgo pkg-config: sox
#include <sox.h>
*/
import "C"

const (
	EOF = C.SOX_EOF
	EHDR = C.SOX_EHDR
	EFMT = C.SOX_EFMT
	ENOMEM = C.SOX_ENOMEM
	EPERM = C.SOX_EPERM
	ENOTSUP = C.SOX_ENOTSUP
	EINVAL = C.SOX_EINVAL

	HAVE_POPEN = C.sox_version_have_popen
	HAVE_MAGIC = C.sox_version_have_magic
	HAVE_THREADS = C.sox_version_have_threads
	HAVE_MEMOPEN = C.sox_version_have_memopen

	ENCODING_UNKNOWN = C.SOX_ENCODING_UNKNOWN    /**< encoding has not yet been determined */
	ENCODING_SIGN2 = C.SOX_ENCODING_SIGN2      /**< signed linear 2's comp: Mac */
	ENCODING_UNSIGNED = C.SOX_ENCODING_UNSIGNED   /**< unsigned linear: Sound Blaster */
	ENCODING_FLOAT = C.SOX_ENCODING_FLOAT      /**< floating point (binary format) */
	ENCODING_FLOAT_TEXT = C.SOX_ENCODING_FLOAT_TEXT /**< floating point (text format) */
	ENCODING_FLAC = C.SOX_ENCODING_FLAC       /**< FLAC compression */
	ENCODING_HCOM = C.SOX_ENCODING_HCOM       /**< Mac FSSD files with Huffman compression */
	ENCODING_WAVPACK = C.SOX_ENCODING_WAVPACK    /**< WavPack with integer samples */
	ENCODING_WAVPACKF = C.SOX_ENCODING_WAVPACKF   /**< WavPack with float samples */
	ENCODING_ULAW = C.SOX_ENCODING_ULAW       /**< u-law signed logs: US telephony SPARC */
	ENCODING_ALAW = C.SOX_ENCODING_ALAW       /**< A-law signed logs: non-US telephony Psion */
	ENCODING_G721 = C.SOX_ENCODING_G721       /**< G.721 4-bit ADPCM */
	ENCODING_G723 = C.SOX_ENCODING_G723       /**< G.723 3 or 5 bit ADPCM */
	ENCODING_CL_ADPCM = C.SOX_ENCODING_CL_ADPCM   /**< Creative Labs 8 --> 234 bit Compressed PCM */
	ENCODING_CL_ADPCM16 = C.SOX_ENCODING_CL_ADPCM16 /**< Creative Labs 16 --> 4 bit Compressed PCM */
	ENCODING_MS_ADPCM = C.SOX_ENCODING_MS_ADPCM   /**< Microsoft Compressed PCM */
	ENCODING_IMA_ADPCM = C.SOX_ENCODING_IMA_ADPCM  /**< IMA Compressed PCM */
	ENCODING_OKI_ADPCM = C.SOX_ENCODING_OKI_ADPCM  /**< Dialogic/OKI Compressed PCM */
	ENCODING_DPCM = C.SOX_ENCODING_DPCM       /**< Differential PCM: Fasttracker 2 (xi) */
	ENCODING_DWVW = C.SOX_ENCODING_DWVW       /**< Delta Width Variable Word */
	ENCODING_DWVWN = C.SOX_ENCODING_DWVWN      /**< Delta Width Variable Word N-bit */
	ENCODING_GSM = C.SOX_ENCODING_GSM        /**< GSM 6.10 33byte frame lossy compression */
	ENCODING_MP3 = C.SOX_ENCODING_MP3        /**< MP3 compression */
	ENCODING_VORBIS = C.SOX_ENCODING_VORBIS     /**< Vorbis compression */
	ENCODING_AMR_WB = C.SOX_ENCODING_AMR_WB     /**< AMR-WB compression */
	ENCODING_AMR_NB = C.SOX_ENCODING_AMR_NB     /**< AMR-NB compression */
	ENCODING_CVSD = C.SOX_ENCODING_CVSD       /**< Continuously Variable Slope Delta modulation */
	ENCODING_LPC10 = C.SOX_ENCODING_LPC10      /**< Linear Predictive Coding */

	LOSSLESS = C.sox_encodings_none
	LOSSY1 = C.sox_encodings_lossy1
	LOSSY2 = C.sox_encodings_lossy2

	LOOP_NONE = C.sox_loop_none
	LOOP_FORWARD = C.sox_loop_forward
	LOOP_FORWARD_BACK = C.sox_loop_forward_back
	LOOP_8 = C.sox_loop_8
	LOOP_SUSTAIN_DECAY = C.sox_loop_sustain_decay

	DEFAULT_CHANNELS = C.SOX_DEFAULT_CHANNELS
	DEFAULT_RATE = C.SOX_DEFAULT_RATE
	DEFAULT_PRECISION = C.SOX_DEFAULT_PRECISION
	DEFAULT_ENCODING = C.SOX_DEFAULT_ENCODING

	MAX_NLOOPS = C.SOX_MAX_NLOOPS

	FILE_NOSTDIO = C.SOX_FILE_NOSTDIO
	FILE_DEVICE = C.SOX_FILE_DEVICE
	FILE_PHONY = C.SOX_FILE_PHONY
	FILE_REWIND = C.SOX_FILE_REWIND

)

func FormatInit() bool {
	ret := C.sox_format_init()
	return ret == C.SOX_SUCCESS
}

func FormatQuit() {
	C.sox_format_quit()
}

func Init() bool {
	ret := C.sox_init()
	return ret == C.SOX_SUCCESS
}

func Quit() bool {
	ret := C.sox_quit()
	return ret == C.SOX_SUCCESS
}


