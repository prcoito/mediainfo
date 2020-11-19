package mediainfo

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Inform returns the Media details (struct Info) from file f
func Inform(f string) (r Info, err error) {
	f, _ = filepath.Abs(f) // set here to avoid short path representation in windows

	mi, err := newMediaInfo()
	if err != nil {
		return
	}

	defer mi.Close()
	err = mi.OpenFile(f)
	if err != nil {
		return
	}

	info, err := mi.Inform()
	if err != nil {
		return
	}

	r.General.CompleteName = f

	for _, track := range info.Media.Tracks {
		switch track.Type {
		case "General":
			r.General.Duration = toFloat(track.Duration)
			r.General.UniqueID = track.UniqueID
			r.General.AudioCount = toUint(track.AudioCount)
			r.General.VideoCount = toUint(track.VideoCount)
			r.General.TextCount = toUint(track.TextCount)
			r.General.MenuCount = toUint(track.MenuCount)
			r.General.FileExtension = strings.ToLower(track.FileExtension)
			r.General.Format = track.Format
			r.General.FormatVersion = track.FormatVersion
			r.General.FileSize = toUint(track.FileSize)
			r.General.OverallBitRate = toFloat(track.OverallBitRate)
			r.General.FrameRate = toFloat(track.FrameRate)
			r.General.FrameCount = toUint(track.FrameCount)
			r.General.IsStreamable = toBool(track.IsStreamable)
			r.General.EncodedDate = toTime(track.EncodedDate)
			r.General.FileCreatedDate = toTime(track.FileCreatedDate)
			r.General.FileModifiedDate = toTime(track.FileModifiedDate)
			r.General.EncodedApplication = track.EncodedApplication
			r.General.EncodedLibrary = fmt.Sprintf("%s", track.EncodedLibrary)
			r.General.EncodedLibraryVersion = track.EncodedLibraryVersion
			r.General.Title = track.Title
		case "Video":
			r.VideoTracks = append(r.VideoTracks, Video{
				StreamOrder:            toUint(track.StreamOrder),
				ID:                     toUint(track.ID),
				UniqueID:               track.UniqueID,
				Format:                 track.Format,
				FormatProfile:          track.FormatProfile,
				FormatLevel:            track.FormatLevel,
				FormatTier:             track.FormatTier,
				CodecID:                track.CodecID,
				Duration:               toFloat(track.Duration),
				BitRate:                toFloat(track.BitRate),
				Width:                  toUint(track.Width),
				Height:                 toUint(track.Height),
				SampledWidth:           toUint(track.SampledWidth),
				SampledHeight:          toUint(track.Height),
				PixelAspectRatio:       toFloat(track.PixelAspectRatio),
				DisplayAspectRatio:     toFloat(track.DisplayAspectRatio),
				FrameRateMode:          track.FrameRateMode,
				FrameRate:              toFloat(track.FrameRate),
				FrameCount:             toUint(track.FrameCount),
				ColorSpace:             track.ColorSpace,
				ChromaSubsampling:      track.ChromaSubsampling,
				BitDepth:               toUint(track.BitDepth),
				StreamSize:             toUint(track.StreamSize),
				StreamSizeProportion:   toFloat(track.StreamSizeProportion),
				EncodedLibrary:         fmt.Sprintf("%s", track.EncodedLibrary),
				EncodedLibraryName:     track.EncodedLibraryName,
				EncodedLibraryVersion:  track.EncodedLibraryVersion,
				EncodedLibrarySettings: track.EncodedLibrarySettings,
				Default:                toBool(track.Default),
				Forced:                 toBool(track.Forced),
				B3D:                    track.MultiViewCount != "",
			})
		case "Audio":
			r.AudioTracks = append(r.AudioTracks, Audio{
				Channels:                 toUint(track.Channels),
				BitRate:                  toFloat(track.BitRate),
				ChannelLayout:            track.ChannelLayout,
				ChannelPositions:         track.ChannelPositions,
				CodecID:                  track.CodecID,
				CompressionMode:          track.CompressionMode,
				Default:                  toBool(track.Default),
				Duration:                 toFloat(track.Duration),
				Forced:                   toBool(track.Forced),
				Format:                   track.Format,
				FormatAdditionalFeatures: track.FormatAdditionalFeatures,
				FormatCommercial:         track.FormatCommercial,
				FrameCount:               toUint(track.FrameCount),
				FrameRate:                toFloat(track.FrameRate),
				ID:                       toUint(track.ID),
				Language:                 track.Language,
				SamplesPerFrame:          toUint(track.SamplesPerFrame),
				SamplingCount:            toUint(track.SamplingCount),
				SamplingRate:             toUint(track.SamplingRate),
				StreamOrder:              toUint(track.StreamOrder),
				StreamSize:               toUint(track.StreamSize),
				StreamSizeProportion:     toFloat(track.StreamSizeProportion),
				UniqueID:                 track.UniqueID,
			})
		case "Text":
			r.TextTracks = append(r.TextTracks, Text{
				ElementCount: toUint(track.ElementCount),
				BitRate:      toFloat(track.BitRate),
				CodecID:      track.CodecID,
				Default:      toBool(track.Default),
				Duration:     toFloat(track.Duration),
				Forced:       toBool(track.Forced),
				Format:       track.Format,
				FrameCount:   toUint(track.FrameCount),
				ID:           toUint(track.ID),
				Language:     track.Language,
				StreamOrder:  toUint(track.StreamOrder),
				StreamSize:   toUint(track.StreamSize),
				Order:        toUint(track.TypeOrder),
				UniqueID:     track.UniqueID,
			})
		case "Menu":
			m := Menu{
				Order:    toUint(track.TypeOrder),
				Duration: r.General.Duration,
			}

			for k, v := range track.Extra {
				if !strings.HasPrefix(k, "_") {
					// cases where menu has extra information
					continue
				}
				// k = time
				// v = lang:title
				ss := formatExtra(k)
				st := formatTime(ss)
				idx := strings.Index(v, ":")

				m.Entries = append(m.Entries, Entry{
					StartTime:    st,
					StartTimeStr: ss,
					Language:     v[:idx],
					Title:        v[idx+1:],
				})
			}
			// sort now to make sure endtime last entry is properly set
			// since track.Extra is a map
			sort.Slice(m.Entries, func(i, j int) bool {
				return m.Entries[i].StartTime < m.Entries[j].StartTime
			})

			for i := 0; i < len(m.Entries)-1; i++ {
				m.Entries[i].EndTime = m.Entries[i+1].StartTime
				m.Entries[i].EndTimeStr = m.Entries[i+1].StartTimeStr
			}

			m.Entries[len(m.Entries)-1].EndTime = m.Duration
			m.Entries[len(m.Entries)-1].EndTimeStr = toFormatTimeStr(m.Duration)

			r.MenuTracks = append(r.MenuTracks, m)
		}
	}

	return
}

