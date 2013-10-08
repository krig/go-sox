package sox

import "unsafe"

/*
#cgo pkg-config: sox
#include <sox.h>
#include <stdlib.h>
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

func StrError(errno int) string {
	return C.GoString(C.sox_strerror(C.int(errno)))
}

type Format struct {
	cFormat *C.sox_format_t
}

type SignalInfo struct {
	cSignal *C.sox_signalinfo_t
}

type EncodingInfo struct {
	cEncoding *C.sox_encodinginfo_t
}

type FormatHandler struct {
	cHandler *C.sox_format_handler_t
}

type EffectsGlobals struct {
	cGlobals *C.sox_effects_globals_t
}

type EffectsChain struct {
	cChain *C.sox_effects_chain_t
}

type EffectHandler struct {
	cHandler *C.sox_effect_handler_t
}

type Effect struct {
	cEffect *C.sox_effect_t
}

func (f *Format) Close() {
	C.sox_close(f.cFormat)
	f.cFormat = nil
}

func (f *Format) Filename() string {
	return C.GoString(f.cFormat.filename)
}

func (f *Format) Signal() *SignalInfo {
	var s SignalInfo
	s.cSignal = &f.cFormat.signal
	return &s
}

func (f *Format) Encoding() *EncodingInfo {
	var e EncodingInfo
	e.cEncoding = &f.cFormat.encoding
	return &e
}

func (f *Format) Type() string {
	return C.GoString(f.cFormat.filetype)
}

func (f *Format) Seekable() bool {
	return f.cFormat.seekable != 0
}

func (f *Format) Read(buffer []int32) int64 {
	return int64(C.sox_read(f.cFormat, (*C.sox_sample_t)(unsafe.Pointer(&buffer[0])), C.size_t(len(buffer))))
}

func (f *Format) Write(buffer []int32) int64 {
	return int64(C.sox_write(f.cFormat, (*C.sox_sample_t)(unsafe.Pointer(&buffer[0])), C.size_t(len(buffer))))
}

func (f *Format) Seek(offset uint64) bool {
	return C.sox_seek(f.cFormat, C.sox_uint64_t(offset), C.SOX_SEEK_SET) == C.SOX_SUCCESS
}

func OpenRead(path string) *Format {
	cpath := C.CString(path)
	var fmt Format
	fmt.cFormat = C.sox_open_read(cpath, nil, nil, nil)
	C.free(unsafe.Pointer(cpath))
	if fmt.cFormat == nil {
		return nil
	}
	return &fmt
}

func OpenMemRead(buffer []byte) *Format {
	var fmt Format
	fmt.cFormat = C.sox_open_mem_read(unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)), nil, nil, nil)
	if fmt.cFormat == nil {
		return nil
	}
	return &fmt
}

func FormatSupportsEncoding(path string, encoding *EncodingInfo) bool {
	cpath := C.CString(path)
	ret := C.sox_format_supports_encoding(cpath, nil, encoding.cEncoding)
	C.free(unsafe.Pointer(cpath))
	return int(ret) != 0
}

func maybeCSignal(signal *SignalInfo) *C.sox_signalinfo_t {
	if signal != nil {
		return signal.cSignal
	}
	return nil
}

func OpenWrite(path string, signal *SignalInfo) *Format {
	cpath := C.CString(path)
	var fmt Format
	fmt.cFormat = C.sox_open_write(cpath, maybeCSignal(signal), nil, nil, nil, nil)
	C.free(unsafe.Pointer(cpath))
	if fmt.cFormat == nil {
		return nil
	}
	return &fmt
}

func OpenMemWrite(buffer []byte, signal *SignalInfo) *Format {
	var fmt Format
	fmt.cFormat = C.sox_open_mem_write(unsafe.Pointer(&buffer[0]),
		C.size_t(len(buffer)),
		maybeCSignal(signal),
		nil,
		nil,
		nil)
	if fmt.cFormat == nil {
		return nil
	}
	return &fmt
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func FindFormat(name string, ignore_devices bool) *FormatHandler {
	var fmt FormatHandler
	cname := C.CString(name)
	fmt.cHandler = C.sox_find_format(cname, C.sox_bool(bool2int(ignore_devices)))
	C.free(unsafe.Pointer(cname))
	return &fmt
}

func GetEffectsGlobals() *EffectsGlobals {
	var g EffectsGlobals
	g.cGlobals = C.sox_get_effects_globals()
	return &g
}

func CreateEffectsChain(in *EncodingInfo, out *EncodingInfo) *EffectsChain {
	var chain EffectsChain
	chain.cChain = C.sox_create_effects_chain(in.cEncoding, out.cEncoding)
	return &chain
}

func (c *EffectsChain) Delete() {
	C.sox_delete_effects_chain(c.cChain)
}

func (c *EffectsChain) AddEffect(effect *Effect, in, out *SignalInfo) bool {
	return C.sox_add_effect(c.cChain, effect.cEffect, in.cSignal, out.cSignal) == C.SOX_SUCCESS
}

func (c *EffectsChain) FlowEffects() bool {
	return C.sox_flow_effects(c.cChain, nil, nil) == C.SOX_SUCCESS
}

func FindEffect(name string) *EffectHandler {
	var h EffectHandler
	cname := C.CString(name)
	h.cHandler = C.sox_find_effect(cname)
	C.free(unsafe.Pointer(cname))
	return &h
}

func CreateEffect(handler *EffectHandler) *Effect {
	var e Effect
	e.cEffect = C.sox_create_effect(handler.cHandler)
	return &e
}

func (e *Effect) Free() {
	C.free(unsafe.Pointer(e.cEffect))
	e.cEffect = nil
}

func (e *Effect) Options(args ...interface{}) int {
	if len(args) == 0 {
		return int(C.sox_effect_options(e.cEffect, 0, nil))
	}
	var cargs [10](*C.char)
	n := len(args)
	for i, v := range args {
		switch v := v.(type) {
		case *Format:
			cargs[i] = (*C.char)(unsafe.Pointer(v.cFormat))
		case string:
			cargs[i] = C.CString(v)
		}
	}

	return int(C.sox_effect_options(e.cEffect, C.int(n), (&cargs[0])))
}
