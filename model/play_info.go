package model

type PlayInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Dash              struct {
			Duration       int     `json:"duration"`
			MinBufferTime  float64 `json:"minBufferTime"`
			MinBufferTime2 float64 `json:"min_buffer_time"`
			Video          []struct {
				ID            int      `json:"id"`
				BaseURL       string   `json:"baseUrl"`
				BaseURL2      string   `json:"base_url"`
				BackupURL     []string `json:"backupUrl"`
				BackupURL2    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType2     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate2    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap2 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase2 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				ID            int      `json:"id"`
				BaseURL       string   `json:"baseUrl"`
				BaseURL2      string   `json:"base_url"`
				BackupURL     []string `json:"backupUrl"`
				BackupURL2    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType2     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate2    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap2 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase2 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby interface{} `json:"dolby"`
			Flac  interface{} `json:"flac"`
		} `json:"dash"`
		SupportFormats []struct {
			Quality        int      `json:"quality"`
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
		} `json:"support_formats"`
		HighFormat interface{} `json:"high_format"`
		Volume     struct {
			MeasuredI         float64 `json:"measured_i"`
			MeasuredLra       int     `json:"measured_lra"`
			MeasuredTp        float64 `json:"measured_tp"`
			MeasuredThreshold float64 `json:"measured_threshold"`
			TargetOffset      int     `json:"target_offset"`
			TargetI           int     `json:"target_i"`
			TargetTp          int     `json:"target_tp"`
		} `json:"volume"`
	} `json:"data"`
	Session string `json:"session"`
}
