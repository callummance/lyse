package models

type CategoryType int
const (
  Dungeon CategoryType = iota
  Trial
  Raid
  SavageRaid
  Job
  Misc
)

type PatchNumber struct {
  Expansion           int
  MajorPatchNumber    int
  MinorPatchNumber    int
}

type TriggerTransmisssion struct {
  Regex         string    `json:"regex" xml:"regex"`
  SoundType     int       `json:"stype" xml:"stype"`
  SoundData     string    `json:"sdata" xml:"sdata"`
  HasTimer      bool      `json:"timer" xml:"timer"`
  TimerName     string    `json:"tname" xml:"tname"`
  HasTab        bool      `json:"tabbed" xml:"tabbed"`
  RestrictZone  bool      `json:"restrict" xml:"restrict"`
}

type Category struct {
  Name            string
  Type            CategoryType
  Patch           *PatchNumber
  RestrictZone    bool
}

type Trigger struct {
  Regex       string
  SoundType   int
  SoundData   string
  HasTimer    bool
  LinkedTimer Timer
  Category    Category
}

type Timer struct {
  Name              string
  OnlyMasterTicks   bool
  TimerDuration     int
  RestrictToMe      bool
  AbsoluteTiming    bool
  StartSoundData    string
  WarningSoundData  string
  WarningTriggerVal int
  RadialDisplay     bool
  Modable           bool
  Tooltip           string
  FillColour        string
  Panel1            bool
  Panel2            bool
}