func toUint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

func toFloat(s string) float32 {
	f, _ := strconv.ParseFloat(s, 64)
	return float32(f)
}

func toBool(s string) bool {
	s = strings.ToUpper(s)
	if s == "YES" {
		return true
	}
	return false
}

func toTime(s string) time.Time {
	// layout example UTC 2020-10-20 19:04:07
	t, _ := time.Parse("MST 2006-01-02 15:04:05", s)
	return t
}

// s in format _00_00_00_000 => 00:00:00.000
func formatExtra(s string) string {
	entries := strings.Split(s[1:], "_")
	lastIndex := len(entries) - 1
	return strings.Join(entries[:lastIndex], ":") + "." + entries[lastIndex]
}

// 00:00:00.000 => (00*60*60)+(00*60)+00,000
func formatTime(s string) float32 {
	t1 := strings.Split(s, ".")
	hhmmss := t1[0]
	t2 := strings.Split(hhmmss, ":")

	hh := toUint(t2[0]) * 60 * 60
	mm := toUint(t2[1]) * 60
	ss := toUint(t2[2])
	ms := float32(toUint(t1[1])) / 1000
	return float32(hh+mm+ss) + ms

}

func toFormatTimeStr(f float32) string {
	// Rounded to int64 as soon as possible to avoid floating point precision errors
	t := time.Unix(0, int64(f*1000)*1000*1000) // to nanoseconds.
	t = t.UTC()
	return t.Format("15:04:05.000")
}
