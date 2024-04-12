package mediainfo

import (
	"encoding/json"
	"os/exec"
)

type VideoInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                      string `json:"@type"`
			Count                     string `json:"Count"`
			StreamCount               string `json:"StreamCount"`
			StreamKind                string `json:"StreamKind"`
			StreamKindString          string `json:"StreamKind_String"`
			StreamKindID              string `json:"StreamKindID"`
			VideoCount                string `json:"VideoCount,omitempty"`
			AudioCount                string `json:"AudioCount,omitempty"`
			TextCount                 string `json:"TextCount,omitempty"`
			VideoFormatList           string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList   string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList            string `json:"Video_Codec_List,omitempty"`
			VideoLanguageList         string `json:"Video_Language_List,omitempty"`
			AudioFormatList           string `json:"Audio_Format_List,omitempty"`
			AudioFormatWithHintList   string `json:"Audio_Format_WithHint_List,omitempty"`
			AudioCodecList            string `json:"Audio_Codec_List,omitempty"`
			AudioLanguageList         string `json:"Audio_Language_List,omitempty"`
			AudioChannelsTotal        string `json:"Audio_Channels_Total,omitempty"`
			TextFormatList            string `json:"Text_Format_List,omitempty"`
			TextFormatWithHintList    string `json:"Text_Format_WithHint_List,omitempty"`
			TextCodecList             string `json:"Text_Codec_List,omitempty"`
			CompleteName              string `json:"CompleteName,omitempty"`
			FileNameExtension         string `json:"FileNameExtension,omitempty"`
			FileName                  string `json:"FileName,omitempty"`
			FileExtension             string `json:"FileExtension,omitempty"`
			Format                    string `json:"Format"`        // "HEVC"
			FormatString              string `json:"Format_String"` // "HEVC"
			FormatExtensions          string `json:"Format_Extensions,omitempty"`
			FormatCommercial          string `json:"Format_Commercial"`
			FormatProfile             string `json:"Format_Profile,omitempty"`
			InternetMediaType         string `json:"InternetMediaType,omitempty"`
			CodecID                   string `json:"CodecID"` // "hvc1"
			CodecIDString             string `json:"CodecID_String,omitempty"`
			CodecIDUrl                string `json:"CodecID_Url,omitempty"`
			CodecIDCompatible         string `json:"CodecID_Compatible,omitempty"`
			FileSize                  string `json:"FileSize,omitempty"`
			FileSizeString            string `json:"FileSize_String,omitempty"`
			FileSizeString1           string `json:"FileSize_String1,omitempty"`
			FileSizeString2           string `json:"FileSize_String2,omitempty"`
			FileSizeString3           string `json:"FileSize_String3,omitempty"`
			FileSizeString4           string `json:"FileSize_String4,omitempty"`
			Duration                  string `json:"Duration"`
			DurationString            string `json:"Duration_String"`
			DurationString1           string `json:"Duration_String1"`
			DurationString2           string `json:"Duration_String2"`
			DurationString3           string `json:"Duration_String3"`
			DurationString4           string `json:"Duration_String4,omitempty"`
			DurationString5           string `json:"Duration_String5"`
			OverallBitRateMode        string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRateModeString  string `json:"OverallBitRate_Mode_String,omitempty"`
			OverallBitRate            string `json:"OverallBitRate,omitempty"`
			OverallBitRateString      string `json:"OverallBitRate_String,omitempty"`
			FrameRate                 string `json:"FrameRate"`        // "46.875"
			FrameRateString           string `json:"FrameRate_String"` // "46.875 FPS (1024 SPF)"
			FrameCount                string `json:"FrameCount"`       // "345543" 帧数
			StreamSize                string `json:"StreamSize"`
			StreamSizeString          string `json:"StreamSize_String"`
			StreamSizeString1         string `json:"StreamSize_String1"`
			StreamSizeString2         string `json:"StreamSize_String2"`
			StreamSizeString3         string `json:"StreamSize_String3"`
			StreamSizeString4         string `json:"StreamSize_String4"`
			StreamSizeString5         string `json:"StreamSize_String5"`
			StreamSizeProportion      string `json:"StreamSize_Proportion"`
			HeaderSize                string `json:"HeaderSize,omitempty"`
			DataSize                  string `json:"DataSize,omitempty"`
			FooterSize                string `json:"FooterSize,omitempty"`
			IsStreamable              string `json:"IsStreamable,omitempty"`
			FileModifiedDate          string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal     string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication        string `json:"Encoded_Application,omitempty"`
			EncodedApplicationString  string `json:"Encoded_Application_String,omitempty"`
			StreamOrder               string `json:"StreamOrder,omitempty"`
			ID                        string `json:"ID,omitempty"`
			IDString                  string `json:"ID_String,omitempty"`
			FormatInfo                string `json:"Format_Info,omitempty"` // "High Efficiency Video Coding"
			FormatUrl                 string `json:"Format_Url,omitempty"`
			FormatLevel               string `json:"Format_Level,omitempty"`
			FormatTier                string `json:"Format_Tier,omitempty"`
			CodecIDInfo               string `json:"CodecID_Info,omitempty"`
			BitRate                   string `json:"BitRate,omitempty"`        // "69618"
			BitRateString             string `json:"BitRate_String,omitempty"` // "69.6 kb/s"
			Width                     string `json:"Width,omitempty"`          // "1920"
			WidthString               string `json:"Width_String,omitempty"`   // "1 920 pixels"
			Height                    string `json:"Height,omitempty"`         // "1080"
			HeightString              string `json:"Height_String,omitempty"`  // "1 080 pixels"
			SampledWidth              string `json:"Sampled_Width,omitempty"`
			SampledHeight             string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio          string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio        string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString  string `json:"DisplayAspectRatio_String,omitempty"`
			Rotation                  string `json:"Rotation,omitempty"`
			FrameRateMode             string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString       string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateNum              string `json:"FrameRate_Num,omitempty"`
			FrameRateDen              string `json:"FrameRate_Den,omitempty"`
			ColorSpace                string `json:"ColorSpace,omitempty"`
			ChromaSubsampling         string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString   string `json:"ChromaSubsampling_String,omitempty"`
			ChromaSubsamplingPosition string `json:"ChromaSubsampling_Position,omitempty"`
			BitDepth                  string `json:"BitDepth,omitempty"`
			BitDepthString            string `json:"BitDepth_String,omitempty"`
			ScanType                  string `json:"ScanType,omitempty"`
			ScanTypeString            string `json:"ScanType_String,omitempty"`
			BitsPixelFrame            string `json:"BitsPixel_Frame,omitempty"`
			EncodedLibrary            string `json:"Encoded_Library,omitempty"`
			EncodedLibraryString      string `json:"Encoded_Library_String,omitempty"`
			EncodedLibraryName        string `json:"Encoded_Library_Name,omitempty"`
			EncodedLibraryVersion     string `json:"Encoded_Library_Version,omitempty"`
			EncodedLibrarySettings    string `json:"Encoded_Library_Settings,omitempty"`
			Language                  string `json:"Language,omitempty"`
			LanguageString            string `json:"Language_String,omitempty"`
			LanguageString1           string `json:"Language_String1,omitempty"`
			LanguageString2           string `json:"Language_String2,omitempty"`
			LanguageString3           string `json:"Language_String3,omitempty"`
			LanguageString4           string `json:"Language_String4,omitempty"`
			ColourRange               string `json:"colour_range,omitempty"`
			ColourRangeSource         string `json:"colour_range_Source,omitempty"`
			Extra                     struct {
				CodecConfigurationBox string `json:"CodecConfigurationBox,omitempty"`
				SourceDelay           string `json:"Source_Delay,omitempty"`
				SourceDelaySource     string `json:"Source_Delay_Source,omitempty"`
			} `json:"extra,omitempty"`
			FormatSettingsSBR          string `json:"Format_Settings_SBR,omitempty"`
			FormatSettingsSBRString    string `json:"Format_Settings_SBR_String,omitempty"`
			FormatAdditionalFeatures   string `json:"Format_AdditionalFeatures,omitempty"`
			SourceDuration             string `json:"Source_Duration,omitempty"`
			SourceDurationString       string `json:"Source_Duration_String,omitempty"`
			SourceDurationString1      string `json:"Source_Duration_String1,omitempty"`
			SourceDurationString2      string `json:"Source_Duration_String2,omitempty"`
			SourceDurationString3      string `json:"Source_Duration_String3,omitempty"`
			SourceDurationString5      string `json:"Source_Duration_String5,omitempty"`
			BitRateMode                string `json:"BitRate_Mode,omitempty"`
			BitRateModeString          string `json:"BitRate_Mode_String,omitempty"`
			Channels                   string `json:"Channels,omitempty"`
			ChannelsString             string `json:"Channels_String,omitempty"`
			ChannelPositions           string `json:"ChannelPositions,omitempty"`
			ChannelPositionsString2    string `json:"ChannelPositions_String2,omitempty"`
			ChannelLayout              string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame            string `json:"SamplesPerFrame,omitempty"`
			SamplingRate               string `json:"SamplingRate,omitempty"`
			SamplingRateString         string `json:"SamplingRate_String,omitempty"`
			SamplingCount              string `json:"SamplingCount,omitempty"`
			SourceFrameCount           string `json:"Source_FrameCount,omitempty"`
			CompressionMode            string `json:"Compression_Mode,omitempty"`
			CompressionModeString      string `json:"Compression_Mode_String,omitempty"`
			SourceStreamSize           string `json:"Source_StreamSize,omitempty"`
			SourceStreamSizeString     string `json:"Source_StreamSize_String,omitempty"`
			SourceStreamSizeString1    string `json:"Source_StreamSize_String1,omitempty"`
			SourceStreamSizeString2    string `json:"Source_StreamSize_String2,omitempty"`
			SourceStreamSizeString3    string `json:"Source_StreamSize_String3,omitempty"`
			SourceStreamSizeString4    string `json:"Source_StreamSize_String4,omitempty"`
			SourceStreamSizeString5    string `json:"Source_StreamSize_String5,omitempty"`
			SourceStreamSizeProportion string `json:"Source_StreamSize_Proportion,omitempty"`
			Default                    string `json:"Default,omitempty"`
			DefaultString              string `json:"Default_String,omitempty"`
			AlternateGroup             string `json:"AlternateGroup,omitempty"`
			AlternateGroupString       string `json:"AlternateGroup_String,omitempty"`
			MuxingMode                 string `json:"MuxingMode,omitempty"`
			Forced                     string `json:"Forced,omitempty"`
			ForcedString               string `json:"Forced_String,omitempty"`
			EventsTotal                string `json:"Events_Total,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

func (bi *BasicInfo) SetVideoInfo(info VideoInfo) {
	bi.VInfo = info
}
func (bi *BasicInfo) InsertVideoInfo() {
	cmd := exec.Command("mediainfo", bi.FullPath, "--Full", "--Output=JSON")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	var v VideoInfo
	json.Unmarshal(output, &v)
	//fmt.Println(v)
	bi.VInfo = v
}
