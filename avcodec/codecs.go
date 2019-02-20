// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// Multiple encoders have the same ID and are able to encode compatible streams.
const (
	AV_CODEC_ID_012V              = int(C.AV_CODEC_ID_012V)
	AV_CODEC_ID_4XM               = int(C.AV_CODEC_ID_4XM)
	AV_CODEC_ID_8BPS              = int(C.AV_CODEC_ID_8BPS)
	AV_CODEC_ID_8SVX_EXP          = int(C.AV_CODEC_ID_8SVX_EXP)
	AV_CODEC_ID_8SVX_FIB          = int(C.AV_CODEC_ID_8SVX_FIB)
	AV_CODEC_ID_A64_MULTI         = int(C.AV_CODEC_ID_A64_MULTI)
	AV_CODEC_ID_A64_MULTI5        = int(C.AV_CODEC_ID_A64_MULTI5)
	AV_CODEC_ID_AAC               = int(C.AV_CODEC_ID_AAC)
	AV_CODEC_ID_AAC_LATM          = int(C.AV_CODEC_ID_AAC_LATM)
	AV_CODEC_ID_AASC              = int(C.AV_CODEC_ID_AASC)
	AV_CODEC_ID_AC3               = int(C.AV_CODEC_ID_AC3)
	AV_CODEC_ID_ADPCM_4XM         = int(C.AV_CODEC_ID_ADPCM_4XM)
	AV_CODEC_ID_ADPCM_ADX         = int(C.AV_CODEC_ID_ADPCM_ADX)
	AV_CODEC_ID_ADPCM_AFC         = int(C.AV_CODEC_ID_ADPCM_AFC)
	AV_CODEC_ID_ADPCM_CT          = int(C.AV_CODEC_ID_ADPCM_CT)
	AV_CODEC_ID_ADPCM_DTK         = int(C.AV_CODEC_ID_ADPCM_DTK)
	AV_CODEC_ID_ADPCM_EA          = int(C.AV_CODEC_ID_ADPCM_EA)
	AV_CODEC_ID_ADPCM_EA_MAXIS_XA = int(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AV_CODEC_ID_ADPCM_EA_R1       = int(C.AV_CODEC_ID_ADPCM_EA_R1)
	AV_CODEC_ID_ADPCM_EA_R2       = int(C.AV_CODEC_ID_ADPCM_EA_R2)
	AV_CODEC_ID_ADPCM_EA_R3       = int(C.AV_CODEC_ID_ADPCM_EA_R3)
	AV_CODEC_ID_ADPCM_EA_XAS      = int(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AV_CODEC_ID_ADPCM_G722        = int(C.AV_CODEC_ID_ADPCM_G722)
	AV_CODEC_ID_ADPCM_G726        = int(C.AV_CODEC_ID_ADPCM_G726)
	AV_CODEC_ID_ADPCM_G726LE      = int(C.AV_CODEC_ID_ADPCM_G726LE)
	AV_CODEC_ID_ADPCM_IMA_AMV     = int(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AV_CODEC_ID_ADPCM_IMA_APC     = int(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AV_CODEC_ID_ADPCM_IMA_DK3     = int(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AV_CODEC_ID_ADPCM_IMA_DK4     = int(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AV_CODEC_ID_ADPCM_IMA_EA_EACS = int(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AV_CODEC_ID_ADPCM_IMA_EA_SEAD = int(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AV_CODEC_ID_ADPCM_IMA_ISS     = int(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AV_CODEC_ID_ADPCM_IMA_OKI     = int(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AV_CODEC_ID_ADPCM_IMA_QT      = int(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AV_CODEC_ID_ADPCM_IMA_RAD     = int(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AV_CODEC_ID_ADPCM_IMA_SMJPEG  = int(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AV_CODEC_ID_ADPCM_IMA_WAV     = int(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AV_CODEC_ID_ADPCM_IMA_WS      = int(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AV_CODEC_ID_ADPCM_MS          = int(C.AV_CODEC_ID_ADPCM_MS)
	AV_CODEC_ID_ADPCM_SBPRO_2     = int(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AV_CODEC_ID_ADPCM_SBPRO_3     = int(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AV_CODEC_ID_ADPCM_SBPRO_4     = int(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AV_CODEC_ID_ADPCM_SWF         = int(C.AV_CODEC_ID_ADPCM_SWF)
	AV_CODEC_ID_ADPCM_THP         = int(C.AV_CODEC_ID_ADPCM_THP)
	AV_CODEC_ID_ADPCM_VIMA        = int(C.AV_CODEC_ID_ADPCM_VIMA)
	AV_CODEC_ID_ADPCM_XA          = int(C.AV_CODEC_ID_ADPCM_XA)
	AV_CODEC_ID_ADPCM_YAMAHA      = int(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AV_CODEC_ID_AIC               = int(C.AV_CODEC_ID_AIC)
	AV_CODEC_ID_ALAC              = int(C.AV_CODEC_ID_ALAC)
	AV_CODEC_ID_ALIAS_PIX         = int(C.AV_CODEC_ID_ALIAS_PIX)
	AV_CODEC_ID_AMR_NB            = int(C.AV_CODEC_ID_AMR_NB)
	AV_CODEC_ID_AMR_WB            = int(C.AV_CODEC_ID_AMR_WB)
	AV_CODEC_ID_AMV               = int(C.AV_CODEC_ID_AMV)
	AV_CODEC_ID_ANM               = int(C.AV_CODEC_ID_ANM)
	AV_CODEC_ID_ANSI              = int(C.AV_CODEC_ID_ANSI)
	AV_CODEC_ID_APE               = int(C.AV_CODEC_ID_APE)
	AV_CODEC_ID_ASS               = int(C.AV_CODEC_ID_ASS)
	AV_CODEC_ID_ASV1              = int(C.AV_CODEC_ID_ASV1)
	AV_CODEC_ID_ASV2              = int(C.AV_CODEC_ID_ASV2)
	AV_CODEC_ID_ATRAC1            = int(C.AV_CODEC_ID_ATRAC1)
	AV_CODEC_ID_ATRAC3            = int(C.AV_CODEC_ID_ATRAC3)
	AV_CODEC_ID_ATRAC3P           = int(C.AV_CODEC_ID_ATRAC3P)
	AV_CODEC_ID_AURA              = int(C.AV_CODEC_ID_AURA)
	AV_CODEC_ID_AURA2             = int(C.AV_CODEC_ID_AURA2)
	AV_CODEC_ID_AVRN              = int(C.AV_CODEC_ID_AVRN)
	AV_CODEC_ID_AVRP              = int(C.AV_CODEC_ID_AVRP)
	AV_CODEC_ID_AVS               = int(C.AV_CODEC_ID_AVS)
	AV_CODEC_ID_AVUI              = int(C.AV_CODEC_ID_AVUI)
	AV_CODEC_ID_AYUV              = int(C.AV_CODEC_ID_AYUV)
	AV_CODEC_ID_BETHSOFTVID       = int(C.AV_CODEC_ID_BETHSOFTVID)
	AV_CODEC_ID_BFI               = int(C.AV_CODEC_ID_BFI)
	AV_CODEC_ID_BINKAUDIO_DCT     = int(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AV_CODEC_ID_BINKAUDIO_RDFT    = int(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AV_CODEC_ID_BINKVIDEO         = int(C.AV_CODEC_ID_BINKVIDEO)
	AV_CODEC_ID_BINTEXT           = int(C.AV_CODEC_ID_BINTEXT)
	AV_CODEC_ID_BIN_DATA          = int(C.AV_CODEC_ID_BIN_DATA)
	AV_CODEC_ID_BMP               = int(C.AV_CODEC_ID_BMP)
	AV_CODEC_ID_BMV_AUDIO         = int(C.AV_CODEC_ID_BMV_AUDIO)
	AV_CODEC_ID_BMV_VIDEO         = int(C.AV_CODEC_ID_BMV_VIDEO)
	AV_CODEC_ID_BRENDER_PIX       = int(C.AV_CODEC_ID_BRENDER_PIX)
	AV_CODEC_ID_C93               = int(C.AV_CODEC_ID_C93)
	AV_CODEC_ID_CAVS              = int(C.AV_CODEC_ID_CAVS)
	AV_CODEC_ID_CDGRAPHICS        = int(C.AV_CODEC_ID_CDGRAPHICS)
	AV_CODEC_ID_CDXL              = int(C.AV_CODEC_ID_CDXL)
	AV_CODEC_ID_CELT              = int(C.AV_CODEC_ID_CELT)
	AV_CODEC_ID_CINEPAK           = int(C.AV_CODEC_ID_CINEPAK)
	AV_CODEC_ID_CLJR              = int(C.AV_CODEC_ID_CLJR)
	AV_CODEC_ID_CLLC              = int(C.AV_CODEC_ID_CLLC)
	AV_CODEC_ID_CMV               = int(C.AV_CODEC_ID_CMV)
	AV_CODEC_ID_COMFORT_NOISE     = int(C.AV_CODEC_ID_COMFORT_NOISE)
	AV_CODEC_ID_COOK              = int(C.AV_CODEC_ID_COOK)
	AV_CODEC_ID_CPIA              = int(C.AV_CODEC_ID_CPIA)
	AV_CODEC_ID_CSCD              = int(C.AV_CODEC_ID_CSCD)
	AV_CODEC_ID_CYUV              = int(C.AV_CODEC_ID_CYUV)
	AV_CODEC_ID_DFA               = int(C.AV_CODEC_ID_DFA)
	AV_CODEC_ID_DIRAC             = int(C.AV_CODEC_ID_DIRAC)
	AV_CODEC_ID_DNXHD             = int(C.AV_CODEC_ID_DNXHD)
	AV_CODEC_ID_DPX               = int(C.AV_CODEC_ID_DPX)
	AV_CODEC_ID_DSD_LSBF          = int(C.AV_CODEC_ID_DSD_LSBF)
	AV_CODEC_ID_DSD_LSBF_PLANAR   = int(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AV_CODEC_ID_DSD_MSBF          = int(C.AV_CODEC_ID_DSD_MSBF)
	AV_CODEC_ID_DSD_MSBF_PLANAR   = int(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AV_CODEC_ID_DSICINAUDIO       = int(C.AV_CODEC_ID_DSICINAUDIO)
	AV_CODEC_ID_DSICINVIDEO       = int(C.AV_CODEC_ID_DSICINVIDEO)
	AV_CODEC_ID_DTS               = int(C.AV_CODEC_ID_DTS)
	AV_CODEC_ID_DVAUDIO           = int(C.AV_CODEC_ID_DVAUDIO)
	AV_CODEC_ID_DVB_SUBTITLE      = int(C.AV_CODEC_ID_DVB_SUBTITLE)
	AV_CODEC_ID_DVB_TELETEXT      = int(C.AV_CODEC_ID_DVB_TELETEXT)
	AV_CODEC_ID_DVD_NAV           = int(C.AV_CODEC_ID_DVD_NAV)
	AV_CODEC_ID_DVD_SUBTITLE      = int(C.AV_CODEC_ID_DVD_SUBTITLE)
	AV_CODEC_ID_DVVIDEO           = int(C.AV_CODEC_ID_DVVIDEO)
	AV_CODEC_ID_DXA               = int(C.AV_CODEC_ID_DXA)
	AV_CODEC_ID_DXTORY            = int(C.AV_CODEC_ID_DXTORY)
	AV_CODEC_ID_EAC3              = int(C.AV_CODEC_ID_EAC3)
	AV_CODEC_ID_EIA_608           = int(C.AV_CODEC_ID_EIA_608)
	AV_CODEC_ID_ESCAPE124         = int(C.AV_CODEC_ID_ESCAPE124)
	AV_CODEC_ID_ESCAPE130         = int(C.AV_CODEC_ID_ESCAPE130)
	AV_CODEC_ID_EVRC              = int(C.AV_CODEC_ID_EVRC)
	AV_CODEC_ID_EXR               = int(C.AV_CODEC_ID_EXR)
	AV_CODEC_ID_FFMETADATA        = int(C.AV_CODEC_ID_FFMETADATA)
	AV_CODEC_ID_FFV1              = int(C.AV_CODEC_ID_FFV1)
	AV_CODEC_ID_FFVHUFF           = int(C.AV_CODEC_ID_FFVHUFF)
	AV_CODEC_ID_FFWAVESYNTH       = int(C.AV_CODEC_ID_FFWAVESYNTH)
	AV_CODEC_ID_FIC               = int(C.AV_CODEC_ID_FIC)
	AV_CODEC_ID_FIRST_AUDIO       = int(C.AV_CODEC_ID_FIRST_AUDIO)
	AV_CODEC_ID_FIRST_SUBTITLE    = int(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AV_CODEC_ID_FIRST_UNKNOWN     = int(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AV_CODEC_ID_FLAC              = int(C.AV_CODEC_ID_FLAC)
	AV_CODEC_ID_FLASHSV           = int(C.AV_CODEC_ID_FLASHSV)
	AV_CODEC_ID_FLASHSV2          = int(C.AV_CODEC_ID_FLASHSV2)
	AV_CODEC_ID_FLIC              = int(C.AV_CODEC_ID_FLIC)
	AV_CODEC_ID_FLV1              = int(C.AV_CODEC_ID_FLV1)
	AV_CODEC_ID_FRAPS             = int(C.AV_CODEC_ID_FRAPS)
	AV_CODEC_ID_FRWU              = int(C.AV_CODEC_ID_FRWU)
	AV_CODEC_ID_G2M               = int(C.AV_CODEC_ID_G2M)
	AV_CODEC_ID_G723_1            = int(C.AV_CODEC_ID_G723_1)
	AV_CODEC_ID_G729              = int(C.AV_CODEC_ID_G729)
	AV_CODEC_ID_GIF               = int(C.AV_CODEC_ID_GIF)
	AV_CODEC_ID_GSM               = int(C.AV_CODEC_ID_GSM)
	AV_CODEC_ID_GSM_MS            = int(C.AV_CODEC_ID_GSM_MS)
	AV_CODEC_ID_H261              = int(C.AV_CODEC_ID_H261)
	AV_CODEC_ID_H263              = int(C.AV_CODEC_ID_H263)
	AV_CODEC_ID_H263I             = int(C.AV_CODEC_ID_H263I)
	AV_CODEC_ID_H263P             = int(C.AV_CODEC_ID_H263P)
	AV_CODEC_ID_H264              = int(C.AV_CODEC_ID_H264)
	AV_CODEC_ID_HDMV_PGS_SUBTITLE = int(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AV_CODEC_ID_HEVC              = int(C.AV_CODEC_ID_HEVC)
	AV_CODEC_ID_HNM4_VIDEO        = int(C.AV_CODEC_ID_HNM4_VIDEO)
	AV_CODEC_ID_HUFFYUV           = int(C.AV_CODEC_ID_HUFFYUV)
	AV_CODEC_ID_IAC               = int(C.AV_CODEC_ID_IAC)
	AV_CODEC_ID_IDCIN             = int(C.AV_CODEC_ID_IDCIN)
	AV_CODEC_ID_IDF               = int(C.AV_CODEC_ID_IDF)
	AV_CODEC_ID_IFF_BYTERUN1      = int(C.AV_CODEC_ID_IFF_BYTERUN1)
	AV_CODEC_ID_IFF_ILBM          = int(C.AV_CODEC_ID_IFF_ILBM)
	AV_CODEC_ID_ILBC              = int(C.AV_CODEC_ID_ILBC)
	AV_CODEC_ID_IMC               = int(C.AV_CODEC_ID_IMC)
	AV_CODEC_ID_INDEO2            = int(C.AV_CODEC_ID_INDEO2)
	AV_CODEC_ID_INDEO3            = int(C.AV_CODEC_ID_INDEO3)
	AV_CODEC_ID_INDEO4            = int(C.AV_CODEC_ID_INDEO4)
	AV_CODEC_ID_INDEO5            = int(C.AV_CODEC_ID_INDEO5)
	AV_CODEC_ID_INTERPLAY_DPCM    = int(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AV_CODEC_ID_INTERPLAY_VIDEO   = int(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AV_CODEC_ID_JACOSUB           = int(C.AV_CODEC_ID_JACOSUB)
	AV_CODEC_ID_JPEG2000          = int(C.AV_CODEC_ID_JPEG2000)
	AV_CODEC_ID_JPEGLS            = int(C.AV_CODEC_ID_JPEGLS)
	AV_CODEC_ID_JV                = int(C.AV_CODEC_ID_JV)
	AV_CODEC_ID_KGV1              = int(C.AV_CODEC_ID_KGV1)
	AV_CODEC_ID_KMVC              = int(C.AV_CODEC_ID_KMVC)
	AV_CODEC_ID_LAGARITH          = int(C.AV_CODEC_ID_LAGARITH)
	AV_CODEC_ID_LJPEG             = int(C.AV_CODEC_ID_LJPEG)
	AV_CODEC_ID_LOCO              = int(C.AV_CODEC_ID_LOCO)
	AV_CODEC_ID_MACE3             = int(C.AV_CODEC_ID_MACE3)
	AV_CODEC_ID_MACE6             = int(C.AV_CODEC_ID_MACE6)
	AV_CODEC_ID_MAD               = int(C.AV_CODEC_ID_MAD)
	AV_CODEC_ID_MDEC              = int(C.AV_CODEC_ID_MDEC)
	AV_CODEC_ID_METASOUND         = int(C.AV_CODEC_ID_METASOUND)
	AV_CODEC_ID_MICRODVD          = int(C.AV_CODEC_ID_MICRODVD)
	AV_CODEC_ID_MIMIC             = int(C.AV_CODEC_ID_MIMIC)
	AV_CODEC_ID_MJPEG             = int(C.AV_CODEC_ID_MJPEG)
	AV_CODEC_ID_MJPEGB            = int(C.AV_CODEC_ID_MJPEGB)
	AV_CODEC_ID_MLP               = int(C.AV_CODEC_ID_MLP)
	AV_CODEC_ID_MMVIDEO           = int(C.AV_CODEC_ID_MMVIDEO)
	AV_CODEC_ID_MOTIONPIXELS      = int(C.AV_CODEC_ID_MOTIONPIXELS)
	AV_CODEC_ID_MOV_TEXT          = int(C.AV_CODEC_ID_MOV_TEXT)
	AV_CODEC_ID_MP1               = int(C.AV_CODEC_ID_MP1)
	AV_CODEC_ID_MP2               = int(C.AV_CODEC_ID_MP2)
	AV_CODEC_ID_MP3               = int(C.AV_CODEC_ID_MP3)
	AV_CODEC_ID_MP3ADU            = int(C.AV_CODEC_ID_MP3ADU)
	AV_CODEC_ID_MP3ON4            = int(C.AV_CODEC_ID_MP3ON4)
	AV_CODEC_ID_MP4ALS            = int(C.AV_CODEC_ID_MP4ALS)
	AV_CODEC_ID_MPEG1VIDEO        = int(C.AV_CODEC_ID_MPEG1VIDEO)
	AV_CODEC_ID_MPEG2TS           = int(C.AV_CODEC_ID_MPEG2TS)
	AV_CODEC_ID_MPEG2VIDEO        = int(C.AV_CODEC_ID_MPEG2VIDEO)
	AV_CODEC_ID_MPEG2VIDEO_XVMC   = int(C.AV_CODEC_ID_MPEG2VIDEO_XVMC)
	AV_CODEC_ID_MPEG4             = int(C.AV_CODEC_ID_MPEG4)
	AV_CODEC_ID_MPEG4SYSTEMS      = int(C.AV_CODEC_ID_MPEG4SYSTEMS)
	AV_CODEC_ID_MPL2              = int(C.AV_CODEC_ID_MPL2)
	AV_CODEC_ID_MSA1              = int(C.AV_CODEC_ID_MSA1)
	AV_CODEC_ID_MSMPEG4V1         = int(C.AV_CODEC_ID_MSMPEG4V1)
	AV_CODEC_ID_MSMPEG4V2         = int(C.AV_CODEC_ID_MSMPEG4V2)
	AV_CODEC_ID_MSMPEG4V3         = int(C.AV_CODEC_ID_MSMPEG4V3)
	AV_CODEC_ID_MSRLE             = int(C.AV_CODEC_ID_MSRLE)
	AV_CODEC_ID_MSS1              = int(C.AV_CODEC_ID_MSS1)
	AV_CODEC_ID_MSS2              = int(C.AV_CODEC_ID_MSS2)
	AV_CODEC_ID_MSVIDEO1          = int(C.AV_CODEC_ID_MSVIDEO1)
	AV_CODEC_ID_MSZH              = int(C.AV_CODEC_ID_MSZH)
	AV_CODEC_ID_MTS2              = int(C.AV_CODEC_ID_MTS2)
	AV_CODEC_ID_MUSEPACK7         = int(C.AV_CODEC_ID_MUSEPACK7)
	AV_CODEC_ID_MUSEPACK8         = int(C.AV_CODEC_ID_MUSEPACK8)
	AV_CODEC_ID_MVC1              = int(C.AV_CODEC_ID_MVC1)
	AV_CODEC_ID_MVC2              = int(C.AV_CODEC_ID_MVC2)
	AV_CODEC_ID_MXPEG             = int(C.AV_CODEC_ID_MXPEG)
	AV_CODEC_ID_NELLYMOSER        = int(C.AV_CODEC_ID_NELLYMOSER)
	AV_CODEC_ID_NONE              = int(C.AV_CODEC_ID_NONE)
	AV_CODEC_ID_NUV               = int(C.AV_CODEC_ID_NUV)
	AV_CODEC_ID_ON2AVC            = int(C.AV_CODEC_ID_ON2AVC)
	AV_CODEC_ID_OPUS              = int(C.AV_CODEC_ID_OPUS)
	AV_CODEC_ID_OTF               = int(C.AV_CODEC_ID_OTF)
	AV_CODEC_ID_PAF_AUDIO         = int(C.AV_CODEC_ID_PAF_AUDIO)
	AV_CODEC_ID_PAF_VIDEO         = int(C.AV_CODEC_ID_PAF_VIDEO)
	AV_CODEC_ID_PAM               = int(C.AV_CODEC_ID_PAM)
	AV_CODEC_ID_PBM               = int(C.AV_CODEC_ID_PBM)
	AV_CODEC_ID_PCM_ALAW          = int(C.AV_CODEC_ID_PCM_ALAW)
	AV_CODEC_ID_PCM_BLURAY        = int(C.AV_CODEC_ID_PCM_BLURAY)
	AV_CODEC_ID_PCM_DVD           = int(C.AV_CODEC_ID_PCM_DVD)
	AV_CODEC_ID_PCM_F32BE         = int(C.AV_CODEC_ID_PCM_F32BE)
	AV_CODEC_ID_PCM_F32LE         = int(C.AV_CODEC_ID_PCM_F32LE)
	AV_CODEC_ID_PCM_F64BE         = int(C.AV_CODEC_ID_PCM_F64BE)
	AV_CODEC_ID_PCM_F64LE         = int(C.AV_CODEC_ID_PCM_F64LE)
	AV_CODEC_ID_PCM_LXF           = int(C.AV_CODEC_ID_PCM_LXF)
	AV_CODEC_ID_PCM_MULAW         = int(C.AV_CODEC_ID_PCM_MULAW)
	AV_CODEC_ID_PCM_S16BE         = int(C.AV_CODEC_ID_PCM_S16BE)
	AV_CODEC_ID_PCM_S16BE_PLANAR  = int(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AV_CODEC_ID_PCM_S16LE         = int(C.AV_CODEC_ID_PCM_S16LE)
	AV_CODEC_ID_PCM_S16LE_PLANAR  = int(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AV_CODEC_ID_PCM_S24BE         = int(C.AV_CODEC_ID_PCM_S24BE)
	AV_CODEC_ID_PCM_S24DAUD       = int(C.AV_CODEC_ID_PCM_S24DAUD)
	AV_CODEC_ID_PCM_S24LE         = int(C.AV_CODEC_ID_PCM_S24LE)
	AV_CODEC_ID_PCM_S24LE_PLANAR  = int(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AV_CODEC_ID_PCM_S32BE         = int(C.AV_CODEC_ID_PCM_S32BE)
	AV_CODEC_ID_PCM_S32LE         = int(C.AV_CODEC_ID_PCM_S32LE)
	AV_CODEC_ID_PCM_S32LE_PLANAR  = int(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AV_CODEC_ID_PCM_S8            = int(C.AV_CODEC_ID_PCM_S8)
	AV_CODEC_ID_PCM_S8_PLANAR     = int(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AV_CODEC_ID_PCM_U16BE         = int(C.AV_CODEC_ID_PCM_U16BE)
	AV_CODEC_ID_PCM_U16LE         = int(C.AV_CODEC_ID_PCM_U16LE)
	AV_CODEC_ID_PCM_U24BE         = int(C.AV_CODEC_ID_PCM_U24BE)
	AV_CODEC_ID_PCM_U24LE         = int(C.AV_CODEC_ID_PCM_U24LE)
	AV_CODEC_ID_PCM_U32BE         = int(C.AV_CODEC_ID_PCM_U32BE)
	AV_CODEC_ID_PCM_U32LE         = int(C.AV_CODEC_ID_PCM_U32LE)
	AV_CODEC_ID_PCM_U8            = int(C.AV_CODEC_ID_PCM_U8)
	AV_CODEC_ID_PCM_ZORK          = int(C.AV_CODEC_ID_PCM_ZORK)
	AV_CODEC_ID_PCX               = int(C.AV_CODEC_ID_PCX)
	AV_CODEC_ID_PGM               = int(C.AV_CODEC_ID_PGM)
	AV_CODEC_ID_PGMYUV            = int(C.AV_CODEC_ID_PGMYUV)
	AV_CODEC_ID_PICTOR            = int(C.AV_CODEC_ID_PICTOR)
	AV_CODEC_ID_PJS               = int(C.AV_CODEC_ID_PJS)
	AV_CODEC_ID_PNG               = int(C.AV_CODEC_ID_PNG)
	AV_CODEC_ID_PPM               = int(C.AV_CODEC_ID_PPM)
	AV_CODEC_ID_PROBE             = int(C.AV_CODEC_ID_PROBE)
	AV_CODEC_ID_PRORES            = int(C.AV_CODEC_ID_PRORES)
	AV_CODEC_ID_PTX               = int(C.AV_CODEC_ID_PTX)
	AV_CODEC_ID_QCELP             = int(C.AV_CODEC_ID_QCELP)
	AV_CODEC_ID_QDM2              = int(C.AV_CODEC_ID_QDM2)
	AV_CODEC_ID_QDMC              = int(C.AV_CODEC_ID_QDMC)
	AV_CODEC_ID_QDRAW             = int(C.AV_CODEC_ID_QDRAW)
	AV_CODEC_ID_QPEG              = int(C.AV_CODEC_ID_QPEG)
	AV_CODEC_ID_QTRLE             = int(C.AV_CODEC_ID_QTRLE)
	AV_CODEC_ID_R10K              = int(C.AV_CODEC_ID_R10K)
	AV_CODEC_ID_R210              = int(C.AV_CODEC_ID_R210)
	AV_CODEC_ID_RALF              = int(C.AV_CODEC_ID_RALF)
	AV_CODEC_ID_RAWVIDEO          = int(C.AV_CODEC_ID_RAWVIDEO)
	AV_CODEC_ID_RA_144            = int(C.AV_CODEC_ID_RA_144)
	AV_CODEC_ID_RA_288            = int(C.AV_CODEC_ID_RA_288)
	AV_CODEC_ID_REALTEXT          = int(C.AV_CODEC_ID_REALTEXT)
	AV_CODEC_ID_RL2               = int(C.AV_CODEC_ID_RL2)
	AV_CODEC_ID_ROQ               = int(C.AV_CODEC_ID_ROQ)
	AV_CODEC_ID_ROQ_DPCM          = int(C.AV_CODEC_ID_ROQ_DPCM)
	AV_CODEC_ID_RPZA              = int(C.AV_CODEC_ID_RPZA)
	AV_CODEC_ID_RV10              = int(C.AV_CODEC_ID_RV10)
	AV_CODEC_ID_RV20              = int(C.AV_CODEC_ID_RV20)
	AV_CODEC_ID_RV30              = int(C.AV_CODEC_ID_RV30)
	AV_CODEC_ID_RV40              = int(C.AV_CODEC_ID_RV40)
	AV_CODEC_ID_S302M             = int(C.AV_CODEC_ID_S302M)
	AV_CODEC_ID_SAMI              = int(C.AV_CODEC_ID_SAMI)
	AV_CODEC_ID_SANM              = int(C.AV_CODEC_ID_SANM)
	AV_CODEC_ID_SGI               = int(C.AV_CODEC_ID_SGI)
	AV_CODEC_ID_SGIRLE            = int(C.AV_CODEC_ID_SGIRLE)
	AV_CODEC_ID_SHORTEN           = int(C.AV_CODEC_ID_SHORTEN)
	AV_CODEC_ID_SIPR              = int(C.AV_CODEC_ID_SIPR)
	AV_CODEC_ID_SMACKAUDIO        = int(C.AV_CODEC_ID_SMACKAUDIO)
	AV_CODEC_ID_SMACKVIDEO        = int(C.AV_CODEC_ID_SMACKVIDEO)
	AV_CODEC_ID_SMC               = int(C.AV_CODEC_ID_SMC)
	AV_CODEC_ID_SMPTE_KLV         = int(C.AV_CODEC_ID_SMPTE_KLV)
	AV_CODEC_ID_SMV               = int(C.AV_CODEC_ID_SMV)
	AV_CODEC_ID_SMVJPEG           = int(C.AV_CODEC_ID_SMVJPEG)
	AV_CODEC_ID_SNOW              = int(C.AV_CODEC_ID_SNOW)
	AV_CODEC_ID_SOL_DPCM          = int(C.AV_CODEC_ID_SOL_DPCM)
	AV_CODEC_ID_SONIC             = int(C.AV_CODEC_ID_SONIC)
	AV_CODEC_ID_SONIC_LS          = int(C.AV_CODEC_ID_SONIC_LS)
	AV_CODEC_ID_SP5X              = int(C.AV_CODEC_ID_SP5X)
	AV_CODEC_ID_SPEEX             = int(C.AV_CODEC_ID_SPEEX)
	AV_CODEC_ID_SRT               = int(C.AV_CODEC_ID_SRT)
	AV_CODEC_ID_SSA               = int(C.AV_CODEC_ID_SSA)
	AV_CODEC_ID_SUBRIP            = int(C.AV_CODEC_ID_SUBRIP)
	AV_CODEC_ID_SUBVIEWER         = int(C.AV_CODEC_ID_SUBVIEWER)
	AV_CODEC_ID_SUBVIEWER1        = int(C.AV_CODEC_ID_SUBVIEWER1)
	AV_CODEC_ID_SUNRAST           = int(C.AV_CODEC_ID_SUNRAST)
	AV_CODEC_ID_SVQ1              = int(C.AV_CODEC_ID_SVQ1)
	AV_CODEC_ID_SVQ3              = int(C.AV_CODEC_ID_SVQ3)
	AV_CODEC_ID_TAK               = int(C.AV_CODEC_ID_TAK)
	AV_CODEC_ID_TARGA             = int(C.AV_CODEC_ID_TARGA)
	AV_CODEC_ID_TARGA_Y216        = int(C.AV_CODEC_ID_TARGA_Y216)
	AV_CODEC_ID_TEXT              = int(C.AV_CODEC_ID_TEXT)
	AV_CODEC_ID_TGQ               = int(C.AV_CODEC_ID_TGQ)
	AV_CODEC_ID_TGV               = int(C.AV_CODEC_ID_TGV)
	AV_CODEC_ID_THEORA            = int(C.AV_CODEC_ID_THEORA)
	AV_CODEC_ID_THP               = int(C.AV_CODEC_ID_THP)
	AV_CODEC_ID_TIERTEXSEQVIDEO   = int(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AV_CODEC_ID_TIFF              = int(C.AV_CODEC_ID_TIFF)
	AV_CODEC_ID_TIMED_ID3         = int(C.AV_CODEC_ID_TIMED_ID3)
	AV_CODEC_ID_TMV               = int(C.AV_CODEC_ID_TMV)
	AV_CODEC_ID_TQI               = int(C.AV_CODEC_ID_TQI)
	AV_CODEC_ID_TRUEHD            = int(C.AV_CODEC_ID_TRUEHD)
	AV_CODEC_ID_TRUEMOTION1       = int(C.AV_CODEC_ID_TRUEMOTION1)
	AV_CODEC_ID_TRUEMOTION2       = int(C.AV_CODEC_ID_TRUEMOTION2)
	AV_CODEC_ID_TRUESPEECH        = int(C.AV_CODEC_ID_TRUESPEECH)
	AV_CODEC_ID_TSCC              = int(C.AV_CODEC_ID_TSCC)
	AV_CODEC_ID_TSCC2             = int(C.AV_CODEC_ID_TSCC2)
	AV_CODEC_ID_TTA               = int(C.AV_CODEC_ID_TTA)
	AV_CODEC_ID_TTF               = int(C.AV_CODEC_ID_TTF)
	AV_CODEC_ID_TWINVQ            = int(C.AV_CODEC_ID_TWINVQ)
	AV_CODEC_ID_TXD               = int(C.AV_CODEC_ID_TXD)
	AV_CODEC_ID_ULTI              = int(C.AV_CODEC_ID_ULTI)
	AV_CODEC_ID_UTVIDEO           = int(C.AV_CODEC_ID_UTVIDEO)
	AV_CODEC_ID_V210              = int(C.AV_CODEC_ID_V210)
	AV_CODEC_ID_V210X             = int(C.AV_CODEC_ID_V210X)
	AV_CODEC_ID_V308              = int(C.AV_CODEC_ID_V308)
	AV_CODEC_ID_V408              = int(C.AV_CODEC_ID_V408)
	AV_CODEC_ID_V410              = int(C.AV_CODEC_ID_V410)
	AV_CODEC_ID_VB                = int(C.AV_CODEC_ID_VB)
	AV_CODEC_ID_VBLE              = int(C.AV_CODEC_ID_VBLE)
	AV_CODEC_ID_VC1               = int(C.AV_CODEC_ID_VC1)
	AV_CODEC_ID_VC1IMAGE          = int(C.AV_CODEC_ID_VC1IMAGE)
	AV_CODEC_ID_VCR1              = int(C.AV_CODEC_ID_VCR1)
	AV_CODEC_ID_VIMA              = int(C.AV_CODEC_ID_VIMA)
	AV_CODEC_ID_VIXL              = int(C.AV_CODEC_ID_VIXL)
	AV_CODEC_ID_VMDAUDIO          = int(C.AV_CODEC_ID_VMDAUDIO)
	AV_CODEC_ID_VMDVIDEO          = int(C.AV_CODEC_ID_VMDVIDEO)
	AV_CODEC_ID_VMNC              = int(C.AV_CODEC_ID_VMNC)
	AV_CODEC_ID_VORBIS            = int(C.AV_CODEC_ID_VORBIS)
	AV_CODEC_ID_VOXWARE           = int(C.AV_CODEC_ID_VOXWARE)
	AV_CODEC_ID_VP3               = int(C.AV_CODEC_ID_VP3)
	AV_CODEC_ID_VP5               = int(C.AV_CODEC_ID_VP5)
	AV_CODEC_ID_VP6               = int(C.AV_CODEC_ID_VP6)
	AV_CODEC_ID_VP6A              = int(C.AV_CODEC_ID_VP6A)
	AV_CODEC_ID_VP6F              = int(C.AV_CODEC_ID_VP6F)
	AV_CODEC_ID_VP7               = int(C.AV_CODEC_ID_VP7)
	AV_CODEC_ID_VP8               = int(C.AV_CODEC_ID_VP8)
	AV_CODEC_ID_VP9               = int(C.AV_CODEC_ID_VP9)
	AV_CODEC_ID_VPLAYER           = int(C.AV_CODEC_ID_VPLAYER)
	AV_CODEC_ID_WAVPACK           = int(C.AV_CODEC_ID_WAVPACK)
	AV_CODEC_ID_WEBP              = int(C.AV_CODEC_ID_WEBP)
	AV_CODEC_ID_WEBVTT            = int(C.AV_CODEC_ID_WEBVTT)
	AV_CODEC_ID_WESTWOOD_SND1     = int(C.AV_CODEC_ID_WESTWOOD_SND1)
	AV_CODEC_ID_WMALOSSLESS       = int(C.AV_CODEC_ID_WMALOSSLESS)
	AV_CODEC_ID_WMAPRO            = int(C.AV_CODEC_ID_WMAPRO)
	AV_CODEC_ID_WMAV1             = int(C.AV_CODEC_ID_WMAV1)
	AV_CODEC_ID_WMAV2             = int(C.AV_CODEC_ID_WMAV2)
	AV_CODEC_ID_WMAVOICE          = int(C.AV_CODEC_ID_WMAVOICE)
	AV_CODEC_ID_WMV1              = int(C.AV_CODEC_ID_WMV1)
	AV_CODEC_ID_WMV2              = int(C.AV_CODEC_ID_WMV2)
	AV_CODEC_ID_WMV3              = int(C.AV_CODEC_ID_WMV3)
	AV_CODEC_ID_WMV3IMAGE         = int(C.AV_CODEC_ID_WMV3IMAGE)
	AV_CODEC_ID_WNV1              = int(C.AV_CODEC_ID_WNV1)
	AV_CODEC_ID_WS_VQA            = int(C.AV_CODEC_ID_WS_VQA)
	AV_CODEC_ID_XAN_DPCM          = int(C.AV_CODEC_ID_XAN_DPCM)
	AV_CODEC_ID_XAN_WC3           = int(C.AV_CODEC_ID_XAN_WC3)
	AV_CODEC_ID_XAN_WC4           = int(C.AV_CODEC_ID_XAN_WC4)
	AV_CODEC_ID_XBIN              = int(C.AV_CODEC_ID_XBIN)
	AV_CODEC_ID_XBM               = int(C.AV_CODEC_ID_XBM)
	AV_CODEC_ID_XFACE             = int(C.AV_CODEC_ID_XFACE)
	AV_CODEC_ID_XSUB              = int(C.AV_CODEC_ID_XSUB)
	AV_CODEC_ID_XWD               = int(C.AV_CODEC_ID_XWD)
	AV_CODEC_ID_Y41P              = int(C.AV_CODEC_ID_Y41P)
	AV_CODEC_ID_YOP               = int(C.AV_CODEC_ID_YOP)
	AV_CODEC_ID_YUV4              = int(C.AV_CODEC_ID_YUV4)
	AV_CODEC_ID_ZEROCODEC         = int(C.AV_CODEC_ID_ZEROCODEC)
	AV_CODEC_ID_ZLIB              = int(C.AV_CODEC_ID_ZLIB)
	AV_CODEC_ID_ZMBV              = int(C.AV_CODEC_ID_ZMBV)
)
