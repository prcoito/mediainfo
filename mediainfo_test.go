package mediainfo

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

var moduleFolder = func() string { w, _ := os.Getwd(); return w }()

func TestMediaInfo(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name    string
		args    args
		want    Info
		wantErr bool
	}{
		{
			name: "1 video 1 audio 1 menu",
			args: args{filepath.Join("testdata", `1_video_1_audio_1_menu.mkv`)},
			want: Info{
				General: General{
					CompleteName:       filepath.Join(moduleFolder, "testdata", `1_video_1_audio_1_menu.mkv`),
					UniqueID:           "96343877327587173645643017013582379057",
					VideoCount:         1,
					AudioCount:         1,
					MenuCount:          1,
					FileExtension:      "mkv",
					Format:             "Matroska",
					FormatVersion:      "4",
					FileSize:           8980,
					OverallBitRate:     18,
					FrameRate:          23.976,
					FrameCount:         0,
					IsStreamable:       true,
					Title:              "This is the title",
					EncodedDate:        time.Date(2020, 10, 20, 17, 38, 13, 0, time.UTC),
					FileCreatedDate:    time.Date(2020, 11, 16, 17, 48, 13, 928000000, time.UTC),
					FileModifiedDate:   time.Date(2020, 11, 17, 17, 19, 54, 510000000, time.UTC),
					EncodedApplication: "mkvmerge v50.0.0 ('Awakenings') 64-bit",
					EncodedLibrary:     "libebml v1.4.0 + libmatroska v1.6.2",
					Duration:           4086.355,
				},
				VideoTracks: []Video{{
					StreamOrder:            0,
					ID:                     1,
					UniqueID:               "4801514838969937575",
					Format:                 "HEVC",
					FormatProfile:          "Main 10",
					FormatLevel:            "4",
					FormatTier:             "Main",
					CodecID:                "V_MPEGH/ISO/HEVC",
					Duration:               4086.355000000,
					BitRate:                0,
					Width:                  1920,
					Height:                 1080,
					SampledWidth:           1920,
					SampledHeight:          1080,
					PixelAspectRatio:       1.000,
					DisplayAspectRatio:     1.778,
					FrameRateMode:          "CFR",
					FrameRate:              23.976,
					FrameCount:             0,
					ColorSpace:             "YUV",
					ChromaSubsampling:      "4:2:0",
					BitDepth:               10,
					StreamSize:             0,
					EncodedLibrary:         "x265 - 2.5+48-bd438ce10843:[Windows][MSVC 1911][64 bit] 10bit",
					EncodedLibraryName:     "x265",
					EncodedLibraryVersion:  "2.5+48-bd438ce10843:[Windows][MSVC 1911][64 bit] 10bit",
					EncodedLibrarySettings: "cpuid=1173503 / frame-threads=3 / numa-pools=8 / wpp / no-pmode / no-pme / no-psnr / no-ssim / log-level=2 / input-csp=1 / input-res=1920x1080 / interlace=0 / total-frames=97974 / level-idc=0 / high-tier=1 / uhd-bd=0 / ref=5 / no-allow-non-conformance / no-repeat-headers / annexb / no-aud / no-hrd / info / hash=0 / no-temporal-layers / open-gop / min-keyint=23 / keyint=250 / bframes=8 / b-adapt=2 / b-pyramid / bframe-bias=0 / rc-lookahead=120 / lookahead-slices=4 / scenecut=40 / no-intra-refresh / ctu=64 / min-cu-size=8 / rect / no-amp / max-tu-size=32 / tu-inter-depth=1 / tu-intra-depth=1 / limit-tu=0 / rdoq-level=2 / dynamic-rd=0.00 / no-ssim-rd / signhide / no-tskip / nr-intra=0 / nr-inter=0 / no-constrained-intra / no-strong-intra-smoothing / max-merge=3 / limit-refs=3 / limit-modes / me=3 / subme=5 / merange=57 / temporal-mvp / weightp / no-weightb / no-analyze-src-pics / deblock=-3:-3 / no-sao / no-sao-non-deblock / rd=4 / no-early-skip / no-rskip / no-fast-intra / no-tskip-fast / no-cu-lossless / no-b-intra / no-splitrd-skip / rdpenalty=0 / psy-rd=1.20 / psy-rdoq=1.50 / no-rd-refine / analysis-reuse-mode=0 / no-lossless / cbqpoffs=0 / crqpoffs=0 / rc=crf / crf=20.0 / qcomp=0.60 / qpstep=4 / stats-write=0 / stats-read=0 / ipratio=1.40 / pbratio=1.30 / aq-mode=3 / aq-strength=0.90 / cutree / zone-count=0 / no-strict-cbr / qg-size=32 / no-rc-grain / qpmax=69 / qpmin=0 / no-const-vbv / sar=0 / overscan=0 / videoformat=5 / range=0 / colorprim=2 / transfer=2 / colormatrix=2 / chromaloc=0 / display-window=0 / max-cll=0,0 / min-luma=0 / max-luma=1023 / log2-max-poc-lsb=8 / vui-timing-info / vui-hrd-info / slices=1 / opt-qp-pps / opt-ref-list-length-pps / no-multi-pass-opt-rps / scenecut-bias=0.05 / no-opt-cu-delta-qp / no-aq-motion / no-hdr / no-hdr-opt / no-dhdr10-opt / analysis-reuse-level=5 / scale-factor=0 / refine-intra=0 / refine-inter=0 / refine-mv=0 / no-limit-sao / ctu-info=0 / no-lowpass-dct / refine-mv-type=0",
					Default:                true,
					Forced:                 false,
				}},
				AudioTracks: []Audio{{
					StreamOrder:              1,
					ID:                       2,
					UniqueID:                 "14945515745438299057",
					Format:                   "AAC",
					FormatAdditionalFeatures: "LC",
					CodecID:                  "A_AAC-2",
					Duration:                 0.000000000,
					BitRate:                  0,
					Channels:                 6,
					ChannelPositions:         "Front: L C R, Side: L R, LFE",
					ChannelLayout:            "C L R Ls Rs LFE",
					SamplesPerFrame:          1024,
					SamplingRate:             48000,
					SamplingCount:            196145040,
					FrameRate:                46.875,
					FrameCount:               0,
					CompressionMode:          "Lossy",
					StreamSize:               0,
					StreamSizeProportion:     0.00000,
					Language:                 "en",
					Default:                  true,
					Forced:                   false,
				}},
				MenuTracks: []Menu{{
					Duration: 4086.355,
					Entries: []Entry{{
						StartTime:    0,
						StartTimeStr: "00:00:00.000",
						EndTime:      107.607,
						EndTimeStr:   "00:01:47.607",
						Language:     "en",
						Title:        "Chapter 01",
					}, {
						StartTime:    107.607,
						StartTimeStr: "00:01:47.607",
						EndTime:      854.895,
						EndTimeStr:   "00:14:14.895",
						Language:     "en",
						Title:        "Chapter 02",
					}, {
						StartTime:    854.895,
						StartTimeStr: "00:14:14.895",
						EndTime:      1361.234,
						EndTimeStr:   "00:22:41.234",
						Language:     "en",
						Title:        "Chapter 03",
					}, {
						StartTime:    1361.234,
						StartTimeStr: "00:22:41.234",
						EndTime:      1994.450,
						EndTimeStr:   "00:33:14.450",
						Language:     "en",
						Title:        "Chapter 04",
					}, {
						StartTime:    1994.450,
						StartTimeStr: "00:33:14.450",
						EndTime:      2581.370,
						EndTimeStr:   "00:43:01.370",
						Language:     "en",
						Title:        "Chapter 05",
					}, {
						StartTime:    2581.370,
						StartTimeStr: "00:43:01.370",
						EndTime:      3284.781,
						EndTimeStr:   "00:54:44.781",
						Language:     "en",
						Title:        "Chapter 06",
					}, {
						StartTime:    3284.781,
						StartTimeStr: "00:54:44.781",
						EndTime:      4011.257,
						EndTimeStr:   "01:06:51.257",
						Language:     "en",
						Title:        "Chapter 07",
					}, {
						StartTime:    4011.257,
						StartTimeStr: "01:06:51.257",
						EndTime:      4086.355,
						EndTimeStr:   "01:08:06.355",
						Language:     "en",
						Title:        "Chapter 08",
					}},
				}},
			},
		}, {
			name: "very long path",
			args: args{filepath.Join("testdata", `this_is_a_folder_used_to_increase_path_len`, "this_is_a_file_with_a_very_very_very_very_very_veryvery_very_very_very_very_very_very_very_very_very_long_name_used_to_verify_if_max_path_limit_on_windows_does_not_break_file_access_and_mediainfo_access.mka")},
			want: Info{
				General: General{
					CompleteName:       filepath.Join(moduleFolder, "testdata", `this_is_a_folder_used_to_increase_path_len`, "this_is_a_file_with_a_very_very_very_very_very_veryvery_very_very_very_very_very_very_very_very_very_long_name_used_to_verify_if_max_path_limit_on_windows_does_not_break_file_access_and_mediainfo_access.mka"),
					UniqueID:           "318531224748573671368333994850674734764",
					VideoCount:         0,
					AudioCount:         1,
					MenuCount:          0,
					FileExtension:      "mka",
					Format:             "Matroska",
					FormatVersion:      "4",
					FileSize:           5821,
					OverallBitRate:     336,
					IsStreamable:       true,
					EncodedDate:        time.Date(2020, 10, 20, 19, 04, 07, 0, time.UTC),
					FileCreatedDate:    time.Date(2020, 11, 17, 17, 12, 56, 125000000, time.UTC),
					FileModifiedDate:   time.Date(2020, 11, 17, 13, 30, 11, 701000000, time.UTC),
					EncodedApplication: "mkvmerge v50.0.0 ('Awakenings') 64-bit",
					EncodedLibrary:     "libebml v1.4.0 + libmatroska v1.6.2",
					Duration:           138.396,
				},
				AudioTracks: []Audio{{
					StreamOrder:          0,
					ID:                   1,
					UniqueID:             "7418062013777177105",
					Format:               "MPEG Audio",
					CodecID:              "A_MPEG/L3",
					Duration:             138.396,
					BitRate:              0,
					Channels:             2,
					SamplingRate:         44100,
					SamplingCount:        6103264,
					FrameCount:           0,
					CompressionMode:      "Lossy",
					StreamSize:           0,
					StreamSizeProportion: 0.00000,
					Default:              true,
					Forced:               false,
				}},
			},
		}, {
			name: "path with unicode",
			args: args{filepath.Join("testdata", `audio_only_with_ùnicode_char.mka`)},
			want: Info{
				General: General{
					CompleteName:       filepath.Join(moduleFolder, "testdata", `audio_only_with_ùnicode_char.mka`),
					UniqueID:           "318531224748573671368333994850674734764",
					VideoCount:         0,
					AudioCount:         1,
					MenuCount:          0,
					FileExtension:      "mka",
					Format:             "Matroska",
					FormatVersion:      "4",
					FileSize:           5821,
					OverallBitRate:     336,
					IsStreamable:       true,
					EncodedDate:        time.Date(2020, 10, 20, 19, 04, 07, 0, time.UTC),
					FileCreatedDate:    time.Date(2020, 11, 16, 17, 48, 13, 902000000, time.UTC),
					FileModifiedDate:   time.Date(2020, 11, 17, 13, 30, 11, 701000000, time.UTC),
					EncodedApplication: "mkvmerge v50.0.0 ('Awakenings') 64-bit",
					EncodedLibrary:     "libebml v1.4.0 + libmatroska v1.6.2",
					Duration:           138.396,
				},
				AudioTracks: []Audio{{
					StreamOrder:          0,
					ID:                   1,
					UniqueID:             "7418062013777177105",
					Format:               "MPEG Audio",
					CodecID:              "A_MPEG/L3",
					Duration:             138.396,
					BitRate:              0,
					Channels:             2,
					SamplingRate:         44100,
					SamplingCount:        6103264,
					FrameCount:           0,
					CompressionMode:      "Lossy",
					StreamSize:           0,
					StreamSizeProportion: 0.00000,
					Default:              true,
					Forced:               false,
				}},
			},
		},
	}

	if !Load() {
		t.Fatalf("Failed to load dll/shared object")
	}
	defer Unload()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Inform(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("mi.MediaInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mi.MediaInfo()\nGot \n%+v\nWant\n%+v", got, tt.want)
			}
		})
	}
}

