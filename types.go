package mediainfo

import "time"

// inform is the struct that represents a call to MediaInfo.Inform
type informStruct struct {
	Media struct {
		Ref    string  `json:"@ref"`
		Tracks []track `json:"track"`
	} `json:"media"`
}

// Info represents the information returned from Inform
type Info struct {
	General     General
	VideoTracks []Video
	AudioTracks []Audio
	TextTracks  []Text
	MenuTracks  []Menu
}

// General represents the general track information present in Info
type General struct {
	UniqueID              string
	AudioCount            uint
	VideoCount            uint
	TextCount             uint
	MenuCount             uint
	FileExtension         string
	Format                string
	FormatVersion         string
	FileSize              uint
	Duration              float32
	OverallBitRate        float32
	FrameRate             float32
	FrameCount            uint
	IsStreamable          bool
	EncodedDate           time.Time
	FileCreatedDate       time.Time
	FileModifiedDate      time.Time
	EncodedApplication    string
	EncodedLibrary        string
	EncodedLibraryVersion string
	Title                 string
	CompleteName          string
}

// Video represents a video track information present in Info
type Video struct {
	StreamOrder            uint
	ID                     uint
	UniqueID               string
	Format                 string
	FormatProfile          string
	FormatLevel            string
	FormatTier             string
	CodecID                string
	Duration               float32
	BitRate                float32
	Width                  uint
	Height                 uint
	SampledWidth           uint
	SampledHeight          uint
	PixelAspectRatio       float32
	DisplayAspectRatio     float32
	FrameRateMode          string
	FrameRate              float32
	FrameCount             uint
	ColorSpace             string
	ChromaSubsampling      string
	BitDepth               uint
	StreamSize             uint
	StreamSizeProportion   float32
	EncodedLibrary         string
	EncodedLibraryName     string
	EncodedLibraryVersion  string
	EncodedLibrarySettings string
	Default                bool
	Forced                 bool
	B3D                    bool
	Title                  string
}

// Audio represents a audio track information present in Info
type Audio struct {
	StreamOrder              uint
	ID                       uint
	UniqueID                 string
	Format                   string
	FormatCommercial         string
	FormatAdditionalFeatures string
	CodecID                  string
	Duration                 float32
	BitRate                  float32
	Channels                 uint
	ChannelPositions         string
	ChannelLayout            string
	SamplesPerFrame          uint
	SamplingRate             uint
	SamplingCount            uint
	FrameRate                float32
	FrameCount               uint
	CompressionMode          string
	StreamSize               uint
	StreamSizeProportion     float32
	Language                 string
	Default                  bool
	Forced                   bool
	Title                    string
}

// Text represents a text track (subtitles) information present in Info
type Text struct {
	Order        uint
	StreamOrder  uint
	ID           uint
	UniqueID     string
	Format       string
	CodecID      string
	Duration     float32
	BitRate      float32
	FrameCount   uint
	ElementCount uint
	StreamSize   uint
	Language     string
	Default      bool
	Forced       bool
	Title        string
}

// Menu represents the Menu track (also known as Chapter) present in Info
type Menu struct {
	Order    uint
	Entries  []Entry
	Duration float32
}

// Entry represents an entry in Menu.Entries
type Entry struct {
	StartTime    float32
	StartTimeStr string
	EndTime      float32
	EndTimeStr   string
	Title        string
	Language     string
}

// track struct represent a media track
type track struct {
	Type                  string `json:"@type"`
	UniqueID              string
	VideoCount            string
	AudioCount            string
	TextCount             string
	MenuCount             string
	FileExtension         string
	Format                string
	FormatVersion         string `json:"Format_Version"`
	CodecID               string
	FileSize              string
	Duration              string
	OverallBitRate        string
	FrameRate             string
	FrameCount            string
	StreamSize            string
	IsStreamable          string
	Title                 string
	Movie                 string
	FileCreatedDate       string      `json:"File_Created_Date"`
	EncodedDate           string      `json:"Encoded_Date"`
	FileModifiedDate      string      `json:"File_Modified_Date"`
	FileModifiedDateLocal string      `json:"File_Modified_Date_Local"`
	EncodedApplication    string      `json:"Encoded_Application"`
	EncodedLibrary        interface{} `json:"Encoded_Library"` // some encoded libraries are json objects
	EncodedLibraryVersion string      `json:"Encoded_Library_Version"`

	StreamOrder        string
	ID                 string
	FormatProfile      string `json:"Format_Profile"`
	FormatLevel        string `json:"Format_Level"`
	FormatTier         string `json:"Format_Tier"`
	FormatCommercial   string `json:"Format_Commercial_IfAny"`
	Width              string
	Height             string
	SampledWidth       string `json:"Sampled_Width"`
	SampledHeight      string `json:"Sampled_Height"`
	PixelAspectRatio   string
	DisplayAspectRatio string
	FrameRateMode      string `json:"FrameRate_Mode"`
	MultiViewCount     string `json:"Multi_View_Count"`

	ColorSpace             string
	ChromaSubsampling      string
	BitDepth               string
	Delay                  string
	EncodedLibraryName     string `json:"Encoded_Library_Name"`
	EncodedLibrarySettings string `json:"Encoded_Library_Settings"`
	Language               string
	Default                string
	Forced                 string

	FormatAdditionalFeatures string `json:"Format_AdditionalFeatures"`
	BitRate                  string
	Channels                 string
	ChannelPositions         string
	ChannelLayout            string
	SamplesPerFrame          string
	SamplingRate             string
	SamplingCount            string
	CompressionMode          string `json:"Compression_Mode"`
	DelaySource              string `json:"Delay_Source"`
	StreamSizeProportion     string `json:"StreamSize_Proportion"`

	ElementCount string
	TypeOrder    string `json:"@typeorder"`

	Extra map[string]string
}
