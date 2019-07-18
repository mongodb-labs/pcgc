package opsmanager

// Net part of the internal Process struct
type Net struct {
	Port int `json:"port,omitempty"`
}

// Storage part of the internal Process struct
type Storage struct {
	DBPath string `json:"dbPath,omitempty"`
}

// SystemLog part of the internal Process struct
type SystemLog struct {
	Destination string `json:"destination,omitempty"`
	Path        string `json:"path,omitempty"`
}

// Args26 part of the internal Process struct
type Args26 struct {
	NET       *Net       `json:"net,omitempty"`
	Storage   *Storage   `json:"storage,omitempty"`
	SystemLog *SystemLog `json:"systemLog,omitempty"`
}

// LogRotate part of the internal Process struct
type LogRotate struct {
	SizeThresholdMB  int `json:"sizeThresholdMB,omitempty"`
	TimeThresholdHrs int `json:"timeThresholdHrs,omitempty"`
}

// Process represents a single process in a deployment
type Process struct {
	Name                        string     `json:"name,omitempty"`
	ProcessType                 string     `json:"processType,omitempty"`
	Version                     string     `json:"version,omitempty"`
	AuthSchemaVersion           int        `json:"authSchemaVersion,omitempty"`
	FeatureCompatibilityVersion string     `json:"featureCompatibilityVersion,omitempty"`
	Disabled                    bool       `json:"disabled,omitempty"`
	ManualMode                  bool       `json:"manualMode,omitempty"`
	Hostname                    string     `json:"hostname,omitempty"`
	Args26                      *Args26    `json:"args2_6,omitempty"`
	LogRotate                   *LogRotate `json:"logRotate,omitempty"`
}