func Test_formatTime(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want float32
	}{
		{name: "trivial", arg: "00:00:00.000", want: 0},
		{name: "1h", arg: "01:00:00.000", want: 3600},
		{name: "1h1m", arg: "01:01:00.000", want: 3660},
		{name: "1h25m30s.675", arg: "01:25:30.675", want: 5130.675},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTime(tt.arg); got != tt.want {
				t.Errorf("formatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFormatTimeStr(t *testing.T) {
	tests := []struct {
		name string
		arg  float32
		want string
	}{
		{name: "trivial", arg: 0, want: "00:00:00.000"},
		{name: "two minutes and seconds", arg: 123.6, want: "00:02:03.600"},
		{name: "hour and minute and second", arg: 3685.005, want: "01:01:25.005"},
		{name: "hour and eight minutes and six second", arg: 4086.355, want: "01:08:06.355"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFormatTimeStr(tt.arg); got != tt.want {
				t.Errorf("toFormatTimeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toTime(t *testing.T) {
	tests := []struct {
		arg  string
		want time.Time
	}{
		{arg: "UTC 2020-10-20 17:38:13", want: time.Date(2020, 10, 20, 17, 38, 13, 0, time.UTC)},
		{arg: "UTC 2020-11-17 13:30:42.535", want: time.Date(2020, 11, 17, 13, 30, 42, 535000000, time.UTC)},
		{arg: "UTC 2020-11-16 17:48:13.928", want: time.Date(2020, 11, 16, 17, 48, 13, 928000000, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := toTime(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
